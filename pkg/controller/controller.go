package controller

import (
	"github.com/appscode/go/log"
	apiext_util "github.com/appscode/kutil/apiextensions/v1beta1"
	"github.com/appscode/kutil/tools/queue"
	pcm "github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
	api "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	cs "github.com/kubedb/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1"
	kutildb "github.com/kubedb/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1/util"
	api_listers "github.com/kubedb/apimachinery/client/listers/kubedb/v1alpha1"
	amc "github.com/kubedb/apimachinery/pkg/controller"
	"github.com/kubedb/apimachinery/pkg/controller/dormantdatabase"
	snapc "github.com/kubedb/apimachinery/pkg/controller/snapshot"
	"github.com/kubedb/apimachinery/pkg/eventer"
	"github.com/kubedb/mongodb/pkg/docker"
	core "k8s.io/api/core/v1"
	crd_api "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/tools/reference"
)

type Controller struct {
	amc.Config
	*amc.Controller

	docker docker.Docker
	// Prometheus client
	promClient pcm.MonitoringV1Interface
	// Cron Controller
	cronController snapc.CronControllerInterface
	// Event Recorder
	recorder record.EventRecorder

	// MongoDB
	mgQueue    *queue.Worker
	mgInformer cache.SharedIndexInformer
	mgLister   api_listers.MongoDBLister
}

var _ amc.Snapshotter = &Controller{}
var _ amc.Deleter = &Controller{}

func New(
	client kubernetes.Interface,
	apiExtKubeClient crd_cs.ApiextensionsV1beta1Interface,
	extClient cs.KubedbV1alpha1Interface,
	promClient pcm.MonitoringV1Interface,
	cronController snapc.CronControllerInterface,
	docker docker.Docker,
	opt amc.Config,
) *Controller {
	return &Controller{
		Controller: &amc.Controller{
			Client:           client,
			ExtClient:        extClient,
			ApiExtKubeClient: apiExtKubeClient,
		},
		Config:         opt,
		docker:         docker,
		promClient:     promClient,
		cronController: cronController,
		recorder:       eventer.NewEventRecorder(client, "MongoDB operator"),
	}
}

// Ensuring Custom Resource Definitions
func (c *Controller) EnsureCustomResourceDefinitions() error {
	log.Infoln("Ensuring CustomResourceDefinition...")
	crds := []*crd_api.CustomResourceDefinition{
		api.MongoDB{}.CustomResourceDefinition(),
		api.DormantDatabase{}.CustomResourceDefinition(),
		api.Snapshot{}.CustomResourceDefinition(),
	}
	return apiext_util.RegisterCRDs(c.ApiExtKubeClient, crds)
}

// Init initializes MongoDB, DormantDB amd Snapshot watcher
func (c *Controller) Init() error {
	labelMap := map[string]string{
		api.LabelDatabaseKind: api.ResourceKindMongoDB,
	}

	tweakListOptions := func(options *metav1.ListOptions) {
		options.LabelSelector = labels.SelectorFromSet(labelMap).String()
	}

	if err := c.EnsureCustomResourceDefinitions(); err != nil {
		return err
	}

	c.initWatcher()
	c.DDBQueue = dormantdatabase.NewController(c.Controller, c, c.Config, tweakListOptions).Init()
	c.SNQueue, c.JobQueue = snapc.NewController(c.Controller, c, c.Config, tweakListOptions).Init()

	return nil
}

// RunControllers runs queue.worker
func (c *Controller) RunControllers(stopCh <-chan struct{}) {
	// Start Cron
	c.cronController.StartCron()

	// Watch x  TPR objects
	c.mgQueue.Run(stopCh)
	c.DDBQueue.Run(stopCh)
	c.SNQueue.Run(stopCh)
	c.JobQueue.Run(stopCh)
}

// Blocks caller. Intended to be called as a Go routine.
func (c *Controller) Run(stopCh <-chan struct{}) {
	go c.StartAndRunControllers(stopCh)

	// Run HTTP server to expose metrics, audit endpoint & debug profiles.
	go c.runHTTPServer()

	<-stopCh
	c.cronController.StopCron()
}

// StartAndRunControllers starts InformetFactory and runs queue.worker
func (c *Controller) StartAndRunControllers(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()

	log.Infoln("Starting KubeDB controller")
	c.KubeInformerFactory.Start(stopCh)
	c.KubedbInformerFactory.Start(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	for t, v := range c.KubeInformerFactory.WaitForCacheSync(stopCh) {
		if !v {
			log.Fatalf("%v timed out waiting for caches to sync\n", t)
			return
		}
	}
	for t, v := range c.KubedbInformerFactory.WaitForCacheSync(stopCh) {
		if !v {
			log.Fatalf("%v timed out waiting for caches to sync\n", t)
			return
		}
	}

	c.RunControllers(stopCh)

	<-stopCh
	log.Infoln("Stopping KubeDB controller")
}

func (c *Controller) pushFailureEvent(mongodb *api.MongoDB, reason string) {
	if ref, rerr := reference.GetReference(clientsetscheme.Scheme, mongodb); rerr == nil {
		c.recorder.Eventf(
			ref,
			core.EventTypeWarning,
			eventer.EventReasonFailedToStart,
			`Fail to be ready MongoDB: "%v". Reason: %v`,
			mongodb.Name,
			reason,
		)
	}

	mg, _, err := kutildb.PatchMongoDB(c.ExtClient, mongodb, func(in *api.MongoDB) *api.MongoDB {
		in.Status.Phase = api.DatabasePhaseFailed
		in.Status.Reason = reason
		return in
	})
	if err != nil {
		if ref, rerr := reference.GetReference(clientsetscheme.Scheme, mongodb); rerr == nil {
			c.recorder.Eventf(
				ref,
				core.EventTypeWarning,
				eventer.EventReasonFailedToUpdate,
				err.Error(),
			)
		}
	}
	mongodb.Status = mg.Status
}

package framework

import (
	"time"

	"github.com/appscode/go/crypto/rand"
	"github.com/appscode/go/encoding/json/types"
	tapi "github.com/k8sdb/apimachinery/apis/kubedb/v1alpha1"
	kutildb "github.com/k8sdb/apimachinery/client/typed/kubedb/v1alpha1/util"
	. "github.com/onsi/gomega"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Invocation) MongoDB() *tapi.MongoDB {
	return &tapi.MongoDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.WithUniqSuffix("mongodb"),
			Namespace: f.namespace,
			Labels: map[string]string{
				"app": f.app,
			},
		},
		Spec: tapi.MongoDBSpec{
			Version: types.StrYo("3.4"),
		},
	}
}

func (f *Framework) CreateMongoDB(obj *tapi.MongoDB) error {
	_, err := f.extClient.MongoDBs(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) GetMongoDB(meta metav1.ObjectMeta) (*tapi.MongoDB, error) {
	return f.extClient.MongoDBs(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
}

func (f *Framework) TryPatchMongoDB(meta metav1.ObjectMeta, transform func(*tapi.MongoDB) *tapi.MongoDB) (*tapi.MongoDB, error) {
	return kutildb.TryPatchMongoDB(f.extClient, meta, transform)
}

func (f *Framework) DeleteMongoDB(meta metav1.ObjectMeta) error {
	return f.extClient.MongoDBs(meta.Namespace).Delete(meta.Name, &metav1.DeleteOptions{})
}

func (f *Framework) EventuallyMongoDB(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.extClient.MongoDBs(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			if err != nil {
				if kerr.IsNotFound(err) {
					return false
				} else {
					Expect(err).NotTo(HaveOccurred())
				}
			}
			return true
		},
		time.Minute*5,
		time.Second*5,
	)
}

func (f *Framework) EventuallyMongoDBRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			mongodb, err := f.extClient.MongoDBs(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return mongodb.Status.Phase == tapi.DatabasePhaseRunning
		},
		time.Minute*5,
		time.Second*5,
	)
}

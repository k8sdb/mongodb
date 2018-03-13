package validator

import (
	"fmt"

	"github.com/appscode/go/types"
	api "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	cs "github.com/kubedb/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1"
	amv "github.com/kubedb/apimachinery/pkg/validator"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
)

var (
	mongodbVersions = sets.NewString("3.4", "3.6")
)

func ValidateMongoDB(client kubernetes.Interface, extClient cs.KubedbV1alpha1Interface, mongodb *api.MongoDB) error {
	if mongodb.Spec.Version == "" {
		return fmt.Errorf(`object 'Version' is missing in '%v'`, mongodb.Spec)
	}

	// check MongoDB version validation
	if !mongodbVersions.Has(string(mongodb.Spec.Version)) {
		return fmt.Errorf(`KubeDB doesn't support MongoDB version: %s`, string(mongodb.Spec.Version))
	}

	if mongodb.Spec.Replicas != nil {
		replicas := types.Int32(mongodb.Spec.Replicas)
		if replicas != 1 {
			return fmt.Errorf(`spec.replicas "%d" invalid. Value must be one`, replicas)
		}
	}

	if mongodb.Spec.Storage != nil {
		var err error
		if err = amv.ValidateStorage(client, mongodb.Spec.Storage); err != nil {
			return err
		}
	}

	backupScheduleSpec := mongodb.Spec.BackupSchedule
	if backupScheduleSpec != nil {
		if err := amv.ValidateBackupSchedule(client, backupScheduleSpec, mongodb.Namespace); err != nil {
			return err
		}
	}

	monitorSpec := mongodb.Spec.Monitor
	if monitorSpec != nil {
		if err := amv.ValidateMonitorSpec(monitorSpec); err != nil {
			return err
		}
	}

	if err := killMatchingDormantDatabase(extClient, mongodb); err != nil {
		return err
	}

	return nil
}

func killMatchingDormantDatabase(extClient cs.KubedbV1alpha1Interface, mongodb *api.MongoDB) error {
	// Check if DormantDatabase exists or not
	if _, err := extClient.DormantDatabases(mongodb.Namespace).Get(mongodb.Name, metav1.GetOptions{}); err != nil {
		if !kerr.IsNotFound(err) {
			return err
		}
		return nil
	}

	// Delete  Matching dormantDatabase
	deletePolicy := metav1.DeletePropagationForeground
	if err := extClient.DormantDatabases(mongodb.Namespace).Delete(mongodb.Name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil && !kerr.IsNotFound(err) {
		return err
	}

	return nil
}

package validator

import (
	"fmt"

	api "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	amv "github.com/kubedb/apimachinery/pkg/validator"
	dr "github.com/kubedb/mongodb/pkg/docker"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	mongodbVersions = []string{"3.4", "3.6"}
)

func ValidateMongoDB(client kubernetes.Interface, mongodb *api.MongoDB, docker *dr.Docker) error {
	if mongodb.Spec.Version == "" {
		return fmt.Errorf(`object 'Version' is missing in '%v'`, mongodb.Spec)
	}

	// check MongoDB version validation
	exists := false
	for _, ver := range mongodbVersions {
		if ver == string(mongodb.Spec.Version) {
			exists = true
		}
	}
	if !exists {
		return fmt.Errorf(`image %s not found`, docker.GetImageWithTag(mongodb))
	}

	if mongodb.Spec.Storage != nil {
		var err error
		if err = amv.ValidateStorage(client, mongodb.Spec.Storage); err != nil {
			return err
		}
	}

	databaseSecret := mongodb.Spec.DatabaseSecret
	if databaseSecret != nil {
		if _, err := client.CoreV1().Secrets(mongodb.Namespace).Get(databaseSecret.SecretName, metav1.GetOptions{}); err != nil {
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
	return nil
}

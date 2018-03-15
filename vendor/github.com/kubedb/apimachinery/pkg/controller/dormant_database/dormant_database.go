package dormant_database

import (
	api "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	"github.com/kubedb/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1/util"
	"github.com/kubedb/apimachinery/pkg/eventer"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Controller) create(ddb *api.DormantDatabase) error {
	if ddb.Status.CreationTime == nil {
		_, _, err := util.PatchDormantDatabase(c.ExtClient, ddb, func(in *api.DormantDatabase) *api.DormantDatabase {
			t := metav1.Now()
			in.Status.CreationTime = &t
			return in
		})
		if err != nil {
			c.recorder.Eventf(ddb.ObjectReference(), core.EventTypeWarning, eventer.EventReasonFailedToUpdate, err.Error())
			return err
		}
	}

	if ddb.Status.Phase == api.DormantDatabasePhasePaused {
		return nil
	}

	_, _, err := util.PatchDormantDatabase(c.ExtClient, ddb, func(in *api.DormantDatabase) *api.DormantDatabase {
		in.Status.Phase = api.DormantDatabasePhasePausing
		return in
	})
	if err != nil {
		c.recorder.Eventf(ddb, core.EventTypeWarning, eventer.EventReasonFailedToUpdate, err.Error())
		return err
	}

	c.recorder.Event(ddb, core.EventTypeNormal, eventer.EventReasonPausing, "Pausing Database")

	// Pause Database workload
	if err := c.deleter.ExDatabaseStatus(ddb); err != nil {
		c.recorder.Eventf(
			ddb.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonFailedToDelete,
			"Failed to pause. Reason: %v",
			err,
		)
		return err
	}

	c.recorder.Event(
		ddb.ObjectReference(),
		core.EventTypeNormal,
		eventer.EventReasonSuccessfulPause,
		"Successfully paused Database workload",
	)

	_, _, err = util.PatchDormantDatabase(c.ExtClient, ddb, func(in *api.DormantDatabase) *api.DormantDatabase {
		t := metav1.Now()
		in.Status.PausingTime = &t
		in.Status.Phase = api.DormantDatabasePhasePaused
		return in
	})
	if err != nil {
		c.recorder.Eventf(ddb, core.EventTypeWarning, eventer.EventReasonFailedToUpdate, err.Error())
		return err
	}

	return nil
}

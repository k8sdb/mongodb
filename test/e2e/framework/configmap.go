package framework

import (
	"strings"

	"github.com/appscode/go/crypto/rand"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Invocation) ConfigMapForInitialization() *core.ConfigMap {
	return &core.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.WithUniqSuffix(i.app + "-local"),
			Namespace: i.namespace,
		},
		Data: map[string]string{
			"init.js": `db = db.getSiblingDB('kubedb');
db.people.insert({"firstname" : "kubernetes", "lastname" : "database" }); `,
			"init-in-existing-db.js": `db.people.insert({"firstname" : "kubernetes", "lastname" : "database" });`,
		},
	}
}

func (i *Invocation) GetCustomConfig(configs []string) *core.ConfigMap {
	configs = append([]string{"net:"}, configs...)
	return &core.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      i.app,
			Namespace: i.namespace,
		},
		Data: map[string]string{
			"mongod.conf": strings.Join(configs, "\n"),
		},
	}
}

func (i *Invocation) CreateConfigMap(obj *core.ConfigMap) error {
	_, err := i.kubeClient.CoreV1().ConfigMaps(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) DeleteConfigMap(meta metav1.ObjectMeta) error {
	err := f.kubeClient.CoreV1().ConfigMaps(meta.Namespace).Delete(meta.Name, deleteInForeground())
	if !kerr.IsNotFound(err) {
		return err
	}
	return nil
}

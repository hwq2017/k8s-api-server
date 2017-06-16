package k8s

import (
// "fmt"
// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type K8s struct {
	router Router
	client Clientset
}

func (k *K8s) Run() {
	k.client.Init()
	k.router.RunServer(&k.client)
}

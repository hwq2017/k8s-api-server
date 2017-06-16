package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/pkg/api/v1"
)

type Namespaces struct {
	client *Clientset
}

func (n *Namespaces) List() (result *v1.NamespaceList) {
	result, err := n.client.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	CheckErr(err)
	return
}

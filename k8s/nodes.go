package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/pkg/api/v1"
)

type Nodes struct {
	client *Clientset
}

func (n *Nodes) List() (result *v1.NodeList) {
	result, err := n.client.clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	CheckErr(err)
	return
}

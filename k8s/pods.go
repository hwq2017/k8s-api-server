package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/pkg/api/v1"
)

type Pods struct {
	client *Clientset
}

func (p *Pods) List(namespaces string) (result *v1.PodList) {
	// result, err := p.client.clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	result, err := p.client.clientset.CoreV1().Pods(namespaces).List(metav1.ListOptions{})
	CheckErr(err)
	return
}

func (p *Pods) Create(name, namespaces, image string) (result *v1.Pod, err error) {
	pod := new(v1.Pod)
	pod.ObjectMeta = metav1.ObjectMeta{Name: name, Namespace: namespaces}
	pod.Spec = v1.PodSpec{
		RestartPolicy: v1.RestartPolicyAlways,
		Containers: []v1.Container{
			v1.Container{
				Name:  name,
				Image: image,
			},
		},
	}
	result, err = p.client.clientset.CoreV1().Pods(namespaces).Create(pod)
	return
}

func (p *Pods) Delete(name, namespaces string) error {
	err := p.client.clientset.CoreV1().Pods(namespaces).Delete(name, &metav1.DeleteOptions{})
	return err
}

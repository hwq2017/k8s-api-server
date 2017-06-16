package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/pkg/api/v1"
)

type Service struct {
	client *Clientset
}

func (s *Service) List() (result *v1.ServiceList) {
	result, err := s.client.clientset.CoreV1().Services("").List(metav1.ListOptions{})
	CheckErr(err)
	return
}

func (s *Service) Create(name, namespaces string, port, nodeport int32) (result *v1.Service, err error) {
	service := new(v1.Service)
	service.ObjectMeta = metav1.ObjectMeta{Name: name, Namespace: namespaces}
	service.Spec = v1.ServiceSpec{
		Selector: map[string]string{
			"app": name,
		},
		Type: "NodePort",
		Ports: []v1.ServicePort{
			v1.ServicePort{
				Port:     port,
				NodePort: nodeport,
			},
		},
	}
	result, err = s.client.clientset.CoreV1().Services(namespaces).Create(service)
	return
}

func (s *Service) Delete(name, namespaces string) error {
	err := s.client.clientset.CoreV1().Services(namespaces).Delete(name, &metav1.DeleteOptions{})
	return err
}

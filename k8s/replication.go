package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/pkg/api/v1"
)

type Replication struct {
	client *Clientset
}

func (r *Replication) List() (result *v1.ReplicationControllerList) {
	result, err := r.client.clientset.CoreV1().ReplicationControllers("").List(metav1.ListOptions{})
	CheckErr(err)
	return
}

func (r *Replication) Create(name, namespaces, image string, replicas *int32, containerport int32) (result *v1.ReplicationController, err error) {
	replication := new(v1.ReplicationController)
	replication.ObjectMeta = metav1.ObjectMeta{Name: name, Namespace: namespaces}
	replication.Spec = v1.ReplicationControllerSpec{
		Replicas: replicas,
		Selector: map[string]string{
			"app": name,
		},
		Template: &v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": name,
				},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					v1.Container{
						Name:  name,
						Image: image,
						Ports: []v1.ContainerPort{
							v1.ContainerPort{
								ContainerPort: containerport,
							},
						},
					},
				},
			},
		},
	}
	result, err = r.client.clientset.CoreV1().ReplicationControllers(namespaces).Create(replication)
	return
}

func (r *Replication) Delete(name, namespaces string) error {
	err := r.client.clientset.CoreV1().ReplicationControllers(namespaces).Delete(name, &metav1.DeleteOptions{})
	return err
}

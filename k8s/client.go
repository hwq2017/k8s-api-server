package k8s

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Clientset struct {
	clientset *kubernetes.Clientset
}

func (c *Clientset) Init() {
	kubeconfig := flag.String("kubeconfig", "/Users/zhitan/dev/go/src/k8s-api-server/k8s/config", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	c.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

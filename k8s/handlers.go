package k8s

type Handler struct {
	pods        Pods
	nodes       Nodes
	namespaces  Namespaces
	replication Replication
	service     Service
}

func (h *Handler) Init(client *Clientset) {
	h.pods.client = client
	h.nodes.client = client
	h.namespaces.client = client
	h.replication.client = client
	h.service.client = client
}

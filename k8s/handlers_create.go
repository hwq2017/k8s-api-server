package k8s

import (
	"net/http"
)

func (h *Handler) PodsCreate(rw http.ResponseWriter, req *http.Request) {
	podspec := JsonUnmarshal(rw, req)

	var name string = podspec["name"].(string)
	var namespaces string = podspec["namespaces"].(string)
	var image string = podspec["image"].(string)

	result, err := h.pods.Create(name, namespaces, image)
	if err != nil {
		panic(err.Error())
	}

	Response(result, rw)
}

func (h *Handler) ReplicationCreate(rw http.ResponseWriter, req *http.Request) {
	replication := JsonUnmarshal(rw, req)

	var name string = replication["name"].(string)
	var namespaces string = replication["namespaces"].(string)
	var image string = replication["image"].(string)
	var replicas int32 = int32(replication["replicas"].(float64))
	var containerport int32 = int32(replication["containerport"].(float64))

	result, err := h.replication.Create(name, namespaces, image, &replicas, containerport)
	if err != nil {
		panic(err.Error())
	}

	Response(result, rw)
}

func (h *Handler) ServiceCreate(rw http.ResponseWriter, req *http.Request) {
	servicespec := JsonUnmarshal(rw, req)

	var name string = servicespec["name"].(string)
	var namespaces string = servicespec["namespaces"].(string)
	var port int32 = int32(servicespec["port"].(float64))
	var nodeport int32 = int32(servicespec["nodeport"].(float64))

	result, err := h.service.Create(name, namespaces, port, nodeport)
	if err != nil {
		panic(err.Error())
	}

	Response(result, rw)
}

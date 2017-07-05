package k8s

import (
	"net/http"
)

func (h *Handler) PodsDelete(rw http.ResponseWriter, req *http.Request) {
	podspec := JsonUnmarshal(rw, req)

	var msg string

	podsname := podspec["name"].([]interface{})
	namespaces := podspec["namespaces"].(string)
	for _, name := range podsname {
		err := h.pods.Delete(name.(string), namespaces)
		if err != nil {
			msg = err.Error()
		} else {
			msg = "Success"
		}
	}
	Response(msg, rw)

}

func (h *Handler) ServiceDelete(rw http.ResponseWriter, req *http.Request) {
	servicespec := JsonUnmarshal(rw, req)

	var name string = servicespec["name"].(string)
	var namespaces string = servicespec["namespaces"].(string)

	result, err := "success", h.service.Delete(name, namespaces)
	if err != nil {
		panic(err.Error())
	}

	Response(result, rw)
}

func (h *Handler) ReplicationDelete(rw http.ResponseWriter, req *http.Request) {
	replicationspec := JsonUnmarshal(rw, req)

	var name string = replicationspec["name"].(string)
	var namespaces string = replicationspec["namespaces"].(string)

	result, err := "success", h.replication.Delete(name, namespaces)
	if err != nil {
		panic(err.Error())
	}

	Response(result, rw)
}

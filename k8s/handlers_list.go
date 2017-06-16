package k8s

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) PodsList(rw http.ResponseWriter, req *http.Request) {
	result := h.pods.List()
	resp, _ := json.Marshal(result.Items)
	fmt.Fprintf(rw, string(resp))
}

func (h *Handler) NodesList(rw http.ResponseWriter, req *http.Request) {
	result := h.nodes.List()
	resp, _ := json.Marshal(result.Items)
	fmt.Fprintf(rw, string(resp))
}

func (h *Handler) NamespacesList(rw http.ResponseWriter, req *http.Request) {
	result := h.namespaces.List()
	resp, _ := json.Marshal(result.Items)
	fmt.Fprintf(rw, string(resp))
}

func (h *Handler) ReplicationList(rw http.ResponseWriter, req *http.Request) {
	result := h.replication.List()
	resp, _ := json.Marshal(result.Items)
	fmt.Fprintf(rw, string(resp))
}

func (h *Handler) ServiceList(rw http.ResponseWriter, req *http.Request) {
	result := h.service.List()
	resp, _ := json.Marshal(result.Items)
	fmt.Fprintf(rw, string(resp))
}

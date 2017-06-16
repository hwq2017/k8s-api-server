package k8s

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	handler Handler
}

func (r *Router) RunServer(client *Clientset) {
	r.handler.Init(client)

	m := mux.NewRouter()
	m.HandleFunc("/pods/list", r.handler.PodsList)
	m.HandleFunc("/pods/create", r.handler.PodsCreate)
	m.HandleFunc("/pods/delete", r.handler.PodsDelete)

	m.HandleFunc("/nodes/list", r.handler.NodesList)

	m.HandleFunc("/replications/list", r.handler.ReplicationList)
	m.HandleFunc("/replications/create", r.handler.ReplicationCreate)
	m.HandleFunc("/replications/delete", r.handler.ReplicationDelete)

	m.HandleFunc("/namespaces/list", r.handler.NamespacesList)

	m.HandleFunc("/services/list", r.handler.ServiceList)
	m.HandleFunc("/services/create", r.handler.ServiceCreate)
	m.HandleFunc("/services/delete", r.handler.ServiceDelete)
	http.ListenAndServe("10.100.34.157:8080", m)
}

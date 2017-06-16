package k8s

type PodSpec struct {
	Name       string `json:name`
	Namespaces string `json:namespaces`
	Image      string `json:image`
}

type ReplicationSpec struct {
	Name          string `json:name`
	Namespaces    string `json:namespaces`
	Image         string `json:image`
	Replicas      int32  `json:replicas`
	Containerport int32  `json:containerport`
}

type ServiceSpec struct {
	Name       string `json:name`
	Namespaces string `json:namespaces`
	Port       int32  `json:port`
	Nodeport   int32  `json:nodeport`
}

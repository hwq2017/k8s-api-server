package main

import (
	"k8s-api-server/k8s"
)

func main() {
	var k k8s.K8s
	k.Run()
}

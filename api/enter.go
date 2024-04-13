package api

import (
	"kube-auto/api/example"
	"kube-auto/api/k8s"
	"kube-auto/api/user"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8SApiGroup     k8s.PodApi
	LoginApiGroup   user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

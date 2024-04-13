package router

import (
	"kube-auto/router/example"
	"kube-auto/router/k8s"
	"kube-auto/router/user"
)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
	K8sRouterGroup     k8s.K8sRouter
	UserLoginGroup     user.LoginRouter
}

var RouterGroupApp = new(RouterGroup)

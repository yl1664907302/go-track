package router

import router "kube-auto/router/example"

type RouterGroup struct {
	ExampleRouterGroup router.ExampleRouter
}

var RouterGroupApp = new(RouterGroup)

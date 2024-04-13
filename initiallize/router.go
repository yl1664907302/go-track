package initiallize

import (
	"github.com/gin-gonic/gin"
	"kube-auto/middleware"
	"kube-auto/router"
)

func Router() *gin.Engine {
	//gin.Default() 返回一个 gin.Engine
	r := gin.Default()
	r.Use(middleware.Cors)
	g1 := router.RouterGroupApp.ExampleRouterGroup
	g1.InitExample(r)
	g2 := router.RouterGroupApp.K8sRouterGroup
	g2.InitK8s(r)
	g3 := router.RouterGroupApp.UserLoginGroup
	g3.Login(r)
	return r
}

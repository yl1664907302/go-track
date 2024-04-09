package initiallize

import (
	"github.com/gin-gonic/gin"
	"kube-auto/router"
)

func Router() *gin.Engine {
	//gin.Default() 返回一个 gin.Engine
	r := gin.Default()
	group := router.RouterGroupApp.ExampleRouterGroup
	group.InitExample(r)
	return r
}

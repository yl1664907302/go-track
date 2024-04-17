package example

import (
	"github.com/gin-gonic/gin"
	"kube-auto/api"
)

type ExampleRouter struct {
}

func (*ExampleRouter) InitExample(r *gin.Engine) {
	//首次默认路径
	group := r.Group("/example")
	apiGroup := api.ApiGroupApp.ExampleApiGroup
	//追加路径
	group.POST("/ping", apiGroup.ExampleTest)
}

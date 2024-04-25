package example

import (
	"github.com/gin-gonic/gin"
	"kube-auto/api"
)

type ExampleRouter struct {
}

func (*ExampleRouter) InitExample(r *gin.Engine) {
	apiGroup := api.ApiGroupApp.ExampleApiGroup
	r.GET("/test", apiGroup.ExampleTest)
}

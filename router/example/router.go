package example

import (
	"github.com/gin-gonic/gin"
	"go-track/api"
)

type ExampleRouter struct {
}

func (*ExampleRouter) InitExample(r *gin.Engine) {
	apiGroup := api.ApiGroupApp.ExampleApiGroup
	r.GET("/test", apiGroup.ExampleTest)
}

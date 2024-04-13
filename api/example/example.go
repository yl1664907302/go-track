package example

import (
	"github.com/gin-gonic/gin"
	"kube-auto/response"
)

type ExampleApi struct {
}

func (*ExampleApi) ExampleTest(c *gin.Context) {
	response.SuccssWithDetailed(c, "请求数据成功", map[string]string{
		"message": "pong",
	})
}

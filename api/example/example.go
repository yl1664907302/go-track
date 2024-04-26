package example

import (
	"github.com/gin-gonic/gin"
	"go-track/response"
)

type ExampleApi struct {
}

func (*ExampleApi) ExampleTest(c *gin.Context) {
	//body, err := c.GetRawData()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//fmt.Println(string(body))
	response.SuccssWithDetailed(c, "测试成功", map[string]string{
		"code": "200",
	})
}

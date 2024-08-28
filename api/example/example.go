package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-track/database/mysql"
	"go-track/response"
	"net/http"
)

type ExampleApi struct {
}

func (*ExampleApi) ExampleTest(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = mysql.InsertReceiver("l1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Println(string(body))
	response.SuccssWithDetailed(c, "测试成功", map[string]string{
		"code": "",
	})
}

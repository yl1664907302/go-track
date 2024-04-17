package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	fmt.Println(string(body))
}

package feishu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kube-auto/global"
	"kube-auto/pojo"
	"net/http"
)

type FeishuApi struct {
}

func (*FeishuApi) GetFeishuMessage(c *gin.Context) {
	var feishu pojo.FeishuMarkdownMessage
	err := c.ShouldBind(&feishu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}
	for _, element := range feishu.Card.Elements {
		for k, v := range global.ActionMessage.ExtractInfo(element.Content) {
			fmt.Println(k, v)
		}
	}
}

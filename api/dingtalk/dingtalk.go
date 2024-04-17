package dingtalk

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kube-auto/global"
	"kube-auto/pojo"
	"net/http"
)

type DingTalkApi struct {
}

func (*DingTalkApi) GetDingTalkMessage(c *gin.Context) {
	var dingtalk pojo.DingtalkMarkdownMessage
	err := c.ShouldBind(&dingtalk)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}
	for _, info := range global.ActionMessage.ExtractInfo(dingtalk.Markdown.Text) {
		fmt.Println(info)
	}
}

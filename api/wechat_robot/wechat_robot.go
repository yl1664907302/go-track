package wechat_robot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kube-auto/global"
	"kube-auto/pojo"
	"net/http"
)

type WeChat_RobotApi struct {
}

func (*WeChat_RobotApi) GetWeChat_RobotMessage(c *gin.Context) {
	var wechat_robot pojo.WeChat_RobotMarkdownMessage
	err := c.ShouldBind(&wechat_robot)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}
	for _, v := range global.ActionMessage.ExtractInfo(wechat_robot.Markdown.Content) {
		fmt.Println(v)
	}
}

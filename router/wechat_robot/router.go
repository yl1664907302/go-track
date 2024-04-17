package wechat_robot

import (
	"github.com/gin-gonic/gin"
	"kube-auto/api"
)

type WeChat_RobotRouter struct {
}

func (*WeChat_RobotRouter) InitWeChat_RobotRouter(r *gin.Engine) {
	group := r.Group("/wechat_robot")
	Wechat_robotApigroup := api.ApiGroupApp.Wechat_robotApiGroup
	group.POST("/message", Wechat_robotApigroup.GetWeChat_RobotMessage)
}

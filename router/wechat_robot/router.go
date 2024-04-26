package wechat_robot

import (
	"github.com/gin-gonic/gin"
	"go-track/api"
)

type WeChat_RobotRouter struct {
}

func (*WeChat_RobotRouter) InitWeChat_RobotRouter(r *gin.Engine) {
	group := r.Group("/wechat_robot")
	Wechat_robotApigroup := api.ApiGroupApp.Wechat_robotApiGroup
	group.POST("/message", Wechat_robotApigroup.PostWeChat_RobotMessage)
	group.GET("getmessage", Wechat_robotApigroup.GetWeChat_RobotMessagebyFenye)
	group.GET("getmessagemohu", Wechat_robotApigroup.GetWeChat_RobotMessagebyMohu)
}

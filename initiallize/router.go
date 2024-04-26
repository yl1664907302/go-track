package initiallize

import (
	"github.com/gin-gonic/gin"
	"go-track/middleware"
	"go-track/router"
)

func Router() *gin.Engine {
	//gin.Default() 返回一个 gin.Engine
	r := gin.Default()
	r.Use(middleware.Cors)
	g1 := router.RouterGroupApp.ExampleRouterGroup
	g1.InitExample(r)
	g3 := router.RouterGroupApp.UserLoginGroup
	g3.Login(r)
	g4 := router.RouterGroupApp.DingtalkRouterGroup
	g4.InitDingTalkRouter(r)
	g5 := router.RouterGroupApp.FeishuFeishuGroup
	g5.InitFeishuRouter(r)
	g6 := router.RouterGroupApp.Wechat_robotWeChat_RobotGroup
	g6.InitWeChat_RobotRouter(r)
	g7 := router.RouterGroupApp.WechatWeChatGroup
	g7.InitWeChatRouter(r)
	return r
}

package wechat

import (
	"github.com/gin-gonic/gin"
	"kube-auto/api"
)

type WeChatRouter struct {
}

func (*WeChatRouter) InitWeChatRouter(r *gin.Engine) {
	group := r.Group("/wechat")
	WechatApigroup := api.ApiGroupApp.WechatApiGroup
	group.POST("/message", WechatApigroup.GetWeChatMessage)
}

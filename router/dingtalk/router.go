package dingtalk

import (
	"github.com/gin-gonic/gin"
	"kube-auto/api"
)

type DingTalkRouter struct {
}

func (*DingTalkRouter) InitDingTalkRouter(r *gin.Engine) {
	group := r.Group("/dingtalk")
	dingTalkApigroup := api.ApiGroupApp.DingtalkApiGroup
	group.POST("/message", dingTalkApigroup.PostDingTalkMessage)
	group.GET("getmessage", dingTalkApigroup.GetDingTalkMessagebyFenye)
}

package feishu

import (
	"github.com/gin-gonic/gin"
	"kube-auto/api"
)

type FeishuRouter struct {
}

func (*FeishuRouter) InitFeishuRouter(r *gin.Engine) {
	group := r.Group("/feishu")
	FeishuApigroup := api.ApiGroupApp.FeishuApiGroup
	group.POST("/message", FeishuApigroup.PostFeishuMessage)
	group.GET("getmessage", FeishuApigroup.GetFeishuMessagebyFenye)
}

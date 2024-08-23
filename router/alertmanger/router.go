package alertmanger

import (
	"github.com/gin-gonic/gin"
	"go-track/api"
)

type AlertMangerRouter struct {
}

func (*AlertMangerRouter) InitAlertMangerRouter(r *gin.Engine) {
	group := r.Group("/alertmanger")
	alertmangerApigroup := api.ApiGroupApp.AlertmangerApiGroup
	group.POST("/post", alertmangerApigroup.PostAlertMangerMessage)
	group.POST("/post/test", alertmangerApigroup.PostTestAlertMangerMessage)
	group.POST("/post/markdown", alertmangerApigroup.PostMarkDownTemplate)
	group.POST("/post/dingtalk", alertmangerApigroup.PostDingTalkRobotConf)
	group.GET("/origin", alertmangerApigroup.GetAlertMangerMessage)
	group.GET("/newmarkdown", alertmangerApigroup.GetMarkDownMessage)
	group.GET("/robot", alertmangerApigroup.GetRobot)
	group.GET("/delrobot/dingtalk", alertmangerApigroup.GetDelRobot)
}

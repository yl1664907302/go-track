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

	//分步表单的接收
	group.POST("post/stepform", alertmangerApigroup.PostStepFormToAlertManger)

	//原始消息接收
	group.POST("/post", alertmangerApigroup.PostAlertMangerMessage)
	group.POST("/post/test", alertmangerApigroup.PostTestAlertMangerMessage)

	//markdown模板
	group.POST("/post/newmarkdowntemplate", alertmangerApigroup.PostMarkDownTemplate)
	group.POST("/post/updatemarkdowntemplate", alertmangerApigroup.PostUpdateMarkDownTemplate)
	group.GET("/newmarkdowntemplate", alertmangerApigroup.GetNewMarkDownTemplate)

	//查询告警消息
	group.GET("/origin", alertmangerApigroup.GetAlertMangerMessage)
	group.GET("/markdown", alertmangerApigroup.GetMarkDownMessage)
	group.GET("/markdownbystatus", alertmangerApigroup.GetMarkDownMessagebyStatus2Api)
	group.GET("/markdownbystatus2mohu", alertmangerApigroup.GetMarkDownMessagebyStatus2Mohu)

	//robot
	group.POST("/post/newrobot", alertmangerApigroup.PostRobotConf)
	group.POST("/post/updaterobot", alertmangerApigroup.PostUpdateRobot)
	group.GET("/robot", alertmangerApigroup.GetRobot)
	group.GET("/delrobot", alertmangerApigroup.GetDelRobot)

	//查询Receiver信息
	group.GET("/receivers", alertmangerApigroup.GetReceivers)

	//删除索引
	group.DELETE("/del/index", alertmangerApigroup.DelReceiver)
}

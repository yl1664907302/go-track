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

	//markdown模板
	group.POST("/post/newmarkdown", alertmangerApigroup.PostMarkDownTemplate)
	group.POST("/post/updatemarkdown", alertmangerApigroup.PostUpdateMarkDownTemplate)
	group.GET("/newmarkdown", alertmangerApigroup.GetNewMarkDownTemplate)

	//查询告警消息
	group.GET("/origin", alertmangerApigroup.GetAlertMangerMessage)
	group.GET("/markdown", alertmangerApigroup.GetMarkDownMessage)

	//robot
	group.POST("/post/newrobot", alertmangerApigroup.PostRobotConf)
	group.POST("/post/updaterobot", alertmangerApigroup.PostUpdateRobot)
	group.GET("/robot", alertmangerApigroup.GetRobot)
	group.GET("/delrobot", alertmangerApigroup.GetDelRobot)

}

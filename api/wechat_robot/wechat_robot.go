package wechat_robot

import (
	"github.com/gin-gonic/gin"
	"go-track/action"
	"go-track/elastics"
	"go-track/pojo"
	"go-track/response"

	"log"
	"net/http"
)

type WeChat_RobotApi struct {
}

func (*WeChat_RobotApi) PostWeChat_RobotMessage(c *gin.Context) {
	var wechat_robot *pojo.WeChat_RobotMarkdownMessage
	var index = "wechat_robot"
	err := c.ShouldBind(&wechat_robot)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}
	message := action.WeChat_Robot_TransForm(wechat_robot)
	err, resp := elastics.CreateIndexES(&message, index)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "wechat_robot消息失败存入es", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "wechat_robot消息成功存入es", map[string]string{
			"code": resp,
		})
	}
}

func (*WeChat_RobotApi) GetWeChat_RobotMessagebyFenye(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	fenye.From = c.Query("from")
	fenye.Size = c.Query("size")
	fenye.SortField = c.Query("sort_field")
	fenye.Asc = c.Query("asc")
	messages, err := elastics.PaginateSearchEsDoc(&fenye)
	if err != nil {
		log.Println(err)
	}
	response.SuccssWithDetailed(c, "分页查询索引成功", messages)
}

func (*WeChat_RobotApi) GetWeChat_RobotMessagebyMohu(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	fenye.From = c.Query("from")
	fenye.Size = c.Query("size")
	fenye.SortField = c.Query("sort_field")
	fenye.Asc = c.Query("asc")
	messages, err := elastics.SelectEsDocByIndex2keyword(&fenye, c.Query("groupname"), c.Query("time"), c.Query("keyword1"))
	if err != nil {
		log.Println(err)
	}
	response.SuccssWithDetailed(c, "模糊查询索引成功", messages)
}

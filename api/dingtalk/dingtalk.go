package dingtalk

import (
	"github.com/gin-gonic/gin"
	"kube-auto/action"
	"kube-auto/elastic"
	"kube-auto/pojo"
	"kube-auto/response"
	"log"
	"net/http"
)

type DingTalkApi struct {
}

func (*DingTalkApi) PostDingTalkMessage(c *gin.Context) {
	var dingtalk *pojo.DingtalkMarkdownMessage
	var index = "dingtalk"
	err := c.ShouldBind(&dingtalk)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}
	message := action.DingTalkTransForm(dingtalk)
	err = elastic.CreateIndexES(&message, index)
	if err != nil {
		log.Print(err)
	}
}

func (*DingTalkApi) GetDingTalkMessagebyFenye(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	fenye.From = c.Query("from")
	fenye.Size = c.Query("size")
	fenye.SortField = c.Query("sort_field")
	fenye.Asc = c.Query("asc")
	messages, err := elastic.PaginateSearchEsDoc(&fenye)
	if err != nil {
		log.Println(err)
	}
	response.SuccssWithDetailed(c, "分页查询索引成功", messages)
}

func (*DingTalkApi) GetDingTalkMessagebyMohu(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	fenye.From = c.Query("from")
	fenye.Size = c.Query("size")
	fenye.SortField = c.Query("sort_field")
	fenye.Asc = c.Query("asc")
	messages, err := elastic.SelectEsDocByIndex2keyword(&fenye, c.Query("groupname"), c.Query("time"), c.Query("keyword1"))
	if err != nil {
		log.Println(err)
	}
	response.SuccssWithDetailed(c, "模糊查询索引成功", messages)
}

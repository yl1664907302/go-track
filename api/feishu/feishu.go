package feishu

import (
	"github.com/gin-gonic/gin"
	"go-track/action"
	"go-track/elastic"
	"go-track/pojo"
	"go-track/response"

	"log"
	"net/http"
)

type FeishuApi struct {
}

func (*FeishuApi) PostFeishuMessage(c *gin.Context) {
	var feishu *pojo.FeishuMarkdownMessage
	var index = "feishu"
	err := c.ShouldBind(&feishu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}
	message := action.FeiShuTransForm(feishu)
	err = elastic.CreateIndexES(&message, index)
	if err != nil {
		log.Print(err)
	}
}

func (*FeishuApi) GetFeishuMessagebyFenye(c *gin.Context) {
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

func (*FeishuApi) GetFeishuMessagebyMohu(c *gin.Context) {
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

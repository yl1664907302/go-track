package alertmanger

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"go-track/database/mysql"
	"go-track/elastics"
	"go-track/going"
	"go-track/pojo"
	"go-track/response"
	"go-track/utils"
	"log"
	"net/http"
	"time"
)

type AlertMangerApi struct{}

// 负责接收alertmanger的告警消息，并存储es
func (*AlertMangerApi) PostAlertMangerMessage(c *gin.Context) {
	body, err := c.GetRawData()
	var alert pojo.Alert
	err = sonic.Unmarshal(body, &alert)
	if err != nil {
		log.Print(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//获取index（首字母大写转小写）
	index := utils.ActionMessages.EditFisrtCharToLower(alert.Receiver)
	alerts := alert.Alerts
	//存入es
	err, resp := elastics.CreateIndexESForAlert(alerts, index)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
	//判断是否存在模板
	found, desc, err := elastics.SearchMarkDown(index + "_t")
	if found {
		log.Println("索引:" + index + "存在模板")
		log.Println("模板为:\n" + desc.Markdown)
		log.Println("模板创建时间为:" + desc.Maketime)

	} else {
		log.Println("索引" + index + "不存在模板")
	}
	//执行函数将模板中的字段替换为json字段对应的值
	for _, a := range alerts {
		markdown, err := utils.ActionMessages.InsertJsonToMarkdown(desc, &a)
		if err != nil {
			log.Println(err)
		}
		//发送给钉钉
		err = going.RobotDingTalkGoing(index, markdown)
		//markdown实例存入es
		newmarkdown := pojo.NewNewmarkdown(a.Fingerprint, a.StartsAt, markdown)
		err, _ = elastics.CreateIndexForNewMarkDown(newmarkdown, index)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("markdown实例已成功写入索引：" + index + "_n")
		}
		//receiver存入mysql
		err = mysql.InsertReceiver(alert.Receiver)
		if err != nil {
			log.Println(err)
		}
	}
}

// 负责接收前端发送过来的Markdown模板（目前只能存储一份模板）
func (*AlertMangerApi) PostMarkDownTemplate(c *gin.Context) {
	body, err := c.GetRawData()
	var markdown pojo.Markdown
	err = sonic.Unmarshal(body, &markdown)
	if err != nil {
		log.Print(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//获取index（首字母大写转小写）
	index := utils.ActionMessages.EditFisrtCharToLower(markdown.Receiver)
	markdown.Desc.Markdown = utils.ActionMessages.TranferSingleToDouble(markdown.Desc.Markdown)
	markdown.Desc.Maketime = time.Now().Format("2006-01-02 15:04:05")
	err, _ = elastics.CreateIndexForMarkDown(&markdown.Desc, index)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "markdown模板失败存入es", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "markdown模板成功存入es", map[string]string{
			"code": "200",
		})
	}
}

func (*AlertMangerApi) GetNewMarkDownTemplate(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	err, markdown := elastics.SelectNewMarkdownTempByIndex(fenye.Index)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "获取最新模板失败", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "获取最新模板成功", markdown)
	}
}

// 负责获取alertmanger去重排序过滤后的告警消息
func (*AlertMangerApi) GetAlertMangerMessage(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	fenye.From = c.Query("from")
	fenye.Size = c.Query("size")
	fenye.SortField = c.Query("sort_field")
	fenye.Asc = c.Query("asc")
	alerts, _, err := elastics.SearchBySortAndUnique(&fenye, false)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "去重排序查询alertmanger原信息索引失败", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "去重排序查询alertmanger原信息索引成功", alerts)
	}
}

// 负责获取alertmanger去重排序过滤后的MarkDown告警消息
func (*AlertMangerApi) GetMarkDownMessage(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	fenye.From = c.Query("from")
	fenye.Size = c.Query("size")
	fenye.SortField = c.Query("sort_field")
	fenye.Asc = c.Query("asc")
	_, markdowns, err := elastics.SearchBySortAndUnique(&fenye, true)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "去重排序查询markdown实例索引失败", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "去重排序查询markdown实例索引成功", markdowns)
	}
}

func (*AlertMangerApi) PostUpdateMarkDownTemplate(c *gin.Context) {
	body, err := c.GetRawData()
	var markdown pojo.Markdown
	err = sonic.Unmarshal(body, &markdown)
	if err != nil {
		log.Print(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	markdown.Desc.Markdown = utils.ActionMessages.TranferSingleToDouble(markdown.Desc.Markdown)
	err, _ = elastics.UpdateIndexForMarkDown(&markdown, markdown.Receiver)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "更新markdown模板失败", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "更新markdown模板成功", markdown)
	}
}

// 测试手动发送消息到go-track
func (*AlertMangerApi) PostTestAlertMangerMessage(c *gin.Context) {
	body, err := c.GetRawData()
	var alerts pojo.Alerts
	err = sonic.Unmarshal(body, &alerts)
	if err != nil {
		log.Print(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err, _ = elastics.TestCreateIndexForAlert(alerts)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "AlertManger测试消息失败存入es", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "AlertManger测试消息成功存入es", map[string]string{
			"code": "200",
		})
	}
}

// **********************robot配置********************************

// 新增钉钉机器人
func (*AlertMangerApi) PostRobotConf(c *gin.Context) {
	body, err := c.GetRawData()
	var robot pojo.Robot
	err = sonic.Unmarshal(body, &robot)
	if err != nil {
		log.Print(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err, _ = elastics.CreateIndexForRobot(&robot, robot.Receiver)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "Robot失败存入es", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "Robot成功存入es", map[string]string{
			"code": "200",
		})
	}
}

// 查询Robot
func (*AlertMangerApi) GetRobot(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	fenye.From = c.Query("from")
	fenye.Size = c.Query("size")
	fenye.SortField = c.Query("sort_field")
	fenye.Asc = c.Query("asc")
	robot, err := elastics.SearchRobot(fenye.Index)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "robot实例获取失败", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "robot实例获取成功", robot)
	}
}

// 删除告警机器人
func (*AlertMangerApi) GetDelRobot(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	doc_id := c.Query("robot_id")
	err := elastics.DelDocByKey(fenye.Index+"_r", "robot_id", doc_id)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "robot删除失败", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "robot删除成功", "")
	}
}

// 更新告警机器人
func (*AlertMangerApi) PostUpdateRobot(c *gin.Context) {
	body, err := c.GetRawData()
	var robot pojo.Robot
	err = sonic.Unmarshal(body, &robot)
	if err != nil {
		log.Print(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = elastics.UpdateDocForRobot(robot.Receiver+"_r", robot)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "Robot更新失败", map[string]string{
			"code": err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "Robot更新成功", map[string]string{
			"code": "200",
		})
	}
}

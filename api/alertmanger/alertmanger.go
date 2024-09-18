package alertmanger

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"go-track/action"
	"go-track/database/mysql"
	"go-track/elastics"
	"go-track/going"
	"go-track/pojo"
	"go-track/response"
	"go-track/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AlertMangerApi struct{}

func (*AlertMangerApi) DelReceiver(c *gin.Context) {
	var pagination pojo.Fenye
	pagination.Index = c.Query("index")
	//执行删除
	indexSuffixes := []string{"", "_r", "_t", "_n"}
	for _, suffix := range indexSuffixes {
		//判断是否存在索引
		index, _ := elastics.JudgeIndex(pagination.Index + suffix)
		if index == 0 {
			continue
		}
		err := elastics.DelIndex(pagination.Index + suffix)
		if err != nil {
			log.Println(err)
			response.FailWithDetailed(c, "", map[string]any{
				"status":  "error",
				"message": "失败删除索引：" + pagination.Index + suffix,
			})
			return
		}

	}
	err := mysql.DelReceiver(pagination.Index)
	if err != nil {
		log.Println(err)
		response.FailWithDetailed(c, "", map[string]any{
			"status":  "error",
			"message": "失败删除receiver：" + pagination.Index,
		})
		return
	}
	response.SuccssWithDetailed(c, "", map[string]any{
		"status":  "success",
		"message": "删除成功",
	})
}

func (*AlertMangerApi) GetReceivers(c *gin.Context) {
	receivers, err := mysql.SelectReceiver()
	if err != nil {
		response.FailWithDetailed(c, "", map[string]any{
			"status":  "error",
			"message": "receiver查询失败",
		})
	} else {
		response.SuccssWithDetailed(c, "", receivers)
	}
}

func (*AlertMangerApi) GetMarkDownMessagebyStatus2Api(c *gin.Context) {
	key, err := action.SelectAlertsByKey(c.Query("index"), "status", c.Query("status"))
	if err != nil {
		response.FailWithDetailed(c, "", map[string]any{
			"status":  "error",
			"message": "正在告警数查询失败:" + err.Error(),
		})
		return
	}
	response.SuccssWithDetailed(c, "", map[string]any{
		"status": "success",
		"number": key,
	})
}

//func (*AlertMangerApi) GetMarkDownMessagebyStatus2Mohu(c *gin.Context) {
//	var fenye pojo.Fenye
//	fenye.Time = c.Query("time")
//	fenye.Index = c.Query("index") + "_n"
//	fenye.Status = c.Query("status")
//}

func (*AlertMangerApi) GetMarkDownMessagebyStatus2Mohu(c *gin.Context) {
	var fenye pojo.Fenye
	//获取index（首字母大写转小写）
	i := utils.ActionMessages.EditFisrtCharToLower(c.Query("index"))
	fenye.Index = i + "_n"
	fenye.From = c.Query("from")
	fenye.Size = c.Query("size")
	fenye.SortField = c.Query("sort_field")
	fenye.Asc = c.Query("asc")
	fenye.Status = c.Query("status")
	fenye.Time_start = c.Query("time_start")
	fenye.Time_end = c.Query("time_end")
	_, markdowns, number, err := elastics.SearchBySortAndUniqueAndByKey2time(&fenye, true, "status", fenye.Status)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "", map[string]any{
			"status":  "error",
			"message": "消息来源不存在",
		})
	} else {
		response.SuccssWithDetailed(c, "", map[string]any{
			"status":    "success",
			"number":    number,
			"markdowns": markdowns,
		})
	}
}

// 分步表单接受api
func (*AlertMangerApi) PostStepFormToAlertManger(c *gin.Context) {
	body, err := c.GetRawData()
	var step pojo.Stepform
	var robot pojo.Robot
	var robot2 pojo.Robot
	var markdown pojo.Markdown

	err = sonic.Unmarshal(body, &step)
	if err != nil {
		log.Print(err)
	}
	if err != nil {
		response.FailWithDetailed(c, "消息来源不存在!", map[string]any{
			"status":  "error",
			"message": "消息来源不存在",
		})
		return
	}
	//获取index（首字母大写转小写）
	receiver := utils.ActionMessages.EditFisrtCharToLower(step.Receiver)

	////验证消息是否存在在于es
	//key, err := elastics.JudgeIndex(receiver)
	//if err != nil {
	//	log.Println(err)
	//}
	//if key == 0 {
	//	response.FailWithDetailed(c, "消息来源不存在!", map[string]any{
	//		"status":  "error",
	//		"message": "消息来源不存在",
	//	})
	//	return
	//}

	//receiver存入mysql(首字母不变小写)
	err = mysql.InsertReceiver(step.Receiver, step.Niname)
	if err != nil {
		response.FailWithDetailed(c, "数据库异常："+err.Error(), map[string]any{
			"status":  "error",
			"message": "数据库异常：" + err.Error(),
		})
		return
	}

	//添加robot
	if step.Robot_ok {
		robot.Receiver = receiver
		robot.Robot_class = step.Robot_class
		robot.Secret = step.Secret
		robot.Switch = step.Switch
		robot.Accesstoken = step.Accesstoken
		robot.Robot_name = step.Robot_name
		byindex, err := elastics.SelectNewDocByindex(robot.Receiver+"_r", "robot_id", &pojo.Robot{})
		err = sonic.Unmarshal(byindex, &robot2)
		robot.Robot_id = robot2.Robot_id + 1
		err, _ = elastics.CreateIndexForRobot(&robot, robot.Receiver)
		if err != nil {
			log.Print(err)
			response.FailWithDetailed(c, "Robot保存失败", map[string]any{
				"status":  "error",
				"message": "Robot保存失败",
			})
			return
		}
	}

	//添加markdown模板
	if step.Markdown_ok {
		markdown.Desc.Markdown = step.Markdown
		markdown.Receiver = receiver
		markdown.Desc.Maketime = time.Now().Format("2006-01-02 15:04:05")
		err, _ = elastics.CreateIndexForMarkDown(&markdown.Desc, receiver)
		if err != nil {
			log.Print(err)
			response.FailWithDetailed(c, "markdown模板保存失败", map[string]any{
				"status":  "error",
				"message": "markdown模板保存失败",
			})
			return
		}
	}

	response.SuccssWithDetailed(c, "成功创建告警通道"+step.Niname, map[string]any{
		"status":  "success",
		"message": "成功创建告警通道" + step.Niname,
	})
}

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
	found, desc, _ := elastics.SearchMarkDown(index + "_t")
	if found {
		log.Println("索引:" + index + "存在模板")
		log.Println("模板为:\n" + desc.Markdown)
		log.Println("模板创建时间为:" + desc.Maketime)
		//执行函数将模板中的字段替换为json字段对应的值
		for _, a := range alerts {
			markdown, err := utils.ActionMessages.InsertJsonToMarkdown(desc, &a)
			if err != nil {
				log.Println(err)
			}
			//发送给钉钉
			err = going.RobotDingTalkGoing(index, markdown)
			//markdown实例存入es
			newmarkdown := pojo.NewNewmarkdown(a.Status, a.Fingerprint, a.StartsAt, markdown, a.EndsAt)
			err, _ = elastics.CreateIndexForNewMarkDown(newmarkdown, index)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("markdown实例已成功写入索引：" + index + "_n")
			}
		}
	} else {
		log.Println("索引" + index + "不存在模板")
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
	//获取index（首字母大写转小写）
	index := utils.ActionMessages.EditFisrtCharToLower(fenye.Index)
	err, markdown := elastics.SelectNewMarkdownTempByIndex(index)
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
	//获取index（首字母大写转小写）
	index := utils.ActionMessages.EditFisrtCharToLower(c.Query("index"))
	fenye.Index = index
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
	//获取index（首字母大写转小写）
	index := utils.ActionMessages.EditFisrtCharToLower(c.Query("index"))
	fenye.Index = index + "_n"
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
	//获取index（首字母大写转小写）
	markdown.Receiver = utils.ActionMessages.EditFisrtCharToLower(markdown.Receiver)
	markdown.Desc.Markdown = utils.ActionMessages.TranferSingleToDouble(markdown.Desc.Markdown)
	num, err := elastics.SelectNumByIndex(markdown.Receiver + "_t")
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "更新markdown模板失败", map[string]string{
			"status":  "success",
			"message": "更新markdown模板失败:" + err.Error(),
		})
	}
	if num != 0 {
		err, _ = elastics.UpdateIndexForMarkDown(&markdown, markdown.Receiver)
		if err != nil {
			log.Print(err)
			response.FailWithDetailed(c, "更新markdown模板失败", map[string]string{
				"status":  "success",
				"message": "更新markdown模板失败:" + err.Error(),
			})
		} else {
			response.SuccssWithDetailed(c, "更新markdown模板成功", map[string]string{
				"status":  "success",
				"message": "完成更新markdown模板:" + markdown.Receiver,
			})
		}
	} else {
		markdown.Desc.Markdown = utils.ActionMessages.TranferSingleToDouble(markdown.Desc.Markdown)
		markdown.Desc.Maketime = time.Now().Format("2006-01-02 15:04:05")
		err, _ = elastics.CreateIndexForMarkDown(&markdown.Desc, markdown.Receiver)
		if err != nil {
			log.Print(err)
			response.FailWithDetailed(c, "更新markdown模板失败", map[string]string{
				"status":  "success",
				"message": "更新markdown模板失败:" + err.Error(),
			})
		} else {
			response.SuccssWithDetailed(c, "更新markdown模板成功", map[string]string{
				"status":  "success",
				"message": "完成更新markdown模板:" + markdown.Receiver,
			})
		}
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
	if err != nil {
		log.Println(err)
	}
	var robot pojo.Robot
	var robot2 pojo.Robot
	err = sonic.Unmarshal(body, &robot)
	if err != nil {
		log.Print(err)
	}
	if err != nil {
		response.FailWithDetailed(c, "Robot失败存入es", map[string]string{
			"status":  "error",
			"message": "失败创建机器人,由于:" + err.Error(),
		})
		return
	}
	//获取index（首字母大写转小写）
	robot.Receiver = utils.ActionMessages.EditFisrtCharToLower(robot.Receiver)
	key, err := elastics.JudgeIndex(robot.Receiver + "_r")
	if err != nil {
		log.Println(err)
	}
	num, err := elastics.SelectNumByIndex(robot.Receiver + "_r")
	if key == 1 && num != 0 {
		byindex, err := elastics.SelectNewDocByindex(robot.Receiver+"_r", "robot_id", &pojo.Robot{})
		if err != nil {
			log.Println(err)
		}
		err = sonic.Unmarshal(byindex, &robot2)
		if err != nil {
			response.FailWithDetailed(c, "Robot失败存入es", map[string]string{
				"status":  "error",
				"message": "失败创建机器人:" + err.Error(),
			})
			return
		}
		//复制最新的id
		robot.Robot_id = robot2.Robot_id + 1
	} else if key == 0 {
		robot.Robot_id = 1
	}
	err, _ = elastics.CreateIndexForRobot(&robot, robot.Receiver)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "Robot失败存入es", map[string]string{
			"status":  "error",
			"message": "失败创建机器人:" + robot.Robot_name,
		})
	} else {
		response.SuccssWithDetailed(c, "Robot成功存入es", map[string]string{
			"status":  "success",
			"message": "成功创建机器人:" + robot.Robot_name,
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
	//获取index（首字母大写转小写）
	fenye.Index = utils.ActionMessages.EditFisrtCharToLower(fenye.Index)
	robot, err := elastics.SearchRobot(fenye.Index)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "robot实例获取失败", map[string]string{
			"status":  "error",
			"message": "robot实例获取失败",
		})
		return
	} else {
		response.SuccssWithDetailed(c, "robot实例获取成功", map[string]any{
			"status":  "success",
			"message": "成功查询机器人",
			"robots":  robot,
		})
	}
}

// 删除告警机器人
func (*AlertMangerApi) DelRobot(c *gin.Context) {
	var fenye pojo.Fenye
	fenye.Index = c.Query("index")
	doc_id := c.Query("robot_id")
	atoi, _ := strconv.Atoi(doc_id)
	//获取index（首字母大写转小写）
	fenye.Index = utils.ActionMessages.EditFisrtCharToLower(fenye.Index)
	err := elastics.DelDocByKey(fenye.Index+"_r", "robot_id", atoi)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "robot删除失败", map[string]string{
			"status":  "error",
			"message": "robot实例获取失败:" + err.Error(),
		})
	} else {
		response.SuccssWithDetailed(c, "robot删除成功", map[string]string{
			"status":  "success",
			"message": "成功删除机器人",
		})
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
	//获取index（首字母大写转小写）
	robot.Receiver = utils.ActionMessages.EditFisrtCharToLower(robot.Receiver)
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

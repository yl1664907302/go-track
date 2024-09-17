package action

import (
	"github.com/bytedance/sonic"
	"go-track/global"
	"go-track/pojo"
	"io"
	"log"
	"net/http"
	"time"
)

func WeChat_Robot_TransForm(WeChat_RobotMessage *pojo.WeChat_RobotMarkdownMessage) pojo.Message {
	// 获取当前时间
	currentTime := time.Now()
	var contexts []pojo.Context
	platForm := "企业微信"
	Info, groupName := global.ActionMessage.ExtractInfo(WeChat_RobotMessage.Markdown.Content, true)
	for k, v := range Info {
		context := pojo.Context{
			Num:         k,
			LineName:    v.Key,
			LineContext: v.Value,
		}
		contexts = append(contexts, context)
	}

	message := pojo.Message{
		Time:      currentTime.Format("2006-01-02 15:04:05"),
		PlatForm:  platForm,
		GroupName: groupName,
		Contests:  contexts,
	}
	return message
}

func DingTalkTransForm(DingtalkMessage *pojo.DingtalkMarkdownMessage) pojo.Message {
	var contexts []pojo.Context
	// 获取当前时间
	currentTime := time.Now()

	platForm := "钉钉"

	Info, _ := global.ActionMessage.ExtractInfo(DingtalkMessage.Markdown.Text, false)
	for k, v := range Info {
		context := pojo.Context{
			Num:         k,
			LineName:    v.Key,
			LineContext: v.Value,
		}
		contexts = append(contexts, context)
	}

	message := pojo.Message{
		Time:      currentTime.Format("2006-01-02 15:04:05"),
		PlatForm:  platForm,
		GroupName: DingtalkMessage.Markdown.Title,
		Contests:  contexts,
	}
	return message
}

func FeiShuTransForm(FeishuMessage *pojo.FeishuMarkdownMessage) pojo.Message {
	// 获取当前时间
	currentTime := time.Now()
	//目前只只存读取最后一个message
	var contexts []pojo.Context
	platForm := "飞书"
	for _, element := range FeishuMessage.Card.Elements {
		info, _ := global.ActionMessage.ExtractInfo(element.Content, false)
		for k, v := range info {
			context := pojo.Context{
				Num:         k,
				LineName:    v.Key,
				LineContext: v.Value,
			}
			contexts = append(contexts, context)
		}
	}
	message := pojo.Message{
		Time:      currentTime.Format("2006-01-02 15:04:05"),
		PlatForm:  platForm,
		GroupName: FeishuMessage.Card.Header.Title.Content,
		Contests:  contexts,
	}
	return message
}

// 创建http请求
func SelectAlertsByKey(receiver string, key string, value string) (int, error) {
	var num int
	var alerts []pojo.Alert2
	req, err := http.Get("http://" + global.CONF.System.Alertmanger_api.Host + ":" + global.CONF.System.Alertmanger_api.Port + "/api/v2/alerts")
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)
	all, err := io.ReadAll(req.Body)
	if err != nil {
		return 0, err
	}

	err = sonic.Unmarshal(all, &alerts)
	if err != nil {
		return 0, err
	}

	for _, a := range alerts {
		key2 := false
		for _, r := range a.Receivers {
			if r.Name == receiver {
				key2 = true
			}
		}
		if key2 && key == "status" {
			if a.Status.State == value {
				num++
			}
		}
	}
	log.Println(num)
	return num, err
}

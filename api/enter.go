package api

import (
	"kube-auto/api/dingtalk"
	"kube-auto/api/example"
	"kube-auto/api/feishu"
	"kube-auto/api/user"
	"kube-auto/api/wechat"
	"kube-auto/api/wechat_robot"
)

type ApiGroup struct {
	ExampleApiGroup      example.ApiGroup
	LoginApiGroup        user.ApiGroup
	DingtalkApiGroup     dingtalk.ApiGroup
	FeishuApiGroup       feishu.ApiGroup
	WechatApiGroup       wechat.ApiGroup
	Wechat_robotApiGroup wechat_robot.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

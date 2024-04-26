package api

import (
	"go-track/api/dingtalk"
	"go-track/api/example"
	"go-track/api/feishu"
	"go-track/api/user"
	"go-track/api/wechat"
	"go-track/api/wechat_robot"
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

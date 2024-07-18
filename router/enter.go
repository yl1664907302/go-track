package router

import (
	"go-track/router/alertmanger"
	"go-track/router/dingtalk"
	"go-track/router/example"
	"go-track/router/feishu"
	"go-track/router/user"
	"go-track/router/wechat"
	"go-track/router/wechat_robot"
)

type RouterGroup struct {
	ExampleRouterGroup            example.ExampleRouter
	UserLoginGroup                user.LoginRouter
	DingtalkRouterGroup           dingtalk.DingTalkRouter
	FeishuFeishuGroup             feishu.FeishuRouter
	WechatWeChatGroup             wechat.WeChatRouter
	Wechat_robotWeChat_RobotGroup wechat_robot.WeChat_RobotRouter
	AlertMangerRouterGroup        alertmanger.AlertMangerRouter
}

var RouterGroupApp = new(RouterGroup)

package router

import (
	"kube-auto/router/dingtalk"
	"kube-auto/router/example"
	"kube-auto/router/feishu"
	"kube-auto/router/user"
	"kube-auto/router/wechat"
	"kube-auto/router/wechat_robot"
)

type RouterGroup struct {
	ExampleRouterGroup            example.ExampleRouter
	UserLoginGroup                user.LoginRouter
	DingtalkRouterGroup           dingtalk.DingTalkRouter
	FeishuFeishuGroup             feishu.FeishuRouter
	WechatWeChatGroup             wechat.WeChatRouter
	Wechat_robotWeChat_RobotGroup wechat_robot.WeChat_RobotRouter
}

var RouterGroupApp = new(RouterGroup)

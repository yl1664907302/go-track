package router

import (
	"go-track/router/alertmanger"
	"go-track/router/example"
	"go-track/router/user"
)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
	UserLoginGroup     user.LoginRouter
	//DingtalkRouterGroup           dingtalk.DingTalkRouter
	//FeishuFeishuGroup             feishu.FeishuRouter
	//WechatWeChatGroup             wechat.WeChatRouter
	//Wechat_robotWeChat_RobotGroup wechat_robot.WeChat_RobotRouter
	AlertMangerRouterGroup alertmanger.AlertMangerRouter
}

var RouterGroupApp = new(RouterGroup)

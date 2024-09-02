package going

import (
	"github.com/blinkbean/dingtalk"
	"github.com/wanghuiyt/ding"
	"go-track/elastics"
)

// 发生文本消息到钉钉
func RobotDingTalkGoing2(index string, markdown string) error {
	robots, _ := elastics.SearchRobot(index)
	for _, robot := range robots {
		if robot.Robot_class == "dingtalk" && robot.Switch == "on" {
			webhook := ding.Webhook{AccessToken: robot.Accesstoken, Secret: robot.Secret}
			err := webhook.SendMessageText(markdown)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 发生markdown消息到钉钉
func RobotDingTalkGoing(index string, markdown string) error {
	robots, _ := elastics.SearchRobot(index)
	for _, robot := range robots {
		if robot.Robot_class == "dingtalk" && robot.Switch == "on" {
			secret := dingtalk.InitDingTalkWithSecret(robot.Accesstoken, robot.Secret)
			//webhook := ding.Webhook{AccessToken: robot.Accesstoken, Secret: robot.Secret}
			//err := webhook.SendMessageText(markdown)
			err := secret.SendMarkDownMessage("告警消息", markdown, dingtalk.WithAtAll())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

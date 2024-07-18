package going

import (
	"github.com/wanghuiyt/ding"
	"go-track/elastics"
)

func RobotDingTalkGoing(index string, markdown string) error {
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

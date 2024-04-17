package global

import (
	"gorm.io/gorm"
	"kube-auto/config"
	"kube-auto/utils"
)

var (
	CONF             config.Server
	MysqlDataConnect *gorm.DB
	ActionMessage    utils.ActionMessage
)

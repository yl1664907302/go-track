package global

import (
	"go-track/config"
	"go-track/utils"
	"gorm.io/gorm"
)

var (
	CONF             config.Server
	MysqlDataConnect *gorm.DB
	ActionMessage    utils.ActionMessage
)

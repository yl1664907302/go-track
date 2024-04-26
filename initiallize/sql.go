package initiallize

import (
	"fmt"
	"go-track/database/dao"
	"go-track/database/dao/mysqldatabase"
	"go-track/database/service"
	"go-track/database/service/servicelmpl"
	"go-track/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
)

func InitMysqlDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.CONF.System.Database.MysqlUser, global.CONF.System.Database.MysqlPassword, global.CONF.System.Database.MysqlHost, global.CONF.System.Database.MysqlPort, global.CONF.System.Database.MysqlDatabasename)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err)
	}
	global.MysqlDataConnect = db
	InitgetuserDAO()
	InitgetuserServicelmpl()
}

func InitgetuserDAO() {
	dao.UserDAO = mysqldatabase.NewMysqlUserDAO(global.MysqlDataConnect)
}

func InitgetuserServicelmpl() {
	service.UserServiceImpl = servicelmpl.NewUserServiceImpl(dao.UserDAO)
}

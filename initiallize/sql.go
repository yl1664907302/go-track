package initiallize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kube-auto/database/dao"
	"kube-auto/database/dao/mysqldatabase"
	"kube-auto/database/service"
	"kube-auto/database/service/servicelmpl"
	"kube-auto/global"
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

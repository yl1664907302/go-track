package mysqldatabase

import (
	"gorm.io/gorm"
	"kube-auto/global"
	"kube-auto/pojo"
	"log"
)

type MysqlUserDAO struct {
	Db *gorm.DB
}

func NewMysqlUserDAO(db *gorm.DB) *MysqlUserDAO {
	return &MysqlUserDAO{Db: db}
}

func (d *MysqlUserDAO) GetByNameAndPasswd(username string, password string) (*pojo.User, error) {
	var user pojo.User
	d.Db = global.MysqlDataConnect
	err := d.Db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		log.Print(err)
	}
	return &user, err
}

package mysql

import (
	"go-track/global"
	"go-track/pojo"
	"log"
)

func LoginUser(username string, password string) (pojo.User, error) {
	var user pojo.User
	//err := global.MysqlDataConnect.Select("username,password").Find(&p).Error
	err := global.MysqlDataConnect.Where("username = ? AND password = ?", username, password).First(&user).Error
	log.Println(user)
	return user, err
}

func InsertReceiver(receiver string, niname string) error {
	var r pojo.Receiver
	r.Niname = niname
	r.Receiver_name = receiver
	err := global.MysqlDataConnect.Create(&r).Error
	return err
}

package mysql

import (
	"go-track/global"
	"go-track/pojo"
)

func LoginUser(p *pojo.User) error {
	err := global.MysqlDataConnect.Select("username,password").Find(&p).Error
	return err
}

func InsertReceiver(receiver string, niname string) error {
	var r pojo.Receiver
	r.Niname = niname
	r.Receiver_name = receiver
	err := global.MysqlDataConnect.Create(&r).Error
	return err
}

package mysql

import (
	"go-track/global"
	"go-track/pojo"
)

func SelectReceivers() ([]pojo.Receiver, error) {
	var r []pojo.Receiver
	err := global.MysqlDataConnect.Select("id,receiver_name").Find(&r).Error
	return r, err
}

func InsertReceiver(receiver string) error {
	var r pojo.Receiver
	r.Receiver_name = receiver
	err := global.MysqlDataConnect.Create(&r).Error
	return err
}

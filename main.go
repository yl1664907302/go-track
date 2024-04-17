package main

import (
	"kube-auto/global"
	"kube-auto/initiallize"
)

func main() {
	//读取yaml配置文件
	initiallize.Viper()
	initiallize.InitMysqlDB()
	r := initiallize.Router()
	panic(r.Run(global.CONF.System.Addr))
}

package main

import (
	"kube-auto/global"
	"kube-auto/initiallize"
)

func main() {
	r := initiallize.Router()
	initiallize.Viper()
	panic(r.Run(global.CONF.System.Addr))
}

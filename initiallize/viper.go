package initiallize

import (
	"github.com/spf13/viper"
	"go-track/global"
	"log"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Print(err.Error())
	}
	err = v.Unmarshal(&global.CONF)
	if err != nil {
		log.Print(err.Error())
	}
}

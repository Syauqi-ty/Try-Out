package config

import (
	"time"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	localTime, _ := time.LoadLocation(viper.GetString("timezone"))
	viper.Set("TimeZone", localTime)
}

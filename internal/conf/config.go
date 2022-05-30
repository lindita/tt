package conf

import (
	"github.com/spf13/viper"
)

var config Config

func GetConfig() *Config {
	return &config
}

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./configs/env.yml")
	//v.AddConfigPath("./configs/")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}
}

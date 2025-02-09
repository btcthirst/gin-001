package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	return initViper()
}

func initViper() error {
	viper.SetConfigFile(".env")
	return viper.ReadInConfig()
}

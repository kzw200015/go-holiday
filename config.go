package main

import (
	"fmt"
	"github.com/spf13/viper"
)

const remoteUrl = "https://raw.githubusercontent.com/NateScarlet/holiday-cn/master/{year}.json"

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("HOLIDAY")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("holiday.remote_url", remoteUrl)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

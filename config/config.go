package config

import "github.com/spf13/viper"

var Version string

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Version = viper.GetString("version")
}

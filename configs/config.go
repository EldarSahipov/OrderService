package configs

import "github.com/spf13/viper"

func InitConf() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("appconfig")
	return viper.ReadInConfig()
}

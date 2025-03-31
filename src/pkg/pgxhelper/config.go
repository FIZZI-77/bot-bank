package pgxhelper

import "github.com/spf13/viper"

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}

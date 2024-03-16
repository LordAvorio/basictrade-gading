package utils

import "github.com/spf13/viper"

func ReadConfigEnvironment() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if errReadConfig := viper.ReadInConfig(); errReadConfig != nil {
		panic(errReadConfig)
	}
}
package config

import (
	"github.com/spf13/viper"
)

func ReadConfig() error {
	viper.AddConfigPath("./config/config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

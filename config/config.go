package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Database Database
	Redis    Redis
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		config = nil
		return config, err
	}

	return config, nil
}

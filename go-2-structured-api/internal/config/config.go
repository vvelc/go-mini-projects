package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port    int16 `mapstructure:"PORT"`
	IsDebug bool  `mapstructure:"DEBUG"`
}

var config *Config = &Config{
	Port:    3000,
	IsDebug: true,
}

func LoadConfig() (c *Config, err error) {
	viper.AddConfigPath(".")
    viper.SetConfigName(".env")
    viper.SetConfigType("env")

	viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)

    return config, err
}

func GetConfig() *Config {
	return config
}
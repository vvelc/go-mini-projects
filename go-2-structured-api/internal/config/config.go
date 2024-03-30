package config

type Config struct {
	PORT int16
	IS_DEBUG bool
}

var config *Config = &Config{
	PORT: 3000,
	IS_DEBUG: true,
}

func GetConfig() *Config {
	return config
}
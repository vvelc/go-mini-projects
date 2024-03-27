package config

type Config struct {
	PORT int16
}

var config *Config = &Config{
	PORT: 3000,
}

func GetConfig() *Config {
	return config
}
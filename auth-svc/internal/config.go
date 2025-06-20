package internal

import "light-ms/pkg/common/config"

type Config struct {
	ServerPort string `env:"SERVER_PORT"`
	JwtKey     string `env:"JWT_KEY"`
}

func LoadConf() Config {
	conf := Config{}

	config.Load(&conf, "auth-svc/.env")

	return conf
}

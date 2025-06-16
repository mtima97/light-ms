package internal

import "light-ms/pkg/common/config"

type Config struct {
	config.Svc
	ServerPort string `env:"SERVER_PORT"`
}

func LoadConf() Config {
	conf := Config{}

	config.Load(&conf, "auth-svc/.env")

	return conf
}

package config

import (
	"light-ms/pkg/common/config"
)

type Config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     string `env:"DB_PORT"`
	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbName     string `env:"DB_NAME"`
	JwtKey     string `env:"JWT_KEY"`
	ServerPort string `env:"SERVER_PORT"`
}

func LoadConf() Config {
	conf := Config{}

	config.Load(&conf, "order/.env")

	return conf
}

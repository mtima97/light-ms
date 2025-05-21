package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	DbUser         string   `env:"DB_USER"`
	DbPassword     string   `env:"DB_PASSWORD"`
	DbHost         string   `env:"DB_HOST" envDefault:"localhost"`
	DbName         string   `env:"DB_NAME" envDefault:"postgres"`
	DbPort         string   `env:"CUSTOM_DB_PORT"`
	ServerPort     string   `env:"ORDERS_SERVER_PORT" envDefault:":8080"`
	TrustedProxies []string `env:"TRUSTED_PROXIES"`
}

func LoadConf() Config {
	_ = godotenv.Load()

	conf := Config{}

	if err := env.Parse(&conf); err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return conf
}

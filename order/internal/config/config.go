package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     string `env:"CUSTOM_DB_PORT"`
}

func LoadConf() Config {
	_ = godotenv.Load()

	conf := Config{}

	if err := env.Parse(&conf); err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return conf
}

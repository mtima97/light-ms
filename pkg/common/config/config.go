package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

type Dsn struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func Load(conf any, path string) {
	_ = godotenv.Load(path)

	if err := env.Parse(conf); err != nil {
		log.Fatal(err)
	}
}

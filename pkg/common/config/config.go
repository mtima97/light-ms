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

type Svc struct {
	DbUser         string   `env:"DB_USER"`
	DbPassword     string   `env:"DB_PASSWORD"`
	DbPort         string   `env:"DB_PORT"`
	DbHost         string   `env:"DB_HOST" envDefault:"localhost"`
	DbName         string   `env:"DB_NAME"`
	TrustedProxies []string `env:"TRUSTED_PROXIES"`
	JwtKey         string   `env:"JWT_KEY"`
}

func Load(conf any, path string) {
	_ = godotenv.Load(path, ".env")

	if err := env.Parse(conf); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	_ "light-ms/order/docs"
	"light-ms/order/internal/config"
	"light-ms/order/internal/repository"
	"light-ms/order/internal/routes"
	"light-ms/order/internal/usecase"
	cmnconf "light-ms/pkg/common/config"
	"log"
)

// @title Мой АПИ
// @version 1.0
// @host http://localhost:9001
// @BasePath /
func main() {
	ctx := context.Background()
	cfg := config.LoadConf()

	db := repository.NewOrderRepository(ctx, dsn(cfg))
	defer db.Close()

	router := routes.NewRouter(cfg)
	router.Register(usecase.NewOrdersUcase(db))

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

func dsn(c config.Config) cmnconf.Dsn {
	return cmnconf.Dsn{
		Host:     c.DbHost,
		Port:     c.DbPort,
		User:     c.DbUser,
		Password: c.DbPassword,
		Database: c.DbName,
	}
}

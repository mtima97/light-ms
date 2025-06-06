package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"light-ms/order/internal/config"
	"light-ms/order/internal/repository"
	"light-ms/order/internal/routes"
	"light-ms/order/internal/usecase"
	cmnconf "light-ms/pkg/common/config"
	"log"
)

func main() {
	ctx := context.Background()

	conf := config.LoadConf()

	db := repository.NewOrderRepository(ctx, dsn(conf))
	defer db.Close()

	r := setupRouter(usecase.NewOrdersUcase(db))

	if err := r.SetTrustedProxies(conf.TrustedProxies); err != nil {
		log.Fatal(err)
	}

	log.Printf("server started on port %s", conf.ServerPort)

	if err := r.Run(conf.ServerPort); err != nil {
		log.Fatal(err)
	}
}

func setupRouter(ordersUcase usecase.OrdersUcase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	routes.RegisterRoutes(r, ordersUcase)

	return r
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

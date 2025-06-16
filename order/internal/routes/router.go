package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"light-ms/order/internal/config"
	"light-ms/order/internal/handlers"
	"light-ms/order/internal/usecase"
	"light-ms/pkg/common/middleware"
)

type Router struct {
	conf   config.Config
	engine *gin.Engine
}

func NewRouter(conf config.Config) Router {
	return Router{
		conf:   conf,
		engine: gin.Default(),
	}
}

func (r Router) Register(ucase usecase.OrdersUcase) {
	handler := handlers.NewOrdersHandler(ucase)

	r.engine.POST("/orders", middleware.Auth(r.conf.JwtKey), handler.CreateOrder)
	r.engine.GET("/orders", handler.Get)
	r.engine.GET("/orders/:id", handler.GetById)
	r.engine.PATCH("/orders/:id/status", middleware.Auth(r.conf.JwtKey), handler.UpdateStatus)

	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r Router) Run() error {
	if err := r.engine.SetTrustedProxies(r.conf.TrustedProxies); err != nil {
		return err
	}

	return r.engine.Run(r.conf.ServerPort)
}

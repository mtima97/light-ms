package internal

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Router struct {
	config Config
	engine *gin.Engine
}

func NewRouter(conf Config) Router {
	return Router{
		config: conf,
		engine: gin.Default(),
	}
}

func (r Router) Register() {
	handler := NewHandler(r.config)

	r.engine.POST("/auth/login", handler.Login)
}

func (r Router) Run() error {
	if err := r.engine.Run(r.config.ServerPort); err != nil {
		return err
	}

	log.Printf("Listening on port %s", r.config.ServerPort)

	return nil
}

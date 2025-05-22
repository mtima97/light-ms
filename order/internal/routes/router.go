package routes

import (
	"github.com/gin-gonic/gin"
	"light-ms/order/internal/handlers"
	"light-ms/order/internal/usecase"
)

func RegisterRoutes(r *gin.Engine, ordersUcase usecase.OrdersUcase) {
	ordersH := handlers.NewOrdersHandler(ordersUcase)

	r.POST("/orders", ordersH.CreateOrder)
	r.GET("/orders/:id", ordersH.GetById)
	r.PATCH("/orders/:id/status", ordersH.UpdateStatus)
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"light-ms/order/internal/usecase"
)

type OrdersHandler struct {
	ucase usecase.OrdersUcase
}

func NewOrdersHandler(ucase usecase.OrdersUcase) OrdersHandler {
	return OrdersHandler{ucase: ucase}
}

func (h OrdersHandler) CreateOrder(ctx *gin.Context) {
	//
}

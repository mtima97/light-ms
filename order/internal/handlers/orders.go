package handlers

import (
	"github.com/gin-gonic/gin"
	"light-ms/order/internal/models/requests"
	umodels "light-ms/order/internal/models/usecase"
	"light-ms/order/internal/usecase"
	"net/http"
)

type OrdersHandler struct {
	BaseHandler
	ucase usecase.OrdersUcase
}

func NewOrdersHandler(ucase usecase.OrdersUcase) OrdersHandler {
	return OrdersHandler{ucase: ucase}
}

func (h OrdersHandler) CreateOrder(ctx *gin.Context) {
	var req requests.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, h.errorResponse(err.Error()))
		return
	}

	dto := umodels.CreateOrderDto{
		UserId: req.UserId,
		Amount: req.Amount,
	}

	if err := h.ucase.CreateOrder(ctx, dto); err != nil {
		ctx.JSON(http.StatusBadRequest, h.errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, h.successResponse(nil, http.StatusText(http.StatusCreated)))
}

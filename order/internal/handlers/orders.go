package handlers

import (
	"github.com/gin-gonic/gin"
	"light-ms/order/internal/models/requests"
	umodels "light-ms/order/internal/models/usecase"
	"light-ms/order/internal/usecase"
	"net/http"
	"strconv"
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

func (h OrdersHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, h.errorResponse(err.Error()))
		return
	}

	order, err := h.ucase.GetById(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, h.errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, h.successResponse(order, http.StatusText(http.StatusOK)))
}

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
	ucase usecase.OrdersUcase
}

func NewOrdersHandler(ucase usecase.OrdersUcase) OrdersHandler {
	return OrdersHandler{ucase: ucase}
}

func (h OrdersHandler) CreateOrder(ctx *gin.Context) {
	var req requests.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	dto := umodels.CreateOrderDto{
		UserId: req.UserId,
		Amount: req.Amount,
	}

	if err := h.ucase.CreateOrder(ctx, dto); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, successResponse(nil, http.StatusText(http.StatusCreated)))
}

func (h OrdersHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := h.ucase.GetById(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(order, http.StatusText(http.StatusOK)))
}

func (h OrdersHandler) Get(ctx *gin.Context) {
	orders, err := h.ucase.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(orders, http.StatusText(http.StatusOK)))
}

func (h OrdersHandler) UpdateStatus(ctx *gin.Context) {
	var req requests.UpdateStatusRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	dto := umodels.UpdateStatusDto{
		OrderId: req.OrderId,
		Status:  req.Status,
	}

	if err := h.ucase.UpdateStatus(ctx, dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(nil, http.StatusText(http.StatusOK)))
}

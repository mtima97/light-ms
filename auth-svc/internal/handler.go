package internal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	config Config
}

func NewHandler(config Config) Handler {
	return Handler{
		config: config,
	}
}

func (h Handler) Login(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorMsg(err))
		return
	}

	if req.UserId != 1 {
		ctx.JSON(http.StatusUnprocessableEntity, errorMsgFromStr(http.StatusText(http.StatusUnprocessableEntity)))
		return
	}

	token, err := GenerateTkn(req.UserId, h.config.JwtKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(token))
}

func errorMsg(err error) gin.H {
	return gin.H{
		"error": fmt.Sprintf("Ошибка: %v", err),
	}
}

func errorMsgFromStr(msg string) gin.H {
	return errorMsg(errors.New(msg))
}

func successResponse(token string) gin.H {
	return gin.H{
		"token": token,
	}
}

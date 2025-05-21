package handlers

import "github.com/gin-gonic/gin"

type BaseHandler struct {
	//
}

func (h BaseHandler) errorResponse(message string) gin.H {
	return gin.H{
		"error": message,
	}
}

func (h BaseHandler) successResponse(data any, message string) gin.H {
	return gin.H{
		"data":    data,
		"message": message,
	}
}

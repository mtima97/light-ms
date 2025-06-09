package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{
		"message": fmt.Sprintf("error: %v", err),
		"data":    nil,
	}
}

func successResponse(data any, message string) gin.H {
	return gin.H{
		"message": message,
		"data":    data,
	}
}

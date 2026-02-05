package helpers

import "github.com/gin-gonic/gin"

func SuccessResponse(ctx *gin.Context, message string, data any) {
	ctx.JSON(200, gin.H{
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string, err any) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"error":   err,
	})
}

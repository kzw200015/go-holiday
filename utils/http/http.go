package http

import "github.com/gin-gonic/gin"

func Error(context *gin.Context, err error) {
	context.JSON(400, gin.H{
		"error": err.Error(),
	})
}

func Ok(context *gin.Context, body gin.H) {
	context.JSON(200, body)
}

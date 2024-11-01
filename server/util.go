package server

import "github.com/gin-gonic/gin"

func sendErr(context *gin.Context, err error) {
	context.JSON(400, gin.H{
		"error": err.Error(),
	})
}

func sendOk(context *gin.Context, body gin.H) {
	context.JSON(200, body)
}

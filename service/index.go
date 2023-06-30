package service

import "github.com/gin-gonic/gin"

func GetIndex(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello",
	})
}

package service

import "github.com/gin-gonic/gin"

// GetIndex godoc
// @Tags         index
// @Success      200  {string}  welcome
// @Router       /index [get]
func GetIndex(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world",
	})
}

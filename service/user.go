package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	user := models.GetUserList()
	ctx.JSON(200, gin.H{"message": user})
}

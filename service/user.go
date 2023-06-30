package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Tags         getUser
// @Success      200  {string} json{"message"}
// @Router       /user/getUser [get]
func GetUser(ctx *gin.Context) {
	user := models.GetUserList()
	ctx.JSON(200, gin.H{"message": user})
}

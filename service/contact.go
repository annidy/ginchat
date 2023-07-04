package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"

	"github.com/gin-gonic/gin"
)

// SearchFriends godoc
// @Summary      获取用户好友信息
// @Tags         用户
// @Param        userId query string false "userId"
// @Success      200  {string} json{"message"}
// @Router       /user/searchFriends [post]
func SearchFriends(ctx *gin.Context) {
	userId := utils.Atou(ctx.PostForm("userId"))
	// TODO: 校验token
	users := models.SearchFriends(userId)
	utils.RespOkList(ctx.Writer, users, users)
}

func AddFriend(ctx *gin.Context) {
	userId := utils.Atou(ctx.PostForm("userId"))
	targetName := ctx.PostForm("targetName")

	targetUser, err := models.FindUserByName(targetName)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": fmt.Sprintf("user %s not found", targetName)})
		return
	}
	res := models.AddFriend(userId, targetUser.ID)
	if res.Error != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": res.Error.Error()})
		return
	}
	ctx.JSON(200, gin.H{"code": 0, "message": "success"})
}

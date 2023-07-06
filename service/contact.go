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
	utils.RespOkList(ctx.Writer, users, len(users))
}

func AddFriend(ctx *gin.Context) {
	userId := utils.Atou(ctx.PostForm("userId"))
	targetName := ctx.PostForm("targetName")

	targetUser, err := models.FindUserByName(targetName)
	if err != nil {
		utils.RespFail(ctx.Writer, fmt.Sprintf("user %s not found", targetName))
		return
	}
	if targetUser.ID == userId {
		utils.RespFail(ctx.Writer, "can't add yourself as friend")
		return
	}
	friends := models.SearchFriends(userId)
	for _, f := range friends {
		if f.ID == targetUser.ID {
			utils.RespFail(ctx.Writer, "already friends")
			return
		}
	}
	if err := models.AddFriend(userId, targetUser.ID); err != nil {
		utils.RespFail(ctx.Writer, err.Error())
		return
	}
	utils.RespOk(ctx.Writer, nil, "success")
}

func CreateCommunity(ctx *gin.Context) {

}

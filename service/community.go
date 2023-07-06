package service

import (
	"ginchat/models"
	"ginchat/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCommunity(ctx *gin.Context) {
	community := models.Community{}
	community.OwnerId = utils.Atou(ctx.PostForm("ownerId"))
	community.Name = ctx.PostForm("name")
	community.Icon = ctx.PostForm("icon")
	community.Desc = ctx.PostForm("desc")
	community.Cate, _ = strconv.Atoi(ctx.PostForm("cate"))
	community.Memo = ctx.PostForm("memo")
	if err := models.CreateCommunity(&community); err != nil {
		utils.RespFail(ctx.Writer, err.Error())
		return
	}
	utils.RespOk(ctx.Writer, community, "success")
}

func LoadCommunity(ctx *gin.Context) {
	ownerId := utils.Atou(ctx.PostForm("ownerId"))
	communities, _ := models.LoadCommunity(ownerId)
	utils.RespOkList(ctx.Writer, communities, len(communities))
}

package service

import (
	"fmt"
	"net/http"
	"text/template"

	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIndex(ctx *gin.Context) {
	tpl, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	tpl.Execute(ctx.Writer, nil)
}

func ToRegister(ctx *gin.Context) {
	tpl, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	tpl.Execute(ctx.Writer, "register")
}

func ToChat(c *gin.Context) {
	tpl, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	//fmt.Println("ToChat>>>>>>>>", user)
	tpl.Execute(c.Writer, user)
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

// func Chat(c *gin.Context) {
// 	models.Chat(c.Writer, c.Request)
// }

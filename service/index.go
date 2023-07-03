package service

import (
	"fmt"
	"net/http"
	"text/template"

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

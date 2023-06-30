package router

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUser", service.GetUser)
	return r
}

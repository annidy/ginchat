package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/getUser", service.GetUser)
	r.POST("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)

	r.POST("/searchFriends", service.SearchFriends)
	r.POST("/contact/addfriend", service.AddFriend)
	r.POST("/attach/upload", service.Upload)

	r.POST("/contact/createCommunity", service.CreateCommunity)
	r.POST("/contact/loadcommunity", service.LoadCommunity)

	r.GET("/chat", service.Chat)

	// 静态文件
	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("views/**/*")
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)

	return r
}

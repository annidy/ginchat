package service

import (
	"ginchat/models"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList godoc
// @Summary      获取用户列表
// @Tags         用户
// @Success      200  {string} json{"message"}
// @Router       /user/getUserList [get]
func GetUserList(ctx *gin.Context) {
	user := models.GetUserList()
	ctx.JSON(200, gin.H{"message": user})
}

// CreateUser godoc
// @Summary      创建用户
// @Tags         用户
// @Param        name query string false "name"
// @Param        password query string false "password"
// @Success      200  {string} json{code, message}
// @Router       /user/createUser [get]
func CreateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Query("name")
	user.Password = ctx.Query("password")
	if user.Name == "" || user.Password == "" {
		ctx.JSON(200, gin.H{"code": 1, "message": "name or password is empty"})
		return
	}
	result := models.CreateUser(&user)
	if result.Error != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": result.Error.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": user})
}

// CreateUser godoc
// @Summary      删除用户
// @Tags         用户
// @Param        id query string false "id"
// @Success      200  {string} json{code, message}
// @Router       /user/deleteUser [get]
func DeleteUser(ctx *gin.Context) {
	user := models.UserBasic{}
	ID, _ := strconv.Atoi(ctx.Query("id"))
	user.ID = uint(ID)
	result := models.DeleteUser(&user)
	if result.Error != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

// UpdateUser godoc
// @Summary      修改用户
// @Tags         用户
// @Param        id formData integer false "id"
// @Param        name formData string false "name"
// @Param        password formData string false "password"
// @Param        phone formData string false "phone"
// @Param        email formData string false "email"
// @Success      200  {string} json{code, message}
// @Router       /user/updateUser [post]
func UpdateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	ID, _ := strconv.Atoi(ctx.PostForm("id"))
	user.ID = uint(ID)
	user.Name = ctx.PostForm("name")
	user.Password = ctx.PostForm("password")
	user.Phone = ctx.PostForm("phone")
	user.Email = ctx.PostForm("email")
	_, err := govalidator.ValidateStruct(&user)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	models.UpdateUser(&user)
	ctx.JSON(200, gin.H{"message": user})
}

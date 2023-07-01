package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Identity   string
	Name       string
	Password   string
	Phone      string
	Email      string
	ClientIp   string
	ClientPort string
	LoginTime  time.Time
	LogoutTime *time.Time
	IsLogout   bool
	DeviceId   string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func (user *UserBasic) BeforeCreate(tx *gorm.DB) error {
	user.LoginTime = time.Now()
	fmt.Println("before create", user)
	return nil
}

func Init() {
	utils.Db.AutoMigrate(&UserBasic{})
}

func GetUserList() []*UserBasic {
	userBasic := make([]*UserBasic, 10)
	utils.Db.Find(&userBasic)
	return userBasic
}

func CreateUser(user *UserBasic) *gorm.DB {
	return utils.Db.Create(user)
}

func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.Db.Delete(user)
}

func UpdateUser(user *UserBasic) *gorm.DB {
	return utils.Db.Save(user)
}

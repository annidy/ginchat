package models

import (
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
	LogoutTime time.Time
	IsLogout   bool
	DeviceId   string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	userBasic := make([]*UserBasic, 10)
	utils.Db.Find(&userBasic)
	return userBasic
}

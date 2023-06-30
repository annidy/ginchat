package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity   string
	Name       string
	Password   string
	Phone      string
	Email      string
	ClientIp   string
	ClientPort string
	LoginTime  uint64
	LogoutTime uint64
	IsLogout   bool
	DeviceId   string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

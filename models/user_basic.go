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
	Phone      string `valid:"matches(^1[0-9]{10}$)"`
	Email      string `valid:"email"`
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
	if FindUserByName(user.Name) != nil {
		return fmt.Errorf("user name %s already exists", user.Name)
	}
	if FindUserByPhone(user.Phone) != nil {
		return fmt.Errorf("user phone %s already exists", user.Phone)
	}
	user.LoginTime = time.Now()
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

func FindUserByName(name string) *gorm.DB {
	return utils.Db.Where("name = ?", name).First(&UserBasic{})
}

func FindUserByPhone(phone string) *gorm.DB {
	return utils.Db.Where("phone = ?", phone).First(&UserBasic{})
}

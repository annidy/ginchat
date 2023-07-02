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

func (user *UserBasic) SaveIdentity() *gorm.DB {
	return utils.Db.Model(user).Where("id = ?", user.ID).Update("identity", user.Identity)
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

func FindUserByNameAndPassword(name, password string) (*UserBasic, error) {
	user := UserBasic{}
	if utils.Db.Where("name = ? AND password = ?", name, password).First(&user).Error != nil {
		return nil, fmt.Errorf("user name or password is wrong")
	}
	return &user, nil
}

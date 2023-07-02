package models

import "ginchat/utils"

func Init() {
	utils.Db.AutoMigrate(&UserBasic{})
	utils.Db.AutoMigrate(&Message{})
	utils.Db.AutoMigrate(&GroupBasic{})
}

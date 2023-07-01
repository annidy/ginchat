package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitConfig() {
	viper.SetConfigFile("./config/app.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.ConfigFileUsed())
}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			Colorful:                  true,
			IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
		},
	)
	db, err := gorm.Open(mysql.Open("root:1234@tcp(mysql8019:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	Db = db
}

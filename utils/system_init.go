package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db  *gorm.DB
	Rdb *redis.Client
)

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

func InitRedies() {
	addr := fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))
	fmt.Println("redis addr:", addr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: viper.GetString("redies.password"), // no password set
		DB:       viper.GetInt("redies.db"),          // use default DB
		Username: viper.GetString("redies.username"),
		PoolSize: viper.GetInt("redies.pool_size"),
	})
	var ctx = context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	Rdb = rdb
}

const (
	PublishKey = "websocket"
)

func Publish(ctx context.Context, channel string, msg interface{}) error {
	fmt.Println("Publish", channel, msg)
	return Rdb.Publish(ctx, channel, msg).Err()
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	msg, err := Rdb.Subscribe(ctx, channel).ReceiveMessage(ctx)
	if err != nil {
		return "", err
	}
	fmt.Println("Subscribe", channel, msg.Payload)
	return msg.Payload, err
}

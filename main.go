package main

import (
	"fmt"
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	models.Init()

	r := router.Router()
	fmt.Println("Start :8080")
	r.Run(":8080")
}

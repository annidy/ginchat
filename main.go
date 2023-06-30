package main

import (
	"fmt"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	fmt.Println("Start :8080")
	r.Run(":8080")
}

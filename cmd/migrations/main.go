package main

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
)

func main() {
	err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	db := configs.ConnectDB()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Url{})
}

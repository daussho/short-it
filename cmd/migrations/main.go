package main

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
)

func main() {
	db := configs.ConnectDB()

	db.Migrator().DropTable(models.Url{})
	db.AutoMigrate(&models.Url{})
}

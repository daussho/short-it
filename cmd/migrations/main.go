package main

import (
	"fmt"
	"log"

	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	err := configs.InitDB()
	if err != nil {
		panic(err)
	}

	db = configs.ConnectDB()

	log.Print("Input migration number: ")
	var migrationNumber string
	_, err = fmt.Scanln(&migrationNumber)
	if err != nil {
		panic(err)
	}

	switch migrationNumber {
	case "0001":
		migration0001()
	case "0002":
		migration0002()
	default:
		log.Fatal("Migration number not found")
	}
}

func migration0001() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Url{})
}

func migration0002() {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	db.Create(&models.User{
		Username: "admin",
		Password: string(bcryptPassword),
	})
}

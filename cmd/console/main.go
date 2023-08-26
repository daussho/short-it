package main

import (
	"fmt"
	"log"

	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	configs.InitDB()

	log.Println("1. Create user")
	log.Println("2. Change user password")
	log.Print("Input menu: ")
	var menuNumber string
	_, err := fmt.Scanln(&menuNumber)
	if err != nil {
		panic(err)
	}

	switch menuNumber {
	case "1":
		createUser()
	case "2":
		changeUserPassword()
	default:
		log.Fatal("Menu number not found")
	}
}

func createUser() {
	db := configs.ConnectDB()

	var username string
	log.Print("Input username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		panic(err)
	}

	var user models.User
	db.Find(&user, "username = ?", username)
	if user.ID > 0 {
		log.Fatal("User already exists")
	}

	var password string
	log.Print("Input password: ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		panic(err)
	}

	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	db.Create(&models.User{
		Username: username,
		Password: string(bcryptPassword),
	})

	log.Println("User created", username, password)
}

func changeUserPassword() {
	db := configs.ConnectDB()

	var username string
	log.Print("Input username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		panic(err)
	}

	var user models.User
	db.Find(&user, "username = ?", username)
	if user.ID == 0 {
		log.Fatal("User not found")
	}

	var password string
	log.Print("Input new password: ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		panic(err)
	}

	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(bcryptPassword)
	user.Token = ""
	db.Save(&user)

	log.Println("User password changed", username, password)
}

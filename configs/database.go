package configs

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var (
	db *gorm.DB
)

func InitDB() error {
	dbTemp, err := gorm.Open(sqlite.Open("short-it.db"), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	db = dbTemp

	return nil
}

func ConnectDB() *gorm.DB {
	if db == nil {
		log.Fatal("db is nil")
	}

	return db
}

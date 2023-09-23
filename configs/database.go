package configs

import (
	"log"
	"os"

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
	isProduction := os.Getenv("APP_ENV") == "production"
	gormLogLevel := logger.Error
	if !isProduction {
		gormLogLevel = logger.Info
	}

	var err error
	db, err = gorm.Open(sqlite.Open("short-it.db"), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(gormLogLevel),
	})

	return err
}

func ConnectDB() *gorm.DB {
	if db == nil {
		log.Fatal("db is nil")
	}

	return db
}

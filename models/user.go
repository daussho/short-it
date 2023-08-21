package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
	Token    string `gorm:"uniqueIndex"`
}

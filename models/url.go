package models

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey"`
	Url         string
	ShortUrl    string `gorm:"uniqueIndex"`
	ViewCount   int64
	OwnerUserID int64 `gorm:"index"`
}

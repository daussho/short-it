package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Url         string         `json:"url"`
	ShortUrl    string         `json:"shortUrl" gorm:"uniqueIndex"`
	ViewCount   int64          `json:"viewCount"`
	OwnerUserID uint           `json:"ownerUserId" gorm:"index"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

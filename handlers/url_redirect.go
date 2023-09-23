package handlers

import (
	"context"
	"sync"

	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var m = &sync.Mutex{}

func UrlRedirect(ctx *fiber.Ctx) error {
	var url models.Url
	db := configs.ConnectDB()

	db.Find(&url, "short_url = ?", ctx.Params("shortUrl"))
	if url.ID == 0 {
		return ctx.Redirect("/")
	}

	go updateCount(url)

	return ctx.Redirect(url.Url)
}

func updateCount(url models.Url) {
	m.Lock()
	defer m.Unlock()

	db := configs.ConnectDB()

	db.WithContext(context.Background()).Model(&url).Update("view_count", gorm.Expr("view_count + ?", 1))
}

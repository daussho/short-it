package handlers

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UrlRedirect(ctx *fiber.Ctx) error {
	var url models.Url
	db := configs.ConnectDB()

	db.Find(&url, "short_url = ?", ctx.Params("shortUrl"))
	if url.ID == 0 {
		return ctx.Redirect("/")
	}

	db.Model(&url).Update("view_count", gorm.Expr("view_count + ?", 1))

	return ctx.Redirect(url.Url)
}

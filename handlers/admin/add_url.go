package admin

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/daussho/short-it/utils"
	"github.com/gofiber/fiber/v2"
)

func AddUrl(ctx *fiber.Ctx) error {
	url := ctx.FormValue("url")
	if url == "" {
		return ctx.Redirect("/admin?error=URL cannot be empty")
	}

	user := ctx.Locals("user").(models.User)

	shortUrl := utils.RandStringRunes(6)

	db := configs.ConnectDB()
	db.Create(&models.Url{
		Url:         url,
		ShortUrl:    shortUrl,
		ViewCount:   0,
		OwnerUserID: user.ID,
	})

	return ctx.Redirect("/admin")
}

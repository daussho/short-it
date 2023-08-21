package admin

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/gofiber/fiber/v2"
)

func Dashboard(ctx *fiber.Ctx) error {
	var urls []models.Url
	db := configs.ConnectDB()

	userID := ctx.Locals("user").(models.User).ID
	db.Where("owner_user_id = ?", userID).Find(&urls)

	return ctx.Render("admin/dashboard", fiber.Map{
		"title": "Dashboard",
		"urls":  urls,
	}, "layouts/main")
}

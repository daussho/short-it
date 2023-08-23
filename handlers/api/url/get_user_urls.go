package url

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/gofiber/fiber/v2"
)

func GetUserUrls(ctx *fiber.Ctx) error {
	// get user from session
	user := ctx.Locals("user").(models.User)

	db := configs.ConnectDB()
	var urls []models.Url
	db.Where("owner_user_id = ?", user.ID).Order("id DESC").Find(&urls)

	return ctx.JSON(fiber.Map{
		"data": urls,
	})
}

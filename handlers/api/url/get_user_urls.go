package url

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/gofiber/fiber/v2"
)

func GetUserUrls(ctx *fiber.Ctx) error {
	// get user from session
	user := ctx.Locals("user").(models.User)

	// get param
	search := ctx.Query("search")

	db := configs.ConnectDB()
	var urls []models.Url
	db = db.Where("owner_user_id = ?", user.ID)

	if search != "" {
		db = db.Where("url LIKE ? OR short_url LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	db = db.Order("id DESC").Find(&urls)

	return ctx.JSON(fiber.Map{
		"data": urls,
	})
}

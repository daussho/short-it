package url

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/gofiber/fiber/v2"
)

// fiber handler
func DeleteUrl(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(models.User)

	db := configs.ConnectDB()

	db.Where("owner_user_id = ?", user.ID).
		Where("id = ?", ctx.Params("id")).
		Delete(&models.Url{})

	return ctx.JSON(fiber.Map{
		"message": "URL deleted successfully",
	})
}

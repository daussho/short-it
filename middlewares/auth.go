package middlewares

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	// check if user is logged in
	token := ctx.Cookies("token")
	if token == "" {
		return ctx.Redirect("/admin/login")
	}

	// check if token is valid
	db := configs.ConnectDB()
	var user models.User
	db.Find(&user, "token = ?", token)
	if user.ID == 0 {
		return ctx.Redirect("/admin/login")
	}

	ctx.Locals("user", user)
	return ctx.Next()
}

func AuthJSON(ctx *fiber.Ctx) error {
	// check if user is logged in
	token := ctx.Cookies("token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	// check if token is valid
	db := configs.ConnectDB()
	var user models.User
	db.Find(&user, "token = ?", token)
	if user.ID == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	ctx.Locals("user", user)
	return ctx.Next()
}

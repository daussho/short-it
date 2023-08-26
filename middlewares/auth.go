package middlewares

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	// check if user is logged in
	token := ctx.Cookies("token")
	user, valid := validateToken(token)
	if !valid {
		return ctx.Redirect("/admin/login")
	}

	ctx.Locals("user", user)
	return ctx.Next()
}

func AuthJSON(ctx *fiber.Ctx) error {
	// check if user is logged in
	token := ctx.Cookies("token")
	user, valid := validateToken(token)
	if !valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	ctx.Locals("user", user)
	return ctx.Next()
}

func validateToken(token string) (models.User, bool) {
	if token == "" {
		return models.User{}, false
	}

	if len(token) != models.TOKEN_LENGTH {
		return models.User{}, false
	}

	// check if token is valid
	db := configs.ConnectDB()
	var user models.User
	db.Find(&user, "token = ?", token)

	return user, user.ID > 0
}

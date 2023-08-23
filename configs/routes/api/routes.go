package api

import "github.com/gofiber/fiber/v2"

func Init(app *fiber.App) {
	urlRoutes(app)
}

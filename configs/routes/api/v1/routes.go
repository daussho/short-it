package v1

import "github.com/gofiber/fiber/v2"

func Init(app *fiber.App) {
	v1Routes := app.Group("/api/v1")
	urlRoutes(v1Routes)
}

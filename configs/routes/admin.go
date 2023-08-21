package routes

import (
	"github.com/daussho/short-it/handlers/admin"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	adminRoute := app.Group("/admin")

	// Login
	adminRoute.Get("/login", admin.Login)
	adminRoute.Post("/login", admin.Login)
}

package routes

import (
	"github.com/daussho/short-it/handlers/admin"
	"github.com/daussho/short-it/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	app.Get("/admin/login", admin.Login)
	app.Post("/admin/login", admin.Login)

	adminRoute := app.Group("/admin", middlewares.Auth)

	// Login
	adminRoute.Get("/", admin.Dashboard)
}

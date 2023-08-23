package web

import (
	"github.com/daussho/short-it/handlers/admin"
	"github.com/daussho/short-it/middlewares"
	"github.com/gofiber/fiber/v2"
)

func adminRoutes(app *fiber.App) {
	app.Get("/admin/login", admin.Login)
	app.Post("/admin/login", admin.Login)

	adminRoute := app.Group("/admin", middlewares.Auth)

	// Login
	adminRoute.Get("/", admin.Dashboard)
	adminRoute.Post("/add-url", admin.AddUrl)
}

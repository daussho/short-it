package web

import (
	"github.com/daussho/short-it/handlers"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	adminRoutes(app)

	app.Get("/:shortUrl<len(6)>", handlers.UrlRedirect)

	// 404 page
	app.Get("/404", func(c *fiber.Ctx) error {
		return c.Render("404", fiber.Map{
			"title": "404",
		}, "layouts/main")
	})
}

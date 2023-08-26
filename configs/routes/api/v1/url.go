package v1

import (
	urlHandler "github.com/daussho/short-it/handlers/api/url"
	"github.com/daussho/short-it/middlewares"
	"github.com/gofiber/fiber/v2"
)

func urlRoutes(app fiber.Router) {
	route := app.Group("/url", middlewares.AuthJSON)

	route.Get("/", urlHandler.GetUserUrls)
	route.Post("/", urlHandler.AddUrl)
	route.Delete("/:id", urlHandler.DeleteUrl)
}

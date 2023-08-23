package api

import (
	urlHandler "github.com/daussho/short-it/handlers/api/url"
	"github.com/daussho/short-it/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UrlRoutes(app *fiber.App) {

	route := app.Group("/api/v1/url", middlewares.AuthJSON)

	route.Get("/", urlHandler.GetUserUrls)
}

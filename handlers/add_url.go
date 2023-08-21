package handlers

import "github.com/gofiber/fiber/v2"

func AddUrl(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

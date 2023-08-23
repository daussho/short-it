package admin

import (
	"github.com/gofiber/fiber/v2"
)

func Dashboard(ctx *fiber.Ctx) error {
	return ctx.Render("admin/dashboard", fiber.Map{
		"title": "Dashboard",
	}, "layouts/main")
}

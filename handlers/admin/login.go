package admin

import (
	"log"
	"os"

	"github.com/daussho/short-it/configs"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	method := ctx.Method()
	if method == "GET" {
		return loginGet(ctx)
	} else if method == "POST" {
		return loginPost(ctx)
	}

	return fiber.ErrBadRequest
}

func loginGet(ctx *fiber.Ctx) error {
	return ctx.Render("admin/login", fiber.Map{
		"title": "Login",
	}, "layouts/main")
}

func loginPost(ctx *fiber.Ctx) error {
	// get username and password from form
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	log.Println(username, password)

	// check if username and password is correct
	envUsername := os.Getenv("ADMIN_USER")
	envPassword := os.Getenv("ADMIN_PASS")

	if username == envUsername && password == envPassword {
		// set session
		session, err := configs.GetSession(ctx)
		if err != nil {
			return err
		}

		session.Set("isLoggedIn", true)

		return ctx.Redirect("/admin")
	}

	return ctx.Redirect("/admin/login")
}

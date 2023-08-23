package admin

import (
	"time"

	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/daussho/short-it/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
		"error": ctx.Query("error"),
	}, "layouts/main")
}

func loginPost(ctx *fiber.Ctx) error {
	// get username and password from form
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	// check if username and password is correct
	db := configs.ConnectDB()
	var user models.User
	db.Find(&user, "username = ?", username)
	if user.ID == 0 {
		return ctx.Redirect("/admin/login?error=username not found")
	}

	// check if password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return ctx.Redirect("/admin/login?error=incorrect password")
	}

	token := utils.RandStringRunes(64)
	user.Token = token
	db.Save(&user)

	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(7 * 24 * time.Hour),
	})

	return ctx.Redirect("/admin")
}

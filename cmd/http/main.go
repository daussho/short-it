package main

import (
	"errors"
	"log"
	"os"

	"github.com/daussho/short-it/configs"
	apiV1Routes "github.com/daussho/short-it/configs/routes/api/v1"
	webRoutes "github.com/daussho/short-it/configs/routes/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	// init session
	configs.InitSession()

	// init env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// init db
	err = configs.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// init html template engine
	engine := html.New("./views", ".html")

	// init fiber app
	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: errorHandler,
	})
	apiV1Routes.Init(app)
	webRoutes.Init(app)

	// get port from env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	if code == fiber.StatusNotFound {
		return ctx.Redirect("/404")
	}

	log.Fatalf("Error: %v", err)
	return ctx.Status(code).SendString("Internal Server Error")
}

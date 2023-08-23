package main

import (
	"log"
	"os"

	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/configs/routes"
	"github.com/daussho/short-it/handlers"
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
		Views: engine,
	})

	// admin routes
	routes.AdminRoutes(app)

	// 404 page
	app.Get("/404", func(c *fiber.Ctx) error {
		return c.Render("404", fiber.Map{
			"title": "404",
		}, "layouts/main")
	})

	app.Get("/:shortUrl<len(6)>", handlers.UrlRedirect)

	// catch all routes to 404
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.Redirect("/404")
	})

	// get port from env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}

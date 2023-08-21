package main

import (
	"log"
	"os"

	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/configs/routes"
	"github.com/daussho/short-it/models"
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

	// routes
	routes.AdminRoutes(app)

	app.Get("/:shortUrl", func(c *fiber.Ctx) error {

		db := configs.ConnectDB()

		var url models.Url
		db.Where("short_url = ?", c.Params("shortUrl")).First(&url)

		return c.Redirect(url.Url)
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}

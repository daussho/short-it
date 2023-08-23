package url

import (
	"github.com/daussho/short-it/configs"
	"github.com/daussho/short-it/models"
	"github.com/daussho/short-it/utils"
	"github.com/gofiber/fiber/v2"
)

type AddUrlRequest struct {
	Url string `json:"url"`
}

func AddUrl(ctx *fiber.Ctx) error {
	// get url from json body
	var req AddUrlRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if req.Url == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "URL cannot be empty",
		})
	}

	user := ctx.Locals("user").(models.User)

	shortUrl := utils.RandStringRunes(6)

	db := configs.ConnectDB()
	db.Create(&models.Url{
		Url:         req.Url,
		ShortUrl:    shortUrl,
		ViewCount:   0,
		OwnerUserID: user.ID,
	})

	return ctx.JSON(fiber.Map{
		"message": "URL created successfully",
	})
}

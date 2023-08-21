package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store *session.Store
)

func InitSession() {
	store = session.New()
}

func GetSession(ctx *fiber.Ctx) (*session.Session, error) {
	return store.Get(ctx)
}

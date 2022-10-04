package api

import (
	v1api "tinyurl/pkg/api/v1"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/tiny", v1api.Tiny)
}

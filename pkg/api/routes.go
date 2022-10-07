package api

import (
	v1api "tinyurl/pkg/api/v1"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/:tiny_url", Redirect)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/create", v1api.Create)
}

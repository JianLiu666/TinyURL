package main

import (
	"tinyurl/pkg/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api.SetRoutes(app)
	app.Listen(":3000")
}

package api

import "github.com/gofiber/fiber/v2"

func Redirect(c *fiber.Ctx) error {
	return c.SendString(c.Params("tiny_url"))
}

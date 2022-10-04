package v1

import "github.com/gofiber/fiber/v2"

func Tiny(c *fiber.Ctx) error {
	return c.SendString("tiny")
}

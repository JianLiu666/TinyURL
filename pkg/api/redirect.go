package api

import (
	"tinyurl/pkg/storage/mysql"

	"github.com/gofiber/fiber/v2"
)

func Redirect(c *fiber.Ctx) error {
	// 1. select from mysql
	url, err := mysql.GetUrl(c.Params("tiny_url"))
	if err != nil {
		return err
	}

	// TODO: 用戶資料分析
	// 2. redirection
	return c.Redirect(url.Origin, fiber.StatusFound)
}

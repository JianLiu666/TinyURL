package v1

import (
	"errors"
	"tinyurl/pkg/storage/mysql"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Redirect(c *fiber.Ctx) error {
	// 1. select from mysql
	url, err := mysql.GetUrl(c.Params("tiny_url"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusBadRequest)
		} else {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	// TODO: 用戶資料分析
	// 2. redirection
	return c.Redirect(url.Origin, fiber.StatusFound)
}

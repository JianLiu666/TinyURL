package v1

import (
	"errors"
	"tinyurl/pkg/storage/mysql"
	"tinyurl/pkg/storage/redis"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// // TODO: 用戶資料分析
func Redirect(c *fiber.Ctx) error {
	tiny := c.Params("tiny_url")

	// get origin url cache from redis
	origin, status := redis.GetOriginUrl(c.Params("tiny_url"))

	// 短網址命中時的處理流程
	if status == redis.ErrNotFound {
		if origin == "" {
			return c.Status(fiber.StatusBadRequest).SendString("tinyurl not found.")
		}
		return c.Redirect(origin, fiber.StatusFound)
	}

	// 短網址未命中時的處理流程
	url, err := mysql.GetUrl(tiny)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 寫入 redis cache
			if code := redis.SetTinyUrl(&url); code != redis.ErrNotFound {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			return c.Status(fiber.StatusBadRequest).SendString("tinyurl not found.")
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// 寫入 redis cache
	if code := redis.SetTinyUrl(&url); code != redis.ErrNotFound {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Redirect(url.Origin, fiber.StatusFound)
}

package v1

import (
	"errors"
	"time"
	"tinyurl/internal/storage/kvstore"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// TODO: 用戶資料分析
// @Summary      Redirect
// @Description  Redirect user by given shorten url
// @Tags         api/v1
// @Accept       json
// @Produce      text/html
//
// @Param  tiny  path  string true "shorten url"
//
// @Success  302
// @Failure  400  {string}  string
// @Failure  500  {string}  string
//
// @Header 302 {string} Location "Original URL"
//
// @Router  /api/v1/{tiny} [get]
func (h *handler) Redirect(c *fiber.Ctx) error {
	tiny := c.Params("tiny_url")

	// get origin url cache from redis
	origin, status := h.kvStore.GetOriginUrl(c.UserContext(), c.Params("tiny_url"))

	// 短網址命中時的處理流程
	if status == kvstore.ErrNotFound {
		if origin == "" {
			return c.Status(fiber.StatusBadRequest).SendString("tinyurl not found.")
		}
		return c.Redirect(origin, fiber.StatusFound)
	}

	// 短網址未命中時的處理流程
	url, err := h.rdb.GetUrl(c.UserContext(), tiny)
	if err != nil {
		// TODO: remove gorm constant
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 寫入 redis cache
			if code := h.kvStore.SetTinyUrl(c.UserContext(), &url, time.Duration(h.serverConfig.TinyUrlCacheExpired)*time.Second); code != kvstore.ErrNotFound {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			return c.Status(fiber.StatusBadRequest).SendString("tinyurl not found.")
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// 寫入 redis cache
	if code := h.kvStore.SetTinyUrl(c.UserContext(), &url, time.Duration(h.serverConfig.TinyUrlCacheExpired)*time.Second); code != kvstore.ErrNotFound {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Redirect(url.Origin, fiber.StatusFound)
}

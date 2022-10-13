package v1

import (
	"encoding/json"
	"fmt"
	"time"
	"tinyurl/config"
	"tinyurl/pkg/storage/mysql"
	"tinyurl/util"

	"github.com/gofiber/fiber/v2"
	"github.com/spaolacci/murmur3"
)

func Create(c *fiber.Ctx) error {
	reqBody := new(CreateReqBody)
	if err := c.BodyParser(reqBody); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// 1. validation
	if reqBody.Url == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: tiny url 碰撞檢查
	// 2. business logic
	tiny := encode(reqBody.Url)
	data := &mysql.Url{
		Hash:      tiny,
		Origin:    reqBody.Url,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	if err := mysql.CreateUrl(data); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// 3. response
	respBody := &CreateRespBody{
		Origin:    data.Origin,
		Tiny:      fmt.Sprintf("%s%s/%s", config.Env().Server.Domain, config.Env().Server.Port, data.Hash),
		CreateAt:  data.CreatedAt.Unix(),
		ExpiresAt: data.ExpiresAt.Unix(),
	}
	b, err := json.Marshal(respBody)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Response().BodyWriter().Write(b)
	return nil
}

// encode origin url to tiny url
// @param origin url
//
// @return string tiny url
func encode(origin string) string {
	hasher := murmur3.New32()
	hasher.Write([]byte(origin))
	return util.Base10ToBase62(uint64(hasher.Sum32()))
}

type CreateReqBody struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

type CreateRespBody struct {
	Origin    string `json:"origin"`
	Tiny      string `json:"tiny"`
	CreateAt  int64  `json:"created_at"`
	ExpiresAt int64  `json:"expires_at"`
}

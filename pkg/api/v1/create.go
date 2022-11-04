package v1

import (
	"encoding/json"
	"fmt"
	"time"
	"tinyurl/config"
	"tinyurl/pkg/storage"
	"tinyurl/pkg/storage/mysql"
	"tinyurl/pkg/storage/redis"
	"tinyurl/util"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CreateReqBody struct {
	Url   string `json:"url"`   // 原始網址
	Alias string `json:"alias"` // 指定短網址格式
}

type CreateRespBody struct {
	Origin    string `json:"origin"`     // 原始網址
	Tiny      string `json:"tiny"`       // 短網址
	CreateAt  int64  `json:"created_at"` // 短網址產生時間
	ExpiresAt int64  `json:"expires_at"` // 短網址有效時間
}

func Create(c *fiber.Ctx) error {
	// parsing request body
	reqBody := new(CreateReqBody)
	if err := c.BodyParser(reqBody); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// validation
	if reqBody.Url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("field 'url' is empty.")
	}
	if len(reqBody.Alias) > 20 {
		return c.Status(fiber.StatusBadRequest).SendString("field 'alias is invalid.")
	}

	// generate tiny url by custom alias or hash method
	tiny := reqBody.Alias
	if tiny == "" {
		tiny = util.EncodeUrlByHash(reqBody.Url)
	}

	data := &storage.Url{
		Tiny:      tiny,
		Origin:    reqBody.Url,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	// check whether tiny url exists or not from redis
	if code := redis.CheckTinyUrl(data, tiny == reqBody.Alias); code != redis.ErrNotFound {
		if code == redis.ErrInvalidData {
			return c.Status(fiber.StatusBadRequest).SendString("alias dunplicated.")
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// create url record into mysql
	urlAlreadyExists, err := mysql.CreateUrl(c.UserContext(), data, tiny == reqBody.Alias)
	if err != nil {
		logrus.Errorf("Failed to run sql: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// set tiny url cache into redis
	if code := redis.SetTinyUrl(data); code != redis.ErrNotFound {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// initial reponse body
	respBody := &CreateRespBody{
		Origin:    data.Origin,
		Tiny:      fmt.Sprintf("%s%s/api/v1/%s", config.Env().Server.Domain, config.Env().Server.Port, data.Tiny),
		CreateAt:  data.CreatedAt.Unix(),
		ExpiresAt: data.ExpiresAt.Unix(),
	}
	b, err := json.Marshal(respBody)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = c.Response().BodyWriter().Write(b)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if urlAlreadyExists {
		return c.Status(fiber.StatusBadRequest).SendString("alias dunplicated.")
	}
	return c.SendStatus(fiber.StatusOK)
}

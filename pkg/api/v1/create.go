package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"tinyurl/config"
	"tinyurl/pkg/storage/mysql"
	"tinyurl/util"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spaolacci/murmur3"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx) error {
	// 1. parsing request body
	reqBody := new(CreateReqBody)
	if err := c.BodyParser(reqBody); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// 2. validation
	if reqBody.Url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("field 'url' is empty.")
	}
	if len(reqBody.Alias) > 20 {
		return c.Status(fiber.StatusBadRequest).SendString("field 'alias is invalid.")
	}

	// 3. create tiny url by custom alias or hash method
	tiny := reqBody.Alias
	if tiny == "" {
		tiny = encode(reqBody.Url)
	}

	// 4. create or update url metadata into database
	data := &mysql.Url{
		Tiny:      tiny,
		Origin:    reqBody.Url,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	if err := mysql.CreateUrl(data, tiny == reqBody.Alias); err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			return c.Status(fiber.StatusBadRequest).SendString("alias dunplicated.")
		}
		logrus.Errorf("Failed to run sql: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// 5. initial reponse body
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
	Url   string `json:"url"`   // 原始網址
	Alias string `json:"alias"` // 指定短網址格式
}

type CreateRespBody struct {
	Origin    string `json:"origin"`     // 原始網址
	Tiny      string `json:"tiny"`       // 短網址
	CreateAt  int64  `json:"created_at"` // 短網址產生時間
	ExpiresAt int64  `json:"expires_at"` // 短網址有效時間
}

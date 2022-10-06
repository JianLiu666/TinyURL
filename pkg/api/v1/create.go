package v1

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"tinyurl/pkg/storage/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/spaolacci/murmur3"
)

func Create(c *fiber.Ctx) error {
	reqBody := new(createReqBody)
	if err := c.BodyParser(reqBody); err != nil {
		return err
	}

	if reqBody.Url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	tiny := encode(reqBody.Url)
	data := &mysql.Url{
		Hash:      tiny,
		Origin:    reqBody.Url,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	if err := mysql.CreateUrl(data); err != nil {
		fmt.Println(err)
	}

	respBody := &createRespBody{
		Origin:    reqBody.Url,
		Tiny:      tiny,
		CreateAt:  data.CreatedAt.Unix(),
		ExpiresAt: data.ExpiresAt.Unix(),
	}
	b, err := json.Marshal(respBody)
	if err != nil {
		return err
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
	return strconv.FormatUint(uint64(hasher.Sum32()), 16)
}

type createReqBody struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

type createRespBody struct {
	Origin    string `json:"origin"`
	Tiny      string `json:"tiny"`
	CreateAt  int64  `json:"created_at"`
	ExpiresAt int64  `json:"expires_at"`
}

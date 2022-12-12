package v1

import (
	"encoding/json"
	"fmt"
	"time"
	"tinyurl/internal/storage"
	"tinyurl/internal/storage/kvstore"
	"tinyurl/tools"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CreateReqBody struct {
	Url   string `json:"url"   example:"https://github.com/JianLiu666"` // 原始網址
	Alias string `json:"alias" example:"jian"`                          // 指定短網址格式
}

type CreateRespBody struct {
	Origin    string `json:"origin"     example:"https://github.com/JianLiu666"`     // 原始網址
	Tiny      string `json:"tiny"       example:"http://localhost:6600/api/v1/jian"` // 短網址
	CreateAt  int64  `json:"created_at" example:"1669229019"`                        // 短網址產生時間
	ExpiresAt int64  `json:"expires_at" example:"1670936510"`                        // 短網址有效時間
}

// @Summary      Create a shorten url
// @Description  Generate shortenl url by user's original url
// @Tags         api/v1
// @Accept       json
// @Produce      json
//
// @Param  url    body  CreateReqBody true "original url"
// @Param  alias  body  CreateReqBody true "alias name"
//
// @Success  200  {object}  CreateRespBody
// @Failure  400  {string}  string
// @Failure  500  {string}  string
//
// @Router  /api/v1/create [post]
func (h *handler) Create(c *fiber.Ctx) error {
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
		tiny = tools.EncodeUrlByHash(reqBody.Url)
	}

	data := &storage.Url{
		Tiny:      tiny,
		Origin:    reqBody.Url,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	// check whether tiny url exists or not from redis
	if code := h.kvStore.CheckTinyUrl(c.UserContext(), data, tiny == reqBody.Alias, h.serverConfig.TinyUrlRetry); code != kvstore.ErrNotFound {
		if code == kvstore.ErrInvalidData {
			return c.Status(fiber.StatusBadRequest).SendString("alias dunplicated.")
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// create url record into mysql
	urlAlreadyExists, err := h.rdb.CreateUrl(c.UserContext(), data, tiny == reqBody.Alias)
	if err != nil {
		logrus.Errorf("Failed to run sql: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// set tiny url cache into redis
	if code := h.kvStore.SetTinyUrl(c.UserContext(), data, time.Duration(h.serverConfig.TinyUrlCacheExpired)*time.Second); code != kvstore.ErrNotFound {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// initial reponse body
	respBody := &CreateRespBody{
		Origin:    data.Origin,
		Tiny:      fmt.Sprintf("%s%s/api/v1/%s", h.serverConfig.Domain, h.serverConfig.Port, data.Tiny),
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

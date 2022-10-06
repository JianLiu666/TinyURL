package v1

import (
	"fmt"
	"strconv"
	"time"
	"tinyurl/pkg/storage/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/spaolacci/murmur3"
)

func Create(c *fiber.Ctx) error {
	origin := c.Params("origin_url")
	tiny := encode(origin)

	data := &mysql.Url{
		Hash:      tiny,
		Origin:    origin,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}
	if err := mysql.CreateUrl(data); err != nil {
		fmt.Println(err)
	}

	return c.SendString(tiny)
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

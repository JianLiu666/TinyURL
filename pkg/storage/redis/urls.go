package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"tinyurl/config"
	"tinyurl/pkg/storage"
	"tinyurl/util"

	"github.com/go-redis/redis/v9"
)

func SetTinyUrl(data *storage.Url) int {

	if err := instance.SetEx(context.TODO(), "tiny:"+data.Tiny, data.Origin, time.Duration(config.Env().Server.TinyUrlCacheExpired)*time.Second).Err(); err != nil {
		fmt.Println(err)
		return ErrUnexpected
	}

	return ErrNotFound
}

func CheckTinyUrl(data *storage.Url, isCustomAlias bool) int {

	for i := 0; i < config.Env().Server.TinyUrlRetry; i++ {
		// 檢查 redis 中是否存在相同的短網址
		res, err := instance.Get(context.TODO(), "tiny:"+data.Tiny).Result()

		// 可以使用的短網址
		if err == redis.Nil || res == "" {
			return ErrNotFound
		}

		// 例外錯誤
		if err != nil {
			fmt.Println(err)
			return ErrUnexpected
		}

		// 短網址發生碰撞, 不處理的情境
		//  1. 相同的自定義短網址代碼
		//  2. 相同的原始網址
		if isCustomAlias || res == data.Origin {
			return ErrInvalidData
		}

		// 原始網址加上後綴改變雜湊結果
		data.Tiny = util.EncodeUrlByHash(data.Origin + strconv.Itoa(i))
	}

	return ErrUnexpected
}

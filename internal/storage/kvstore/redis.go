package kvstore

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"tinyurl/internal/storage"
	"tinyurl/tools"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type redisClient struct {
	conn *redis.Client
}

func NewRedisClient(ctx context.Context, addr, password string, db int) KvStore {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		// PoolSize: 10, // TODO: pool size 也需要控制
	})

	ct, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if _, err := conn.Ping(ct).Result(); err != nil {
		logrus.Panicf("failed to ping reids server: %v", err)
	}

	return &redisClient{
		conn: conn,
	}
}

func (c *redisClient) SetTinyUrl(ctx context.Context, data *storage.Url, expiration time.Duration) int {
	if err := c.conn.SetEX(ctx, getTinyKey(data.Tiny), data.Origin, expiration).Err(); err != nil {
		fmt.Println(err)
		return ErrUnexpected
	}

	return ErrNotFound
}

func (c *redisClient) GetOriginUrl(ctx context.Context, tiny string) (string, int) {
	res, err := c.conn.Get(ctx, getTinyKey(tiny)).Result()
	// 短網址未命中
	if err == redis.Nil {
		return "", ErrKeyNotFound
	}

	// 例外錯誤
	if err != nil {
		fmt.Println(err)
		return "", ErrUnexpected
	}

	return res, ErrNotFound
}

func (c *redisClient) CheckTinyUrl(ctx context.Context, data *storage.Url, isCustomAlias bool, retryCount int) int {
	for i := 0; i < retryCount; i++ {
		// 檢查 redis 中是否存在相同的短網址
		res, err := c.conn.Get(ctx, getTinyKey(data.Tiny)).Result()

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
		data.Tiny = tools.EncodeUrlByHash(data.Origin + strconv.Itoa(i))
	}

	return ErrUnexpected
}

func (c *redisClient) Shutdown(ctx context.Context) {
	if err := c.conn.Close(); err != nil {
		logrus.Errorf("failed to close redis client: %v", err)
	}
}

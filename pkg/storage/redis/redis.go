package redis

import (
	"fmt"
	"sync"
	"tinyurl/internal/config"

	"github.com/go-redis/redis/v8"
)

var once sync.Once
var instance *redis.Client

func GetInstance() *redis.Client {
	return instance
}

func Init() {
	once.Do(func() {
		instance = redis.NewClient(&redis.Options{
			Addr:     config.Env().Redis.Address,
			Password: config.Env().Redis.Password,
			DB:       config.Env().Redis.DB,
		})

		fmt.Println("connect to redis successful.")
	})
}

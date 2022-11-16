package redis

import (
	"sync"
	"tinyurl/internal/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var once sync.Once
var instance *redis.Client

func GetInstance() *redis.Client {
	return instance
}

func Init() {
	once.Do(func() {
		defer logrus.Infof("connect to redis successful.")

		instance = redis.NewClient(&redis.Options{
			Addr:     config.Env().Redis.Address,
			Password: config.Env().Redis.Password,
			DB:       config.Env().Redis.DB,
		})
	})
}

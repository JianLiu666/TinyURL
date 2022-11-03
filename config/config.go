package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var env *environment
var once sync.Once
var initialized bool

func Env() *environment {
	if initialized {
		return env
	}
	return nil
}

func LoadFromViper() {
	once.Do(func() {
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		fmt.Printf("read config file from: %v\n", viper.ConfigFileUsed())

		env = &environment{}
		err = viper.Unmarshal(env)
		if err != nil {
			panic(fmt.Errorf("fatal unmarshal config file: %w", err))
		}

		initialized = true
	})
}

type environment struct {
	Server server `yaml:"server"`
	MySQL  mysql  `yaml:"mysql"`
	Redis  redis  `yaml:"redis"`
	Jaeger jaeger `yaml:"jaeger"`
}

type server struct {
	Name                string `mapstructure:"name" yaml:"name"`
	Domain              string `mapstructure:"domain" yaml:"domain"`
	Port                string `mapstructure:"port" yaml:"port"`
	TinyUrlCacheExpired int    `mapstructure:"tinyurl_cache_expired" yaml:"tinyurl_cache_expired"`
	TinyUrlRetry        int    `mapstructure:"tinyurl_retry" yaml:"tinyurl_retry"`
}

type mysql struct {
	Address         string `mapstructure:"address" yaml:"address"`
	UserName        string `mapstructure:"username" yaml:"username"`
	Password        string `mapstructure:"password" yaml:"password"`
	DBName          string `mapstructure:"dbname" yaml:"dbname"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" yaml:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime"`
}

type redis struct {
	Address  string `mapstructure:"address" yaml:"address"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
}

type jaeger struct {
	Address string `mapstructure:"address" yaml:"address"`
}

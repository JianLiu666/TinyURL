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
}

type server struct {
	Domain string `yaml:"domain"`
	Port   string `yaml:"port"`
}

type mysql struct {
	Address         string `yaml:"address"`
	UserName        string `yaml:"username"`
	Password        string `yaml:"password"`
	DBName          string `yaml:"dbname"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

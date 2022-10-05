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

		env = &environment{}
		viper.Unmarshal(env)
		if err != nil {
			panic(fmt.Errorf("fatal unmarshal config file: %w", err))
		}

		initialized = true
	})
}

type environment struct {
	MySQL mysql `yaml:"mysql"`
}

type mysql struct {
	Address  string `yaml:"address"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

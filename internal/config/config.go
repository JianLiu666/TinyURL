package config

import (
	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	Server ServerOpts `mapstructure:"server" yaml:"server"`
	MySQL  MysqlOpts  `mapstructure:"mysql" yaml:"mysql"`
	Redis  RedisOpts  `mapstructure:"redis" yaml:"redis"`
	Jaeger JaegerOpts `mapstructure:"jaeger" yaml:"jaeger"`
}

func NewFromViper() *Config {
	err := viper.ReadInConfig()
	if err != nil {
		return NewFromDefault()
	}

	cfg = &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return NewFromDefault()
	}

	return cfg
}

func NewFromDefault() *Config {
	server := ServerOpts{
		Name:                "server",
		Domain:              "localhost",
		Port:                "6600",
		TinyUrlCacheExpired: 3600,
		TinyUrlRetry:        10,
	}

	mysql := MysqlOpts{
		Address:         "mysql:3306",
		UserName:        "root",
		Password:        "0",
		DBName:          "tinyurl",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 60,
	}
	redis := RedisOpts{
		Address:  "redis:6379",
		Password: "",
		DB:       0,
	}

	jaeger := JaegerOpts{
		RPCMetrics: true,
		Sampler: jaegerSampler{
			Type:  "const",
			Param: 1,
		},
		Reporter: jaegerReporter{
			LogSpans:            true,
			BufferFlushInterval: 1,
			LocalAgentHostPort:  "jaeger:6831",
		},
		Headers: jaegerHeaders{
			TraceBaggageHeaderPrefix: "ctx-",
			TraceContextHeaderName:   "headerName",
		},
	}

	cfg := &Config{
		Server: server,
		MySQL:  mysql,
		Redis:  redis,
		Jaeger: jaeger,
	}

	return cfg
}

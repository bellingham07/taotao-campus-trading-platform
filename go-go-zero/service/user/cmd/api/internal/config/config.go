package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	MysqlConfig struct {
		Dsn string
	}

	RedisConfig struct {
		Addr     string `yaml:"url"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	}
}

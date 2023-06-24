package config

import (
	"github.com/zeromicro/go-zero/rest"
	"go-go-zero/common/utils"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		Dsn string
	}

	Redis utils.Redis

	Mongo utils.Mongo

	RabbitMQ utils.RabbitMQConf
}

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/common/utils"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf

	Mysql struct {
		Dsn string
	}

	Redis struct {
		Addr     string
		Password string
		Db       int
	}

	Mongo struct {
		Url string
	}

	RabbitMQ utils.RabbitMQConf
}

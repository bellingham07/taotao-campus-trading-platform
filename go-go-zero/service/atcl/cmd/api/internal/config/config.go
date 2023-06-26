package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/common/utils"
)

type Config struct {
	rest.RestConf
	Mysql    utils.Mysql
	Redis    utils.Redis
	Mongo    utils.Mongo
	RabbitMQ utils.RabbitMQConf
	UserRpc  zrpc.RpcClientConf
}

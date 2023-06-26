package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"go-go-zero/common/utils"
)

type Config struct {
	zrpc.RpcServerConf
	Consul consul.Conf
	Mysql  utils.Mysql
}

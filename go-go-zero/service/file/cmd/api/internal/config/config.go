package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRpc  zrpc.RpcClientConf
	CmdtyRpc zrpc.RpcClientConf
	AtclRpc  zrpc.RpcClientConf

	Mysql struct {
		Dsn string
	}

	Oss struct {
		AccessKeyId     string
		AccessKeySecret string
		EndPoint        string
		BucketName      string
	}
}

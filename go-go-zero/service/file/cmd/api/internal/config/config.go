package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/common/utils"
)

type Config struct {
	rest.RestConf
	UserRpc  zrpc.RpcClientConf
	CmdtyRpc zrpc.RpcClientConf
	AtclRpc  zrpc.RpcClientConf

	Mysql utils.Mysql
	Idgen struct {
		WorkerId uint16
	}

	Oss struct {
		AccessKeyId     string
		AccessKeySecret string
		EndPoint        string
		BucketName      string
	}
}

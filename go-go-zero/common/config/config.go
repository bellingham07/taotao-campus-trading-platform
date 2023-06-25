package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type (
	UserApi struct {
		TaoTaoApi
	}

	CmdtyApi struct {
		TaoTaoApi
	}

	ChatApi struct {
		TaoTaoApi
	}

	TradeApi struct {
		TaoTaoApi
		UserRpc  zrpc.RpcClientConf
		CmdtyRpc zrpc.RpcClientConf
	}

	FileApi struct {
		TaoTaoApi
		UserRpc  zrpc.RpcClientConf
		CmdtyRpc zrpc.RpcClientConf
		AtclRpc  zrpc.RpcClientConf

		Oss struct {
			AccessKeyId     string
			AccessKeySecret string
			EndPoint        string
			BucketName      string
		}
	}

	AtclApi struct {
		TaoTaoApi
		UserRpc zrpc.RpcClientConf
	}
)

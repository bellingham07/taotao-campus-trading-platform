package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	sysConfig "go-go-zero/common/config"

	"go-go-zero/service/atcl/cmd/rpc/internal/config"
	"go-go-zero/service/atcl/cmd/rpc/internal/server"
	"go-go-zero/service/atcl/cmd/rpc/internal/svc"
	"go-go-zero/service/atcl/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/atcl/cmd/rpc/etc/atcl.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	c.Consul = *sysConfig.LoadConsulConf("service/atcl/cmd/rpc/etc/atcl.yaml")
	c.TaoTaoRpc = *sysConfig.LoadTaoTaoRpc(&c.Consul)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.TaoTaoRpc.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterAtclServiceServer(grpcServer, server.NewAtclServiceServer(ctx))

		if c.TaoTaoRpc.Mode == service.DevMode || c.TaoTaoRpc.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	_ = consul.RegisterService(c.TaoTaoRpc.ListenOn, c.TaoTaoRpc.Consul)

	fmt.Printf("Starting rpc server at %s...\n", c.TaoTaoRpc.ListenOn)
	s.Start()
}

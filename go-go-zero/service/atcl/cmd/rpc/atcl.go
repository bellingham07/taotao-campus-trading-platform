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
	cc := sysConfig.LoadConsulConf("service/atcl/cmd/rpc/etc/atcl.yaml")
	sysConfig.LoadTaoTaoRpc(cc, &c)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterAtclServiceServer(grpcServer, server.NewAtclServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	_ = consul.RegisterService(c.ListenOn, c.Consul)

	fmt.Printf("Starting rpc gateway at %s...\n", c.ListenOn)
	s.Start()
}

package main

import (
	"flag"
	"fmt"

	"go-go-zero/service/atcl/cmd/rpc/internal/config"
	"go-go-zero/service/atcl/cmd/rpc/internal/server"
	"go-go-zero/service/atcl/cmd/rpc/internal/svc"
	"go-go-zero/service/atcl/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/atcl/cmd/rpc/etc/atcl.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterAtclServiceServer(grpcServer, server.NewAtclServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

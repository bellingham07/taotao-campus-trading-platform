package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	sysConfig "go-go-zero/common/config"
	"go-go-zero/service/user/cmd/api/internal/config"
	"go-go-zero/service/user/cmd/api/internal/handler"
	"go-go-zero/service/user/cmd/api/internal/svc"
)

var configFile = flag.String("f", "service/user/cmd/api/etc/trade-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	cc := sysConfig.LoadConsulConf("service/user/cmd/api/etc/trade-api.yaml")
	sysConfig.LoadTaoTaoApi(cc, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

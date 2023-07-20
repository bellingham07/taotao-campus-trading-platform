package main

import (
	"flag"
	"fmt"
	sysConfig "go-go-zero/common/config"

	"go-go-zero/service/atcl/cmd/api/internal/config"
	"go-go-zero/service/atcl/cmd/api/internal/handler"
	"go-go-zero/service/atcl/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "service/atcl/cmd/api/etc/atcl-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	cc := sysConfig.LoadConsulConf("service/atcl/cmd/api/etc/atcl-api.yaml")
	sysConfig.LoadTaoTaoRpc(cc, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

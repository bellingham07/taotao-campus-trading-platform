package main

import (
	"flag"
	"fmt"
	sysConfig "go-go-zero/common/config"

	"go-go-zero/service/trade/cmd/api/internal/config"
	"go-go-zero/service/trade/cmd/api/internal/handler"
	"go-go-zero/service/trade/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "service/trade/cmd/api/etc/trade-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	c.Consul = *sysConfig.LoadConsulConf("service/trade/cmd/api/etc/trade-api.yaml")
	c.TradeApi = *sysConfig.LoadTaoTaoApi(&c.Consul, &c.TradeApi).(*sysConfig.TradeApi)

	server := rest.MustNewServer(c.TradeApi.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.TradeApi.Host, c.TradeApi.Port)
	server.Start()
}

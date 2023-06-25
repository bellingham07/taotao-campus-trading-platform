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

var configFile = flag.String("f", "service/user/cmd/api/etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	c.Consul = *sysConfig.LoadConsulConf("service/user/cmd/api/etc/user-api.yaml")
	c.UserApi = *sysConfig.LoadTaoTaoApi(&c.Consul, &c.UserApi).(*sysConfig.UserApi)

	server := rest.MustNewServer(c.UserApi.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.UserApi.Host, c.UserApi.Port)
	server.Start()
}

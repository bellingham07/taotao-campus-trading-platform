package main

import (
	"flag"
	"fmt"
	sysConfig "go-go-zero/common/config"

	"go-go-zero/service/file/cmd/api/internal/config"
	"go-go-zero/service/file/cmd/api/internal/handler"
	"go-go-zero/service/file/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "service/file/cmd/api/etc/file-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	c.Consul = *sysConfig.LoadConsulConf("service/file/cmd/api/etc/file-api.yaml")
	c.FileApi = *sysConfig.LoadTaoTaoApi(&c.Consul, &c.FileApi).(*sysConfig.FileApi)

	server := rest.MustNewServer(c.FileApi.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.FileApi.Host, c.FileApi.Port)
	server.Start()
}

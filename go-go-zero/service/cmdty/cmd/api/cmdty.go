package main

import (
	"flag"
	"fmt"
	sysConfig "go-go-zero/common/config"
	"go-go-zero/common/middleware"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/cron"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/mq"

	"go-go-zero/service/cmdty/cmd/api/internal/config"
	"go-go-zero/service/cmdty/cmd/api/internal/handler"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "service/cmdty/cmd/api/etc/cmdty-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	cc := sysConfig.LoadConsulConf("service/cmdty/cmd/api/etc/cmdty-api.yaml")
	sysConfig.LoadTaoTaoApi(cc, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(middleware.HandleCors) // 全局中间件

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	mq.InitRabbitMQ(ctx)
	go cron.InitCronJob(ctx)

	fmt.Printf("Starting gateway at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

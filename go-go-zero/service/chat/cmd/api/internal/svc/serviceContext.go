package svc

import (
	"github.com/gorilla/websocket"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/chat/cmd/api/internal/config"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm *xorm.Engine

	Upgrader websocket.Upgrader
}

func NewServiceContext(c config.Config) *ServiceContext {
	options := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(options)

	engine, err := xorm.NewEngine("mysql", c.Mysql.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())

	}

	return &ServiceContext{
		Config: c,
		Xorm:   engine,
	}
}

package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"go-go-zero/service/cmdty/cmd/rpc/internal/config"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm      *xorm.Engine
	CmdtyInfo *xorm.Session
	//FileCmdty  *xorm.Session
	//FileAvatar *xorm.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine, err := xorm.NewEngine("mysql", c.Mysql.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())

	}

	return &ServiceContext{
		Config:    c,
		Xorm:      engine,
		CmdtyInfo: engine.Table("cmdty_info"),
	}
}

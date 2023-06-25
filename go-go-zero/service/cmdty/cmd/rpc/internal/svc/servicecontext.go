package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/rpc/internal/config"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm         *xorm.Engine
	CmdtyInfo    *xorm.Session
	CmdtyCollect *xorm.Session
	CmdtyDone    *xorm.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.TaoTaoRpc.Mysql)

	return &ServiceContext{
		Config:       c,
		Xorm:         engine,
		CmdtyInfo:    engine.Table("cmdty_info"),
		CmdtyCollect: engine.Table("cmdty_collect"),
		CmdtyDone:    engine.Table("cmdty_done"),
	}
}

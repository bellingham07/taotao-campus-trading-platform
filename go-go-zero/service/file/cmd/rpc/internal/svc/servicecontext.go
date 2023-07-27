package svc

import (
	"go-go-zero/common/utils"
	"go-go-zero/service/file/cmd/rpc/internal/config"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm       *xorm.Engine
	FileAtcl   *xorm.Session
	FileCmdty  *xorm.Session
	FileAvatar *xorm.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	return &ServiceContext{
		Config:     c,
		Xorm:       engine,
		FileAtcl:   engine.Table("file_atcl"),
		FileCmdty:  engine.Table("file_cmdty"),
		FileAvatar: engine.Table("file_avatar"),
	}
}

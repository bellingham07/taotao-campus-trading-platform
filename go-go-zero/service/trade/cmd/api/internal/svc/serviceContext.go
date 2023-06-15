package svc

import (
	"go-go-zero/service/cmdty/model"
	"go-go-zero/service/trade/cmd/api/internal/config"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	// mysql
	CmdtyInfo    model.CmdtyInfoModel
	CmdtyCollect model.CmdtyCollectModel
	CmdtyTag     model.CmdtyTagModel
	Xorm         *xorm.Engine
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

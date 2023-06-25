package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"go-go-zero/common/utils"
	"go-go-zero/service/user/cmd/rpc/internal/config"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	// model
	Xorm         *xorm.Engine
	UserInfo     *xorm.Session
	UserFollow   *xorm.Session
	UserCall     *xorm.Session
	UserLocation *xorm.Session
	UserOpt      *xorm.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.TaoTaoRpc.Mysql)

	return &ServiceContext{
		Config:       c,
		Xorm:         engine,
		UserInfo:     engine.Table("user_info"),
		UserFollow:   engine.Table("user_follow"),
		UserCall:     engine.Table("user_call"),
		UserLocation: engine.Table("user_location"),
		UserOpt:      engine.Table("user_opt"),
	}
}

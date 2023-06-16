package svc

import (
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
	engine, err := xorm.NewEngine("mysql", c.Mysql.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())
	}

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

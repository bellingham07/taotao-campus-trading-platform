package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-go-zero/service/user/cmd/rpc/internal/config"
	"go-go-zero/service/user/model"
)

type ServiceContext struct {
	Config config.Config

	// model
	UserInfo   model.UserInfoModel
	UserFollow model.UserFollowModel
	UserCall   model.UserCallModel
	UserDorm   model.UserLocationModel
	UserOpt    model.UserOptModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("mysql", c.Mysql.Dsn)
	return &ServiceContext{
		Config:     c,
		UserInfo:   model.NewUserInfoModel(conn),
		UserFollow: model.NewUserFollowModel(conn),
		UserCall:   model.NewUserCallModel(conn),
		UserDorm:   model.NewUserLocationModel(conn),
		UserOpt:    model.NewUserOptModel(conn),
	}
}

package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-go-zero/service/user/cmd/api/internal/config"
	"go-go-zero/service/user/model"
)

type ServiceContext struct {
	// local
	Config config.Config

	// model
	UserInfo   model.UserInfoModel
	UserFollow model.UserFollowModel
	UserCall   model.UserCallModel
	UserDorm   model.UserDormModel
	UserOpt    model.UserOptModel

	// redis
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("mysql", c.Mysql.Dsn)
	return &ServiceContext{
		Config:     c,
		UserInfo:   model.NewUserInfoModel(conn),
		UserFollow: model.NewUserFollowModel(conn),
		UserCall:   model.NewUserCallModel(conn),
		UserDorm:   model.NewUserDormModel(conn),
		UserOpt:    model.NewUserOptModel(conn),
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.Db,
		}),
	}
}

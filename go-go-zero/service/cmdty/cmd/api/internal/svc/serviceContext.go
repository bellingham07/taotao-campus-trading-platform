package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-go-zero/service/cmdty/cmd/api/internal/config"
	"go-go-zero/service/cmdty/model"
)

type ServiceContext struct {
	Config config.Config

	// mysql
	CmdtyInfo    model.CmdtyInfoModel
	CmdtyCollect model.CmdtyCollectModel
	CmdtyTag     model.CmdtyTagModel

	// redis
	RedisClient *redis.Client

	// rabbitMQ
	MQUrl string
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.Dsn)
	return &ServiceContext{
		Config:       c,
		CmdtyInfo:    model.NewCmdtyInfoModel(conn),
		CmdtyCollect: model.NewCmdtyCollectModel(conn),
		CmdtyTag:     model.NewCmdtyTagModel(conn),
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.Db,
		}),
		MQUrl: c.RabbitMQ.MQUrl,
	}
}

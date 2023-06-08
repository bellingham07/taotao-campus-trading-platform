package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
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
	MQUrl        string
	RabbitMQConn *amqp.Connection
}

func NewServiceContext(c config.Config) *ServiceContext {
	c1 := sqlx.NewMysql(c.Mysql.Dsn)
	c2, err := amqp.Dial(c.RabbitMQ.MQUrl)
	if err != nil {
		panic("连接不到rabbitmq")
	}
	return &ServiceContext{
		Config:       c,
		CmdtyInfo:    model.NewCmdtyInfoModel(c1),
		CmdtyCollect: model.NewCmdtyCollectModel(c1),
		CmdtyTag:     model.NewCmdtyTagModel(c1),
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.Db,
		}),
		MQUrl:        c.RabbitMQ.MQUrl,
		RabbitMQConn: c2,
	}
}

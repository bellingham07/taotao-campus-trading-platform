package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-go-zero/common/utils"
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
	RmqCore *utils.RabbitmqCore
}

func NewServiceContext(c config.Config) *ServiceContext {
	c1 := sqlx.NewMysql(c.Mysql.Dsn)
	c2, err := amqp.Dial(c.RabbitMQ.RmqUrl)
	if err != nil {
		panic("[RABBITMQ ERROR] NewServiceContext 连接不到rabbitmq")
	}
	channel, err := c2.Channel()
	if err != nil {
		panic("[RABBITMQ ERROR] NewServiceContext 获取rabbitmq通道失败")
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
		RmqCore: &utils.RabbitmqCore{
			Conn:    c2,
			Channel: channel,
		},
	}
}

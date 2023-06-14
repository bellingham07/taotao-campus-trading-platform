package svc

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/config"
	"go-go-zero/service/cmdty/model"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	// mysql
	CmdtyInfo    model.CmdtyInfoModel
	CmdtyCollect model.CmdtyCollectModel
	CmdtyTag     model.CmdtyTagModel
	CmdtyCmt     model.CmdtyCmtModel
	Xorm         *xorm.Engine

	// redis
	Redis *redis.Client

	// rabbitMQ
	RmqCore *utils.RabbitmqCore

	Json jsoniter.API
}

func NewServiceContext(c config.Config) *ServiceContext {
	c1 := sqlx.NewMysql(c.Mysql.Dsn)

	engine, err := xorm.NewEngine("mysql", c.Mysql.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())
	}

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
		CmdtyCmt:     model.NewCmdtyCmtModel(c1),
		Xorm:         engine,
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.Db,
		}),
		RmqCore: &utils.RabbitmqCore{
			Conn:    c2,
			Channel: channel,
		},
		Json: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

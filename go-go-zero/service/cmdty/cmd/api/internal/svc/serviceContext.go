package svc

import (
	_ "github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/config"
	"go-go-zero/service/cmdty/cmd/api/internal/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	// mysql
	Xorm         *xorm.Engine
	CmdtyInfo    *xorm.Session
	CmdtyDone    *xorm.Session
	CmdtyCollect *xorm.Session
	CmdtyTag     *xorm.Session
	CmdtyCmt     *mongo.Collection

	// redis
	Redis *redis.Client

	// rabbitMQ
	RmqCore *utils.RabbitmqCore

	Json jsoniter.API

	JwtAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	mc := utils.InitMongo(c.Mongo)

	rc, channel := utils.InitRabbitMQ(c.RabbitMQ)

	return &ServiceContext{
		Config:       c,
		Xorm:         engine,
		CmdtyInfo:    engine.Table("cmdty_info"),
		CmdtyDone:    engine.Table("cmdty_done"),
		CmdtyCollect: engine.Table("cmdty_info"),
		CmdtyTag:     engine.Table("cmdty_info"),
		CmdtyCmt:     mc.Database("taotao_trading_cmdty").Collection("cmdty_cmt"),
		Json:         jsoniter.ConfigCompatibleWithStandardLibrary,
		JwtAuth:      middleware.NewJwtAuthMiddleware().Handle,
		Redis:        utils.InitRedis(c.Redis),
		RmqCore: &utils.RabbitmqCore{
			Conn:    rc,
			Channel: channel,
		},
	}
}

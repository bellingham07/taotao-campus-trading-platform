package svc

import (
	_ "github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/common/utils"
	"go-go-zero/service/atcl/cmd/api/internal/config"
	"go-go-zero/service/atcl/cmd/api/internal/middleware"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	// mysql
	Xorm         *xorm.Engine
	AtclContent  *xorm.Session
	AtclBulletin *xorm.Session
	AtclCollect  *xorm.Session
	AtclCmt      *mongo.Collection

	// redis
	Redis *redis.Client

	// rabbitMQ
	RmqCore *utils.RabbitmqCore

	Json jsoniter.API

	UserRpc userservice.UserService

	JwtAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	mc := utils.InitMongo(c.Mongo)

	rc, channel := utils.InitRabbitMQ(c.RabbitMQ)

	return &ServiceContext{
		Config:       c,
		Xorm:         engine,
		AtclContent:  engine.Table("cmdty_info"),
		AtclBulletin: engine.Table("cmdty_info"),
		AtclCollect:  engine.Table("cmdty_info"),
		AtclCmt:      mc.Database("taotao_trading_chat").Collection("atcl_cmt"),
		Json:         jsoniter.ConfigCompatibleWithStandardLibrary,
		UserRpc:      userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		JwtAuth:      middleware.NewJwtAuthMiddleware().Handle,
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.Db,
		}),
		RmqCore: &utils.RabbitmqCore{
			Conn:    rc,
			Channel: channel,
		},
	}
}

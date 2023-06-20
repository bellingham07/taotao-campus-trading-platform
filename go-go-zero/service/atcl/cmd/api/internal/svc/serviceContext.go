package svc

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/common/utils"
	"go-go-zero/service/atcl/cmd/api/internal/config"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	clientOptions := options.Client().ApplyURI(c.Mongo.Url) // 设置客户端连接配置
	client, err := mongo.NewClient(clientOptions)           // 创建客户端
	if err != nil {
		panic("[MONGO ERROR] NewServiceContext mongodb 连接失败" + err.Error())
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic("[MONGO ERROR] NewServiceContext mongodb 连接失败" + err.Error())
	}
	acCollection := client.Database("taotao_trading_chat").Collection("atcl_cmt")

	c2, err := amqp.Dial(c.RabbitMQ.RmqUrl)
	if err != nil {
		panic("[RABBITMQ ERROR] NewServiceContext 连接不到rabbitmq" + err.Error())
	}
	channel, err := c2.Channel()
	if err != nil {
		panic("[RABBITMQ ERROR] NewServiceContext 获取rabbitmq通道失败")
	}

	return &ServiceContext{
		Config:       c,
		Xorm:         engine,
		AtclContent:  engine.Table("cmdty_info"),
		AtclBulletin: engine.Table("cmdty_info"),
		AtclCollect:  engine.Table("cmdty_info"),
		AtclCmt:      acCollection,
		Json:         jsoniter.ConfigCompatibleWithStandardLibrary,
		UserRpc:      userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		Redis: redis.NewClient(&redis.Options{
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

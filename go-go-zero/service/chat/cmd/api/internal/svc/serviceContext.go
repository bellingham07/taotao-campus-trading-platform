package svc

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/chat/cmd/api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm *xorm.Engine

	ChatMessage *mongo.Collection

	Upgrader websocket.Upgrader
}

func NewServiceContext(c config.Config) *ServiceContext {
	idgenops := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(idgenops)

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
	cmCollection := client.Database("taotao_trading_chat").Collection("chat_message")

	return &ServiceContext{
		Config:      c,
		Xorm:        engine,
		ChatMessage: cmCollection,
	}
}

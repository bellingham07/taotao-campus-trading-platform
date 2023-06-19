package svc

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/rest"
	"go-go-zero/service/chat/cmd/api/internal/config"
	"go-go-zero/service/chat/cmd/api/internal/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm        *xorm.Engine
	ChatRoom    *xorm.Session
	ChatMessage *mongo.Collection

	Upgrader websocket.Upgrader
	Conn     *Conn

	JwtAuth rest.Middleware
}

type Conn struct {
	ConnPool map[string]*websocket.Conn
	Lock     sync.RWMutex
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

	jwtAuth := middleware.NewJwtAuthMiddleware()
	return &ServiceContext{
		Config:      c,
		Xorm:        engine,
		ChatRoom:    engine.Table("chat_room"),
		ChatMessage: cmCollection,
		JwtAuth:     jwtAuth.Handle,
		Conn: &Conn{
			ConnPool: make(map[string]*websocket.Conn),
		},
	}
}

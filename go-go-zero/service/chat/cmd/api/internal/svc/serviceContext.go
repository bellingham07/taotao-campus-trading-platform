package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/rest"
	"go-go-zero/common/utils"
	"go-go-zero/service/chat/cmd/api/internal/config"
	"go-go-zero/service/chat/cmd/api/internal/middleware"
	"go.mongodb.org/mongo-driver/mongo"
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

	Json jsoniter.API

	JwtAuth rest.Middleware
}

type Conn struct {
	ConnPool map[string]*websocket.Conn
	Lock     sync.RWMutex
}

func NewServiceContext(c config.Config) *ServiceContext {
	idgenops := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(idgenops)

	engine := utils.InitXorm("mysql", c.Mysql)

	mc := utils.InitMongo(c.Mongo)

	return &ServiceContext{
		Config:      c,
		Xorm:        engine,
		ChatRoom:    engine.Table("chat_room"),
		ChatMessage: mc.Database("taotao_trading_chat").Collection("chat_message"),
		JwtAuth:     middleware.NewJwtAuthMiddleware().Handle,
		Json:        jsoniter.ConfigCompatibleWithStandardLibrary,
		Conn: &Conn{
			ConnPool: make(map[string]*websocket.Conn),
		},
	}
}

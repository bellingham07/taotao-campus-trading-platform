package svc

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/service/cmdty/cmd/rpc/cmdtyservice"
	"go-go-zero/service/trade/cmd/api/internal/config"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Xorm      *xorm.Engine
	TradeInfo *xorm.Session
	TradeDone *xorm.Session
	TradeCmt  *mongo.Collection

	UserRpc  userservice.UserService
	CmdtyRpc cmdtyservice.CmdtyService
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine, err := xorm.NewEngine("mysql", c.Mysql.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败 " + err.Error())
	}

	clientOptions := options.Client().ApplyURI(c.Mongo.Url) // 设置客户端连接配置
	client, err := mongo.NewClient(clientOptions)           // 创建客户端
	if err != nil {
		panic("[MONGO ERROR] NewServiceContext mongodb 连接失败 " + err.Error())
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic("[MONGO ERROR] NewServiceContext mongodb 连接失败 " + err.Error())
	}
	tcCollection := client.Database("taotao_trading_trade").Collection("trade_cmt")

	return &ServiceContext{
		Config:    c,
		Xorm:      engine,
		TradeInfo: engine.Table("trade_info"),
		TradeDone: engine.Table("trade_done"),
		TradeCmt:  tcCollection,
		UserRpc:   userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		CmdtyRpc:  cmdtyservice.NewCmdtyService(zrpc.MustNewClient(c.CmdtyRpc)),
	}
}

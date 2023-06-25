package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/rpc/cmdtyservice"
	"go-go-zero/service/trade/cmd/api/internal/config"
	"go-go-zero/service/trade/cmd/api/internal/middleware"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"go.mongodb.org/mongo-driver/mongo"
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

	JwtAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.TradeApi.Mysql)

	mc := utils.InitMongo(c.TradeApi.Mongo)

	return &ServiceContext{
		Config:    c,
		Xorm:      engine,
		TradeInfo: engine.Table("trade_info"),
		TradeDone: engine.Table("trade_done"),
		TradeCmt:  mc.Database("taotao_trading_trade").Collection("trade_cmt"),
		UserRpc:   userservice.NewUserService(zrpc.MustNewClient(c.TradeApi.UserRpc)),
		CmdtyRpc:  cmdtyservice.NewCmdtyService(zrpc.MustNewClient(c.TradeApi.CmdtyRpc)),
		JwtAuth:   middleware.NewJwtAuthMiddleware().Handle,
	}
}

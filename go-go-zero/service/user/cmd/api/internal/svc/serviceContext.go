package svc

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/rest"
	"go-go-zero/common/utils"
	"go-go-zero/service/user/cmd/api/internal/config"
	"go-go-zero/service/user/cmd/api/internal/middleware"
	"xorm.io/xorm"
)

type ServiceContext struct {
	// local
	Config config.Config

	// model
	Xorm         *xorm.Engine
	UserInfo     *xorm.Session
	UserFollow   *xorm.Session
	UserCall     *xorm.Session
	UserLocation *xorm.Session
	UserOpt      *xorm.Session

	// redis
	Redis *redis.Client

	Json jsoniter.API

	JwtAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	options := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(options)

	engine := utils.InitXorm("mysql", c.Mysql)

	return &ServiceContext{
		Config:       c,
		Xorm:         engine,
		UserInfo:     engine.Table("user_info"),
		UserFollow:   engine.Table("user_follow"),
		UserCall:     engine.Table("user_call"),
		UserLocation: engine.Table("user_location"),
		UserOpt:      engine.Table("user_opt"),
		Json:         jsoniter.ConfigCompatibleWithStandardLibrary,
		JwtAuth:      middleware.NewJwtAuthMiddleware().Handle,
		Redis:        utils.InitRedis(c.Redis),
	}
}

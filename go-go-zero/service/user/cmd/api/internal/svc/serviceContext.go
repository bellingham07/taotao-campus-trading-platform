package svc

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/user/cmd/api/internal/config"
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
}

func NewServiceContext(c config.Config) *ServiceContext {
	options := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(options)

	engine, err := xorm.NewEngine("mysql", c.Mysql.Dsn)
	if err != nil {
		panic("[XORM ERROR] NewServiceContext 获取mysql连接错误 " + err.Error())
	}
	err = engine.Ping()
	if err != nil {
		panic("[XORM ERROR] NewServiceContext ping mysql 失败" + err.Error())
	}

	return &ServiceContext{
		Config:       c,
		UserInfo:     engine.Table("user_info"),
		UserFollow:   engine.Table("user_follow"),
		UserCall:     engine.Table("user_call"),
		UserLocation: engine.Table("user_location"),
		UserOpt:      engine.Table("user_opt"),
		Xorm:         engine,
		Json:         jsoniter.ConfigCompatibleWithStandardLibrary,
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.Db,
		}),
	}
}

package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type Redis struct {
	Addr     string
	Password string
	Db       int
}

const (

	// user
	UserLogin    = "user:login:"
	UserInfo     = "user:info:"
	UserLocation = "user:location"

	// commodity
	CmdtySellNewest = "cmdty:sell:newest:"
	CmdtyWantNewest = "cmdty:want:newest:"
	CmdtyInfo       = "cmdty:info:"
	CmdtyHistory    = "cmdty:history:"
	CmdtyCategory   = "cmdty:category:"
	CmdtyCollect    = "cmdty:collect:"
	CmdtyView       = "cmdty:view:"
	CmdtyTag        = "cmdty:tag"

	// article
	AtclContent = "atcl:content:"
	AtclCollect = "atcl:collect:"
	AtclView    = "atcl:view:"
	AtclLike    = "atcl:like:"

	// trade
	TradeInfo = "trade:info:"
)

var UserServiceRedis *redis.Client

func UserServiceInit(ctx context.Context, client *redis.Client) {
	UserServiceRedis = client
	err := client.Ping(ctx).Err()
	if err != nil {
		panic("[REDIS ERROR] 连接redis失败 " + err.Error())
	}
}

func InitRedis(rc Redis) *redis.Client {
	var ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Password: rc.Password,
		DB:       rc.Db,
	})
	logx.Infof("[REDIS CONNECTING] InitRedis client: %v\n", client)

	err := client.Ping(ctx).Err()
	if err != nil {
		panic("[REDIS ERROR] 连接redis失败 " + err.Error())
	}
	UserServiceInit(ctx, client)
	return client
}

package cache

import (
	"com.xpwk/go-gin/config"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	USERLOGIN = "user:login:"
	USERINFO  = "user:info:"
)

var (
	RedisClient _RedisClient
	ctx         = context.Background()
)

type _RedisClient struct {
	*redis.Client
}

func InitRedis(config config.RedisConfig) {
	RedisClient = _RedisClient{
		redis.NewClient(&redis.Options{
			Addr:     config.Url,
			Password: config.Password,
			DB:       config.Db,
		}),
	}
	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	ctx.Value(config.Password)
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic("连接redis失败：" + err.Error())
	}
}

func (rc *_RedisClient) Get(key string) (result string, err error) {
	result, err = rc.Client.Get(ctx, key).Result()
	return
}

func (rc *_RedisClient) Set(key string, value any, expiration time.Duration) (err error) {
	jsonStr, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = rc.Client.Set(ctx, key, jsonStr, expiration).Err()
	return
}

func (rc *_RedisClient) Expire(key string, expiration time.Duration) (err error) {
	err = rc.Client.Expire(ctx, key, expiration).Err()
	return
}

package cache

import (
	"com.xpwk/go-gin/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	USERID = "user:id:"
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
	err = rc.Client.Set(ctx, key, value, expiration).Err()
	return
}

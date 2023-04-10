package initial

import (
	"com.xpwk/go-gin/internal/config"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func initRedis(config config.RedisConfig) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Url,
		Password: config.Password,
		DB:       0,
	})
}

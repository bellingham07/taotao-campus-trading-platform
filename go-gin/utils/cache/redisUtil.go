package cache

import (
	"com.xpdj/go-gin/config"
	"com.xpdj/go-gin/utils/jsonUtil"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	UserLogin    = "user:login:"
	UserInfo     = "user:info:"
	UserLocation = "user:location:"

	CommodityInfo     = "cmdty:info:"
	CommodityHistory  = "cmdty:history:"
	CommodityCategory = "cmdty:category:"
	CommodityCollect  = "cmdty:collect:"
	CommodityView     = "cmdty:view:"
	CommodityLike     = "cmdty:like:"

	ArticleContent = "atcl:content:"
	ArticleCollect = "atcl:collect:"
	ArticleView    = "atcl:view:"
	ArticleLike    = "atcl:like:"

	OrderInfo = "trade:info:"
)

var (
	RedisUtil _RedisClient
	ctx       = context.Background()
)

type _RedisClient struct {
	Client *redis.Client
}

func InitRedis(config *config.RedisConfig) {
	RedisUtil = _RedisClient{
		redis.NewClient(&redis.Options{
			Addr:     config.Url,
			Password: config.Password,
			DB:       config.Db,
		}),
	}
	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	ctx.Value(config.Password)
	_, err := RedisUtil.Client.Ping(ctx).Result()
	if err != nil {
		panic("连接redis失败：" + err.Error())
	}
}

func (rc *_RedisClient) GET(key string) (string, error) {
	result, err := rc.Client.Get(ctx, key).Result()
	return result, err
}

func (rc *_RedisClient) SET2JSON(key string, value any, expiration time.Duration) error {
	jsonStr, err := jsonUtil.Json.Marshal(value)
	if err != nil {
		return err
	}
	err = rc.Client.Set(ctx, key, jsonStr, expiration).Err()
	return err
}

func (rc *_RedisClient) EXPIRE(key string, expiration time.Duration) error {
	err := rc.Client.Expire(ctx, key, expiration).Err()
	return err
}
func (rc *_RedisClient) DEL(key string) (err error) {
	err = rc.Client.Del(ctx, key).Err()
	return
}

func (rc *_RedisClient) HGETALL(key string) map[string]string {
	resultMap, err := rc.Client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil
	}
	return resultMap
}

func (rc *_RedisClient) HSET(key string, value any) error {
	if err := rc.Client.HSet(ctx, key, value).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) HSET1(key string, field string, value any) error {
	if err := rc.Client.HSet(ctx, key, field, value).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) HSETNX(key string, field string, value any) error {
	if err := rc.Client.HSetNX(ctx, key, field, value).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) HSETNXPX(key string, field string, value any, expiration time.Duration) error {
	if err := rc.Client.HSetNX(ctx, key, field, value).Err(); err != nil {
		return err
	}
	if err := rc.Client.Expire(ctx, key, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) HDEL(key string, field string) error {
	if err := rc.Client.HDel(ctx, key, field).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) ZADD(key string, members ...redis.Z) error {
	if err := rc.Client.ZAdd(ctx, key, members...).Err(); err != nil {
		return nil
	}
	return nil
}

func (rc *_RedisClient) ZADDNX(key string, members ...redis.Z) error {
	if err := rc.Client.ZAddNX(ctx, key, members...).Err(); err != nil {
		return nil
	}
	return nil
}

func (rc *_RedisClient) ZREM(key string, members ...interface{}) error {
	if err := rc.Client.ZRem(ctx, key, members).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) ZREVRANGEWITHSCORES(key string, start, stop int64) (result []redis.Z) {
	result, err := rc.Client.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil
	}
	return result
}

func (rc *_RedisClient) ZREVRANGE(key string, start, stop int64) []string {
	result, err := rc.Client.ZRevRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil
	}
	return result
}

func (rc *_RedisClient) SADD(key string, members ...interface{}) (affect int64) {
	if affect, _ = rc.Client.SAdd(ctx, key, members).Result(); affect == 0 {
		return 0
	}
	return affect
}

func (rc *_RedisClient) SISMEMBER(key string, member interface{}) (ok bool) {
	if ok, _ = rc.Client.SIsMember(ctx, key, member).Result(); !ok {
		return ok
	}
	return ok
}

func (rc *_RedisClient) SREM(key string, member ...interface{}) (affect int64) {
	if affect, _ = rc.Client.SRem(ctx, key, member).Result(); affect == 0 {
		return 0
	}
	return affect
}

func (rc *_RedisClient) INCR(key string) (affect int64) {
	if affect, _ = rc.Client.Incr(ctx, key).Result(); affect == 0 {
		return 0
	}
	return
}

func (rc *_RedisClient) HINCRBY1(key, field string) (affect int64) {
	if affect, _ = rc.Client.HIncrBy(ctx, key, field, 1).Result(); affect == 0 {
		return 0
	}
	return
}

func (rc *_RedisClient) HINCRBY(key, field string, incr int64) (affect int64) {
	if affect, _ = rc.Client.HIncrBy(ctx, key, field, incr).Result(); affect == 0 {
		return 0
	}
	return
}

func (rc *_RedisClient) HGET(key, field string) string {
	affect, err := rc.Client.HGet(ctx, key, field).Result()
	if err != nil {
		return ""
	}
	return affect
}

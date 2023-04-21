package cache

import (
	"com.xpwk/go-gin/config"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"reflect"
	"time"
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

func (rc *_RedisClient) GET(key string) (result string, err error) {
	result, err = rc.Client.Get(ctx, key).Result()
	return
}

func (rc *_RedisClient) SET(key string, value any, expiration time.Duration) (err error) {
	jsonStr, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = rc.Client.Set(ctx, key, jsonStr, expiration).Err()
	return
}

func (rc *_RedisClient) EXPIRE(key string, expiration time.Duration) (err error) {
	err = rc.Client.Expire(ctx, key, expiration).Err()
	return
}

func (rc *_RedisClient) HGETALL(key string) (resultMap map[string]string, err error) {
	resultMap, err = rc.Client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return resultMap, nil
}

func (rc *_RedisClient) HSET(key string, value any) (err error) {
	if err = rc.Client.HSet(ctx, key, struct2map(value)).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) ZADD(key string, members ...*redis.Z) (err error) {
	if err = rc.Client.ZAdd(ctx, key, members...).Err(); err != nil {
		return nil
	}
	return nil
}

func (rc *_RedisClient) ZADDNX(key string, members ...*redis.Z) (err error) {
	if err = rc.Client.ZAddNX(ctx, key, members...).Err(); err != nil {
		return nil
	}
	return nil
}

func (rc *_RedisClient) ZREM(key string, id string) (err error) {
	if err = rc.Client.ZRem(ctx, key, id).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) ZREVRANGE(key string, start, stop int64) (result []string, err error) {
	if result, err = rc.Client.ZRevRange(ctx, key, start, stop).Result(); err != nil {
		return nil, err
	}
	return result, nil
}

func (rc *_RedisClient) HSET1(key string, field string, value any) (err error) {
	if err = rc.HSet(ctx, key, field, value).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *_RedisClient) HMSET(key string, value any, expiration time.Duration) (err error) {
	valueMap := struct2map(value)
	if err = rc.HSet(ctx, key, valueMap).Err(); err != nil {
		return err
	}
	return nil
}

func struct2map(value any) map[string]interface{} {
	valueMap := make(map[string]interface{})
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if key := field.Tag.Get("json"); key != "" {
			valueMap[key] = v.Field(i).Interface()
		}
	}
	return valueMap
}

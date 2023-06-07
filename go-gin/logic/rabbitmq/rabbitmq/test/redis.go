package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic(err)
	}

	return RedisClient{client: client}
}

func (r RedisClient) Like(userID int, postID string) error {
	// 记录点赞信息
	err := r.client.ZAdd(r.client.Context(), fmt.Sprintf("likes:%s", postID), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: strconv.Itoa(userID),
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r RedisClient) Unlike(userID int, postID string) error {
	// 删除点赞信息
	err := r.client.ZRem(r.client.Context(), fmt.Sprintf("likes:%s", postID), userID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisClient) IsLike(userID int, postID string) (bool, error) {
	// 检查是否点过赞
	result, err := r.client.ZScore(r.client.Context(), fmt.Sprintf("likes:%s", postID), strconv.Itoa(userID)).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	if result > 0 {
		return true, nil
	}

	return false, nil
}

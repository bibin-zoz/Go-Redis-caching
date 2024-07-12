package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	*redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"), // Example: "localhost:6379"
		Password: "8596",                     // Example: "your_redis_password"
		DB:       0,
	})

	return &RedisClient{client}
}

func (c *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := c.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s in Redis: %v", key, err)
	}
	return nil
}

func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key %s does not exist in Redis", key)
		}
		return "", fmt.Errorf("failed to get key %s from Redis: %v", key, err)
	}
	return val, nil
}

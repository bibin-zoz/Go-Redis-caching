package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis-caching/config"
	"go-redis-caching/pkg/model"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserService struct {
	redisClient *config.RedisClient
}

func NewUserService() *UserService {
	return &UserService{
		redisClient: config.NewRedisClient(),
	}
}

func (s *UserService) AddUser(user model.User) error {
	ctx := context.Background()

	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %v", err)
	}

	err = s.redisClient.Set(ctx, user.ID, userJSON, time.Hour)
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("key %s does not exist in Redis", user.ID)
		}
		return fmt.Errorf("failed to store user in Redis: %v", err)
	}

	return nil
}

func (s *UserService) GetUser(userID string) (*model.User, error) {
	ctx := context.Background()

	userJSON, err := s.redisClient.Get(ctx, userID)
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("key %s does not exist in Redis", userID)
		}
		return nil, fmt.Errorf("failed to get user from Redis: %v", err)
	}

	var user model.User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %v", err)
	}

	return &user, nil
}

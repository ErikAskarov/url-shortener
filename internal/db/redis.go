package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr, password string) *RedisStorage {
	return &RedisStorage{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (s *RedisStorage) Ping(ctx context.Context) (string, error) {
	ping, err := s.client.Ping(ctx).Result()
	if err != nil {
		return "", err
	}
	return ping, nil
}

func (s *RedisStorage) SaveUrl(ctx context.Context, url string, code string, expire time.Duration) error {
	return s.client.Set(ctx, code, url, expire).Err()
}

func (s *RedisStorage) GetUrl(ctx context.Context, code string) (string, error) {
	return s.client.Get(ctx, code).Result()
}

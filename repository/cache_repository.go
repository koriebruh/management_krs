package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheRepository interface {
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type RedisCacheRepository struct {
	client *redis.Client
}

func NewRedisCacheRepository(client *redis.Client) *RedisCacheRepository {
	return &RedisCacheRepository{client: client}
}

func (r RedisCacheRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, string(jsonValue), expiration).Err()
}

func (r RedisCacheRepository) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r RedisCacheRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

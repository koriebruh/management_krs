package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
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
	log.Printf("DEBUG - Attempting to save: Key=%s", key) // tambahkan ini

	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal cache value: %v", err)
	}

	// Simpan ke Redis
	err = r.client.Set(ctx, key, jsonData, expiration).Err()
	if err != nil {
		log.Printf("DEBUG - Error saving to Redis: %v", err) // tambahkan ini
		return fmt.Errorf("failed to set cache: %v", err)
	}

	// Verifikasi langsung
	savedValue, err := r.client.Get(ctx, key).Result()
	log.Printf("DEBUG - Immediate verification: Key=%s, Value=%s, Error=%v", key, savedValue, err) // tambahkan ini

	return nil
}

func (r RedisCacheRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("cache miss for key %s", key)
		}
		return "", err
	}

	log.Printf("Cache hit. Key: %s, Value: %s", key, val)
	return val, nil
}

func (r RedisCacheRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

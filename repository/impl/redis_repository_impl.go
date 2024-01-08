package impl

import (
	"context"
	"time"

	"github.com/ariwiraa/my-gram/repository"
	"github.com/redis/go-redis/v9"
)

type redisRepositoryImpl struct {
	client *redis.Client
}

func NewRedisRepositoryImpl(client *redis.Client) repository.RedisRepository {
	return &redisRepositoryImpl{
		client: client,
	}
}

// Get implements repository.RedisRepository.
func (r *redisRepositoryImpl) Get(ctx context.Context, key string) (interface{}, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return value, err
	}

	return value, nil
}

// Set implements repository.RedisRepository.
func (r *redisRepositoryImpl) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	err := r.client.Set(ctx, key, value, ttl)
	if err != nil {
		return err.Err()
	}

	return nil
}

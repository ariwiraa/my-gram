package repository

import (
	"context"
	"time"
)

type RedisRepository interface {
	Set(ctx context.Context, key, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
}

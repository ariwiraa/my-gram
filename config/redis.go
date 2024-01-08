package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func ConnectRedis(cfg *Config) (*redis.Client, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
	})

	err := client.Ping(ctx)
	if err.Err() != nil {
		return nil, err.Err()
	}

	return client, nil
}

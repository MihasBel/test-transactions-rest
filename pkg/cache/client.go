package cache

import (
	"github.com/go-redis/redis"
)

// New create redis client
func New(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: cfg.RedisPass,
		DB:       0,
	})
}

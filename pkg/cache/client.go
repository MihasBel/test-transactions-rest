package cache

import (
	"github.com/MihasBel/test-transactions-rest/internal/app"
	"github.com/go-redis/redis"
)

// New create redis client
func New(cfg app.Configuration) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURL,
		Password: cfg.RedisPass,
		DB:       0,
	})
}

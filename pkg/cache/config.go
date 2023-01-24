package cache

// Config configuration
type Config struct {
	RedisURL  string `json:"redis_url"`
	RedisPass string `json:"redis_pass"`
}

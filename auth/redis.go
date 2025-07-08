package auth

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

// SetupRedis initializes the Redis client with default settings.
func SetupRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

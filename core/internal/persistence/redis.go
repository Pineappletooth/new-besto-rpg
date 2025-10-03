package persistence

import (
	"context"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

var redisClient = getClient()
var ctx = context.Background()

func getClient() *redis.Client {
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opts)
}

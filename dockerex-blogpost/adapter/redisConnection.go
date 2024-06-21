package adapter

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func RedisConnection() *redis.Client {
	address := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	return client
}

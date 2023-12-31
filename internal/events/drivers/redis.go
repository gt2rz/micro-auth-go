package drivers

import (
	"os"

	"github.com/go-redis/redis"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	err := client.Ping().Err()

	if err != nil {
		return nil, err
	}

	return client, nil
}

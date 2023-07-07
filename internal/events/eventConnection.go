package events

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/gt2rz/micro-auth/internal/events/drivers"
)

type EventConnection interface {
	Publish(ctx context.Context, channel string, message string) error
}

type RedisPublisher struct {
	client *redis.Client
}

func NewRedisPublisher() (*RedisPublisher, error) {
	client, err := drivers.NewRedisClient()
	if err != nil {
		return nil, err
	}

	return &RedisPublisher{
		client: client,
	}, nil
}

func (r *RedisPublisher) Publish(ctx context.Context, channel string, message string) error {
	err := r.client.Publish(channel, message).Err()
	if err != nil {
		return err
	}

	return nil
}

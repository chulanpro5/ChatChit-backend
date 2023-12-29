package redis

import (
	"context"
	"fmt"
	GoRedis "github.com/redis/go-redis/v9"
	"test-chat/config"
)

type Client struct {
	Client *GoRedis.Client
}

func NewRedisClient(config *config.Config) (*Client, error) {
	client := GoRedis.NewClient(&GoRedis.Options{
		Addr:     fmt.Sprint(config.Redis.Host, ":", config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})
	// dummy root context
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Client{
		client,
	}, nil
}

func (r *Client) CleanUp() {
	r.Client.Close()
}

func (r *Client) GetValue(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *Client) SetValue(ctx context.Context, key string, val string) error {
	err := r.Client.Set(ctx, key, val, 0).Err()

	return err
}

func (r *Client) Publish(ctx context.Context, channel string, message string) error {
	err := r.Client.Publish(ctx, channel, message).Err()

	return err
}

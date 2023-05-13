package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *redis.Client
	ctx    context.Context
}

func (c *Client) Set(key string, value interface{}) error {
	return c.client.Set(c.ctx, key, value, 0).Err()
}

func (c *Client) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

func NewClient(ctx context.Context, addr string, port int, password string, db int) *Client {
	return &Client{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", addr, port),
			Password: password,
			DB:       db,
		}),
		ctx: ctx,
	}
}

func TestRedisConnection(host string, port int) bool {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return false
	}
	return true
}

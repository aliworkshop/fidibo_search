package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (c *client) Store(key string, data interface{}, dur time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = c.redis.Set(context.Background(), key, b, dur).Err(); err != nil {
		return err
	}
	return nil
}

func (c *client) Exists(key string) bool {
	return c.redis.Exists(context.Background(), key).Val() == 1
}

func (c *client) Get(key string) ([]byte, error) {
	b, err := c.redis.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *client) Delete(key string) error {
	return c.redis.Del(context.Background(), key).Err()
}

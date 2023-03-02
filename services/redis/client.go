package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type client struct {
	redis *redis.Client
}

type Config struct {
	Addr     string
	Password string
	DB       int
	Timeout  time.Duration
}

type Client interface {
	Store(key string, data interface{}, dur time.Duration) error
	Exists(key string) bool
	Get(key string) ([]byte, error)
	Delete(key string) error
}

func NewRedisClient(cfg Config) Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, _ := context.WithTimeout(context.Background(), cfg.Timeout)
	_, errConnect := rdb.Ping(ctx).Result()
	if errConnect != nil {
		panic(errConnect)
	}

	return &client{redis: rdb}
}

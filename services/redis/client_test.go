package redis

import (
	"encoding/json"
	"fidibo/services/fidibo"
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestRedisClient(t *testing.T) {
	f, err := fidibo.NewFidiboClient(config().Sub("fidibo"))
	assert.Equal(t, nil, err)

	redis := NewRedisClient(Config{
		Addr:    "localhost:6379",
		Timeout: time.Second,
	})

	redisKey := fmt.Sprintf("search_data_for_%s", "test")
	resp, err := f.Search("test")
	assert.Equal(t, nil, err)

	err = redis.Delete(redisKey)
	assert.Equal(t, nil, err)

	exists := redis.Exists(redisKey)
	assert.Equal(t, false, exists)

	err = redis.Store(redisKey, resp, 10*time.Minute)
	assert.Equal(t, nil, err)

	exists = redis.Exists(redisKey)
	assert.Equal(t, true, exists)

	var data fidibo.BookResponse
	b, err := redis.Get(redisKey)
	assert.Equal(t, nil, err)

	err = json.Unmarshal(b, &data)
	assert.Equal(t, nil, err)
}

func config() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	f, err := os.Open("../../config.yaml")
	if err != nil {
		panic("cannot read config: " + err.Error())
	}
	err = v.ReadConfig(f)
	if err != nil {
		panic("cannot read config" + err.Error())
	}

	return v
}

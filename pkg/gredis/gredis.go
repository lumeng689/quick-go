package gredis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func SetUp(addr string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10,
	})
}

func Set(key string, data interface{}, ts int) error {

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = rdb.Set(context.Background(), key, value, 5*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

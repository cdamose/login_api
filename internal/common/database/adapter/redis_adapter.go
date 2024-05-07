package adapter

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisAdapter struct {
	Client *redis.Client
}

func (ra *RedisAdapter) Get(key string) string {
	return ""
}

func (ra *RedisAdapter) Set(key string, value interface{}) (bool, error) {
	err := ra.Client.Set(key, value, 0)
	if err != nil {
		return false, err.Err()
	}
	return true, nil
}

func (ra *RedisAdapter) SetWithTTL(key string, value string, ttl time.Duration) (bool, error) {
	err := ra.Client.Set(key, value, ttl)
	if err != nil {
		return false, err.Err()
	}
	return true, nil
}

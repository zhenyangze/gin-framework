// Package persistence provides ...
package persistence

import (
	"gitee.com/zhenyangze/gin-framework/app/providers"
	"time"
)

type RedisStore struct {
}

func (c *RedisStore) Set(key string, value interface{}, expires time.Duration) error {
	return providers.Redis.Set(key, value, expires).Err()
}

func (c *RedisStore) Get(key string) (string, error) {
	val, err := providers.Redis.Get(key).Result()
	return val, err
}

func (c *RedisStore) Delete(key string) error {
	return providers.Redis.Del(key).Err()
}

// Package providers provides ...
package providers

import (
	"github.com/go-redis/redis"
	"gitee.com/zhenyangze/gin-framework/configs"
)

var (
	Redis *redis.Client
)

func init() {
	redisConfig := configs.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisConfig["Addr"].(string),
		Password: redisConfig["Password"].(string), // no password set
		DB:       redisConfig["DB"].(int),          // use default DB
	})
}

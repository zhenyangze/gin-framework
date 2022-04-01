// Package configs provides ...
package configs

import "gitee.com/zhenyangze/gin-framework/helpers"

func GetRedisConfig() map[string]interface{} {
	redisConfig := make(map[string]interface{})
	var config *helpers.Config

	redisConfig["Addr"] = config.GetStringByDefault("redis.addr", "127.0.0.1:3306")
	redisConfig["Password"] = config.GetStringByDefault("redis.password", "")
	redisConfig["DB"] = config.GetInt("redis.db")

	return redisConfig
}

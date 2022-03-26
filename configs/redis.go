// Package configs provides ...
package configs

import "gitee.com/zhenyangze/gin-framework/helpers"

func GetRedisConfig() map[string]interface{} {
	redisConfig := make(map[string]interface{})
	var config *helpers.Config

	redisConfig["Addr"] = config.GetString("redis.addr")
	redisConfig["Password"] = config.GetString("redis.password")
	redisConfig["DB"] = config.GetInt("redis.db")

	return redisConfig
}

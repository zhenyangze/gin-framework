// Package configs provides ...
package configs

func GetRedisConfig() map[string]interface{} {
	redisConfig := make(map[string]interface{})

	redisConfig["Addr"] = "127.0.0.1:6379"
	redisConfig["Password"] = ""
	redisConfig["DB"] = 0

	return redisConfig
}

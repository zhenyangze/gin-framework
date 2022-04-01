// Package configs provides ...
package configs

import (
	"fmt"

	"gitee.com/zhenyangze/gin-framework/helpers"
)

func GetDBConfig() map[string]string {
	dbConfig := make(map[string]string)

	var config *helpers.Config

	dbConfig["DB_HOST"] = config.GetString("databases.host")
	dbConfig["DB_PORT"] = config.GetString("databases.port")
	dbConfig["DB_NAME"] = config.GetString("databases.database")
	dbConfig["DB_USER"] = config.GetString("databases.username")
	dbConfig["DB_PWD"] = config.GetString("databases.password")

	dbConfig["DB_CHARSET"] = config.GetStringByDefault("databases.charset", "utf8")

	dbConfig["DB_MAX_OPEN_CONNS"] = helpers.IntToString(config.GetInt64ByDefault("databases.max_open_conns", 100))          // 连接池最大连接数
	dbConfig["DB_MAX_IDLE_CONNS"] = helpers.IntToString(config.GetInt64ByDefault("databases.max_idle_conns", 50))           // 连接池最大空闲数
	dbConfig["DB_MAX_LIFETIME_CONNS"] = helpers.IntToString(config.GetInt64ByDefault("databases.max_lifetime_conns", 7200)) // 连接池链接最长生命周期

	dbConfig["DB_DEBUG"] = config.GetStringByDefault("databases.debug", "false")
	return dbConfig
}

func GetDbDSN() string {
	dbConfig := GetDBConfig()
	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
	)

	return connStr
}

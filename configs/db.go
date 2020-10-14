// Package configs provides ...
package configs

import "fmt"

func GetDBConfig() map[string]string {
	dbConfig := make(map[string]string)

	dbConfig["DB_HOST"] = "127.0.0.1"
	dbConfig["DB_PORT"] = "3306"
	dbConfig["DB_NAME"] = "go-admin"
	dbConfig["DB_USER"] = "root"
	dbConfig["DB_PWD"] = "root"
	dbConfig["DB_CHARSET"] = "utf8"

	dbConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	dbConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	dbConfig["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	dbConfig["DB_DEBUG"] = "true" // 打印日志
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

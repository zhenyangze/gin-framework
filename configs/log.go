// Package configs provides ...
package configs

import (
	"os"

	"gitee.com/zhenyangze/gin-framework/helpers"
)

func GetLoggerConfig() map[string]interface{} {
	loggerConfig := make(map[string]interface{})

	var config *helpers.Config
	loggerConfig["path"] = config.GetString("log.path")
	if dir, err := os.Getwd(); err == nil && loggerConfig["path"] == "" {
		loggerConfig["path"] = dir + "/runtime/logs"
	}

	loggerConfig["level"] = config.GetStringByDefault("log.level", "debug")
	loggerConfig["log_name"] = config.GetStringByDefault("log.name", "system")

	return loggerConfig
}

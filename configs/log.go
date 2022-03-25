// Package configs provides ...
package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLoggerConfig() map[string]interface{} {
	loggerConfig := make(map[string]interface{})

	loggerConfig["path"] = "/tmp"
	if dir, err := os.Getwd(); err == nil {
		loggerConfig["path"] = dir + "/runtime/logs"
	}
	loggerConfig["level"] = logrus.DebugLevel
	loggerConfig["type"] = 0
	loggerConfig["log_name"] = "system"

	return loggerConfig
}

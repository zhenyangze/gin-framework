package helpers

import jsoniter "github.com/json-iterator/go"

var (
	APP_PATH    string
	CONFIG_PATH string
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
}

func SetAppPath(path string) {
	APP_PATH = path
}

func GetAppPath() string {
	return APP_PATH
}

func SetConfigPath(path string) {
	CONFIG_PATH = path
}

func GetConfigPath() string {
	return CONFIG_PATH
}

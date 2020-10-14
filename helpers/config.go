// Package helpers provides ...
package helpers

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	configList map[string]*viper.Viper
	configOne  sync.Once
)

type Config struct {
	configList map[string]*viper.Viper
}

func init() {
	configOne.Do(func() {
		if configList == nil {
			configList = make(map[string]*viper.Viper)
		}
		configPath := "configs"
		fileList, err := Glob(configPath+"/*.toml", true)
		if err != nil {
			return
		}
		for _, v := range fileList {
			config := viper.New()
			config.AddConfigPath(configPath)
			// 支持多个配置文件,指定目录
			config.SetConfigFile(v)
			config.SetConfigType("toml")

			configList[strings.TrimRight(Basename(v), Ext(v))] = config
		}

	})
}

func getViper(keyname string) *viper.Viper {
	fileName := strings.Split(keyname, ".")[0]
	viper := configList[fileName]
	return viper
}

func getOriginKey(keyname string) (fileName string, realKeyName string) {
	fileName = strings.Split(keyname, ".")[0]
	realKeyName = strings.TrimLeft(keyname, fileName)
	realKeyName = strings.Trim(realKeyName, ".")
	return
}

// get 一个原始值
func (c *Config) Get(keyname string) interface{} {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).Get(realKeyName)
}

// getstring
func (c *Config) GetString(keyname string) string {
	fileName, realKeyName := getOriginKey(keyname)
	fmt.Println(fileName, realKeyName)
	fmt.Println(getViper(fileName))
	return getViper(fileName).GetString(realKeyName)
}

// getbool
func (c *Config) GetBool(keyname string) bool {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetBool(realKeyName)
}

// getint
func (c *Config) GetInt(keyname string) int {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetInt(realKeyName)
}

// getint32
func (c *Config) GetInt32(keyname string) int32 {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetInt32(realKeyName)
}

// getint64
func (c *Config) GetInt64(keyname string) int64 {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetInt64(realKeyName)
}

// float64
func (c *Config) GetFloat64(keyname string) float64 {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetFloat64(realKeyName)
}

// GetDuration
func (c *Config) GetDuration(keyname string) time.Duration {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetDuration(realKeyName)
}

// GetStringSlice
func (c *Config) GetStringSlice(keyname string) []string {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetStringSlice(realKeyName)
}

// GetStringMap
func (c *Config) GetStringMap(keyname string) map[string]interface{} {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetStringMap(realKeyName)
}

// GetStringMapString
func (c *Config) GetStringMapString(keyname string) map[string]string {
	fileName, realKeyName := getOriginKey(keyname)
	return getViper(fileName).GetStringMapString(realKeyName)
}

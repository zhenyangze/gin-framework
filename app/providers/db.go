// Package providers provides...
package providers

import (
	"strconv"
	"time"

	"gitee.com/zhenyangze/gin-framework/configs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	err error
)

func InitDb() {
	dbConfigs := configs.GetDBConfig()
	DB, err = gorm.Open("mysql", configs.GetDbDSN())
	if err != nil {
		panic("连接数据库失败")
	}
	//defer DB.Close()
	// 打印日志

	dbDebug, _ := strconv.ParseBool(dbConfigs["DB_DEBUG"])
	DB.LogMode(dbDebug)

	// 打开连接池
	dbMaxConnect, _ := strconv.Atoi(dbConfigs["DB_MAX_OPEN_CONNS"])
	DB.DB().SetMaxOpenConns(dbMaxConnect)

	//连接池最大空闲数
	dbMaxIdleConns, _ := strconv.Atoi(dbConfigs["DB_MAX_IDLE_CONNS"])
	DB.DB().SetMaxIdleConns(dbMaxIdleConns)

	// 连接池链接最长生命周期
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfigs["DB_MAX_LIFETIME_CONNS"])
	DB.DB().SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))
}

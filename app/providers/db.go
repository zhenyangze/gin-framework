// Package providers provides...
package providers

import (
	"fmt"
	"strconv"
	"time"

	"gitee.com/zhenyangze/gin-framework/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDb() {
	dbConfigs := configs.GetDBConfig()
	DB, err = gorm.Open(mysql.Open(configs.GetDbDSN()), &gorm.Config{})
	if err != nil {
		fmt.Println(configs.GetDbDSN())
		panic("连接数据库失败")
	}
	//defer DB.Close()
	// 打印日志

	dbDebug, _ := strconv.ParseBool(dbConfigs["DB_DEBUG"])
	if dbDebug {
		DB = DB.Debug()
	}

	// 打开连接池
	sqlDB, _ := DB.DB()
	dbMaxConnect, _ := strconv.Atoi(dbConfigs["DB_MAX_OPEN_CONNS"])
	sqlDB.SetMaxOpenConns(dbMaxConnect)

	//连接池最大空闲数
	dbMaxIdleConns, _ := strconv.Atoi(dbConfigs["DB_MAX_IDLE_CONNS"])
	sqlDB.SetMaxIdleConns(dbMaxIdleConns)

	// 连接池链接最长生命周期
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfigs["DB_MAX_LIFETIME_CONNS"])
	sqlDB.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))
}

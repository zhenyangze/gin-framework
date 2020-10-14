// Package boot provides ...
package boot

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"gitee.com/zhenyangze/gin-framework/app/providers"
	"gitee.com/zhenyangze/gin-framework/routes"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go/extra"
)

var (
	h bool
	t string
	p string
	d bool
)

func init() {
	extra.RegisterFuzzyDecoders() // 兼容PHP
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&t, "t", "web", "Run Type[web,cron,rps,all]")
	flag.StringVar(&p, "p", ":8080", "Port")
	flag.BoolVar(&d, "d", false, "Debug")
}

func Run() {
	flag.Parse()

	if h {
		usage()
		return
	}

	// 加载配置
	runtime.GOMAXPROCS(runtime.NumCPU())
	if d {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化
	InitEvent()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// 添加反射，获取
	switch t {
	case "web":
		providers.Event.Publish("main:init")
		routes.Web(router)
	case "cron":
		routes.Cron(router)
	case "rpc":
		routes.Rpc(router)
	case "all":
		providers.Event.Publish("main:init")
		routes.Web(router)
		routes.Cron(router)
		routes.Rpc(router)
	}

	router.Run(p)
}

func usage() {
	fmt.Fprintf(os.Stderr, `gin-framework: v1.0.0
Usage: main [-h] [-t type] [-d debug] [-p port]

Options:
`)
	flag.PrintDefaults()
}

func Done() {
	providers.DB.Close()
	providers.Redis.Close()
}

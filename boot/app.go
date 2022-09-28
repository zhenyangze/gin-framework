// Package boot provides ...
package boot

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"gitee.com/zhenyangze/gin-framework/app/providers"
	"gitee.com/zhenyangze/gin-framework/helpers"
	"gitee.com/zhenyangze/gin-framework/routes"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go/extra"
)

var (
	h bool
	t string
	p string
	d bool
	c string
)

func init() {
	extra.RegisterFuzzyDecoders() // 兼容PHP
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&t, "t", "web", "Run Type[web,cron,rpc,all]")
	flag.StringVar(&p, "p", ":8080", "Port")
	flag.BoolVar(&d, "d", false, "Debug")
	flag.StringVar(&c, "c", "configs", "Config Path")
}

func Run() {
	var config *helpers.Config
	flag.Parse()

	if h {
		usage()
		return
	}

	if dirs, err := os.Getwd(); err == nil {
		helpers.SetAppPath(dirs)
		helpers.SetConfigPath(c)
	} else {
		panic("cant get the app root")
	}

	// 加载配置
	runtime.GOMAXPROCS(runtime.NumCPU())
	if d {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化
	helpers.LoadConfig()
	providers.InitLogger()
	providers.InitDb()
	providers.InitEvent()
	providers.InitRedis()
	providers.InitPool()
	InitEvent()

	router := gin.New()
	//router.Use(gin.Logger())
	router.Use(providers.LoggerHandler())
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

	//router.Run(p)
	srv := &http.Server{
		Addr:    p,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	waitTimes := config.GetInt64ByDefault("app.wait_time", 5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(waitTimes)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds.\n", waitTimes)
	}
	log.Println("Server exiting")
}

func usage() {
	fmt.Fprintf(os.Stderr, `gin-framework: v1.0.0
Usage: main [-h] [-t type] [-d debug] [-p port]

Options:
`)
	flag.PrintDefaults()
}

func Done() {
	if h {
		return
	}
	providers.Event.WaitAsync()
	providers.Redis.Close()
}

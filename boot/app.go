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
	"syscall"
	"time"

	"gitee.com/zhenyangze/gin-framework/app/bases"
	"gitee.com/zhenyangze/gin-framework/app/providers"
	"gitee.com/zhenyangze/gin-framework/helpers"
	"gitee.com/zhenyangze/gin-framework/routes"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go/extra"
	_ "go.uber.org/automaxprocs"
)

var (
	h bool
	t string
	p string
	d bool
	c string
	w int64
)

func init() {
	extra.RegisterFuzzyDecoders() // 兼容PHP
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&t, "t", "web", "Run Type[web,cron,rpc,all]")
	flag.StringVar(&p, "p", ":8080", "Port")
	flag.BoolVar(&d, "d", false, "Debug")
	flag.StringVar(&c, "c", "configs", "Config Path")
	flag.Int64Var(&w, "w", 5, "Wait time")
}

func Run() {
	flag.Parse()

	if h {
		usage()
		return
	}

	if dirs, err := os.Getwd(); err == nil {
		bases.BasePath = dirs
		bases.ConfigPath = c
	} else {
		panic("cant get the app root")
	}

	// 加载配置
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
	InitEvent()

	bases.Router = gin.New()
	bases.Router.Use(gin.Recovery())
	bases.Router.Use(providers.LoggerHandler())

	// 添加反射，获取
	switch t {
	case "web":
		providers.Event.Publish("main:init")
		routes.Web()
	case "cron":
		routes.Cron()
	case "rpc":
		routes.Rpc()
	case "all":
		providers.Event.Publish("main:init")
		routes.Web()
		routes.Cron()
		routes.Rpc()
	}

	//router.Run(p)
	srv := &http.Server{
		Addr:    p,
		Handler: bases.Router,
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(w)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds.\n", w)
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

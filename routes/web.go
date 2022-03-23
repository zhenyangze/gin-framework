package routes

import (
	"time"

	"gitee.com/zhenyangze/gin-framework/app/http/controllers"
	"gitee.com/zhenyangze/gin-framework/app/http/middleware"
	"gitee.com/zhenyangze/gin-framework/app/http/middleware/persistence"
	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	_ "go.uber.org/automaxprocs"
)

func Web(router *gin.Engine) {
	router.Use(middleware.DefaultRequestID())

	store := &persistence.RedisStore{}

	// 加载模板文件
	router.LoadHTMLGlob("resources/views/**/*")
	v1 := router.Group("/v1", middleware.Logger(), limit.MaxAllowed(20))
	{
		v1.GET("/test", controllers.TestHandle)
		v1.GET("/view/:name/*action", controllers.ViewHandle)
		v1.GET("/view", controllers.ViewHandle)
		v1.POST("/valid", controllers.ValidHandle)
		v1.GET("/ormmiddle", middleware.CachePage(store, time.Minute, controllers.OrmHandle))
		v1.GET("/ormcache", controllers.OrmWithCacheHandle)
		v1.GET("/orm", controllers.OrmHandle)
		v1.GET("/redis", controllers.RedisHandle)
	}

	// auth
	authMiddleware := middleware.AuthMiddleware

	router.POST("/login", authMiddleware.LoginHandler)
	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", controllers.UsersHandler)
	}

	router.GET("/my", controllers.MyHandle)
	router.GET("/", controllers.MyHandle)

	//websocket
	router.GET("/ws", controllers.WebSocketHandle)
}

package index

import (
	"time"

	"gitee.com/zhenyangze/gin-framework/app/bases"
	"gitee.com/zhenyangze/gin-framework/app/middleware"
	"gitee.com/zhenyangze/gin-framework/app/middleware/persistence"
	"gitee.com/zhenyangze/gin-framework/app/modules/index/handlers"
	limit "github.com/aviddiviner/gin-limit"
)

func Initer() {
	router := bases.Router
	store := &persistence.RedisStore{}
	v1 := router.Group("/v1", limit.MaxAllowed(20))
	{
		v1.GET("/test", handlers.TestHandle)
		v1.GET("/view/:name/*action", handlers.ViewHandle)
		v1.GET("/view", handlers.ViewHandle)
		v1.POST("/valid", handlers.ValidHandle)
		v1.GET("/ormmiddle", middleware.CachePage(store, time.Minute, handlers.OrmHandle))
		v1.GET("/ormcache", handlers.OrmWithCacheHandle)
		v1.GET("/orm", handlers.OrmHandle)
		v1.GET("/redis", handlers.RedisHandle)
		v1.GET("/md5/:str", handlers.Md5Handle)
	}

	router.GET("/", handlers.MyHandle)
}

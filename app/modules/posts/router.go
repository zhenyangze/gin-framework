package posts

import (
	"gitee.com/zhenyangze/gin-framework/app/modules/posts/handlers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	post := router.Group("/api/v1/post")
	{

		post.GET("/:id", handlers.PostHandler.ShowHandler)
		post.POST("/:id", handlers.PostHandler.UpdateHandler)
		post.DELETE("/:id", handlers.PostHandler.DeleteHandler)

		post.POST("/", handlers.PostHandler.StoreHandler)
		post.GET("/", handlers.PostHandler.IndexHandler)

		// 批量操作
		post.PATCH("/", handlers.PostHandler.BatchHandler)
	}
}

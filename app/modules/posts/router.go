package posts

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine) {
	post := router.Group("/api/v1/post")
	{
		post.GET("/", handlers.IndexHandler)
	}
}

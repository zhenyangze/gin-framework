package users

import (
	"gitee.com/zhenyangze/gin-framework/app/middleware"
	"gitee.com/zhenyangze/gin-framework/app/modules/users/handlers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	// auth
	authMiddleware := middleware.AuthMiddleware

	router.POST("/login", authMiddleware.LoginHandler)
	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", handlers.UsersHandler)
	}

}

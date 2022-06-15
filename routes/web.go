package routes

import (
	"gitee.com/zhenyangze/gin-framework/app/middleware"
	"gitee.com/zhenyangze/gin-framework/app/modules/index"
	"gitee.com/zhenyangze/gin-framework/app/modules/posts"
	"gitee.com/zhenyangze/gin-framework/app/modules/users"
	websocket "gitee.com/zhenyangze/gin-framework/app/modules/websockets"
	"gitee.com/zhenyangze/gin-framework/helpers"
	"github.com/gin-gonic/gin"
	_ "go.uber.org/automaxprocs"
)

func Web(router *gin.Engine) {
	router.Use(middleware.DefaultRequestID())

	// 加载模板文件
	if helpers.IsDir("resources/views/") {
		router.LoadHTMLGlob("resources/views/**/*")
	}

	//service router
	index.Router(router)

	//users
	users.Router(router)

	//websocket
	websocket.Router(router)

	// posts
	posts.Router(router)
}

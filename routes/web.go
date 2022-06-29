package routes

import (
	"gitee.com/zhenyangze/gin-framework/app/bases"
	"gitee.com/zhenyangze/gin-framework/app/middleware"
	"gitee.com/zhenyangze/gin-framework/app/modules/index"
	"gitee.com/zhenyangze/gin-framework/app/modules/posts"
	"gitee.com/zhenyangze/gin-framework/app/modules/users"
	websocket "gitee.com/zhenyangze/gin-framework/app/modules/websockets"
	"gitee.com/zhenyangze/gin-framework/helpers"
)

func Web() {
	bases.Router.Use(middleware.DefaultRequestID())

	// 加载模板文件
	if helpers.IsDir("resources/views/") {
		bases.Router.LoadHTMLGlob("resources/views/**/*")
	}

	//service router
	index.Initer()

	//users
	users.Initer()

	//websocket
	websocket.Initer()

	// posts
	posts.Initer()
}

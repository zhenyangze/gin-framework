package routes

import (
	"gitee.com/zhenyangze/gin-framework/app/middleware"
	"gitee.com/zhenyangze/gin-framework/app/modules/index"
	"gitee.com/zhenyangze/gin-framework/app/modules/users"
	websocket "gitee.com/zhenyangze/gin-framework/app/modules/websockets"
	"github.com/gin-gonic/gin"
	_ "go.uber.org/automaxprocs"
)

func Web(router *gin.Engine) {
	router.Use(middleware.DefaultRequestID())
	// 加载模板文件
	router.LoadHTMLGlob("resources/views/**/*")

	//service router
	index.Router(router)

	//users
	users.Router(router)

	//websocket
	websocket.Router(router)
}

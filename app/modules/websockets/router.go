package websocket

import (
	"gitee.com/zhenyangze/gin-framework/app/modules/websockets/handlers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	router.GET("/ws", handlers.WebSocketHandle)
}

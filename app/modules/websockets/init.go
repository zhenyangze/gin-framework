package websocket

import (
	"gitee.com/zhenyangze/gin-framework/app/bases"
	"gitee.com/zhenyangze/gin-framework/app/modules/websockets/handlers"
)

func Router() {
	router := bases.Router
	router.GET("/ws", handlers.WebSocketHandle)
}

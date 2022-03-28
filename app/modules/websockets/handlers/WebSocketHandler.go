package handlers

import (
	"gitee.com/zhenyangze/gin-framework/app/bases"
	"github.com/gin-gonic/gin"
)

type MyWebSocket struct {
	bases.BaseWebsocket
}

func (this *MyWebSocket) OnMessage(messageType int, data []byte) []byte {
	return []byte("我收到你的消息了：" + string(data))
}

func WebSocketHandle(c *gin.Context) {
	webSocket := &MyWebSocket{}
	webSocket.Run(c, webSocket)
}

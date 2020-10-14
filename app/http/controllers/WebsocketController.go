// Package controllers provides ...
package controllers

import (
	"github.com/gin-gonic/gin"
)

type MyWebSocket struct {
	BaseWebsocket
}

func (this *MyWebSocket) OnMessage(messageType int, data []byte) []byte {
	return []byte("我收到你的消息了：" + string(data))
}

func WebSocketHandle(c *gin.Context) {
	webSocket := &MyWebSocket{}
	webSocket.Run(c, webSocket)
}

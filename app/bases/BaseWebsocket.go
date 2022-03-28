// Package controllers provides ...
package bases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IWebSocket interface {
	OnOpen()
	OnMessage(messageType int, data []byte) []byte
	OnClose()
}

type BaseWebsocket struct {
	Name string
	Ws   *websocket.Conn
	Err  error
	Self IWebSocket
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *BaseWebsocket) Run(ctx *gin.Context, self IWebSocket) {
	this.Ws, this.Err = upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	self.OnOpen()
	defer self.OnClose()
	if this.Err != nil {
		return
	}
	for {
		//读取Ws中的数据
		mt, message, err := this.Ws.ReadMessage()
		if err != nil {
			break
		}

		if string(message) == "ping" {
			message = []byte("pong")
		} else {
			// 解析message
			message = self.OnMessage(mt, message)
		}
		//写入Ws数据
		err = this.Ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}

}

func (this *BaseWebsocket) OnOpen() {

}

func (this *BaseWebsocket) OnMessage(messageType int, data []byte) []byte {
	return data
}

func (this *BaseWebsocket) OnClose() {
	this.Ws.Close()
}

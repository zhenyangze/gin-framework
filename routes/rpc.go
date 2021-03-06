// Package routes provides ...
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hprose/hprose-golang/rpc"
)

func hello1(name string) string {
	return "Hello " + name + "!"
}

func Rpc(router *gin.Engine) {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello1)
	// 加载模板文件
	router.Any("/path", func(c *gin.Context) {
		service.ServeHTTP(c.Writer, c.Request)
	})
}

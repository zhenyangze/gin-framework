// Package middleware provides ...
package middleware

import (
	"github.com/gin-gonic/gin"
)

func DemoMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

// Package middleware provides ...
package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type (
	// 跳过验证函数
	Skipper func(c *gin.Context) bool
)

// 默认不跳过验证
func DefaultSkipper(c *gin.Context) bool {
	return false
}

type RequestIDConfig struct {
	Skipper   Skipper
	Generator func() string
}

func DefaultRequestID() gin.HandlerFunc {
	return NewRequestID(RequestIDConfig{
		Skipper:   DefaultSkipper,
		Generator: defaultRequestIDGenerator,
	})
}

func NewRequestID(c RequestIDConfig) gin.HandlerFunc {
	if c.Skipper == nil {
		c.Skipper = DefaultSkipper
	}

	if c.Generator == nil {
		c.Generator = defaultRequestIDGenerator
	}

	return func(ctx *gin.Context) {
		if c.Skipper(ctx) {
			ctx.Next()
			return
		}

		HeaderXRequestID := "X-Request-ID"
		requestID := ctx.Request.Header.Get(HeaderXRequestID)
		if requestID == "" {
			requestID = c.Generator()
		}

		ctx.Writer.Header().Set(HeaderXRequestID, requestID)
		ctx.Set(HeaderXRequestID, requestID)
		ctx.Next()
	}
}

func defaultRequestIDGenerator() string {
	u := uuid.NewV4()
	return u.String()
}

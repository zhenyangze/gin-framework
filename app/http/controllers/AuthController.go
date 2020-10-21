package controllers

import (
	"gitee.com/zhenyangze/gin-framework/app/http/requests"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 兼容PHP int string混用情况
var identityKey = "id"

func UsersHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*requests.User).UserName,
		"text":     "Hello World.",
	})
}

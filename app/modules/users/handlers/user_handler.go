package handlers

import (
	"gitee.com/zhenyangze/gin-framework/app/middleware/requests"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

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

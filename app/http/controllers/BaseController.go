package controllers

import "github.com/gin-gonic/gin"

func Json(errorCode int64, errorMsg string, data interface{}) gin.H {
	return gin.H{
		"error_code": errorCode,
		"error_msg":  errorMsg,
		"data":       data,
	}
}

package bases

import "github.com/gin-gonic/gin"

func Json(errorCode int64, errorMsg string, data interface{}) gin.H {
	if data == nil {
		data = make(map[string]interface{})
	}
	return gin.H{
		"error_code": errorCode,
		"error_msg":  errorMsg,
		"data":       data,
	}
}
func JsonError(errorMsg string, data interface{}) gin.H {
	return Json(1, errorMsg, data)
}

func JsonOk(errorMsg string, data interface{}) gin.H {
	return Json(0, errorMsg, data)
}

package bases

import (
	"github.com/gin-gonic/gin"
)

type PageQuery struct {
	Page     int `uri:"page" form:"page"`
	PageSize int `uri:"page_size" form:"page_size"`
}

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

func Page(c *gin.Context) PageQuery {
	var query PageQuery
	c.ShouldBind(&query)
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 15
	}
	if query.PageSize > 100 {
		query.PageSize = 100
	}
	return query
}

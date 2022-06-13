package handlers

import (
	"net/http"

	"gitee.com/zhenyangze/gin-framework/app/bases"
	"github.com/gin-gonic/gin"
)

type PostHandler struct{}

func (h *PostHandler) IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{}))
}

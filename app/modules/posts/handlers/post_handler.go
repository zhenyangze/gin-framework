package handlers

import (
	"net/http"

	"gitee.com/zhenyangze/gin-framework/app/bases"
	"gitee.com/zhenyangze/gin-framework/app/modules/posts/models"
	"gitee.com/zhenyangze/gin-framework/app/providers"
	"gitee.com/zhenyangze/gin-framework/helpers"
	"github.com/gin-gonic/gin"
)

type postHandler struct{}

var PostHandler postHandler

func (h *postHandler) IndexHandler(c *gin.Context) {
	// 查询列表
	var total int64
	var postsList []models.Posts

	query := bases.Page(c)
	providers.DB.Scopes(helpers.Paginate(query.Page, query.PageSize)).Where(map[string]interface{}{
		//"status": query.Status,
	}).Order("id desc").Find(&postsList)
	providers.DB.Model(models.Posts{}).Where(map[string]interface{}{}).Count(&total)

	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{
		"data":      postsList,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	}))
}

func (h *postHandler) ShowHandler(c *gin.Context) {
	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{}))
}

func (h *postHandler) StoreHandler(c *gin.Context) {
	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{}))
}

func (h *postHandler) UpdateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{}))
}

func (h *postHandler) DeleteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{}))
}

func (h *postHandler) StoreBatchHandler(c *gin.Context) {
	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{}))
}

func (h *postHandler) UpdateBatchHandler(c *gin.Context) {
	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{}))
}

package handlers

import (
	"log"
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

	type QueryStruct struct {
		Page     int `uri:"page" form:"page" default:"1"`
		PageSize int `uri:"page_size" form:"page_size" default:"15"`
		Id       int `uri:"id" form:"id"`
		Status   int `uri:"status" form:"status" default:"1"`
	}

	var query QueryStruct
	err := c.ShouldBind(&query)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, bases.JsonError("参数错误", nil))
		return
	}
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 15
	}
	if query.PageSize > 100 {
		query.PageSize = 100
	}

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

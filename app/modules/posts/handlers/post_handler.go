package handlers

import (
	"net/http"

	"gitee.com/zhenyangze/gin-framework/app/bases"
	"gitee.com/zhenyangze/gin-framework/app/modules/posts/models"
	"gitee.com/zhenyangze/gin-framework/app/providers"
	"gitee.com/zhenyangze/gin-framework/helpers"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
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
	providers.DB.Model(models.Posts{}).Where(nil).Count(&total)

	c.JSON(http.StatusOK, bases.JsonOk("获取成功", map[string]interface{}{
		"data":      postsList,
		"total":     total,
		"page":      query.Page,
		"page_size": query.PageSize,
	}))
}

func (h *postHandler) ShowHandler(c *gin.Context) {
	id := c.Param("id")
	var postModel models.Posts
	providers.DB.First(&postModel, id)
	c.JSON(http.StatusOK, bases.JsonOk("获取成功", postModel))
}

func (h *postHandler) DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, bases.JsonError("参数异常", nil))
		return
	}
	providers.DB.Delete(models.Posts{}, id)
	c.JSON(http.StatusOK, bases.JsonOk("删除成功", map[string]interface{}{
		"id": helpers.StringToInt(id),
	}))
}

func (h *postHandler) StoreHandler(c *gin.Context) {
	var postModel models.Posts
	if err := c.ShouldBind(&postModel); err != nil {
		c.JSON(http.StatusOK, bases.JsonError("参数异常", err.Error()))
		return
	}
	// 可以指定字段更新
	providers.DB.Select("*").Create(&postModel)
	c.JSON(http.StatusOK, bases.JsonOk("添加成功", map[string]interface{}{
		"id": postModel.ID,
	}))
}

func (h *postHandler) UpdateHandler(c *gin.Context) {
	id := c.Param("id")
	var postModel models.Posts
	if err := c.ShouldBind(&postModel); err != nil {
		c.JSON(http.StatusOK, bases.JsonError("参数异常", err.Error()))
		return
	}

	providers.DB.Model(&postModel).Omit("id").Where("id = ?", id).Select("*").Updates(postModel)
	c.JSON(http.StatusOK, bases.JsonOk("更新成功", map[string]interface{}{
		"id": helpers.StringToInt(id),
	}))
}

/**
* func 批量更新
* [{"op":"replace","id":12, "value":{"title":"1212112","author":"ceshi"}},{"op":"add","value":{"title":"cehshi","author":"ceshi"}},{"op":"remove","id":"1","addition":{"status":1},"value":{"name":"cehshi","title":"ceshi"}}]
*
* @param
 */
func (h *postHandler) BatchHandler(c *gin.Context) {
	type batchItem struct {
		Op       string                 `json:"op" form:"op"`
		Id       string                 `json:"id" form:"id"`
		Addition map[string]interface{} `json:"Addition" form:"addition"`
		Value    map[string]interface{} `json:"value" form:"value"`
	}

	var batchData []batchItem

	if err := jsoniter.NewDecoder(c.Request.Body).Decode(&batchData); err != nil {
		c.JSON(http.StatusOK, bases.JsonError("参数异常", err.Error()))
		return
	}

	err := providers.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range batchData {
			switch v.Op {
			case "add":
				if err := tx.Model(&models.Posts{}).Create(v.Value).Error; err != nil {
					return err
				}
			case "replace":
				if v.Addition == nil {
					v.Addition = make(map[string]interface{})
				}
				v.Addition["id"] = v.Id
				if err := tx.Model(&models.Posts{}).Where(v.Addition).Updates(v.Value).Error; err != nil {
					return err
				}
			case "remove":
				if v.Addition == nil {
					v.Addition = make(map[string]interface{})
				}
				v.Addition["id"] = v.Id
				if err := tx.Model(&models.Posts{}).Where(v.Addition).Delete(&models.Posts{}).Error; err != nil {
					return err
				}
			default:

			}
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusOK, bases.JsonError("更新失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, bases.JsonOk("更新成功", nil))
}

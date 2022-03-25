package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	model "gitee.com/zhenyangze/gin-framework/app/models"
	"gitee.com/zhenyangze/gin-framework/app/providers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
)

// 兼容PHP int string混用情况
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func MyHandle(c *gin.Context) {
	/*providers.Logger().WithFields(logrus.Fields{
		"name": "hanyun",
	}).Info()*/
	c.String(200, "works!")
}

func TestHandle(c *gin.Context) {
	c.JSON(200, Json(200, "success", gin.H{}))
}

func ViewHandle(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"time": time.Now(),
		"name": c.DefaultQuery("name", "Gin"),
	})
}

func Md5Handle(c *gin.Context) {
	h := md5.New()
	str := c.Param("str")
	h.Write([]byte(str))
	c.JSON(200, hex.EncodeToString(h.Sum(nil)))
}

/**
* ValidHandle
*
* @param gin.Context
*
* @return
 */
func ValidHandle(c *gin.Context) {
	type UserInfo struct {
		// 需要大写
		Users    string `form:"user" binding:"required"`
		Password string `form:"password" binding:"required,min=6,max=12"`
	}

	var userinfo UserInfo
	err := c.ShouldBind(&userinfo)
	if err != nil {
		fmt.Printf("error")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	fmt.Printf("ok")
	c.JSON(http.StatusOK, userinfo)
}
func OrmHandle(c *gin.Context) {
	var user model.UserModel
	providers.DB.First(&user, 1)
	providers.DB.Model(&user).Update("name", "admin")
	c.JSON(200, user)
	return

}

func OrmWithCacheHandle(c *gin.Context) {
	var user model.UserModel

	cacheKey := "user_id_1"
	newValue, err := providers.Redis.Get(cacheKey).Result()
	if err == redis.Nil {
		// 不存在
		providers.DB.First(&user, 1)
		providers.DB.Model(&user).Update("name", "admin")
		cacheValue, err := json.Marshal(user)
		if err != nil {
			c.JSON(400, err)
			return
		}
		providers.Redis.Set(cacheKey, cacheValue, 20*time.Second)
	} else if err != nil {
		c.JSON(400, err)
		return
	}

	json.Unmarshal([]byte(newValue), &user)
	c.JSON(200, user)
	return

}

func RedisHandle(c *gin.Context) {
	str := "123"
	providers.Redis.Set("test_a", str, time.Second*10)
	str2, _ := providers.Redis.Get("test_a").Result()

	providers.Redis.LPush("redis_queue", 1, 2, 3, 4, 5, 6)
	for {
		_, err := providers.Redis.RPop("redis_queue").Result()
		if err != nil {
			//fmt.Println(err.Error())
			break
		}
		//fmt.Println(str3)
	}
	c.JSON(200, str2)
}

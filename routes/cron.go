package routes

import (
	"gitee.com/zhenyangze/gin-framework/app/jobs"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func Cron(router *gin.Engine) {
	c := cron.New()
	c.AddFunc("* * * * * *", jobs.MyJob)
	c.Start()
}

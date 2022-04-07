package routes

import (
	"gitee.com/zhenyangze/gin-framework/app/modules/jobs/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/cron"
)

func Cron(router *gin.Engine) {
	c := cron.New()
	c.AddFunc("* * * * * *", handlers.MyJob, "jobs.myjob")
	c.Start()
}

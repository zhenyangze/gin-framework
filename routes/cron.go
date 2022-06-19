package routes

import (
	"gitee.com/zhenyangze/gin-framework/app/modules/jobs/handlers"
	"github.com/jakecoffman/cron"
)

func Cron() {
	c := cron.New()
	c.AddFunc("* * * * * *", handlers.MyJob, "jobs.myjob")
	c.Start()
}

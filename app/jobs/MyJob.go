package jobs

import (
	"log"

	"gitee.com/zhenyangze/gin-framework/helpers"
)

func MyJob() {
	log.Println(helpers.GetCurrentDate())
}

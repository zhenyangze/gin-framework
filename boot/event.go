// Package boot provides ...
package boot

import (
	"gitee.com/zhenyangze/gin-framework/app/providers"
)

func InitEvent() {
	// 添加自定义监听事件
	providers.Event.SubscribeOnceAsync("main:init", func() {
		providers.InitAsynq()
	})
	providers.Event.Publish("main:init")

	//providers.EventBus.Subscribe()

}

// Package boot provides ...
package boot

import (
	"fmt"

	"gitee.com/zhenyangze/gin-framework/app/providers"
)

func InitEvent() {
	// 添加自定义监听事件
	providers.Event.SubscribeOnceAsync("main:init", func() {
		fmt.Println("init event")
	})
	//providers.EventBus.Subscribe()

}

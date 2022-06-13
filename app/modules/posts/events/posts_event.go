package events

import "gitee.com/zhenyangze/gin-framework/app/providers"

func IntEvent() {
	providers.Event.SubscribeAsync("post:event:demo", func() {
		//do some thing
	}, false)
}

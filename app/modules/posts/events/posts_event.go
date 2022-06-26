package events

import "gitee.com/zhenyangze/gin-framework/app/providers"

func IntEvent() {
	providers.Event.SubscribeAsync("post:event:demo", func() {
		//do some thing
	}, false)

	providers.Event.Subscribe("post:event:job", func(playload []byte) {
		providers.Info(playload)
	})
}

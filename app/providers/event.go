// Package providers provides ...
package providers

import "github.com/asaskevich/EventBus"

var (
	Event EventBus.Bus
)

func InitEvent() {
	Event = EventBus.New()
}

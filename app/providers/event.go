// Package providers provides ...
package providers

import "github.com/asaskevich/EventBus"

var (
	Event EventBus.Bus
)

func init() {
	Event = EventBus.New()
}

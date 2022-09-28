// Package providers provides ...
package providers

import (
	"runtime"

	"github.com/panjf2000/ants/v2"
)

var (
	Pool *ants.Pool
)

func InitPool() {
	numCPUs := runtime.NumCPU()
	Pool, _ = ants.NewPool(numCPUs)
}

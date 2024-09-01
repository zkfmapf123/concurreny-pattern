package utils

import (
	"fmt"
	"runtime"
)

func GetMetric() {
	var s runtime.MemStats

	runtime.ReadMemStats(&s)
	currentMemory := s.Alloc
	RedLogger(fmt.Sprintf("memory : %d", currentMemory))
}

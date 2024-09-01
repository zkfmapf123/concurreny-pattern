package concurrenypattern

import (
	"testing"
)

func Test_mutex_counter(t *testing.T) {
	counter := NewMutexCounter()

	counter.Add()
	counter.Add()
	counter.Add()
	counter.Add()
	counter.Add()
	counter.Add()
}

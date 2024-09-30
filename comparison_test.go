package concurrenypattern

import (
	"fmt"
	"testing"
)

var num = 100

// 0.709s
func Test_RWStream(t *testing.T) {

	consumer := NewConsumer()

	go WriteStream(consumer)
	ReadStream(consumer)
}

func Test_LimitConsumer(t *testing.T) {

	lc := NewLimitConsumer(2)

	go func() {

		lc.Write(1)
		lc.Write(2)
		lc.Write(3)
		lc.Write(4)
		lc.Write(5)
		lc.Done()
	}()

	lc.Read()
	fmt.Println(lc.Cap())
}

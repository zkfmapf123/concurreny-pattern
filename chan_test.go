package concurrenypattern

import (
	"fmt"
	"testing"
)

func Test_chan(t *testing.T) {

	num := 10

	cc := NewParallel().SetChannel(num)

	// go func(num int) {
	// 	for i := 0; i < 10; i++ {
	// 		go cc.Push(i)
	// 	}

	// 	cc.Done().Run(func(num int) {
	// 		fmt.Println(num)
	// 	})
	// }(num)

	for i := 0; i < 10; i++ {
		cc.Push(i)
	}

	cc.Done().Run(func(num int) {
		fmt.Println(num)
	})

	fmt.Println("hdaf")
}

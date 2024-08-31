package concurrenypattern

import (
	"fmt"
	"time"
)

func basicBadPattern() {
	var data int

	go func() {
		data++
	}()

	// sleep 을 쓴다는것자체가 문제가 될 수 있음
	time.Sleep(1 * time.Second)

	if data == 0 {
		fmt.Println("data is zero")
	} else {
		fmt.Printf("data is %d\n", data)
	}
}

func basicGoodPattern() {
	ch := make(chan int, 1)

	go func() {
		ch <- 1
	}()

	fmt.Printf("data is %d\n", <-ch)
}

package concurrenypattern

import (
	"fmt"
	"sync"
)

// 동기식 직압
func NotGoRoutine(msg string, fn func(msg string) string) {
	fmt.Println(fn(msg))
}

// 작업이 먼저 끝날 수 있다. (Bad)
func ProcessGoroutine(msg string, fn func(msg string) string) {

	go fmt.Println(fn(msg))
}

func ProcessGoroutineUseWaitGroup(wg sync.WaitGroup, msg string, fn func(msg string) string) {

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(fn(msg))
	}()
	wg.Wait()
}

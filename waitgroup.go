package concurrenypattern

import (
	"fmt"
	"sync"
)

// WaitGroup을 활용한 동시성
func waitGrouphelloworld(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Printf("hello world %d\n", i)
}

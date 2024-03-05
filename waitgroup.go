package main

import (
	"sync"
)

// WaitGroup
// 동시에 수행되는 연산결과를 기다리기위한 방법으로 사용
// 해당 기능이 부족할때는 -> channel, select로 구현을 해야 함...

// Add Counter를 기점으로 임계영역 구현...
func TestWaitGroup1(arr ...string) []string {
	var wg sync.WaitGroup
	res := make([]string, len(arr))

	for i, v := range arr {
		wg.Add(1)

		go func(i int, v string) {
			defer wg.Done()
			res[i] = v
		}(i, v)
	}
	wg.Wait()

	return res
}

func returnToInt(v int) int {
	return v
}

func TestWaitGroup2() []int {
	var num []int
	var wg sync.WaitGroup

	wg.Add(3) // 3

	go func() {
		defer wg.Done() // --3
		num = append(num, returnToInt(10))
	}()

	go func() {
		defer wg.Done() // --2
		num = append(num, returnToInt(20))
	}()

	go func() {
		defer wg.Done() // --1
		num = append(num, returnToInt(30))
	}()

	wg.Wait() // wait until wg count zero
	return num
}

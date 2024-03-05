package main

import "sync"

// 임계영역을 활용하여 공유리소스의 대한 접근 제어
// WaitGroup으로는 순서가 지켜지지 않는다
// mutex + waitGroup을 사용하여 순서를 지켜줄수 있음
// RWMutex는 기존 Mutex와 다르게 메모리를 조금더 제어할수있다.

// 순서보장 X
func LoopUseWaitGroup(len int) []int {
	var wg sync.WaitGroup
	resource, resArr := 0, []int{}

	wg.Add(len)
	for i := 0; i < len; i++ {
		go func() {
			defer wg.Done()
			resArr = append(resArr, resource)
			resource++
		}()
	}
	wg.Wait()
	return resArr
}

func LoopUseMutex(len int) []int {
	var wg sync.WaitGroup
	var m sync.Mutex

	resource, resArr := 0, []int{}

	wg.Add(len)
	for i := 0; i < len; i++ {
		go func() {
			m.Lock()
			defer func() {
				wg.Done()
				m.Unlock()
			}()

			resArr = append(resArr, resource)
			resource++
		}()
	}

	wg.Wait()
	return resArr
}

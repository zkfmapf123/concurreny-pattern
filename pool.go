package main

import (
	"math"
	"sync"
)

var instnaceCount = 0

// 인스턴스진행됨
var myPool = &sync.Pool{
	New: func() interface{} {
		instnaceCount++
		return struct{}{}
	},
}

// Get 인스턴스로 만든 후
// Put 만든 인스턴스를 반납하면
// 인스턴스상태를 계속사용할 수 있음...
func UsePool() int {
	// 1번 호출
	ins := myPool.Get() // 인스턴스 호출
	myPool.Put(ins)     // 인스턴스를 다시 되돌려놓음

	for i := 0; i < 5; i++ {
		callMyPool(myPool)
	}

	return instnaceCount
}

func callMyPool(m *sync.Pool) {
	// 기존 인스턴스 계속 만듬
	ins := m.Get()
	m.Put(ins)
}

var numCalcsCreated = 0
var calcPool = &sync.Pool{
	New: func() interface{} {
		numCalcsCreated++
		mem := make([]byte, 1024)
		return &mem
	},
}

// 미리 4KB Pool을 생성해서 병렬작업에 대기
// 더 생성될 작업들의 대해서 작업수를 조절할 수 있다...
func UsePoolMem() int {
	// 1KB * 4 = 4KB 생성
	for i := 0; i < 4; i++ {
		calcPool.Put(calcPool.New())
	}

	var wg sync.WaitGroup
	numWorkers := int(math.Pow(1024, 2))
	wg.Add(numWorkers)

	// heavy job
	for i := numWorkers; i > 0; i-- {
		go func() {
			mem := calcPool.Get()
			defer func() {
				wg.Done()
				calcPool.Put(mem)
			}()
		}()
	}

	wg.Wait()
	return numCalcsCreated

}

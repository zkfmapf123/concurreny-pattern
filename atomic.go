package concurrenypattern

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
## Mutex
sync.Mutex
sync.RWMutext 읽기, 쓰기 동시접근 허용 (읽기가 많을 경우)

## Atomic
CPU 하드웨어 레벨 연산 (훨씬 빠름)
*/
type SafetyCount struct {
	key   sync.Mutex
	count int64
}

func (s *SafetyCount) replace(i int) {
	// Mutex
	// s.key.Lock()
	// defer s.key.Unlock()

	// s.count = i

	// Atomic
	atomic.StoreInt64(&s.count, int64(i))
}

func (s *SafetyCount) Read() int {

	// Mutex
	// s.key.Lock()
	// defer s.key.Unlock()

	// return s.count

	// Atomic
	return int(atomic.LoadInt64(&s.count))
}

func atomicExecute() {

	c := SafetyCount{
		count: 0,
	}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			c.replace(idx + 1)
		}(i)

	}

	wg.Wait()
	fmt.Printf("res : %+v\n", c.Read())

}

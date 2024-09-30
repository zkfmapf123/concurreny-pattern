# Golang Concurreny-Pattern

## WaitGroup

- 동시에 수행된 연산의 결과를 신경쓰지 않는 경우
- 수행될 연산집합을 기다릴 경우 

## ...
- [간단한_병렬처리_패턴](./basic_pattern.go)
- [뮤텍스를_이용한_카운터](./mutex_counter.go)
- [고루틴_기초](./goroutine.go)

## Benchmark
- [비교](./comparison.go)

## ...

```sh
    ## 특정 테스트만...
    go test -v -run Test_chan
```

## Chan Tips

### Channel이 Close 되는 구간은 작업이 끝나는 메서드 내에서 진행되어야 함

```go
package concurrenypattern

import "fmt"

func createConsumer(cap int) chan int {

	return make(chan int, cap)
}

// i에 값들이 push 된 후에 Channel은 작업이 완료되었음..
func consumer(ch chan<- int, cap int) {

	defer close(ch)  // 작업이 끝나면 닫아야함... Must Be
	for i := 0; i < cap; i++ {
		ch <- i
	}

}

func execute() {
	c := createConsumer(5)
	consumer(c, 5)

	for val := range c {
		fmt.Println(val)
	}

}
```

### Select

- 동시에 진행되는 것이 아님...
- 순서를 보장하지 않지만 동시성을 제어함...

```go
func selectRace() {

	r1, r2 := make(chan string), make(chan string)
	close(r1)
	close(r2)

	r1Count, r2Count := 0, 0

	for i := 0; i < 1000; i++ {

		select {
		case <-r1:
			r1Count++
		case <-r2:
			r2Count++
		}
	}

	fmt.Printf("Race r1 : %d\tr2 : %d\n", r1Count, r2Count)
}

```
package concurrenypattern

import "fmt"

type chanParams struct {
	ch chan int
}

func NewLimitConsumer(num int) chanParams {
	return chanParams{
		ch: make(chan int, num),
	}
}

func (lc *chanParams) Write(num int) {
	lc.ch <- num
}

func (lc *chanParams) Done() {
	close(lc.ch)
}

func (lc *chanParams) Read() {

	for val := range lc.ch {
		fmt.Println(val)
	}
}

func (lc *chanParams) Cap() int {
	return cap(lc.ch)
}

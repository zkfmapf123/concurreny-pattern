package concurrenypattern

import "fmt"

type ChanParams struct {
	ch   chan int
	size int
	done chan string
}

func NewParallel() *ChanParams {

	return &ChanParams{
		size: 0,
		done: make(chan string),
	}
}

func (c *ChanParams) SetChannel(size int) *ChanParams {
	c.ch = make(chan int, size)
	return c
}

func (c *ChanParams) Push(num int) {
	c.ch <- num
	c.size++
}

func (c *ChanParams) Done() *ChanParams {
	close(c.ch)
	c.done <- "done"
	close(c.done)

	return c
}

func (c *ChanParams) Run(fn func(num int)) {

	for msg := range c.ch {
		fn(msg)
	}

	<-c.done
	fmt.Println("Done...")
}

func (c *ChanParams) Peek() {

}

func (c *ChanParams) Size() int {
	return c.size
}

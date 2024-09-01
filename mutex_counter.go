package concurrenypattern

import "sync"

type Counter struct {
	mu  sync.Mutex
	Num int // Public
}

func NewMutexCounter() Counter {
	return Counter{
		mu:  sync.Mutex{},
		Num: 0,
	}
}

func (c *Counter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Num++
}

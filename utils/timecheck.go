package utils

import (
	"fmt"
	"time"
)

type Chekcer struct {
	t       time.Time
	elapsed time.Duration
}

func NewTimeCheck() Chekcer {

	return Chekcer{}
}

func (c *Chekcer) Start() {
	c.t = time.Now()
}

func (c *Chekcer) Finish() {
	c.elapsed = time.Since(c.t)
}

func (c *Chekcer) Result() {
	red := "\033[31m"
	reset := "\033[0m"

	fmt.Printf("%sTime Duration: %s%s\n", red, c.elapsed, reset)
}

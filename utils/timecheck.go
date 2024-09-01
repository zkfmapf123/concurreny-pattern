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

func (c *Chekcer) GetTime() time.Duration {
	return c.elapsed
}

func (c *Chekcer) Result(msg string) {

	RedLogger(fmt.Sprintf("%s : %s", msg, c.elapsed))
}

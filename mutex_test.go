package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseLoopWaitGroup(t *testing.T) {
	res1 := LoopUseWaitGroup(10)
	res2 := LoopUseMutex(10)

	assert.Equal(t, len(res1) <= 10, true)
	assert.Equal(t, len(res2), 10)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsePool(t *testing.T) {
	count := UsePool()
	assert.Equal(t, count, 1)
}

func TestUseMem(t *testing.T) {
	count := UsePoolMem()
	assert.Equal(t, count < 20, true)
}

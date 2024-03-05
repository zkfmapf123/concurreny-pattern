package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseChannel(t *testing.T) {

	num := UseChannel(10)
	assert.Equal(t, num, 10)
}

func TestUseMutex(t *testing.T) {
	num := UseMutex(10)
	assert.Equal(t, num, 10)
}

func TestUseWaitGroup(t *testing.T) {
	num := UseWaitGroup(10)
	assert.Equal(t, num, 10)
}

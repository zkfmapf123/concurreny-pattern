package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaitGroupTest1(t *testing.T) {
	arr := TestWaitGroup1("a", "b", "c")

	assert.Equal(t, arr, []string{"a", "b", "c"})
}

func TestWaitGroupTest2(t *testing.T) {
	arr := TestWaitGroup2()
	fmt.Println(arr)
}

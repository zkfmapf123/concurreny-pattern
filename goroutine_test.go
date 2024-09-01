package concurrenypattern

import (
	"fmt"
	"sync"
	"testing"
	"zkfmapf123/concurrency/utils"
)

func Test_NotGoRoutine(t *testing.T) {

	timer.Start()

	NotGoRoutine("hello world-1", func(msg string) string { return msg })
	NotGoRoutine("hello world-2", func(msg string) string { return msg })
	NotGoRoutine("hello world-3", func(msg string) string { return msg })
	NotGoRoutine("hello world-4", func(msg string) string { return msg })
	NotGoRoutine("hello world-5", func(msg string) string { return msg })

	timer.Finish()
	timer.Result("NotGoRoutine")
}

func Test_GoRoutine(t *testing.T) {

	timer.Start()

	for i := 0; i < 10; i++ {
		ProcessGoroutine(fmt.Sprintf("%s-%d", "hello world-", i), func(msg string) string { return msg })
	}

	// time.Sleep(1 * time.Second)

	timer.Finish()
	timer.Result("GoRoutine")
}

func Test_WatiGroupGoRoutine(t *testing.T) {
	timer.Start()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		ProcessGoroutineUseWaitGroup(wg, fmt.Sprintf("%s-%d", "hello world-", i), func(msg string) string { return msg })
	}
	timer.Finish()
	utils.GetMetric()
	timer.Result("GoRoutine")
}

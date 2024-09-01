package concurrenypattern

import (
	"sync"
	"testing"
	"zkfmapf123/concurrency/utils"
)

var num = 100

func Test_hello(t *testing.T) {

	tc := utils.NewTimeCheck()
	tc.Start()

	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go waitGrouphelloworld(&wg, i)
	}

	wg.Wait()
	tc.Finish()
	tc.Result("WatiGroup Hello world")
}

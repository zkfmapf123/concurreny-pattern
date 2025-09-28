package concurrenypattern

import (
	"testing"
	"time"
)

func Test_bufferedExec(t *testing.T) {

	go bufferedExecute()
	time.Sleep(time.Second * 60)
}

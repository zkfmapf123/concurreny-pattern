package concurrenypattern

import (
	"testing"
	"time"
)

func Test_mainExecute(t *testing.T) {
	go mainExecute()

	time.Sleep(time.Second * 60)
}

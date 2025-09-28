package concurrenypattern

import (
	"testing"
	"time"
)

func Test_selectPatternExecute(t *testing.T) {
	go selectPatternExecute()
	time.Sleep(time.Second * 60)
}

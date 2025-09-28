package concurrenypattern

import (
	"testing"
)

func Test_mainExecute(t *testing.T) {

	done := make(chan bool)

	go func() {
		mainExecute()
		done <- true
	}()

	<-done
}

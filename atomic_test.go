package concurrenypattern

import "testing"

func Test_atomicExecute(t *testing.T) {

	done := make(chan bool)

	go func() {
		atomicExecute()
		done <- true
	}()

	<-done
}

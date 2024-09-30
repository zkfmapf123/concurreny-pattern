package concurrenypattern

import (
	"fmt"
	"testing"
	"time"
)

func timePrint(num int) {

	time.Sleep(2 * time.Second)
	fmt.Println("Values : ", num)
}

func Test_UsefulChan(t *testing.T) {

	// execute()

	selectRace()
}

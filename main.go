package main

import (
	"fmt"
)

var data int

func main() {
	num := make(chan (int), 1)

	go func() {
		num <- 10
	}()

	fmt.Printf("%d is num", <-num)
}

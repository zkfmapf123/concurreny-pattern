package concurrenypattern

import (
	"fmt"
)

var count = 0

func KKString(ch chan<- string, str ...string) {

	defer close(ch)
	for i := 0; i < len(str); i++ {
		ch <- str[i]
	}
}

func execute() {
	var c1, c2 chan string
	var c3 chan string

	c1 = make(chan string, 2)
	c2 = make(chan string, 3)
	c3 = make(chan string, 1)

	go KKString(c1, "A", "B", "C", "D")
	go KKString(c2, "X", "X", "X", "O")

	for {
		select {

		case cValue := <-c2:
			count++
			name := <-c1

			if cValue == "O" {
				fmt.Println("I Love you ", name)
				c3 <- "done"
			}

		case <-c3:
			fmt.Println("System Done >> Count : ", count)
			return
		}
	}
}

func selectRace() {

	r1, r2 := make(chan string), make(chan string)
	close(r1)
	close(r2)

	r1Count, r2Count := 0, 0

	for i := 0; i < 1000; i++ {

		select {
		case <-r1:
			r1Count++
		case <-r2:
			r2Count++
		}
	}

	fmt.Printf("Race r1 : %d\tr2 : %d\n", r1Count, r2Count)
}

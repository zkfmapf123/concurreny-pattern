package concurrenypattern

import (
	"fmt"
	"time"
)

func bufferedExecute() {

	num := 5
	buf := make(chan string, 1)
	unBuf := make(chan string)

	fmt.Println("buffered")
	go buffered(buf, num)

	time.Sleep(time.Second * 5)

	fmt.Println("unbuffered")
	go buffered(unBuf, num)

}

func buffered(ch chan string, num int) {

	for i := 0; i < num; i++ {
		go func(idx int) {
			ch <- response(idx + 1)
		}(i)
	}

	for i := 0; i < num; i++ {
		fmt.Println(<-ch)
	}

}

func response(num int) string {
	time.Sleep(2 * time.Second)
	fmt.Println("....end")
	return fmt.Sprintf("Hello world %d", num)
}

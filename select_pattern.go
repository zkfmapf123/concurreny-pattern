package concurrenypattern

import (
	"fmt"
	"time"
)

func selectPatternExecute() {
	msgCh := make(chan string, 128)

	for i := 0; i < 3; i++ {
		go func(idx int) {
			msgCh <- fetchResponse(i)
		}(i)
	}

	count := 1
	for {
		select {
		case v := <-msgCh:
			fmt.Println("the messages -> ", v)
			count++

			if count == 3 {
				goto done
			}
		}
	}
done:

	fmt.Println("done")
}

func fetchResponse(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}

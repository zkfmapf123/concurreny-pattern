package concurrenypattern

import "fmt"

/*
	데이터 읽기	<-  chan <- 데이터 쓰기
	타입강제...
*/

// var (
// 	writeStream chan<- interface{} // Write only
// 	readStream  <-chan interface{} // Read only
// )

func NewConsumer() chan int {
	return make(chan int)
}

func WriteStream(consumer chan<- int) {

	for i := 0; i < 10; i++ {
		consumer <- i
	}

	close(consumer)
}

func ReadStream(consumer <-chan int) {

	for v := range consumer {
		fmt.Printf("consumer : %d\n", v)
	}
}

// WO 인데 Read 하는 경우
/*
	cannot range over consumer (variable of type chan<- int): receive from send-only channel
	타입 오류 발생
*/
// func BadReadStream(consumer chan<- int) {
// 	for v := range consumer {
// 		fmt.Printf("consumer : %d\n", v)
// 	}
// }

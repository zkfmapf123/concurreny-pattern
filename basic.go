package main

import "sync"

var data = 0

// Problem (동시성 문제..)
// go func() {
// 	data = v
// }()
// return data

// channle을 사용하셔 동시성 해결
func UseChannel(v int) int {

	num := make(chan int)

	go func() {
		data = v
		num <- data
	}()

	return <-num
}

// 임계영역을 활용하여 해결 -> System 자체가 느려질수 있음..
func UseMutex(v int) int {
	var mutex sync.Mutex

	mutex.Lock()
	data = v
	mutex.Unlock()

	return data
}

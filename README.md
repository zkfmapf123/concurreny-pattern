# Golang Concurrency Pattern

## Files 

- [buffered](./buffered.go)
- [select_pattern](./select_pattern.go)
- [atomic](./atomic.go)
- [context](./main.go)

## Bufferd , Unbuffered

- Buffered - 송신자가 수신자를 기다리지 않음 (생산 > 소비)
- UnBuffered - 송신자가 수신자를 기다림...

```go
## unbuffered example
func main() {
    ch := make(chan string) // 언버퍼드
    
    go func() {
        fmt.Println("고루틴: 데이터 보내는 중...")
        ch <- "안녕"  // 여기서 대기! (받는 쪽이 받을 때까지)
        fmt.Println("고루틴: 데이터 전송 완료!")
    }()
    
    time.Sleep(3 * time.Second) // 3초 기다림
    fmt.Println("메인: 이제 받을게!")
    data := <-ch
    fmt.Println("메인: 받은 데이터:", data)
}

고루틴: 데이터 보내는 중...
(3초 대기...)
메인: 이제 받을게!
고루틴: 데이터 전송 완료!
메인: 받은 데이터: 안녕
	
```

```go
## buffered example
func main() {
    ch := make(chan string, 1) // 버퍼 크기 1
    
    go func() {
        fmt.Println("고루틴: 데이터 보내는 중...")
        ch <- "안녕"  // 버퍼에 저장하고 바로 진행!
        fmt.Println("고루틴: 데이터 전송 완료!")
    }()
    
    time.Sleep(3 * time.Second) // 3초 기다림
    fmt.Println("메인: 이제 받을게!")
    data := <-ch
    fmt.Println("메인: 받은 데이터:", data)
}

고루틴: 데이터 보내는 중...
고루틴: 데이터 전송 완료!  ← 바로 완료!
(3초 대기...)
메인: 이제 받을게!
메인: 받은 데이터: 안녕
```
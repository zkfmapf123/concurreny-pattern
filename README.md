# Concurrency Pattern

- [동시성해결_간단](./basic.go)
- [WaitGroup](./waitgroup.go)
  - WaitGroup Counter를 활용하여 동시성제어
  - Add, Done, Wait 메서드를 사용해서 간단하게 동시성 구현

- [Mutex](./mutex.go) 
  - 임계영역을 사용하여, 순서를 보장할 수 있음 (안정성)

- [***Pool](./pool.go)
  - 풀을 생성하고 활용할 수 있게 해주는 방법 (Cache...)
  - Get, Put을 사용하여 인스턴스를 생성하고 반납하는 패턴을 활용
  - 데이터베이스연결과 같은 비용이 많이드는 연결의 대한 효율적인 패턴을 사용할 수 있음
  - 성능개선 목적으로 사용될 수 있음

  
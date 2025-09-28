run:
	clear
	go run main.go

build:
	go build -o main

test:
	go test ./... -v

unit-test:
	go test -run <Method>

main-test:
	go test -run Test_mainExecute

select-test:
	go test -run Test_selectPatternExecute

benchmakr-test:
	go test -bench=. -cpu=1 <path>
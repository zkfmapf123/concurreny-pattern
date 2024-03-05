run:
	clear
	go run main.go

build:
	go build -o main

test:
	go test ./... -v

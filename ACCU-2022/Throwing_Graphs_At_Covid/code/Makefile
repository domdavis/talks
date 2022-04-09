all: clean build lint test

clean:
	go clean

build:
	go vet ./...
	go build

lint:
	golangci-lint run

test:
	go test -covermode=count -count=1 ./...

run:
	go run main.go

.PHONY: test
.DEFAULT_GOAL := all

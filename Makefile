.PHONY: all install test

all:
	go build

install:
	go install

test:
	go test -race -cover ./...

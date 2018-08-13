.PHONY: all install lint test

all:
	go build

install:
	go install

lint:
	golangci-lint run --enable-all --disable lll

test:
	go test -race -cover ./...

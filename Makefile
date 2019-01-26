.PHONY: all install lint test clean

.EXPORT_ALL_VARIABLES:
GO111MODULE = on

all:
	go build -o bin/latest

install:
	go install

lint:
	golangci-lint run

test:
	go test -race -cover ./...

clean:
	rm -rf bin

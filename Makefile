.PHONY: build
build:
	go build -o bin/latest

.PHONY: install
install:
	go install

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: clean
clean:
	rm -rf bin

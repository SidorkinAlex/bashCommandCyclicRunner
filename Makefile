.PHONY: build
build:
	go build -v ./cmd/bashCommandCyclicRunner/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

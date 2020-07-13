.PHONY: build
build:
	go build -v ./cmd/bot

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: dev
dev:
	go run ./cmd/bot


.DEFAULT_GOAL := build
.PHONY: docs
docs:
	./scripts/make_swager_docs.sh

.PHONY: gitpush
gitpush:
	./scripts/git_push.sh


.PHONY: build
build:
	go build -v ./cmd/bot

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: dev
dev:
	go run main.go

.DEFAULT_GOAL := build
.PHONY: docs
docs:
	./scripts/make_swager_docs.sh

.PHONY: gitpush
gitpush:
	./scripts/git_push.sh

.PHONY: migrate
migrate:
	go run ./cmd/db



.PHONY: build
build:
	go build -v main.go

.PHONY: dev
dev:
	go run ./cmd/api





.PHONY: test
test:
	go test -v -race -timeout 30s ./...


.DEFAULT_GOAL := build
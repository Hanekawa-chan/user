PROJECT_NAME := kanji-user
PROJECT := github.com/Hanekawa-chan/kanji-user

VERSION := $(shell git describe --tags)
COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := "-s -w -X $(PROJECT)/internal/version.Version=$(VERSION) -X $(PROJECT)/internal/version.Commit=$(COMMIT)"
build:
	CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o ./bin/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)

test:
	@go test -v -cover -gcflags=-l --race ./...

GOLANGCI_LINT_VERSION := v1.24.0
lint:
	@golangci-lint run -v

dep:
	@go mod download


models_linux:
	cp -f kanji-go-models/*.go internal/services/models/

models_win:
	copy kanji-go-models\*.go internal\services\models\ /Y

modules_init:
	git submodule add https://github.com/Hanekawa-chan/kanji-go-models

modules_update:
	git submodule update --remote

update_modules_win: modules_update models_win

update_modules_linux: modules_update models_linux
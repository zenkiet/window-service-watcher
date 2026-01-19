APP_NAME=z-attendance

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOPATH=$(shell go env GOPATH)

.PHONY: all build clean run test lint fmt deps

all: build

deps:
	@echo " > Installing dependencies..."
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0
	go install mvdan.cc/gofumpt@latest
	go install github.com/evilmartians/lefthook@latest

	@echo " > Setting up git hooks..."
	lefthook install

build:
	@echo " > Building binary..."
	go build -ldflags="-s -w" -o $(GOBIN)/$(APP_NAME) cmd/api/main.go

run:
	go run cmd/api/main.go

format:
	@echo " > Formatting code..."
	gofumpt -w .

lint:
	@echo " > Linting..."
	golangci-lint run ./...
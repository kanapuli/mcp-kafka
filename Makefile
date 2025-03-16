VERSION := $(shell git describe --tags --always --dirty)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: build tidy clean run help

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.version=$(VERSION)" -o bin/mcp-kafka-$(GOOS)-$(GOARCH) *.go

tidy:
	go mod tidy

clean:
	rm -rf bin/*

run: build
	./bin/mcp-kafka

help:
	@echo "Usage:"
	@echo "  make build    Build the binary"
	@echo "  make tidy     Run go mod tidy"
	@echo "  make clean    Remove the binary"
	@echo "  make run      Run the binary"
	@echo "  make help     Show this help message"

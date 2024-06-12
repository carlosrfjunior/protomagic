# Makefile

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: help
all: help

.PHONY: run
run:
	go run $(GOBASE)/cmd/protomagic

.PHONY: build
build:
	go build -v -ldflags="-X 'main.Version=v1.0.0' -X 'main.User=$(shell id -u -n)' -X 'main.Time=$(shell date)'" $(GOBASE)/cmd/protomagic

.PHONY: version
version:
	go run -ldflags="-X 'main.Version=v1.0.0' -X 'main.User=$(shell id -u -n)' -X 'main.Time=$(shell date)'" $(GOBASE)/cmd/protomagic version

.PHONY: help
help: Makefile
	@echo
	@echo "Usage: make [options]"
	@echo
	@echo "Options:"
	@echo "		build	Create binary file"
	@echo "		Help	"
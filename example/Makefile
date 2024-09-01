# Makefile

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: help
all: help

.PHONY: dev
dev:
	docker compose up --remove-orphans

.PHONY: run
run:
	go run $(GOBASE)/cmd/protomagic

.PHONY: build
build:
	go build -v -ldflags="-X 'main.Version=1.0.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/protomagic

.PHONY: version
version:
	go run -ldflags="-X 'main.version=1.0.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/protomagic version

.PHONY: help
help: Makefile
	@echo
	@echo "Usage: make [options]"
	@echo
	@echo "Options:"
	@echo "    build     Create binary file"
	@echo "    run       Run protomagic"
	@echo "    version   Set version in go application"
	@echo "    Help	"

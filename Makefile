SHELL := /bin/bash

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test

BINARY_NAME=mta-hosting-optimizer
EXEC_ROOT=./cmd/mta-hosting-optimizer

.PHONY: build
build:
	$(GOBUILD) -tags "$(TAGS)" -o $(BINARY_NAME) -v $(EXEC_ROOT)

.PHONY: help
help:
	$(GORUN) $(EXEC_ROOT) help

.PHONY: run
run:
	mkdir -p storage/
	$(GORUN) $(EXEC_ROOT) run

.PHONY: test
test:
	$(GOTEST) -v -cover --shuffle=on ./...

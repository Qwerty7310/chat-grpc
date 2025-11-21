.PHONY: all generate build run clean test help

BINARY_NAME=chat-grpc
PROTO_DIR=api
PB_DIR=$(PROTO_DIR)/pb
SCRIPTS_DIR=scripts

all: build

generate:
	@$(SCRIPTS_DIR)/gen-proto.sh

build: generate
	@go build -o bin/$(BINARY_NAME) .

run: build
	@./bin/$(BINARY_NAME)
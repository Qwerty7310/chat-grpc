#!/bin/bash

# Скрипт генерации gRPC кода из .proto файлов

echo "Protoc version: $(protoc --version)"
echo "Generating gRPC code..."

rm -f api/pb/*
mkdir -p api/pb

protoc --go_out=. \
       --go_opt=module=github.com/Qwerty7310/chat-grpc \
       --go-grpc_out=. \
       --go-grpc_opt=module=github.com/Qwerty7310/chat-grpc \
       api/chat.proto

echo "Generating complete."
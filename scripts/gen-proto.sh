#!/bin/bash

set -e

echo "===== Generating protobuf files ====="
echo "Protoc version: $(protoc --version)"

# Paths
PROTO_DIR="api"
GO_OUT_DIR="api/pb"
WEB_OUT_DIR="client/src/generated"

# CLEAN PREVIOUS OUTPUT
echo "Cleaning previous generated files..."
rm -rf ${GO_OUT_DIR}/*
rm -rf ${WEB_OUT_DIR}/*

mkdir -p ${GO_OUT_DIR}
mkdir -p ${WEB_OUT_DIR}

echo "===== Generating Go server code ====="

protoc \
  --go_out=. \
  --go_opt=module=github.com/Qwerty7310/chat-grpc \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/Qwerty7310/chat-grpc \
  ${PROTO_DIR}/chat.proto


echo "===== Generating gRPC-Web client (JS + TS) ====="

# IMPORTANT:
# grpc_tools_node_protoc MUST be used here, not plain protoc!
grpc_tools_node_protoc \
  --proto_path=${PROTO_DIR} \
  --js_out=import_style=commonjs:${WEB_OUT_DIR} \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:${WEB_OUT_DIR} \
  ${PROTO_DIR}/chat.proto

echo "===== Generation complete! ====="
echo "Go code → ${GO_OUT_DIR}"
echo "Web client → ${WEB_OUT_DIR}"

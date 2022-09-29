#!/bin/sh

set -e

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/../"

cd "${PROJECTPATH}"

go install \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  github.com/ethereum/go-ethereum/cmd/geth

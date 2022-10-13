#!/bin/sh

set -ex

OUTPUT=${OUTPUT:-"./build/grid"}

mkdir -p "$(dirname "$OUTPUT")"
GOOS=darwin GOARCH=amd64 go build -o "${OUTPUT}-darwin-amd64" ./cmd
GOOS=darwin GOARCH=arm64 go build -o "${OUTPUT}-darwin-arm64" ./cmd
GOOS=freebsd GOARCH=amd64 go build -o "${OUTPUT}-freebsd-amd64" ./cmd
GOOS=freebsd GOARCH=arm64 go build -o "${OUTPUT}-freebsd-arm64" ./cmd
GOOS=linux GOARCH=amd64 go build -o "${OUTPUT}-linux-amd64" ./cmd
GOOS=linux GOARCH=arm64 go build -o "${OUTPUT}-linux-arm64" ./cmd
GOOS=linux GOARCH=mips64 go build -o "${OUTPUT}-linux-mips64" ./cmd
GOOS=linux GOARCH=mips64le go build -o "${OUTPUT}-linux-mips64le" ./cmd
GOOS=linux GOARCH=ppc64 go build -o "${OUTPUT}-linux-ppc64" ./cmd
GOOS=linux GOARCH=ppc64le go build -o "${OUTPUT}-linux-ppc64le" ./cmd
GOOS=linux GOARCH=riscv64 go build -o "${OUTPUT}-linux-riscv64" ./cmd
GOOS=linux GOARCH=s390x go build -o "${OUTPUT}-linux-s390x" ./cmd
GOOS=windows GOARCH=amd64 go build -o "${OUTPUT}-windows-amd64".exe ./cmd

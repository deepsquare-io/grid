#!/bin/sh

set -ex

OUTPUT=${OUTPUT:-"./build/provider-ssh-authorized-keys"}

mkdir -p "$(dirname "$OUTPUT")"
GOOS=darwin GOARCH=amd64 go build -o "${OUTPUT}-darwin-amd64" ./
GOOS=darwin GOARCH=arm64 go build -o "${OUTPUT}-darwin-arm64" ./
GOOS=freebsd GOARCH=amd64 go build -o "${OUTPUT}-freebsd-amd64" ./
GOOS=freebsd GOARCH=arm64 go build -o "${OUTPUT}-freebsd-arm64" ./
GOOS=linux GOARCH=amd64 go build -o "${OUTPUT}-linux-amd64" ./
GOOS=linux GOARCH=arm64 go build -o "${OUTPUT}-linux-arm64" ./
GOOS=linux GOARCH=mips64 go build -o "${OUTPUT}-linux-mips64" ./
GOOS=linux GOARCH=mips64le go build -o "${OUTPUT}-linux-mips64le" ./
GOOS=linux GOARCH=ppc64 go build -o "${OUTPUT}-linux-ppc64" ./
GOOS=linux GOARCH=ppc64le go build -o "${OUTPUT}-linux-ppc64le" ./
GOOS=linux GOARCH=riscv64 go build -o "${OUTPUT}-linux-riscv64" ./
GOOS=linux GOARCH=s390x go build -o "${OUTPUT}-linux-s390x" ./
GOOS=windows GOARCH=amd64 go build -o "${OUTPUT}-windows-amd64".exe ./

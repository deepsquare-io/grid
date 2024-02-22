GO_SRCS := $(shell find . -type f -name '*.go' -a ! \( -name 'zz_generated*' -o -name '*_test.go' \))
GO_TESTS := $(shell find . -type f -name '*_test.go')
TAG_NAME = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
TAG_NAME_DEV = $(shell git describe --tags --abbrev=0 2>/dev/null)
GIT_COMMIT = $(shell git rev-parse --short=7 HEAD)
VERSION = $(or ${TAG_NAME},$(TAG_NAME_DEV)-dev)
ifeq ($(golint),)
golint := $(shell go env GOPATH)/bin/golangci-lint
endif

bins := grid-logger-reader-linux-amd64 grid-logger-writer-linux-amd64 grid-logger-writer-linux-arm64 grid-logger-writer-linux-riscv64 grid-logger-server-linux-amd64 grid-logger-server-linux-arm64 grid-logger-server-linux-riscv64

.PHONY: all
bin: bin/grid-logger-reader bin/grid-logger-writer bin/grid-logger-server

bin/grid-logger-reader:
	CGO_ENABLED=0 go build -ldflags "-s -w" -tags reader -o "$@" ./reader/main.go

bin/grid-logger-writer:
	CGO_ENABLED=0 go build -ldflags "-s -w" -tags writer -o "$@" ./writer/main.go

bin/grid-logger-server:
	CGO_ENABLED=0 go build -ldflags "-s -w" -tags server -o "$@" ./server/main.go

.PHONY: all
all: $(addprefix bin/,$(bins))

bin/grid-logger-reader-linux-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -tags reader -o "$@" ./reader/main.go

bin/grid-logger-writer-linux-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -tags writer -o "$@" ./writer/main.go

bin/grid-logger-writer-linux-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -tags writer -o "$@" ./writer/main.go

bin/grid-logger-writer-linux-riscv64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags "-s -w" -tags writer -o "$@" ./writer/main.go

bin/grid-logger-server-linux-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -tags server -o "$@" ./server/main.go

bin/grid-logger-server-linux-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -tags server -o "$@" ./server/main.go

bin/grid-logger-server-linux-riscv64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags "-s -w" -tags server -o "$@" ./server/main.go

bin/checksums.txt: $(addprefix bin/,$(bins))
	sha256sum -b $(addprefix bin/,$(bins)) | sed 's/bin\///' > $@

bin/checksums.md: bin/checksums.txt
	@echo "### SHA256 Checksums" > $@
	@echo >> $@
	@echo "\`\`\`" >> $@
	@cat $< >> $@
	@echo "\`\`\`" >> $@

.PHONY: build-all
build-all: $(addprefix bin/,$(bins)) bin/checksums.md

.PHONY: unit
unit: $(GO_TESTS)
	go test -race -covermode=atomic -tags=unit -timeout=30s ./...

.PHONY: coverage
coverage: $(GO_TESTS)
	go test -race -covermode=atomic -tags=unit -timeout=30s -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

.PHONY: integration
integration:
	go test -race -covermode=atomic -tags=integration -timeout=300s ./...

$(golint):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint
lint: $(golint)
	$(golint) run ./...

.PHONY: mocks
mocks:
	mockery --all

.PHONY: grpc
grpc:
	cd ../protos/loggerapis && buf generate

.PHONY: clean
clean:
	rm -rf bin/

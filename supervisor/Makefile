GO_SRCS := $(shell find . -type f -name '*.go' -a ! \( -name 'zz_generated*' -o -name '*_test.go' \))
GO_TESTS := $(shell find . -type f -name '*_test.go')
TAG_NAME = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
TAG_NAME_DEV = $(shell git describe --tags --abbrev=0 --match "supervisor/*" 2>/dev/null)
VERSION_CORE = $(shell echo $(TAG_NAME) | sed 's/supervisor\/\(.*\)/\1/')
VERSION_CORE_DEV = $(shell echo $(TAG_NAME_DEV) | sed 's/supervisor\/\(.*\)/\1/')
GIT_COMMIT = $(shell git rev-parse --short=7 HEAD)
VERSION = $(or $(and $(TAG_NAME),$(VERSION_CORE)),$(and $(TAG_NAME_DEV),$(VERSION_CORE_DEV)-dev),$(GIT_COMMIT))

ifeq ($(golint),)
golint := $(shell go env GOPATH)/bin/golangci-lint
endif

ifeq ($(mockery),)
mockery := $(shell go env GOPATH)/bin/mockery
endif

.PHONY: bin/supervisor
bin/supervisor: $(GO_SRCS)
	go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/checksums.txt: $(addprefix bin/,$(bins))
	sha256sum -b $(addprefix bin/,$(bins)) | sed 's/bin\///' > $@

bin/checksums.md: bin/checksums.txt
	@echo "### SHA256 Checksums" > $@
	@echo >> $@
	@echo "\`\`\`" >> $@
	@cat $< >> $@
	@echo "\`\`\`" >> $@

bin/supervisor-darwin-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-darwin-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-freebsd-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-freebsd-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-linux-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-linux-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-linux-mips64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-linux-mips64le: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-linux-ppc64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=ppc64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-linux-ppc64le: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-linux-riscv64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-linux-s390x: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bin/supervisor-windows-amd64.exe: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/main.go

bins := supervisor-darwin-amd64 supervisor-darwin-arm64 supervisor-freebsd-arm64 supervisor-freebsd-arm64 supervisor-linux-amd64 supervisor-linux-arm64 supervisor-linux-mips64 supervisor-linux-mips64le supervisor-linux-ppc64 supervisor-linux-ppc64le supervisor-linux-riscv64 supervisor-linux-s390x supervisor-windows-amd64.exe

.PHONY: build-all
build-all: $(addprefix bin/,$(bins)) bin/checksums.md

.PHONY: lint
lint: $(golint)
	$(golint) run ./...

.PHONY: unit
unit:
	go test -race -covermode=atomic -tags=unit -timeout=30s ./...

.PHONY: coverage
coverage:
	go test -race -covermode=atomic -tags=unit -timeout=30s -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

.PHONY: integration
integration:
	go test -race -covermode=atomic -tags=integration -timeout=30s ./...

.PHONY: mocks
mocks:
	$(mockery)

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: version
version:
	@echo VERSION_CORE=${VERSION_CORE}
	@echo VERSION_CORE_DEV=${VERSION_CORE_DEV}
	@echo VERSION=${VERSION}

$(golint):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(mockery):
	go install github.com/vektra/mockery/v2@latest

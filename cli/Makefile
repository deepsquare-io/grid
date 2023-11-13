GO_SRCS := $(shell find . -type f -name '*.go' -a ! \( -name 'zz_generated*' -o -name '*_test.go' \))
GO_TESTS := $(shell find . -type f -name '*_test.go')
TAG_NAME = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
TAG_NAME_DEV = $(shell git describe --tags --abbrev=0 --match "cli/*" 2>/dev/null)
VERSION_CORE = $(shell echo $(TAG_NAME) | sed 's/cli\/\(.*\)/\1/')
VERSION_CORE_DEV = $(shell echo $(TAG_NAME_DEV) | sed 's/cli\/\(.*\)/\1/')
GIT_COMMIT = $(shell git rev-parse --short=7 HEAD)
VERSION = $(or $(and $(TAG_NAME),$(VERSION_CORE)),$(and $(TAG_NAME_DEV),$(VERSION_CORE_DEV)-dev),$(GIT_COMMIT))

ifeq ($(golint),)
golint := $(shell go env GOPATH)/bin/golangci-lint
endif

ifeq ($(pkgsite),)
pkgsite := $(shell go env GOPATH)/bin/pkgsite
endif

bin/deepsquaretui: $(GO_SRCS)
	CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/checksums.txt: $(addprefix bin/,$(bins))
	sha256sum -b $(addprefix bin/,$(bins)) | sed 's/bin\///' > $@

bin/checksums.md: bin/checksums.txt
	@echo "### SHA256 Checksums" > $@
	@echo >> $@
	@echo "\`\`\`" >> $@
	@cat $< >> $@
	@echo "\`\`\`" >> $@

bin/deepsquaretui-darwin-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-darwin-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-freebsd-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-freebsd-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-linux-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-linux-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-linux-mips64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-linux-mips64le: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-linux-ppc64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=ppc64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-linux-ppc64le: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-linux-riscv64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-linux-s390x: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bin/deepsquaretui-windows-amd64.exe: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "$@" ./cmd/deepsquaretui/main.go

bins := deepsquaretui-darwin-amd64 deepsquaretui-darwin-arm64 deepsquaretui-freebsd-arm64 deepsquaretui-freebsd-arm64 deepsquaretui-linux-amd64 deepsquaretui-linux-arm64 deepsquaretui-linux-mips64 deepsquaretui-linux-mips64le deepsquaretui-linux-ppc64 deepsquaretui-linux-ppc64le deepsquaretui-linux-riscv64 deepsquaretui-linux-s390x deepsquaretui-windows-amd64.exe

.PHONY: build-all
build-all: $(addprefix bin/,$(bins)) bin/checksums.md

.PHONY: unit
unit:
	go test -race -covermode=atomic -tags=unit -timeout=30s ./...

.PHONY: generate
generate:
	go generate ./...

$(golint):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(pkgsite):
	go install golang.org/x/pkgsite/cmd/pkgsite@latest

.PHONY: license
license:
	go run ./licensing.go


.PHONY: lint
lint: $(golint)
	$(golint) run ./...

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: version
version:
	@echo VERSION_CORE=${VERSION_CORE}
	@echo VERSION_CORE_DEV=${VERSION_CORE_DEV}
	@echo VERSION=${VERSION}

.PHONY: list
list:
	@LC_ALL=C $(MAKE) -pRrq -f $(firstword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/(^|\n)# Files(\n|$$)/,/(^|\n)# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | grep -E -v -e '^[^[:alnum:]]' -e '^$@$$'

.PHONY: doc
doc:
	$(pkgsite) -open .

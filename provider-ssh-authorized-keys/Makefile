GO_SRCS := $(shell find . -type f -name '*.go' -a ! \( -name 'zz_generated*' -o -name '*_test.go' \))
GO_TESTS := $(shell find . -type f -name '*_test.go')
TAG_NAME = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
TAG_NAME_DEV = $(shell git describe --tags --abbrev=0 2>/dev/null)
GIT_COMMIT = $(shell git rev-parse --short=7 HEAD)
VERSION = $(or ${TAG_NAME},$(TAG_NAME_DEV)-dev)
golint := $(shell which golangci-lint)
ifeq ($(golint),)
golint := $(shell go env GOPATH)/bin/golangci-lint
endif

bin/provider-ssh-authorized-keys: $(GO_SRCS)
	go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-darwin-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-darwin-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-freebsd-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-freebsd-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-linux-amd64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-linux-arm64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-linux-mips64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-linux-mips64le: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-linux-ppc64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=ppc64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-linux-ppc64le: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=ppc64le go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-linux-riscv64: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-linux-s390x: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=s390x go build -o "$@" ./main.go

bin/provider-ssh-authorized-keys-windows-amd64.exe: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "$@" ./main.go

bins := provider-ssh-authorized-keys-darwin-amd64 provider-ssh-authorized-keys-darwin-arm64 provider-ssh-authorized-keys-freebsd-arm64 provider-ssh-authorized-keys-freebsd-arm64 provider-ssh-authorized-keys-linux-amd64 provider-ssh-authorized-keys-linux-arm64 provider-ssh-authorized-keys-linux-mips64 provider-ssh-authorized-keys-linux-mips64le provider-ssh-authorized-keys-linux-ppc64 provider-ssh-authorized-keys-linux-ppc64le provider-ssh-authorized-keys-linux-riscv64 provider-ssh-authorized-keys-linux-s390x provider-ssh-authorized-keys-windows-amd64.exe

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
unit:
	go test -race -covermode=atomic -tags=unit -timeout=30s ./...

.PHONY: coverage
coverage:
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

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: grpc
grpc:
	cd "../protos/supervisorapis" && buf generate

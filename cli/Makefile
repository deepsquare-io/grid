GO_SRCS := $(shell find . -type f -name '*.go' -a -name '*.tpl' -a ! \( -name 'zz_generated*' -o -name '*_test.go' \))
GO_TESTS := $(shell find . -type f -name '*_test.go')

ifeq ($(golint),)
golint := $(shell go env GOPATH)/bin/golangci-lint
endif
bins := dps

.PHONY: all
all: $(addprefix bin/,$(bins))

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

.PHONY: bin/dps
bin/dps: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=test" -o "$@" ./tui/main.go

.PHONY: unit
unit:
	go test -race -covermode=atomic -tags=unit -timeout=30s ./...

$(golint):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint
lint: $(golint)
	$(golint) run ./...

.PHONY: clean
clean:
	rm -rf bin/

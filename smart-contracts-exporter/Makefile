GO_SRCS := $(shell find . -type f -name '*.go' -a -name '*.tpl' -a ! \( -name 'zz_generated*' -o -name '*_test.go' \))
GO_TESTS := $(shell find . -type f -name '*_test.go')

ifeq ($(golint),)
golint := $(shell go env GOPATH)/bin/golangci-lint
endif
bins := smart-contracts-exporter

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

.PHONY: bin/smart-contracts-exporter
bin/smart-contracts-exporter: $(GO_SRCS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=test" -o "$@" ./main.go

.PHONY: graphql
graphql: graph/model/models_gen.go

.PHONY: protos
protos:
	protoc --go_out=. --go_opt=module=github.com/deepsquare-io/grid/smart-contracts-exporter metrics/v1/metrics.proto

graph/model/models_gen.go: ../schemas/statapi/schema.graphqls gqlgen.yml
	go run github.com/99designs/gqlgen generate

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

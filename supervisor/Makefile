.PHONY: all
all: build

.PHONY: build
build: supervisor

.PHONY: test
test: unit

.PHONY: supervisor
supervisor:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o $@ ./cmd

.PHONY: unit
unit:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -race -covermode=atomic -tags=unit -timeout=30s ./...

.PHONY: coverage
coverage:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -race -covermode=atomic -tags=unit -timeout=30s -coverprofile=coverage.out ./...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go tool cover -html coverage.out -o coverage.html

.PHONY: integration
integration:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -race -covermode=atomic -tags=integration -timeout=30s ./...

.PHONY: mocks
mocks:
	mockery --all

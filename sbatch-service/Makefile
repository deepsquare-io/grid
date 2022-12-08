ifeq ($(golint),)
golint := $(shell go env GOPATH)/bin/golangci-lint
endif

bin/sbatch-service:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags reader -o "$@" ./server.go

.PHONY: graphql
graphql: graph/model/models_gen.go

graph/model/models_gen.go: ../schemas/sbatchapi/schema.graphqls
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

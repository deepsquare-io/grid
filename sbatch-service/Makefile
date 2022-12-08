bin/sbatch-service:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags reader -o "$@" ./server.go

.PHONY: graphql
graphql: graph/model/models_gen.go

graph/model/models_gen.go: ../schemas/sbatchapi/schema.graphqls
	go run github.com/99designs/gqlgen generate

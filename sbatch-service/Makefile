.PHONY: graphql
graphql: graph/model/models_gen.go

graph/model/models_gen.go: ../schemas/sbatchapi/schema.graphqls
	go run github.com/99designs/gqlgen generate

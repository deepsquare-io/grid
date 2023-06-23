//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/ethereum/go-ethereum"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

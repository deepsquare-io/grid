//go:build tools

package tools

import (
	_ "github.com/Khan/genqlient"
	_ "github.com/ethereum/go-ethereum"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

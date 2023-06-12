package graphql

import "github.com/vektah/gqlparser/v2/gqlerror"

type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type Response[T any] struct {
	Data   T                `json:"data,omitempty"`
	Errors []gqlerror.Error `json:"errors,omitempty"`
}

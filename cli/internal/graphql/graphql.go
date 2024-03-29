// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package graphql provides wrappers for graphQL API objects.
package graphql

import "github.com/vektah/gqlparser/v2/gqlerror"

// Request is used to build a GraphQL query.
type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// Response is the standard object returned by a GraphQL API.
type Response[T any] struct {
	Data   T                `json:"data,omitempty"`
	Errors []gqlerror.Error `json:"errors,omitempty"`
}

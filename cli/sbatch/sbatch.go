// Copyright (C) 2023 DeepSquare Association
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

//go:generate go run generate.go

// Package sbatch provides implementations of the SBatchService client.
package sbatch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/deepsquare-io/grid/cli/internal/graphql"
	"gopkg.in/yaml.v3"
)

// The Service is used to interact with the sbatch hosting service.
type Service interface {
	// Submit a job in the batch service.
	Submit(ctx context.Context, job *Job) (string, error)
}

type service struct {
	*http.Client
	endpoint string
}

// NewService creates a new Service.
func NewService(client *http.Client, endpoint string) Service {
	return &service{
		Client:   client,
		endpoint: endpoint,
	}
}

const submitMutation = `
mutation Submit($job: Job!) {
  submit(job: $job)
}
`

type submitResponseData struct {
	Submit string `json:"submit"`
}

func (s *service) Submit(ctx context.Context, job *Job) (string, error) {
	r := graphql.Request{
		Query: submitMutation,
		Variables: map[string]interface{}{
			"job": job,
		},
	}

	payload, err := json.Marshal(r)
	if err != nil {
		return "", fmt.Errorf("failed to encode body: %w", err)
	}

	req, err := http.NewRequest("POST", s.endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("failed create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)

	resp, err := s.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed send request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return "", fmt.Errorf(
			"server responsed with non-ok status code: %d, %s",
			resp.StatusCode,
			string(body),
		)
	}

	var result graphql.Response[submitResponseData]
	if err = json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if len(result.Errors) > 0 {
		out, err := yaml.Marshal(result.Errors)
		if err != nil {
			panic(err)
		}
		fmt.Println("sbatch failure:")
		fmt.Println(string(out))
		errs := make([]error, 0, len(result.Errors))
		for _, err := range result.Errors {
			errs = append(errs, errors.New(err.Error()))
		}
		return "", errors.Join(errs...)
	}
	return result.Data.Submit, nil
}

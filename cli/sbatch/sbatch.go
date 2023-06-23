//go:generate go run generate.go

package sbatch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deepsquare-io/the-grid/cli/internal/graphql"
)

type Service interface {
	// Submit a job in the batch service.
	Submit(ctx context.Context, job *Job) (string, error)
}

type service struct {
	*http.Client
	endpoint string
}

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
	return result.Data.Submit, err
}

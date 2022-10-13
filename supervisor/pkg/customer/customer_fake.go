package customer

import (
	"context"
	"io"
	"net/http"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

// FakeDataSource fetches the resources linked to the smart-contract from transfer.sh.
type FakeDataSource struct {
}

// Fetch a sbatch script based on the hash.
func (d *FakeDataSource) Fetch(ctx context.Context, hash string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", hash, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.I.Error("resp closed failed", zap.Error(err))
		}
	}()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

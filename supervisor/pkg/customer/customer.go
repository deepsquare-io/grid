package customer

import (
	"context"
	"io"

	customerv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/customer/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// DataSource fetches the resources linked to the smart-contract.
type DataSource struct {
	endpoint string
	opts     []grpc.DialOption
}

func New(
	endpoint string,
	tls bool,
	caFile string,
	serverHostOverride string,
) *DataSource {
	var opts []grpc.DialOption
	if tls {
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			logger.I.Fatal("Failed to create TLS credentials", zap.Error(err))
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &DataSource{
		endpoint: endpoint,
		opts:     opts,
	}
}

// Fetch a sbatch script based on the hash.
func (o *DataSource) Fetch(ctx context.Context, hash string) (string, error) {
	conn, err := grpc.Dial(o.endpoint, o.opts...)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := conn.Close(); err != nil && err != io.EOF {
			logger.I.Warn("closing oracle client thrown an error", zap.Error(err))
		}
	}()
	jobs := customerv1alpha1.NewJobAPIClient(conn)

	resp, err := jobs.FetchJobBatch(ctx, &customerv1alpha1.FetchJobBatchRequest{
		JobLocationHash: hash,
	})
	return resp.Body, err
}

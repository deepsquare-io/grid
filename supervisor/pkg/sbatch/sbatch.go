package sbatch

import (
	"context"
	cryptotls "crypto/tls"
	"io"
	"net/http"

	sbatchv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type API struct {
	endpoint string
	creds    credentials.TransportCredentials
	insecure bool
}

func NewAPI(
	endpoint string,
	tls bool,
	insecure bool,
	caFile string,
	serverHostOverride string,
) *API {
	var creds credentials.TransportCredentials
	if tls {
		if !insecure {
			c, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
			if err != nil {
				logger.I.Fatal("Failed to create TLS credentials", zap.Error(err))
			}
			creds = c
		} else {
			tlsConfig := &cryptotls.Config{
				InsecureSkipVerify: true,
			}
			creds = credentials.NewTLS(tlsConfig)
		}
	}

	d := &API{
		endpoint: endpoint,
		creds:    creds,
		insecure: insecure,
	}

	logger.I.Info("Healthcheck sbatch service")
	_, conn, err := d.dial()
	defer func() {
		_ = conn.Close()
	}()
	if err != nil {
		logger.I.Fatal("sbatch service initial healthcheck failed, exiting...")
	}
	return d
}

func (d *API) dial() (sbatchv1alpha1.SBatchAPIClient, *grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(d.creds)}
	conn, err := grpc.Dial(d.endpoint, opts...)
	if err != nil {
		return nil, nil, err
	}
	return sbatchv1alpha1.NewSBatchAPIClient(conn), conn, nil
}

func (d *API) Fetch(ctx context.Context, hash string) (string, error) {
	res, err := d.grpcFetch(ctx, hash)
	if err != nil {
		logger.I.Error("Failed to fetch sbatch from sbatch API, trying with transfer.sh", zap.Error(err))
		return d.transfershFetch(ctx, hash)
	}
	return res, nil
}

func (d *API) grpcFetch(ctx context.Context, hash string) (string, error) {
	client, conn, err := d.dial()
	if err != nil {
		return "", err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			if err == io.EOF {
				logger.I.Debug("sbatchAPI closed", zap.Error(err))
				return
			}
			logger.I.Error("sbatchAPI closed with error", zap.Error(err))
		}
	}()
	resp, err := client.GetSBatch(ctx, &sbatchv1alpha1.GetSBatchRequest{
		BatchLocationHash: hash,
	})
	return resp.GetSbatch(), err
}

// transfershFetch a sbatch script based on the hash from transfer.sh.
func (d *API) transfershFetch(ctx context.Context, hash string) (string, error) {
	logger.I.Warn("Calling transfershFetch. Fetching via transfer.sh is deprecated!")
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

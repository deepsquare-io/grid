package sbatch

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"

	sbatchv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type API struct {
	endpoint           string
	tls                bool
	insecure           bool
	caFile             string
	serverHostOverride string
}

func NewAPI(
	endpoint string,
	tls bool,
	insecure bool,
	caFile string,
	serverHostOverride string,
) *API {
	return &API{
		endpoint:           endpoint,
		tls:                tls,
		insecure:           insecure,
		caFile:             caFile,
		serverHostOverride: serverHostOverride,
	}
}

func (d *API) dial() (sbatchv1alpha1.SBatchAPIClient, *grpc.ClientConn, error) {
	opts := []grpc.DialOption{}
	if d.tls {
		if !d.insecure {
			creds, err := credentials.NewClientTLSFromFile(d.caFile, d.serverHostOverride)
			if err != nil {
				logger.I.Fatal("Failed to create TLS credentials", zap.Error(err))
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			tlsConfig := &tls.Config{
				InsecureSkipVerify: true,
			}
			opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
		}
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.Dial(d.serverHostOverride, opts...)
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

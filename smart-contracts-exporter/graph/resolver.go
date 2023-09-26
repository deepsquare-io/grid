package graph

import (
	"crypto/tls"
	"net/http"

	"github.com/deepsquare-io/grid/smart-contracts-exporter/logger"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"go.uber.org/zap"
)

type Resolver struct {
	PromAPI              v1.API
	metaschedulerAddress string
}

func NewResolver(prometheusURL string, metaschedulerAddress string) *Resolver {
	client, err := api.NewClient(api.Config{
		Address: prometheusURL,
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	})
	if err != nil {
		logger.I.Panic("failed to initialize prometheus client", zap.Error(err))
	}
	return &Resolver{
		PromAPI:              v1.NewAPI(client),
		metaschedulerAddress: metaschedulerAddress,
	}
}

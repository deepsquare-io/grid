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

package deepsquare

import (
	"context"
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deepsquare-io/grid/cli/logger"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/sbatch"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/deepsquare-io/grid/cli/types/allowance"
	"github.com/deepsquare-io/grid/cli/types/credit"
	"github.com/deepsquare-io/grid/cli/types/event"
	"github.com/deepsquare-io/grid/cli/types/job"
	"github.com/deepsquare-io/grid/cli/types/provider"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

// Default values for the client.
const (
	DefaultRPCEndpoint    = "https://testnet.deepsquare.run/rpc"
	DefaultWSEndpoint     = "wss://testnet.deepsquare.run/ws"
	DefaultSBatchEndpoint = "https://sbatch.deepsquare.run/graphql"
	DefaultLoggerEndpoint = "https://grid-logger.deepsquare.run"
)

// ClientConfig is used to configure the Client's services.
type ClientConfig struct {
	*http.Client
	// MetaschedulerAddress is the address of the smart-contract.
	MetaschedulerAddress common.Address
	// RPCEndpoint is the URL of the network API of the Ethereum Virtual Machine (EVM). The parameter is optional.
	RPCEndpoint string
	// SBatchEndpoint is the URL of the SBatch service. The parameter is optional.
	SBatchEndpoint string
	// SBatchEndpoint is the URL of the grid logger. The parameter is optional.
	LoggerEndpoint string
	// UserPrivateKey is the ECDSA private key of an ethereum wallet. This permits
	// authenticated requests if specified.
	UserPrivateKey *ecdsa.PrivateKey
	// MetaschedulerOracleEndpoint is the URL of the meta-scheduler oracle. The parameter is optional.
	MetaschedulerOracleEndpoint string
	// TLSConfig of the HTTP and WS client used for all the connections. This parameter is optional.
	TLSConfig *tls.Config
}

func (c *ClientConfig) applyDefault() {
	if c == nil {
		c = &ClientConfig{}
	}
	if c.Client == nil {
		c.Client = http.DefaultClient
	}
	if c.RPCEndpoint == "" {
		c.RPCEndpoint = DefaultRPCEndpoint
	}
	if c.SBatchEndpoint == "" {
		c.SBatchEndpoint = DefaultSBatchEndpoint
	}
	if c.LoggerEndpoint == "" {
		c.LoggerEndpoint = DefaultLoggerEndpoint
	}
	if c.MetaschedulerOracleEndpoint == "" {
		c.MetaschedulerOracleEndpoint = metascheduler.DefaultOracleURL
	}
	if c.TLSConfig == nil {
		caCertPool, err := x509.SystemCertPool()
		if err != nil {
			caCertPool = x509.NewCertPool()
		}
		c.TLSConfig = &tls.Config{
			RootCAs: caCertPool,
		}
	}
	if c.Client.Transport == nil {
		c.Client.Transport = &http.Transport{
			TLSClientConfig: c.TLSConfig,
		}
	}
}

// Client implements all the services required to make unary calls to DeepSquare APIs.
//
// Users must call Close() at the end of the application to avoid pending connections.
type Client struct {
	types.Logger
	job.Scheduler
	job.Fetcher
	job.ByProviderFetcher
	CreditManager    credit.Manager
	AllowanceManager allowance.Manager
	ProviderManager  provider.Manager
	loggerConn       *grpc.ClientConn
	rpcClient        *rpc.Client
}

// NewClient creates a new Client for the given ClientConfig.
func NewClient(ctx context.Context, c *ClientConfig) (*Client, error) {
	c.applyDefault()
	rpcClient, err := rpc.DialOptions(ctx, c.RPCEndpoint, rpc.WithHTTPClient(c.Client))
	if err != nil {
		return nil, err
	}
	ethClientRPC := ethclient.NewClient(rpcClient)
	chainID, err := ethClientRPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	rpcClientSet := metascheduler.NewRPCClientSet(metascheduler.Backend{
		MetaschedulerAddress: c.MetaschedulerAddress,
		ChainID:              chainID,
		EthereumBackend:      ethClientRPC,
		UserPrivateKey:       c.UserPrivateKey,
	})
	sbatch := sbatch.NewService(c.Client, c.SBatchEndpoint)
	jobScheduler := rpcClientSet.JobScheduler(sbatch)
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(c.TLSConfig)),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    20 * time.Second,
			Timeout: 30 * time.Second,
		}),
	}
	u, err := url.Parse(c.LoggerEndpoint)
	if err != nil {
		return nil, err
	}
	port := u.Port()
	if port == "" {
		// If the URL doesn't explicitly specify a port, use the default port for the scheme.
		switch strings.ToLower(u.Scheme) {
		case "http":
			port = "80"
		case "https":
			port = "443"
		default:
			port = "443"
		}
	}
	logger, conn, err := logger.DialContext(
		ctx,
		net.JoinHostPort(u.Hostname(), port),
		c.UserPrivateKey,
		dialOptions...)
	if err != nil {
		return nil, err
	}
	oracle := metascheduler.NewOracle(c.MetaschedulerOracleEndpoint, metascheduler.OracleOptions{
		Client: c.Client,
	})
	fetcher := rpcClientSet.JobFetcher()
	runningJobsByProviderFetcher := metascheduler.NewJobsByProviderFetcher(oracle, fetcher)
	return &Client{
		Fetcher:           fetcher,
		Scheduler:         jobScheduler,
		ByProviderFetcher: runningJobsByProviderFetcher,
		CreditManager:     rpcClientSet.CreditManager(),
		AllowanceManager:  rpcClientSet.AllowanceManager(),
		ProviderManager:   rpcClientSet.ProviderManager(),
		Logger:            logger,
		loggerConn:        conn,
		rpcClient:         rpcClient,
	}, nil
}

// Close closes the underlying connections.
func (c *Client) Close() error {
	c.rpcClient.Close()
	return c.loggerConn.Close()
}

// WatcherConfig is used to configure the Watcher's services.
type WatcherConfig struct {
	http.Client
	MetaschedulerAddress common.Address
	RPCEndpoint          string
	WSEndpoint           string
	UserPrivateKey       *ecdsa.PrivateKey
	TLSConfig            *tls.Config
}

func (c *WatcherConfig) applyDefault() {
	if c == nil {
		c = &WatcherConfig{}
	}
	if c.RPCEndpoint == "" {
		c.RPCEndpoint = DefaultRPCEndpoint
	}
	if c.WSEndpoint == "" {
		c.WSEndpoint = DefaultWSEndpoint
	}
	if c.TLSConfig == nil {
		caCertPool, err := x509.SystemCertPool()
		if err != nil {
			caCertPool = x509.NewCertPool()
		}
		c.TLSConfig = &tls.Config{
			RootCAs: caCertPool,
		}
	}
	if c.Client.Transport == nil {
		c.Client.Transport = &http.Transport{
			TLSClientConfig: c.TLSConfig,
		}
	}
}

// Watcher implements all the services required to make streaming calls to DeepSquare APIs.
//
// Users must call Close() at the end of the application to avoid pending connections.
type Watcher struct {
	event.Subscriber
	rpcClient *rpc.Client
	wsClient  *rpc.Client
}

// NewWatcher creates a new Watcher for the given WatcherConfig.
func NewWatcher(ctx context.Context, c *WatcherConfig) (*Watcher, error) {
	c.applyDefault()
	rpcClient, err := rpc.DialOptions(ctx, c.RPCEndpoint, rpc.WithHTTPClient(&c.Client))
	if err != nil {
		return nil, err
	}
	ethClientRPC := ethclient.NewClient(rpcClient)
	chainID, err := ethClientRPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	metaschedulerRPC := metascheduler.Backend{
		MetaschedulerAddress: c.MetaschedulerAddress,
		ChainID:              chainID,
		EthereumBackend:      ethClientRPC,
		UserPrivateKey:       c.UserPrivateKey,
	}
	wsClient, err := rpc.DialOptions(ctx, c.WSEndpoint, rpc.WithHTTPClient(&c.Client))
	if err != nil {
		return nil, err
	}
	ethClientWS := ethclient.NewClient(wsClient)
	metaschedulerWS := metascheduler.Backend{
		MetaschedulerAddress: c.MetaschedulerAddress,
		ChainID:              chainID,
		EthereumBackend:      ethClientWS,
		UserPrivateKey:       c.UserPrivateKey,
	}
	return &Watcher{
		Subscriber: metascheduler.NewEventSubscriber(metaschedulerRPC, metaschedulerWS),
		rpcClient:  rpcClient,
		wsClient:   wsClient,
	}, nil
}

// Close closes the underlying connections.
func (c *Watcher) Close() error {
	c.rpcClient.Close()
	c.wsClient.Close()
	return nil
}

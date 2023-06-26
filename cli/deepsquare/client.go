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

	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	DefaultEndpointRPC    = "https://testnet.deepsquare.run/rpc"
	DefaultEndpointWS     = "wss://testnet.deepsquare.run/ws"
	DefaultSBatchEndpoint = "https://sbatch.deepsquare.run/graphql"
	DefaultLoggerEndpoint = "https://grid-logger.deepsquare.run"
)

var (
	DefaultMetaSchedulerAddress = common.HexToAddress("0xc9AcB97F1132f0FB5dC9c5733B7b04F9079540f0")
)

type Client interface {
	types.LoggerDialer
	types.JobScheduler
	types.JobFetcher
	types.CreditManager
	types.AllowanceManager
}

type ClientConfig struct {
	http.Client
	MetaschedulerAddress common.Address
	EndpointRPC          string
	SBatchEndpoint       string
	LoggerEndpoint       string
	UserPrivateKey       *ecdsa.PrivateKey
	TLSConfig            *tls.Config
}

func (c *ClientConfig) applyDefault() {
	if c == nil {
		c = &ClientConfig{}
	}
	if c.EndpointRPC == "" {
		c.EndpointRPC = DefaultEndpointRPC
	}
	if (c.MetaschedulerAddress == common.Address{}) {
		c.MetaschedulerAddress = DefaultMetaSchedulerAddress
	}
	if c.SBatchEndpoint == "" {
		c.SBatchEndpoint = DefaultSBatchEndpoint
	}
	if c.LoggerEndpoint == "" {
		c.LoggerEndpoint = DefaultLoggerEndpoint
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

type client struct {
	types.LoggerDialer
	types.JobScheduler
	types.JobFetcher
	types.CreditManager
	types.AllowanceManager
}

func NewClient(ctx context.Context, c *ClientConfig) (Client, error) {
	c.applyDefault()
	rpcClient, err := rpc.DialOptions(ctx, c.EndpointRPC, rpc.WithHTTPClient(&c.Client))
	if err != nil {
		return nil, err
	}
	ethClientRPC := ethclient.NewClient(rpcClient)
	chainID, err := ethClientRPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	metaschedulerRPC := metascheduler.Client{
		MetaschedulerAddress: c.MetaschedulerAddress,
		ChainID:              chainID,
		EthereumBackend:      ethClientRPC,
		UserPrivateKey:       c.UserPrivateKey,
	}
	fetcher, err := metascheduler.NewJobFetcher(metaschedulerRPC)
	if err != nil {
		return nil, err
	}
	sbatch := sbatch.NewService(&c.Client, c.SBatchEndpoint)
	jobScheduler, err := metascheduler.NewJobScheduler(metaschedulerRPC, sbatch)
	if err != nil {
		return nil, err
	}
	credits, err := metascheduler.NewCreditManager(ctx, metaschedulerRPC)
	if err != nil {
		return nil, err
	}
	allowance, err := metascheduler.NewAllowanceManager(ctx, metaschedulerRPC)
	if err != nil {
		return nil, err
	}
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(c.TLSConfig)),
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
	logDialer := logger.NewDialer(
		net.JoinHostPort(u.Hostname(), port),
		c.UserPrivateKey,
		dialOptions...,
	)
	return &client{
		JobFetcher:       fetcher,
		JobScheduler:     jobScheduler,
		CreditManager:    credits,
		AllowanceManager: allowance,
		LoggerDialer:     logDialer,
	}, nil
}

type Watcher interface {
	types.EventSubscriber
	types.JobFilterer
	types.CreditFilterer
	types.AllowanceFilterer
}

type WatcherConfig struct {
	http.Client
	MetaschedulerAddress common.Address
	EndpointRPC          string
	EndpointWS           string
	UserPrivateKey       *ecdsa.PrivateKey
	TLSConfig            *tls.Config
}

func (c *WatcherConfig) applyDefault() {
	if c == nil {
		c = &WatcherConfig{}
	}
	if c.EndpointRPC == "" {
		c.EndpointRPC = DefaultEndpointRPC
	}
	if c.EndpointWS == "" {
		c.EndpointWS = DefaultEndpointWS
	}
	if (c.MetaschedulerAddress == common.Address{}) {
		c.MetaschedulerAddress = DefaultMetaSchedulerAddress
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

type watcher struct {
	types.EventSubscriber
	types.JobFilterer
	types.CreditFilterer
	types.AllowanceFilterer
}

func NewWatcher(ctx context.Context, c *WatcherConfig) (Watcher, error) {
	c.applyDefault()
	rpcClient, err := rpc.DialOptions(ctx, c.EndpointRPC, rpc.WithHTTPClient(&c.Client))
	if err != nil {
		return nil, err
	}
	ethClientRPC := ethclient.NewClient(rpcClient)
	chainID, err := ethClientRPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	metaschedulerRPC := metascheduler.Client{
		MetaschedulerAddress: c.MetaschedulerAddress,
		ChainID:              chainID,
		EthereumBackend:      ethClientRPC,
		UserPrivateKey:       c.UserPrivateKey,
	}
	wsClient, err := rpc.DialOptions(ctx, c.EndpointWS, rpc.WithHTTPClient(&c.Client))
	if err != nil {
		return nil, err
	}
	ethClientWS := ethclient.NewClient(wsClient)
	metaschedulerWS := metascheduler.Client{
		MetaschedulerAddress: c.MetaschedulerAddress,
		ChainID:              chainID,
		EthereumBackend:      ethClientWS,
		UserPrivateKey:       c.UserPrivateKey,
	}
	eventSubscriber, err := metascheduler.NewEventSubscriber(metaschedulerWS)
	if err != nil {
		return nil, err
	}
	jobFilterer, err := metascheduler.NewJobFilterer(metaschedulerWS)
	if err != nil {
		return nil, err
	}
	credits, err := metascheduler.NewCreditManager(ctx, metaschedulerRPC)
	if err != nil {
		return nil, err
	}
	allowance, err := metascheduler.NewAllowanceManager(ctx, metaschedulerRPC)
	if err != nil {
		return nil, err
	}
	creditFilterer, err := metascheduler.NewCreditFilterer(ctx, metaschedulerWS, credits)
	if err != nil {
		return nil, err
	}
	allowanceFilterer, err := metascheduler.NewAllowanceFilterer(ctx, metaschedulerWS, allowance)
	if err != nil {
		return nil, err
	}
	return &watcher{
		EventSubscriber:   eventSubscriber,
		JobFilterer:       jobFilterer,
		CreditFilterer:    creditFilterer,
		AllowanceFilterer: allowanceFilterer,
	}, nil
}

// Package deepsquare provides a simple Client that implement all the DeepSquare services.
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

// Default values for the client.
const (
	DefaultRPCEndpoint    = "https://testnet.deepsquare.run/rpc"
	DefaultWSEndpoint     = "wss://testnet.deepsquare.run/ws"
	DefaultSBatchEndpoint = "https://sbatch.deepsquare.run/graphql"
	DefaultLoggerEndpoint = "https://grid-logger.deepsquare.run"
)

var (
	DefaultMetaSchedulerAddress = common.HexToAddress("0x3707aB457CF457275b7ec32e203c54df80C299d5")
)

// Client implements all the services required to make unary calls to DeepSquare APIs.
//
// Users must call Close() at the end of the application to avoid pending connections.
type Client interface {
	types.Logger
	types.JobScheduler
	types.JobFetcher
	types.CreditManager
	types.AllowanceManager
	// Close all connections.
	Close() error
}

// ClientConfig is used to configure the Client's services.
type ClientConfig struct {
	http.Client
	MetaschedulerAddress common.Address
	RPCEndpoint          string
	SBatchEndpoint       string
	LoggerEndpoint       string
	UserPrivateKey       *ecdsa.PrivateKey
	TLSConfig            *tls.Config
}

func (c *ClientConfig) applyDefault() {
	if c == nil {
		c = &ClientConfig{}
	}
	if c.RPCEndpoint == "" {
		c.RPCEndpoint = DefaultRPCEndpoint
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
	types.Logger
	types.JobScheduler
	types.JobFetcher
	types.CreditManager
	types.AllowanceManager
	loggerConn *grpc.ClientConn
	rpcClient  *rpc.Client
}

// NewClient creates a new Client for the given ClientConfig.
func NewClient(ctx context.Context, c *ClientConfig) (Client, error) {
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
	fetcher, err := metascheduler.NewJobFetcher(ctx, metaschedulerRPC)
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
	logger, conn, err := logger.DialContext(
		ctx,
		net.JoinHostPort(u.Hostname(), port),
		c.UserPrivateKey,
		dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{
		JobFetcher:       fetcher,
		JobScheduler:     jobScheduler,
		CreditManager:    credits,
		AllowanceManager: allowance,
		Logger:           logger,
		loggerConn:       conn,
		rpcClient:        rpcClient,
	}, nil
}

func (c *client) Close() error {
	c.rpcClient.Close()
	return c.loggerConn.Close()
}

// Watcher implements all the services required to make streaming calls to DeepSquare APIs.
//
// Users must call Close() at the end of the application to avoid pending connections.
type Watcher interface {
	types.EventSubscriber
	types.JobFilterer
	types.CreditFilterer
	types.AllowanceFilterer
	// Close all connections.
	Close() error
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
	rpcClient *rpc.Client
	wsClient  *rpc.Client
}

// NewWatcher creates a new Watcher for the given WatcherConfig.
func NewWatcher(ctx context.Context, c *WatcherConfig) (Watcher, error) {
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
		rpcClient:         rpcClient,
		wsClient:          wsClient,
	}, nil
}

func (c *watcher) Close() error {
	c.rpcClient.Close()
	c.wsClient.Close()
	return nil
}

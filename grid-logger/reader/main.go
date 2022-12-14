package main

import (
	"crypto/ecdsa"
	"crypto/tls"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/deepsquare-io/the-grid/grid-logger/reader/api"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverEndpoint     string
	serverTLS          bool
	serverTLSInsecure  bool
	serverCAFile       string
	serverHostOverride string

	ethHexPK string
	pk       *ecdsa.PrivateKey
	pub      ecdsa.PublicKey
	address  string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "server.endpoint",
		Value:       "127.0.0.1:3000",
		Usage:       "Server to send logs.",
		Destination: &serverEndpoint,
		EnvVars:     []string{"SERVER_ENDPOINT"},
	},
	&cli.BoolFlag{
		Name:        "server.tls",
		Value:       true,
		Usage:       "Enable TLS for the Customer API.",
		Destination: &serverTLS,
		EnvVars:     []string{"SERVER_TLS_ENABLE"},
	},
	&cli.BoolFlag{
		Name:        "server.tls.insecure",
		Value:       false,
		Usage:       "Skip TLS verification. By enabling it, server.tls.ca and server.tls.server-host-override are ignored.",
		Destination: &serverTLSInsecure,
		EnvVars:     []string{"SERVER_TLS_INSECURE"},
	},
	&cli.StringFlag{
		Name:        "server.tls.ca",
		Value:       "",
		Usage:       "Path to CA certificate for TLS verification.",
		Destination: &serverCAFile,
		EnvVars:     []string{"SERVER_CA"},
	},
	&cli.StringFlag{
		Name:        "server.tls.server-host-override",
		Value:       "logging.deepsquare.io",
		Usage:       "The server name used to verify the hostname returned by the TLS handshake.",
		Destination: &serverHostOverride,
		EnvVars:     []string{"SERVER_HOST_OVERRIDE"},
	},

	&cli.StringFlag{
		Name:        "eth.pk",
		Usage:       "Private key (hexadecimal format: 0xabcdef).",
		Required:    true,
		Destination: &ethHexPK,
		Action: func(ctx *cli.Context, s string) (err error) {
			pk, err = crypto.HexToECDSA(ethHexPK)
			if err != nil {
				logger.I.Fatal("couldn't decode private key", zap.Error(err))
			}
			pub = pk.PublicKey
			address = strings.ToLower(crypto.PubkeyToAddress(pub).Hex())
			return nil
		},
		EnvVars: []string{"ETH_PK"},
	},
	&cli.BoolFlag{
		Name:    "debug",
		EnvVars: []string{"DEBUG"},
		Value:   false,
		Action: func(ctx *cli.Context, s bool) error {
			if s {
				logger.EnableDebug()
			}
			return nil
		},
	},
}

var app = &cli.App{
	Name:      "grid-logger-reader",
	Usage:     "Read logs",
	ArgsUsage: "<log name>",

	Flags:   flags,
	Suggest: true,

	Commands: []*cli.Command{
		{
			Name: "list",
			Action: func(cCtx *cli.Context) error {
				ctx := cCtx.Context

				// Dial server
				conn, err := dial()
				if err != nil {
					logger.I.Error("grpc dial failed", zap.Error(err))
					return err
				}
				defer func() {
					if err := conn.Close(); err != nil {
						if err != io.EOF {
							logger.I.Error("grpc close failed", zap.Error(err))
						}
					}
				}()
				client := api.New(conn, pk)

				return client.List(ctx, address)
			},
		},
	},

	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context

		if cCtx.NArg() < 1 {
			return errors.New("not enough arguments")
		}

		logName := cCtx.Args().First()

		// Dial server
		conn, err := dial()
		if err != nil {
			logger.I.Error("grpc dial failed", zap.Error(err))
			return err
		}
		defer func() {
			if err := conn.Close(); err != nil {
				if err != io.EOF {
					logger.I.Error("grpc close failed", zap.Error(err))
				}
			}
		}()
		client := api.New(conn, pk)

		return client.ReadAndWatch(ctx, address, logName)
	},
}

func dial() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{}
	if serverTLS {
		if !serverTLSInsecure {
			creds, err := credentials.NewClientTLSFromFile(serverCAFile, serverHostOverride)
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
	conn, err := grpc.Dial(serverEndpoint, opts...)
	return conn, err
}

func main() {
	_ = godotenv.Load(".reader.env", ".reader.env.local")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}

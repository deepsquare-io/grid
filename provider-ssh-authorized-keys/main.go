package main

import (
	"context"
	"fmt"
	"io"
	"os"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/provider-ssh-authorized-keys/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/provider-ssh-authorized-keys/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tls                bool
	caFile             string
	serverHostOverride string
	supervisorEndpoint string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "supervisor.endpoint",
		Usage:       "The supervisor endpoint. (supervisor.example.com:3000)",
		Required:    true,
		Destination: &supervisorEndpoint,
		EnvVars:     []string{"SUPERVISOR_ENDPOINT"},
	},
	&cli.BoolFlag{
		Name:        "supervisor.tls",
		Value:       false,
		Destination: &tls,
		Usage:       "Enable TLS for GRPC. Disabling it will enable insecure mode.",
		EnvVars:     []string{"SUPERVISOR_TLS_ENABLE"},
	},
	&cli.StringFlag{
		Name:        "supervisor.tls.ca",
		Value:       "",
		Usage:       "Path to CA certificate.",
		Destination: &caFile,
		EnvVars:     []string{"SUPERVISOR_CA"},
	},
	&cli.StringFlag{
		Name:        "supervisor.tls.server-host-override",
		Value:       "supervisor.example.com",
		Usage:       "The server name used to verify the hostname returned by the TLS handshake.",
		Destination: &serverHostOverride,
		EnvVars:     []string{"SUPERVISOR_SERVER_HOST_OVERRIDE"},
	},
}

var app = &cli.App{
	Name:  "provider-ssh-authorized-keys",
	Usage: "Fetch the public ssh key from the supervisor",
	Flags: flags,
	Action: func(ctx *cli.Context) error {
		out, err := Fetch(ctx.Context)
		if err != nil {
			logger.I.Fatal("failed to fetch authorized_keys", zap.Error(err))
		}
		fmt.Println(out)
		return nil
	},
}

func Fetch(ctx context.Context) (string, error) {
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
	conn, err := grpc.Dial(supervisorEndpoint, opts...)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := conn.Close(); err != nil && err != io.EOF {
			logger.I.Warn("closing grpc client thrown an error", zap.Error(err))
		}
	}()
	jobs := supervisorv1alpha1.NewSshAPIClient(conn)

	resp, err := jobs.FetchAuthorizedKeys(ctx, &supervisorv1alpha1.FetchAuthorizedKeysRequest{})
	if err != nil {
		return "", err
	}
	return resp.GetAuthorizedKeys(), err
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
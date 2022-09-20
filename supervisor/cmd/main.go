package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	supervisorv1alpha1 "github.com/deepsquare/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare/supervisor/logger"
	"github.com/deepsquare/supervisor/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.I.Info(
			"recv",
			zap.String("host", r.Host),
			zap.String("method", r.Method),
			zap.String("content-type", r.Header.Get("content-type")),
		)
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("content-type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	var grpcListenAddress string
	var grpcListenPort int
	var tls bool
	var certFile string
	var keyFile string

	logger.Init()

	app := &cli.App{
		Name:  "supervisor",
		Usage: "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "grpc.listen-address",
				Value:       "0.0.0.0",
				Usage:       "Address to listen on.",
				Destination: &grpcListenAddress,
				EnvVars:     []string{"GRPC_LISTEN_ADDRESS"},
			},
			&cli.IntFlag{
				Name:        "grpc.listen-port",
				Value:       9000,
				Usage:       "Port to listen on.",
				Destination: &grpcListenPort,
				EnvVars:     []string{"GRPC_LISTEN_PORT"},
			},
			&cli.BoolFlag{
				Name:        "tls",
				Value:       false,
				Usage:       "Enable TLS.",
				Destination: &tls,
				EnvVars:     []string{"TLS_ENABLE"},
			},
			&cli.StringFlag{
				Name:        "tls.cert",
				Value:       "",
				Usage:       "TLS Certificate file.",
				Destination: &certFile,
				EnvVars:     []string{"TLS_CERT"},
			},
			&cli.StringFlag{
				Name:        "tls.key",
				Value:       "",
				Usage:       "TLS Key file.",
				Destination: &keyFile,
				EnvVars:     []string{"TLS_KEY"},
			},
		},
		Action: func(ctx *cli.Context) error {
			// gRPC Server Options
			var opts []grpc.ServerOption
			if tls {
				creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
				if err != nil {
					logger.I.Fatal("failed to load certificates", zap.Error(err))
				}
				opts = []grpc.ServerOption{grpc.Creds(creds)}
			}
			grpcServer := grpc.NewServer(opts...)
			supervisorv1alpha1.RegisterRouteGuideAPIServer(grpcServer, server.New())

			// Register reflection service on gRPC server.
			reflection.Register(grpcServer)

			// gRPC Gateway options
			var dopts []grpc.DialOption
			if tls {
				creds, err := credentials.NewClientTLSFromFile(certFile, "")
				if err != nil {
					logger.I.Fatal("failed to load certificates", zap.Error(err))
				}
				dopts = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
			} else {
				dopts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
			}
			mux := runtime.NewServeMux()
			err := supervisorv1alpha1.RegisterRouteGuideAPIHandlerFromEndpoint(ctx.Context, mux, fmt.Sprintf("127.0.0.1:%d", grpcListenPort), dopts)
			if err != nil {
				logger.I.Fatal("failed to register RouteGuideAPIHandler", zap.Error(err))
			}

			// HTTP Setup
			logger.I.Info(
				"serving gRPC and HTTP server",
				zap.String("grpc.listen-address", grpcListenAddress),
				zap.Int("grpc.listen-port", grpcListenPort),
			)
			addr := fmt.Sprintf("%s:%d", grpcListenAddress, grpcListenPort)
			h2 := &http2.Server{}
			grpcHandler := h2c.NewHandler(grpcHandlerFunc(grpcServer, mux), h2)

			if tls {
				logger.I.Info(
					"using TLS",
					zap.String("tls.key-file", keyFile),
					zap.String("tls.cert-file", certFile),
				)
				return http.ListenAndServeTLS(addr, certFile, keyFile, grpcHandler)
			} else {
				return http.ListenAndServe(addr, grpcHandler)
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}

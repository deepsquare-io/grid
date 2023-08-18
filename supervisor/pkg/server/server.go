package server

import (
	"context"
	"net/http"
	"strings"
	"time"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/generated/grpc/health/v1"
	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/hpl"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/speedtest"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/health"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/jobapi"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/sshapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func New(
	ms metascheduler.MetaScheduler,
	resourceManager *lock.ResourceManager,
	launcher benchmark.Launcher,
	pkB64 string,
	hplOpts []benchmark.Option,
	opts ...grpc.ServerOption,
) *http.Server {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Route("/benchmark", func(r chi.Router) {
		r.Use(secret.Guard)
		r.Put(
			"/speedtest",
			benchmark.NewSpeedTestHandler(func(res *speedtest.Result) error {
				benchmark.DefaultStore.SetUploadBandwidth(res.Upload.Bandwidth)
				benchmark.DefaultStore.SetDownloadBandwidth(res.Download.Bandwidth)
				return nil
			}),
		)
		r.Route("/osu", func(r chi.Router) {
			r.Put("/pt2pt-latency", benchmark.NewOSUHandler(func(res float64) error {
				benchmark.DefaultStore.SetP2PLatency(res)
				return nil
			}))
			r.Put("/pt2pt-bibw", benchmark.NewOSUHandler(func(res float64) error {
				benchmark.DefaultStore.SetP2PBidirectionalBandwidth(res)
				return nil
			}))
			r.Put("/alltoall", benchmark.NewOSUHandler(func(res float64) error {
				benchmark.DefaultStore.SetAllToAllCollectiveLatency(res)
				return nil
			}))
		})
		r.Route("/hpl", func(r chi.Router) {
			r.Put(
				"/phase1",
				benchmark.NewHPLPhase1Handler(
					func(optimal *hpl.Result, opts ...benchmark.Option) error {
						opts = append(hplOpts, opts...)
						b, err := benchmark.GeneratePhase2HPLBenchmark(
							optimal.P,
							optimal.Q,
							optimal.ProblemSize,
							optimal.NB,
							opts...,
						)
						if err != nil {
							logger.I.Error(
								"failed to generate hpl phase 2 benchmark",
								zap.Error(err),
							)
							return err
						}

						go func() {
							ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
							defer cancel()

							if err := launcher.Launch(ctx, "hpl-phase2", b); err != nil {
								logger.I.Fatal(
									"hpl-phase2 benchmark failed or failed to be tracked",
									zap.Error(err),
								)
							}
							logger.I.Info(
								"benchmark hpl phase 2 succeeded",
								zap.Uint64(
									"p",
									optimal.P,
								),
								zap.Uint64("q", optimal.Q),
								zap.Uint64("n", optimal.ProblemSize),
								zap.Uint64("nb", optimal.NB),
							)
						}()

						return nil
					},
				),
			)
			r.Put("/phase2", benchmark.NewHPLPhase2Handler(
				func(gflops float64) error {
					benchmark.DefaultStore.SetGFLOPS(gflops)
					return nil
				}))
		})
	})
	g := grpc.NewServer(opts...)
	supervisorv1alpha1.RegisterJobAPIServer(
		g,
		jobapi.New(ms, resourceManager),
	)
	supervisorv1alpha1.RegisterSshAPIServer(
		g,
		sshapi.New(pkB64),
	)
	healthv1.RegisterHealthServer(
		g,
		health.New(),
	)

	rg := mixedHandler(r, g)
	http2Server := &http2.Server{}
	http1Server := &http.Server{Handler: h2c.NewHandler(rg, http2Server)}

	return http1Server
}

func mixedHandler(httpHand http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 &&
			strings.HasPrefix(r.Header.Get("content-type"), "application/grpc") {
			grpcHandler.ServeHTTP(w, r)
			return
		}
		httpHand.ServeHTTP(w, r)
	})
}

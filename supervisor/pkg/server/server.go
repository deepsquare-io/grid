package server

import (
	"net/http"
	"strings"

	healthv1 "github.com/deepsquare-io/grid/supervisor/generated/grpc/health/v1"
	supervisorv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/hpl"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/ior"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/speedtest"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/server/health"
	"github.com/deepsquare-io/grid/supervisor/pkg/server/jobapi"
	"github.com/deepsquare-io/grid/supervisor/pkg/server/sshapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
		r.Post(
			"/machine",
			benchmark.NewMachineHandler(func(spec *benchmark.MachineSpec, err error) error {
				if err != nil {
					benchmark.DefaultStore.SetFailure(err)
					return err
				}
				benchmark.DefaultStore.SetMachineSpec(spec)
				return nil
			}),
		)
		r.Put(
			"/speedtest",
			benchmark.NewSpeedTestHandler(func(res *speedtest.Result, err error) error {
				if err != nil {
					benchmark.DefaultStore.SetFailure(err)
					return err
				}
				benchmark.DefaultStore.SetUploadBandwidth(res.Upload.Bandwidth)
				benchmark.DefaultStore.SetDownloadBandwidth(res.Download.Bandwidth)
				return nil
			}),
		)
		r.Route("/osu", func(r chi.Router) {
			r.Put("/pt2pt-latency", benchmark.NewOSUHandler(func(res float64, err error) error {
				if err != nil {
					benchmark.DefaultStore.SetFailure(err)
					return err
				}
				benchmark.DefaultStore.SetP2PLatency(res)
				return nil
			}))
			r.Put("/pt2pt-bibw", benchmark.NewOSUHandler(func(res float64, err error) error {
				if err != nil {
					benchmark.DefaultStore.SetFailure(err)
					return err
				}
				benchmark.DefaultStore.SetP2PBidirectionalBandwidth(res)
				return nil
			}))
			r.Put("/alltoall", benchmark.NewOSUHandler(func(res float64, err error) error {
				if err != nil {
					benchmark.DefaultStore.SetFailure(err)
					return err
				}
				benchmark.DefaultStore.SetAllToAllCollectiveLatency(res)
				return nil
			}))
		})
		r.Route("/ior", func(r chi.Router) {
			r.Put(
				"/scratch",
				benchmark.NewIORHandler(func(avgr, avgw *ior.Result, err error) error {
					if err != nil {
						benchmark.DefaultStore.SetFailure(err)
						return err
					}
					benchmark.DefaultStore.SetScratchResult(avgr, avgw)
					return nil
				}),
			)
			r.Put(
				"/shared-world-tmp",
				benchmark.NewIORHandler(func(avgr, avgw *ior.Result, err error) error {
					if err != nil {
						benchmark.DefaultStore.SetFailure(err)
						return err
					}
					benchmark.DefaultStore.SetSharedWorldTmpResult(avgr, avgw)
					return nil
				}),
			)
			r.Put(
				"/shared-tmp",
				benchmark.NewIORHandler(func(avgr, avgw *ior.Result, err error) error {
					if err != nil {
						benchmark.DefaultStore.SetFailure(err)
						return err
					}
					benchmark.DefaultStore.SetSharedTmpResult(avgr, avgw)
					return nil
				}),
			)
			r.Put(
				"/disk-tmp",
				benchmark.NewIORHandler(func(avgr, avgw *ior.Result, err error) error {
					if err != nil {
						benchmark.DefaultStore.SetFailure(err)
						return err
					}
					benchmark.DefaultStore.SetDiskTmpResult(avgr, avgw)
					return nil
				}),
			)
			r.Put(
				"/disk-world-tmp",
				benchmark.NewIORHandler(func(avgr, avgw *ior.Result, err error) error {
					if err != nil {
						benchmark.DefaultStore.SetFailure(err)
						return err
					}
					benchmark.DefaultStore.SetDiskWorldTmpResult(avgr, avgw)
					return nil
				}),
			)
		})
		r.Route("/hpl", func(r chi.Router) {
			r.Put(
				"/phase1",
				benchmark.NewHPLPhase1Handler(
					func(optimal *hpl.Result, err error) error {
						if err != nil {
							benchmark.DefaultStore.SetFailure(err)
							return err
						}

						benchmark.DefaultStore.SetGFLOPS(optimal.Gflops)

						return nil
					},
				),
			)
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

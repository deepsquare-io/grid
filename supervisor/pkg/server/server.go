package server

import (
	"net"
	"net/http"
	"strings"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/generated/grpc/health/v1"
	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/health"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/jobapi"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/sshapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type Server struct {
	http1Router *chi.Mux
	http2       *http2.Server
	http1       *http.Server
}

func New(
	ms metascheduler.MetaScheduler,
	resourceManager *lock.ResourceManager,
	pkB64 string,
	opts ...grpc.ServerOption,
) *Server {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
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

	return &Server{
		http1Router: r,
		http1:       http1Server,
		http2:       http2Server,
	}
}

func (s *Server) AddBenchmarkRoutes(
	ms metascheduler.MetaScheduler,
	scheduler scheduler.Scheduler,
	benchmarkLauncher benchmark.Launcher,
) {
	s.http1Router.Route("/benchmark", func(r chi.Router) {
		r.Put("/phase1", benchmark.NewPhase1Handler(benchmarkLauncher))
		r.Put("/phase2", benchmark.NewPhase2Handler(benchmarkLauncher, scheduler, ms))
	})
}

func (s *Server) Serve(l net.Listener) error {
	return s.http1.Serve(l)
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

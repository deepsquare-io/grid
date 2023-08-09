package benchmark

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/result"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"go.uber.org/zap"
)

func NewPhase1Handler(benchmark Launcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretB64 := r.Header.Get("X-Secret")
		data, err := base64.StdEncoding.DecodeString(secretB64)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("bad request: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		if !benchmark.Verify(data) {
			http.Error(
				w,
				fmt.Sprintf("invalid secret: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		nodesStr := r.URL.Query().Get("nodes")
		nodes, err := strconv.ParseUint(nodesStr, 10, 64)
		if err != nil {
			logger.I.Error("failed to convert query nodes to int", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}
		reader := result.NewReader(r.Body)

		optimal, err := result.FindMaxGflopsResult(reader)
		if err != nil {
			logger.I.Error("failed to find max Gflops", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			if err := benchmark.RunPhase2(ctx, optimal.P, optimal.Q, optimal.ProblemSize, optimal.NB, nodes); err != nil {
				logger.I.Fatal(
					"phase2 benchmark failed or failed to be tracked",
					zap.Error(err),
				)
			}
			logger.I.Info(
				"benchmark phase 2 succeeded",
				zap.Error(err),
				zap.Uint64(
					"p",
					optimal.P,
				),
				zap.Uint64("q", optimal.Q),
				zap.Uint64("n", optimal.ProblemSize),
				zap.Uint64("nb", optimal.NB),
				zap.Uint64("nodes", nodes),
			)
		}()

		fmt.Fprint(w, "success")
	}
}

func NewPhase2Handler(
	benchmark Launcher,
	scheduler scheduler.Scheduler,
	ms metascheduler.MetaScheduler,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretB64 := r.Header.Get("X-Secret")
		data, err := base64.StdEncoding.DecodeString(secretB64)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("bad request: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		if !benchmark.Verify(data) {
			http.Error(
				w,
				fmt.Sprintf("invalid secret: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		reader := result.NewReader(r.Body)
		flops, err := result.ComputeAvgGflopsResult(reader)
		if err != nil {
			logger.I.Error("failed to compute avg", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			nodes, err := scheduler.FindTotalNodes(ctx)
			if err != nil {
				logger.I.Fatal("failed to check total number of nodes", zap.Error(err))
			}
			cpus, err := scheduler.FindTotalCPUs(ctx)
			if err != nil {
				logger.I.Fatal("failed to find total number of cpus", zap.Error(err))
			}
			gpus, err := scheduler.FindTotalGPUs(ctx)
			if err != nil {
				logger.I.Fatal("failed to find total number of gpus", zap.Error(err))
			}
			mem, err := scheduler.FindTotalMem(ctx)
			if err != nil {
				logger.I.Fatal("failed to find total mem", zap.Error(err))
			}
			if err := ms.Register(ctx, nodes, cpus, gpus, mem, flops); err != nil {
				logger.I.Fatal("failed to register", zap.Error(err))
			}
		}()

	}
}

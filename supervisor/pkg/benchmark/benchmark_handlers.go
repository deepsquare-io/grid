package benchmark

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/hpl"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/ior"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/osu"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/speedtest"
	"go.uber.org/zap"
)

func NewIORHandler(
	next func(avgr *ior.Result, avgw *ior.Result) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		avgr, avgw, err := ior.ComputeAvgReadWrite(ior.NewReader(r.Body))
		if err != nil {
			logger.I.Error("failed to parse osu logs", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		if err := next(avgr, avgw); err != nil {
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		fmt.Fprint(w, "success")
	}
}

func NewOSUHandler(
	next func(res float64) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := osu.ParseOSULog(r.Body)
		if err != nil {
			logger.I.Error("failed to parse osu logs", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		if err := next(res); err != nil {
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		fmt.Fprint(w, "success")
	}
}

func NewSpeedTestHandler(
	next func(res *speedtest.Result) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res speedtest.Result
		if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
			logger.I.Error("failed to parse body as speedtest.Result", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		if err := next(&res); err != nil {
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		fmt.Fprint(w, "success")
	}
}

func NewHPLPhase1Handler(
	next func(
		optimal *hpl.Result,
		opts ...Option,
	) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nodes, err := strconv.ParseUint(r.URL.Query().Get("nodes"), 10, 64)
		if err != nil {
			logger.I.Error("failed to parse nodes", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("bad request: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		cpusPerNode, err := strconv.ParseUint(r.URL.Query().Get("cpusPerNode"), 10, 64)
		if err != nil {
			logger.I.Error("failed to parse cpusPerNode", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("bad request: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		gpusPerNode, err := strconv.ParseUint(r.URL.Query().Get("gpusPerNode"), 10, 64)
		if err != nil {
			logger.I.Error("failed to parse gpusPerNode", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("bad request: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		memPerNode, err := strconv.ParseUint(r.URL.Query().Get("memPerNode"), 10, 64)
		if err != nil {
			logger.I.Error("failed to parse memPerNode", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("bad request: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		reader := hpl.NewReader(r.Body)

		optimal, err := hpl.FindMaxGflopsResult(reader)
		if err != nil {
			logger.I.Error("failed to find max Gflops", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		opts := []Option{
			WithClusterSpecs(nodes, cpusPerNode, gpusPerNode, memPerNode),
		}

		if err := next(optimal, opts...); err != nil {
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		fmt.Fprint(w, "success")
	}
}

func NewHPLPhase2Handler(
	next func(gflops float64) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reader := hpl.NewReader(r.Body)
		gflops, err := hpl.ComputeAvgGflopsResult(reader)
		if err != nil {
			logger.I.Error("failed to compute avg", zap.Error(err))
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		if err := next(gflops); err != nil {
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

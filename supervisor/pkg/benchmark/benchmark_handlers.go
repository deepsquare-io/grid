package benchmark

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/hpl"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/ior"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/osu"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/speedtest"
	"go.uber.org/zap"
)

type MachineSpec struct {
	MicroArch string `json:"microarch"`
	OS        string `json:"os"`
	CPU       string `json:"cpu"`
	Arch      string `json:"arch"`
}

func NewMachineHandler(
	next func(res *MachineSpec, err error) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var spec MachineSpec
		err := json.NewDecoder(r.Body).Decode(&spec)
		if err != nil {
			logger.I.Error("failed to decode machine specifications", zap.Error(err))
		}

		if err := next(&spec, err); err != nil {
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

func NewIORHandler(
	next func(avgr *ior.Result, avgw *ior.Result, err error) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		avgr, avgw, err := ior.ComputeAvgReadWrite(ior.NewReader(r.Body))
		if err != nil {
			logger.I.Error("failed to parse osu logs", zap.Error(err))
		}

		if err := next(avgr, avgw, err); err != nil {
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
	next func(res float64, err error) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := osu.ParseOSULog(r.Body)
		if err != nil {
			logger.I.Error("failed to parse osu logs", zap.Error(err))
		}

		if err := next(res, err); err != nil {
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
	next func(res *speedtest.Result, err error) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res speedtest.Result
		err := json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			logger.I.Error("failed to parse body as speedtest.Result", zap.Error(err))
		}

		if err := next(&res, err); err != nil {
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
	next func(optimal *hpl.Result, err error) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		reader := hpl.NewReader(r.Body)

		optimal, err := hpl.FindMaxGflopsResult(reader)
		if err != nil {
			logger.I.Error("failed to find max Gflops", zap.Error(err))
		}

		if math.IsNaN(optimal.Gflops) || math.IsInf(optimal.Gflops, 0) {
			err = errors.New("gflops is an invalid value")
		}

		if err := next(optimal, err); err != nil {
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

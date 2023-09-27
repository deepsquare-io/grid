// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package benchmark

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/hpl"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/ior"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/osu"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/speedtest"
	"go.uber.org/zap"
)

type MachineSpec struct {
	MicroArch string `json:"microarch"`
	OS        string `json:"os"`
	CPU       string `json:"cpu"`
	Arch      string `json:"arch"`
	GPU       string `json:"gpu"`
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

		fmt.Fprintln(w, "success")
	}
}

func NewIORHandler(
	next func(avgr *ior.Result, avgw *ior.Result, err error) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		avgr, avgw, err := ior.ComputeAvgReadWrite(ior.NewReader(r.Body))
		if err != nil {
			logger.I.Error("failed to parse ior logs", zap.Error(err))
		}

		if err := next(avgr, avgw, err); err != nil {
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		fmt.Fprintln(w, "success")
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

		if res == 0 {
			err = errors.New("OSU benchmark received a 0 value")
		}

		if err := next(res, err); err != nil {
			http.Error(
				w,
				fmt.Sprintf("internal server error: %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		fmt.Fprintln(w, "success")
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

		fmt.Fprintln(w, "success")
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

		fmt.Fprintln(w, "success")
	}
}

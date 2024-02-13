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

//go:build integration

package benchmark_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"testing"
	"time"

	_ "embed"

	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/hpl"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/ior"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/speedtest"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed hpl/fixtures/1n2gpu16cpu.log
	hplFixture []byte
	//go:embed osu/fixtures/alltoall.log
	osuFixture []byte
	//go:embed speedtest/fixtures/result.json
	speedtestFixture []byte
	//go:embed ior/fixtures/result.csv
	iorFixture []byte
)

func TestHPLHandler(t *testing.T) {
	// Arrange
	// Prepare fixture
	file, err := os.CreateTemp("", "test-tmp")
	require.NoError(t, err)
	file.Write(hplFixture)
	err = file.Close()
	require.NoError(t, err)
	defer os.Remove(file.Name())
	fmt.Printf("created tmp fixture: %s", file.Name())

	impl := benchmark.NewHPLHandler(
		func(optimal *hpl.Result, err error) error {
			require.NotEmpty(t, optimal)
			require.NoError(t, err)
			return nil
		},
	)

	// Arrangements to handle race conditions and asserts
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	http.HandleFunc("/benchmark/hpl", func(w http.ResponseWriter, r *http.Request) {
		impl(w, r)
		out, err := httputil.DumpRequest(r, false)
		require.NoError(t, err)
		fmt.Printf("received request\n%s", string(out))
		done <- struct{}{}
	})

	// Act
	port := generateRandomPort()
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
		require.NoError(t, err)
	}()

	cmd := exec.Command(
		"curl",
		"-fsSL",
		"-H",
		fmt.Sprintf("X-Secret: %s", base64.StdEncoding.EncodeToString(secret.Get())),
		"--upload-file",
		file.Name(),
		fmt.Sprintf(
			"http://localhost:%d/benchmark/hpl?nodes=1&cpusPerNode=16&gpusPerNode=2&memPerNode=100000",
			port,
		),
	)
	_, err = cmd.CombinedOutput()
	require.NoError(t, err)

	select {
	case <-done:
		// Pass
	case <-ctx.Done():
		require.NoError(t, ctx.Err())
	}
}

func TestSpeedTestHandler(t *testing.T) {
	// Arrange
	// Prepare fixture
	file, err := os.CreateTemp("", "test-tmp")
	require.NoError(t, err)
	file.Write(speedtestFixture)
	err = file.Close()
	require.NoError(t, err)
	defer os.Remove(file.Name())
	fmt.Printf("created tmp fixture: %s", file.Name())

	// Mocks
	impl := benchmark.NewSpeedTestHandler(func(res *speedtest.Result, err error) error {
		require.NotEmpty(t, res)
		require.NoError(t, err)
		return nil
	})

	// Arrangements to handle race conditions and asserts
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	http.HandleFunc("/benchmark/speedtest", func(w http.ResponseWriter, r *http.Request) {
		impl(w, r)
		out, err := httputil.DumpRequest(r, false)
		require.NoError(t, err)
		fmt.Printf("received request\n%s", string(out))
		done <- struct{}{}
	})

	// Act
	port := generateRandomPort()
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
		require.NoError(t, err)
	}()

	cmd := exec.Command(
		"curl",
		"-fsSL",
		"-H",
		fmt.Sprintf("X-Secret: %s", base64.StdEncoding.EncodeToString(secret.Get())),
		"--upload-file",
		file.Name(),
		fmt.Sprintf("http://localhost:%d/benchmark/speedtest", port),
	)
	out, err := cmd.CombinedOutput()
	require.NoError(t, err)
	fmt.Println(string(out))

	select {
	case <-done:
		// Pass
	case <-ctx.Done():
		require.NoError(t, ctx.Err())
	}
}

func TestOSUHandler(t *testing.T) {
	// Arrange
	// Prepare fixture
	file, err := os.CreateTemp("", "test-tmp")
	require.NoError(t, err)
	file.Write(osuFixture)
	err = file.Close()
	require.NoError(t, err)
	defer os.Remove(file.Name())
	fmt.Printf("created tmp fixture: %s", file.Name())

	// Mocks
	impl := benchmark.NewOSUHandler(func(res float64, err error) error {
		require.NotEmpty(t, res)
		require.NoError(t, err)
		return nil
	})

	// Arrangements to handle race conditions and asserts
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	http.HandleFunc("/benchmark/osu", func(w http.ResponseWriter, r *http.Request) {
		impl(w, r)
		out, err := httputil.DumpRequest(r, false)
		require.NoError(t, err)
		fmt.Printf("received request\n%s", string(out))
		done <- struct{}{}
	})

	// Act
	port := generateRandomPort()
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
		require.NoError(t, err)
	}()

	cmd := exec.Command(
		"curl",
		"-fsSL",
		"-H",
		fmt.Sprintf("X-Secret: %s", base64.StdEncoding.EncodeToString(secret.Get())),
		"--upload-file",
		file.Name(),
		fmt.Sprintf("http://localhost:%d/benchmark/osu", port),
	)
	out, err := cmd.CombinedOutput()
	require.NoError(t, err)
	fmt.Println(string(out))

	select {
	case <-done:
		// Pass
	case <-ctx.Done():
		require.NoError(t, ctx.Err())
	}
}

func TestIORHandler(t *testing.T) {
	// Arrange
	// Prepare fixture
	file, err := os.CreateTemp("", "test-tmp")
	require.NoError(t, err)
	file.Write(iorFixture)
	err = file.Close()
	require.NoError(t, err)
	defer os.Remove(file.Name())
	fmt.Printf("created tmp fixture: %s", file.Name())

	// Mocks
	impl := benchmark.NewIORHandler(func(avgr, avgw *ior.Result, err error) error {
		require.NotEmpty(t, avgr)
		require.NotEmpty(t, avgw)
		require.NoError(t, err)
		return nil
	})

	// Arrangements to handle race conditions and asserts
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	http.HandleFunc("/benchmark/ior", func(w http.ResponseWriter, r *http.Request) {
		impl(w, r)
		out, err := httputil.DumpRequest(r, false)
		require.NoError(t, err)
		fmt.Printf("received request\n%s", string(out))
		done <- struct{}{}
	})

	// Act
	port := generateRandomPort()
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
		require.NoError(t, err)
	}()

	cmd := exec.Command(
		"curl",
		"-fsSL",
		"-H",
		fmt.Sprintf("X-Secret: %s", base64.StdEncoding.EncodeToString(secret.Get())),
		"--upload-file",
		file.Name(),
		fmt.Sprintf("http://localhost:%d/benchmark/ior", port),
	)
	out, err := cmd.CombinedOutput()
	require.NoError(t, err)
	fmt.Println(string(out))

	select {
	case <-done:
		// Pass
	case <-ctx.Done():
		require.NoError(t, ctx.Err())
	}
}

func generateRandomPort() int {
	minPort := 49152
	maxPort := 65535
	return rand.Intn(maxPort-minPort+1) + minPort
}

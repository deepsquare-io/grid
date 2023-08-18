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

	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/hpl"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/speedtest"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed hpl/fixtures/1n2gpu16cpu.log
	hplFixture []byte
	//go:embed osu/fixtures/alltoall.log
	osuFixture []byte
	//go:embed speedtest/fixtures/result.json
	speedtestFixture []byte
)

func TestPhase1Handler(t *testing.T) {
	// Arrange
	// Prepare fixture
	file, err := os.CreateTemp("", "test-tmp")
	require.NoError(t, err)
	file.Write(hplFixture)
	err = file.Close()
	require.NoError(t, err)
	defer os.Remove(file.Name())
	fmt.Printf("created tmp fixture: %s", file.Name())

	impl := benchmark.NewHPLPhase1Handler(
		func(optimal *hpl.Result, opts ...benchmark.BenchmarkOption) error {
			require.NotEmpty(t, optimal)
			return nil
		},
	)

	// Arrangements to handle race conditions and asserts
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	http.HandleFunc("/benchmark/hpl/phase1", func(w http.ResponseWriter, r *http.Request) {
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
			"http://localhost:%d/benchmark/hpl/phase1?nodes=1&cpusPerNode=16&gpusPerNode=2&memPerNode=100000",
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

func TestPhase2Handler(t *testing.T) {
	// Arrange
	// Prepare fixture
	file, err := os.CreateTemp("", "test-tmp")
	require.NoError(t, err)
	file.Write(hplFixture)
	err = file.Close()
	require.NoError(t, err)
	defer os.Remove(file.Name())
	fmt.Printf("created tmp fixture: %s", file.Name())

	// Mocks
	impl := benchmark.NewHPLPhase2Handler(func(gflops float64) error {
		require.NotEmpty(t, gflops)
		return nil
	})

	// Arrangements to handle race conditions and asserts
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	http.HandleFunc("/benchmark/hpl/phase2", func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Sprintf("http://localhost:%d/benchmark/hpl/phase2", port),
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
	impl := benchmark.NewSpeedTestHandler(func(res *speedtest.Result) error {
		require.NotEmpty(t, res)
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
	impl := benchmark.NewOSUHandler(func(res float64) error {
		require.NotEmpty(t, res)
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

func generateRandomPort() int {
	minPort := 49152
	maxPort := 65535
	return rand.Intn(maxPort-minPort+1) + minPort
}

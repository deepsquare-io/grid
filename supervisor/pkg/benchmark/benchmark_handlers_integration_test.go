//go:build integration

package benchmark_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"testing"
	"time"

	_ "embed"

	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockbenchmark"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockmetascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockscheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:embed fixtures/1n2gpu16cpu.log
var fixture []byte

func TestPhase1Handler(t *testing.T) {
	// Arrange
	// Prepare fixture
	file, err := os.CreateTemp("", "test-tmp")
	require.NoError(t, err)
	file.Write(fixture)
	err = file.Close()
	require.NoError(t, err)
	defer os.Remove(file.Name())
	fmt.Printf("created tmp fixture: %s", file.Name())

	// Mocks
	launcher := mockbenchmark.NewLauncher(t)
	impl := benchmark.NewPhase1Handler(launcher)
	// Expect verify
	launcher.EXPECT().Verify([]byte("SECRET")).Return(true)
	// Expect to launch phase 2
	launcher.EXPECT().
		RunPhase2(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	// Arrangements to handle race conditions and asserts
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	http.HandleFunc("/benchmark/phase1", func(w http.ResponseWriter, r *http.Request) {
		impl(w, r)
		out, err := httputil.DumpRequest(r, false)
		require.NoError(t, err)
		fmt.Printf("received request\n%s", string(out))
		done <- struct{}{}
	})

	// Act
	go http.ListenAndServe(":3000", nil)

	cmd := exec.Command(
		"curl",
		"-sS",
		"-H",
		"X-Secret: U0VDUkVU",
		"--upload-file",
		file.Name(),
		"http://localhost:3000/benchmark/phase1",
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
	file.Write(fixture)
	err = file.Close()
	require.NoError(t, err)
	defer os.Remove(file.Name())
	fmt.Printf("created tmp fixture: %s", file.Name())

	// Mocks
	doneRegister := make(chan struct{})
	launcher := mockbenchmark.NewLauncher(t)
	ms := mockmetascheduler.NewMetaScheduler(t)
	scheduler := mockscheduler.NewScheduler(t)
	impl := benchmark.NewPhase2Handler(launcher, scheduler, ms)
	// Expect verify
	launcher.EXPECT().Verify([]byte("SECRET")).Return(true)
	scheduler.EXPECT().FindTotalNodes(mock.Anything).Return(1, nil)
	scheduler.EXPECT().FindTotalCPUs(mock.Anything).Return(16, nil)
	scheduler.EXPECT().FindTotalGPUs(mock.Anything).Return(2, nil)
	scheduler.EXPECT().FindTotalMem(mock.Anything).Return(128000, nil)
	ms.EXPECT().
		Register(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, u1, u2, u3, u4 uint64, f float64) error {
			doneRegister <- struct{}{}
			return nil
		})

	// Arrangements to handle race conditions and asserts
	done := make(chan struct{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	http.HandleFunc("/benchmark/phase2", func(w http.ResponseWriter, r *http.Request) {
		impl(w, r)
		out, err := httputil.DumpRequest(r, false)
		require.NoError(t, err)
		fmt.Printf("received request\n%s", string(out))
		done <- struct{}{}
	})

	// Act
	go http.ListenAndServe(":3000", nil)

	cmd := exec.Command(
		"curl",
		"-sS",
		"-H",
		"X-Secret: U0VDUkVU",
		"--upload-file",
		file.Name(),
		"http://localhost:3000/benchmark/phase2",
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

	select {
	case <-doneRegister:
		// Pass
	case <-ctx.Done():
		require.NoError(t, ctx.Err())
	}
}

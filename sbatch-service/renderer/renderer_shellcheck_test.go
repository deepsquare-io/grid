package renderer_test

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func shellcheck(t *testing.T, script string) {
	_, err := exec.LookPath("shellcheck")
	if err != nil {
		logger.I.Warn("shellcheck is disabled, test is not complete")
		return
	}
	if err := os.WriteFile("test.sh", []byte(script), 0o777); err != nil {
		logger.I.Panic("failed to write", zap.Error(err))
	}
	out, err := exec.Command("shellcheck", "-S", "warning", "-s", "bash", "test.sh").CombinedOutput()
	if err != nil {
		logger.I.Error(string(out))
		require.NoError(t, errors.New("shellcheck failed"))
	}

	_ = os.Remove("test.sh")
}

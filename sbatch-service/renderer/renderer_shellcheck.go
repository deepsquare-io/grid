package renderer

import (
	"errors"
	"os"
	"os/exec"

	"github.com/deepsquare-io/grid/sbatch-service/logger"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
)

func Shellcheck(script string) error {
	name := utils.GenerateRandomString(8)
	_, err := exec.LookPath("shellcheck")
	if err != nil {
		logger.I.Warn("shellcheck is disabled, test is not complete")
		return nil
	}
	if err := os.WriteFile(name, []byte(script), 0o777); err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(name)
	}()
	out, err := exec.Command("shellcheck", "-S", "warning", "-s", "bash", name).CombinedOutput()
	if err != nil {
		logger.I.Error(string(out))
		return errors.New("shellcheck failed")
	}
	if string(out) != "" {
		logger.I.Warn(string(out))
	}
	return nil
}

// Copyright (C) 2024 DeepSquare Association
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

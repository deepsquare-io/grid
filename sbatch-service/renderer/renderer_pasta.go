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

package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
)

//go:embed renderer_pasta.sh.tpl
var pastaTpl string

func RenderPastaNS(
	run *model.StepRun,
	job *model.Job,
	command string,
) (string, error) {
	for _, nic := range run.CustomNetworkInterfaces {
		if err := validate.I.Struct(nic); err != nil {
			return "", err
		}
	}

	tmpl, err := engine().Parse(pastaTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Run     *model.StepRun
		Job     *model.Job
		Command string
	}{
		Run:     run,
		Job:     job,
		Command: command,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

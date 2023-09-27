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

//go:embed renderer_slirp4netns.sh.tpl
var slirp4netnsTpl string

func RenderSlirp4NetNS(
	i []*model.NetworkInterface,
	dns []string,
	command string,
	shell *string,
) (string, error) {
	for _, nic := range i {
		if err := validate.I.Struct(nic); err != nil {
			return "", err
		}
	}

	tmpl, err := engine().Parse(slirp4netnsTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		NICs    []*model.NetworkInterface
		DNS     []string
		Command string
		Shell   *string
	}{
		NICs:    i,
		DNS:     dns,
		Command: command,
		Shell:   shell,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

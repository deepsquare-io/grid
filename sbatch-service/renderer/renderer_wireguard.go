// Copyright (C) 2023 DeepSquare Asociation
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

//go:embed renderer_wireguard.sh.tpl
var wireguardTpl string

func RenderWireguard(w *model.Wireguard, interfaceName string) (string, error) {
	if err := validate.I.Struct(w); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(wireguardTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Wireguard     *model.Wireguard
		InterfaceName string
	}{
		Wireguard:     w,
		InterfaceName: interfaceName,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

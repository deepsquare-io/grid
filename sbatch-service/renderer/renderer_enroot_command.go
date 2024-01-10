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
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

//go:embed renderer_enroot_command.sh.tpl
var enrootTpl string

func RenderEnrootCommand(r *model.StepRun) (string, error) {
	if err := validate.I.Struct(r); err != nil {
		return "", err
	}

	var image *ocispec.Image
	if r.Container != nil && r.Container.Image != "" &&
		(r.Container.DeepsquareHosted == nil || !*r.Container.DeepsquareHosted) {
		i, err := validate.DefaultImageFetcher.FetchContainerImage(
			utils.SafeDeref(r.Container.Username),
			utils.SafeDeref(r.Container.Password),
			utils.SafeDeref(r.Container.Registry),
			r.Container.Image,
		)
		if err != nil {
			return "", err
		}
		image = &i
	}

	tmpl, err := engine().Parse(enrootTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Run   *model.StepRun
		Image *ocispec.Image
	}{
		Run:   r,
		Image: image,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

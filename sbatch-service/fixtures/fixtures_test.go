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

package fixtures_test

import (
	"testing"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/renderer"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func init() {
	utils.MockRandomString("RANDOM_STRING")
}

var (
	//go:embed tdp.yaml
	fixtureTDP string
	//go:embed tdp.txt
	expectedTDP string
	//go:embed urs.yaml
	fixtureURS string
	//go:embed urs.txt
	expectedURS string
	//go:embed blender-batch-job.yaml
	fixtureBlenderBatchJob string
	//go:embed blender-batch-job.txt
	expectedBlenderBatchJob string
	//go:embed upscale.yaml
	fixtureUpscale string
	//go:embed upscale.txt
	expectedUpscale string
	//go:embed stable-diffusion.yaml
	fixtureStableDiffusion string
	//go:embed stable-diffusion.txt
	expectedStableDiffusion string
	//go:embed minecraft.yaml
	fixtureMinecraft string
	//go:embed minecraft.txt
	expectedMinecraft string
	//go:embed map-uid.yaml
	fixtureMapUID string
	//go:embed map-uid.txt
	expectedMapUID string
	//go:embed cs2.yaml
	fixtureCS2 string
	//go:embed cs2.txt
	expectedCS2 string
	//go:embed virtual-networks.yaml
	fixtureVirtualNetworks string
	//go:embed virtual-networks.txt
	expectedVirtualNetworks string
	//go:embed tgi.yaml
	fixtureTGI string
	//go:embed tgi.txt
	expectedTGI string

	r = renderer.NewJobRenderer(
		"logger.example.com:443",
		"/usr/local/bin/grid-logger-writer",
	)
)

func TestFixtures(t *testing.T) {
	tests := []struct {
		name     string
		fixture  string
		expected string
	}{
		{
			name:     "tdp",
			fixture:  fixtureTDP,
			expected: expectedTDP,
		},
		{
			name:     "urs",
			fixture:  fixtureURS,
			expected: expectedURS,
		},
		{
			name:     "blender",
			fixture:  fixtureBlenderBatchJob,
			expected: expectedBlenderBatchJob,
		},
		{
			name:     "upscale",
			fixture:  fixtureUpscale,
			expected: expectedUpscale,
		},
		{
			name:     "stable-diffusion",
			fixture:  fixtureStableDiffusion,
			expected: expectedStableDiffusion,
		},
		{
			name:     "minecraft",
			fixture:  fixtureMinecraft,
			expected: expectedMinecraft,
		},
		{
			name:     "MapUID",
			fixture:  fixtureMapUID,
			expected: expectedMapUID,
		},
		{
			name:     "CS2",
			fixture:  fixtureCS2,
			expected: expectedCS2,
		},
		{
			name:     "virtual-network",
			fixture:  fixtureVirtualNetworks,
			expected: expectedVirtualNetworks,
		},
		{
			name:     "tgi",
			fixture:  fixtureTGI,
			expected: expectedTGI,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := struct {
				Job model.Job
			}{}
			err := yaml.Unmarshal([]byte(tt.fixture), &j)
			require.NoError(t, err)

			out, err := r.RenderJob(&j.Job)
			require.NoError(t, err)
			require.Equal(t, tt.expected, out)
			require.NoError(t, renderer.Shellcheck(out))
		})
	}
}

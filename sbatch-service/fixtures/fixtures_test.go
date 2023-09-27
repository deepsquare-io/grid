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

package fixtures_test

import (
	"fmt"
	"testing"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/renderer"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

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

	r = renderer.NewJobRenderer(
		"logger.example.com:443",
		"/usr/local/bin/grid-logger-writer",
	)
)

func TestRenderTDP(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureTDP), &j)
	require.NoError(t, err)

	out, err := r.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedTDP, out)
	require.NoError(t, renderer.Shellcheck(out))
}

func TestRenderURS(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureURS), &j)
	require.NoError(t, err)

	out, err := r.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedURS, out)
	require.NoError(t, renderer.Shellcheck(out))
}

func TestRenderBlenderBatchJob(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureBlenderBatchJob), &j)
	require.NoError(t, err)

	out, err := r.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedBlenderBatchJob, out)
	require.NoError(t, renderer.Shellcheck(out))
}

func TestUpscaleJob(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureUpscale), &j)
	require.NoError(t, err)

	out, err := r.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedUpscale, out)
	require.NoError(t, renderer.Shellcheck(out))
}

func TestStableDiffusionJob(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureStableDiffusion), &j)
	require.NoError(t, err)

	out, err := r.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedStableDiffusion, out)
	require.NoError(t, renderer.Shellcheck(out))
}

func TestMinecraftJob(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureMinecraft), &j)
	require.NoError(t, err)

	out, err := r.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedMinecraft, out)
	require.NoError(t, renderer.Shellcheck(out))
}

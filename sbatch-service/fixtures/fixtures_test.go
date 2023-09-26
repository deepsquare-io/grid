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

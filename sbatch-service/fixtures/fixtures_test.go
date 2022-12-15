package fixtures_test

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
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

func TestRenderTDP(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureTDP), &j)
	require.NoError(t, err)

	out, err := renderer.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedTDP, out)
	shellcheck(t, out)
}

func TestRenderURS(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureURS), &j)
	require.NoError(t, err)

	out, err := renderer.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedURS, out)
	shellcheck(t, out)
}

func TestRenderBlenderBatchJob(t *testing.T) {
	j := struct {
		Job model.Job
	}{}
	err := yaml.Unmarshal([]byte(fixtureBlenderBatchJob), &j)
	require.NoError(t, err)

	out, err := renderer.RenderJob(&j.Job)
	require.NoError(t, err)
	fmt.Println(out)
	require.Equal(t, expectedBlenderBatchJob, out)
	shellcheck(t, out)
}

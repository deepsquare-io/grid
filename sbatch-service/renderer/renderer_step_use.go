package renderer

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/module"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
	"github.com/lithammer/shortuuid/v4"
)

//go:embed renderer_step_use.sh.tpl
var stepUseTpl string

type StepUseRenderer struct {
	enc shortuuid.Encoder
}

func NewStepUseRenderer(enc shortuuid.Encoder) *StepUseRenderer {
	return &StepUseRenderer{
		enc: enc,
	}
}

func (r *StepUseRenderer) Render(
	job *model.Job,
	step *model.Step,
	use *model.StepUse,
) (string, error) {
	ctx := context.Background()

	if err := validate.I.Struct(job); err != nil {
		return "", err
	}

	if err := validate.I.Struct(step); err != nil {
		return "", err
	}

	if err := validate.I.Struct(use); err != nil {
		return "", err
	}

	repository, ref, ok := strings.Cut(use.Source, "@")
	if !ok {
		repository = use.Source
		ref = ""
	}

	m, err := module.Resolve(ctx, job, step, repository, ref)
	if err != nil {
		return "", err
	}

	if err := validate.I.Struct(m); err != nil {
		return "", err
	}

	// Assert requirements
	if err := assertRequirements(job, use, m); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(stepUseTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Job    *model.Job
		Step   *model.Step
		Use    *model.StepUse
		Module *model.Module
		UUID   string
	}{
		Job:    job,
		Step:   step,
		Use:    use,
		Module: m,
		UUID:   shortuuid.NewWithEncoder(r.enc),
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

func assertRequirements(j *model.Job, u *model.StepUse, m *model.Module) error {
	if j.Resources.CpusPerTask < m.MinimumResources.CpusPerTask {
		return fmt.Errorf("not enough cpu per task to use the module %s, minimum=%d, actual=%d", u.Source, m.MinimumResources.CpusPerTask, j.Resources.CpusPerTask)
	}
	if j.Resources.MemPerCPU < m.MinimumResources.MemPerCPU {
		return fmt.Errorf("not enough mem per cpu to use the module %s, minimum=%d, actual=%d", u.Source, m.MinimumResources.MemPerCPU, j.Resources.MemPerCPU)
	}
	if j.Resources.GpusPerTask < m.MinimumResources.GpusPerTask {
		return fmt.Errorf("not enough gpu per task to use the module %s, minimum=%d, actual=%d", u.Source, m.MinimumResources.GpusPerTask, j.Resources.GpusPerTask)
	}
	if j.Resources.Tasks < m.MinimumResources.Tasks {
		return fmt.Errorf("not enough tasks to use the module %s, minimum=%d, actual=%d", u.Source, m.MinimumResources.Tasks, j.Resources.Tasks)
	}
	return nil
}

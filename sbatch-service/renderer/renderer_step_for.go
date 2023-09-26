package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
)

//go:embed renderer_step_for.sh.tpl
var forTpl string

func RenderStepFor(j *model.Job, f *model.StepFor) (string, error) {
	if err := validate.I.Struct(j); err != nil {
		return "", err
	}

	if err := validate.I.Struct(f); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(forTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Job *model.Job
		For *model.StepFor
	}{
		Job: j,
		For: f,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

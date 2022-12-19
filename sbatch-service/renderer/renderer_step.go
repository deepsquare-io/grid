package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
)

//go:embed renderer_step.sh.tpl
var stepTpl string

func RenderStep(j *model.Job, s *model.Step) (string, error) {
	if err := validate.I.Struct(j); err != nil {
		return "", err
	}

	if err := validate.I.Struct(s); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(stepTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Job  *model.Job
		Step *model.Step
	}{
		Job:  j,
		Step: s,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}
package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
)

//go:embed renderer_step_async_launch.sh.tpl
var asyncLaunchTpl string

func RenderStepAsyncLaunch(j *model.Job, l *model.StepAsyncLaunch) (string, error) {
	if err := validate.I.Struct(j); err != nil {
		return "", err
	}

	if err := validate.I.Struct(l); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(asyncLaunchTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Job    *model.Job
		Launch *model.StepAsyncLaunch
	}{
		Job:    j,
		Launch: l,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

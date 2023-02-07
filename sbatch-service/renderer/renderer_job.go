package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
)

//go:embed renderer_job.sh.tpl
var jobTpl string

func RenderJob(j *model.Job) (string, error) {
	if err := validate.I.Struct(j); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(jobTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Job *model.Job
	}{
		Job: j,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

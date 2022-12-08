package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/template"
)

var (
	//go:embed step.sh.tpl
	stepTpl string
)

func RenderStep(s *model.Step) (string, error) {
	if err := s.Validate(); err != nil {
		return "", err
	}

	tmpl, err := template.Init().Parse(stepTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, s); err != nil {
		return "", err
	}
	return out.String(), nil
}

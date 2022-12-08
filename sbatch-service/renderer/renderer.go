package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
)

var (
	//go:embed step.sh.tpl
	stepTpl string
)

func RenderStep(s *model.Step) (string, error) {
	if err := s.Validate(); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(stepTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, s); err != nil {
		return "", err
	}
	return out.String(), nil
}

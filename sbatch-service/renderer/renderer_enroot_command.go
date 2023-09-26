package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
)

//go:embed renderer_enroot_command.sh.tpl
var enrootTpl string

func RenderEnrootCommand(r *model.StepRun) (string, error) {
	if err := validate.I.Struct(r); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(enrootTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Run *model.StepRun
	}{
		Run: r,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

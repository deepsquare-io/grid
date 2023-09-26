package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
)

//go:embed renderer_wireguard.sh.tpl
var wireguardTpl string

func RenderWireguard(w *model.Wireguard, interfaceName string) (string, error) {
	if err := validate.I.Struct(w); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(wireguardTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Wireguard     *model.Wireguard
		InterfaceName string
	}{
		Wireguard:     w,
		InterfaceName: interfaceName,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

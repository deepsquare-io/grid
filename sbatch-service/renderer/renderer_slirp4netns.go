package renderer

import (
	"bytes"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
)

//go:embed renderer_slirp4netns.sh.tpl
var slirp4netnsTpl string

func RenderSlirp4NetNS(
	i []*model.NetworkInterface,
	dns []string,
	command string,
	shell *string,
) (string, error) {
	for _, nic := range i {
		if err := validate.I.Struct(nic); err != nil {
			return "", err
		}
	}

	tmpl, err := engine().Parse(slirp4netnsTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		NICs    []*model.NetworkInterface
		DNS     []string
		Command string
		Shell   *string
	}{
		NICs:    i,
		DNS:     dns,
		Command: command,
		Shell:   shell,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

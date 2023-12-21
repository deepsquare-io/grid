package renderer

import (
	"bytes"
	"net"
	"slices"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
)

//go:embed renderer_virtual_network.sh.tpl
var networkTpl string

func RenderVirtualNetwork(
	vnet *model.VirtualNetwork,
	index int,
	steps []*model.Step,
) (string, error) {
	if err := validate.I.Struct(vnet); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(networkTpl)
	if err != nil {
		return "", err
	}

	peers, err := searchPeers(vnet.Name, steps)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		VirtualNetwork *model.VirtualNetwork
		Index          int
		Peers          []string
	}{
		VirtualNetwork: vnet,
		Index:          index,
		Peers:          peers,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

func searchPeers(name string, steps []*model.Step) ([]string, error) {
	peers := make(map[string]bool)
	if err := recurseSearchPeers(name, steps, peers); err != nil {
		return []string{}, err
	}
	keys := make([]string, 0, len(peers))
	for key := range peers {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys, nil
}

func recurseSearchPeers(name string, steps []*model.Step, opeers map[string]bool) error {
	for _, step := range steps {
		if err := recurseSearchPeers(name, step.Catch, opeers); err != nil {
			return err
		}
		if err := recurseSearchPeers(name, step.Finally, opeers); err != nil {
			return err
		}
		if step.Launch != nil {
			if err := recurseSearchPeers(name, step.Launch.Steps, opeers); err != nil {
				return err
			}
		}
		if step.For != nil {
			if err := recurseSearchPeers(name, step.For.Steps, opeers); err != nil {
				return err
			}
		}
		if err := recurseSearchPeers(name, step.Steps, opeers); err != nil {
			return err
		}
		// TODO: fix step use search
		if step.Use != nil {
			if err := recurseSearchPeers(name, step.Use.Steps, opeers); err != nil {
				return err
			}
		}

		if step.Run != nil {
			for _, nic := range step.Run.CustomNetworkInterfaces {
				if nic.VNet != nil && nic.VNet.Name == name {
					ip, _, err := net.ParseCIDR(nic.VNet.Address)
					if err != nil {
						return err
					}
					opeers[ip.String()] = true
				}
			}
		}
	}
	return nil
}

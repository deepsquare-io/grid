// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package renderer

import (
	"bytes"
	"errors"
	"fmt"
	"net"

	_ "embed"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
)

//go:embed renderer_vnet.sh.tpl
var vnetTpl string

var (
	ErrMissingNetwork = errors.New("missing virtual network")
)

func RenderVNet(vnet *model.VNet, interfaceName string, job *model.Job) (string, error) {
	if err := validate.I.Struct(vnet); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(vnetTpl)
	if err != nil {
		return "", err
	}

	// Search network
	vnetIndex := -1
	for i, vn := range job.VirtualNetworks {
		if vn.Name == vnet.Name {
			vnetIndex = i
		}
	}
	if vnetIndex == -1 {
		return "", fmt.Errorf("%w: %s", ErrMissingNetwork, vnet.Name)
	}

	ip, net, err := net.ParseCIDR(vnet.Address)
	if err != nil {
		return "", err
	}

	peers, err := searchPeers(vnet.Name, job.Steps)
	if err != nil {
		return "", err
	}
	peerIndex := -1
	for i, peer := range peers {
		if peer == ip.String() {
			peerIndex = i
		}
	}
	if peerIndex == -1 {
		panic("peer not found, vnet should not exist")
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		VNet          *model.VNet
		InterfaceName string
		VNetIndex     int
		PeerIndex     int
		AllowedIPs    string
	}{
		VNet:          vnet,
		InterfaceName: interfaceName,
		VNetIndex:     vnetIndex,
		PeerIndex:     peerIndex,
		AllowedIPs:    net.String(),
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

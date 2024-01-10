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

package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/renderer"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanWireguardPeer = model.WireguardPeer{
	PublicKey:           "pub",
	AllowedIPs:          []string{"0.0.0.0/0", "172.10.0.0/32"},
	PreSharedKey:        utils.Ptr("sha"),
	Endpoint:            utils.Ptr("10.0.0.0:30"),
	PersistentKeepalive: utils.Ptr(20),
}

var cleanWireguard = model.Wireguard{
	Address:    []string{"10.0.0.1/32"},
	PrivateKey: "abc",
	Peers: []*model.WireguardPeer{
		&cleanWireguardPeer,
	},
}

func TestRenderWireguard(t *testing.T) {
	tests := []struct {
		input struct {
			model.Wireguard
			InterfaceName string
		}
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: struct {
				model.Wireguard
				InterfaceName string
			}{
				Wireguard:     cleanWireguard,
				InterfaceName: "wg0",
			},
			expected: `/usr/bin/cat << 'EOFwireguard' > "$(pwd)/wg0.conf"
[Interface]
Address = 10.0.0.1/32
PrivateKey = abc
[Peer]
PublicKey = pub
AllowedIPs = 0.0.0.0/0,172.10.0.0/32
Endpoint = 10.0.0.0:30
PresharedKey = sha
PersistentKeepalive = 20
EOFwireguard
/usr/bin/chmod 600 "$(pwd)/wg0.conf"
/usr/bin/wg-quick up "$(pwd)/wg0.conf"
`,
			title: "Positive test with wireguard tunnel",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderWireguard(&tt.input.Wireguard, tt.input.InterfaceName)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
				require.NoError(t, renderer.Shellcheck(actual))
			}
		})
	}
}

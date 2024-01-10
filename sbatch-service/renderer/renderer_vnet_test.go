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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanVNet = model.VNet{
	Name:    "test",
	Address: "10.0.0.2/24",
}

func TestRenderVNet(t *testing.T) {
	tests := []struct {
		input struct {
			model.VNet
			InterfaceName string
			model.Job
		}
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: struct {
				model.VNet
				InterfaceName string
				model.Job
			}{
				VNet:          cleanVNet,
				InterfaceName: "vnet0",
				Job: model.Job{
					VirtualNetworks: []*model.VirtualNetwork{
						&cleanVirtualNetwork,
					},
					Steps: []*model.Step{
						cleanStepWithRun(
							&model.StepRun{
								CustomNetworkInterfaces: []*model.NetworkInterface{
									{
										VNet: &cleanVNet,
									},
								},
							},
						),
					},
				},
			},
			expected: `/usr/bin/mkdir -p "$(pwd)/peer0"
/usr/bin/cat << EOFwireguard > "$(pwd)/peer0/vnet0.conf"
[Interface]
Address = 10.0.0.2/24
PrivateKey = $(cat "$STORAGE_PATH/vnet0_peer0_pk")

[Peer]
PublicKey = $(cat "$STORAGE_PATH/vnet0_pub")
AllowedIPs = 10.0.0.0/24
Endpoint = $(cat "$STORAGE_PATH/vnet0_endpoint"):$(cat "$STORAGE_PATH/vnet0_port")
PersistentKeepalive = 20
EOFwireguard
/usr/bin/chmod 600 "$(pwd)/peer0/vnet0.conf"
/usr/bin/wg-quick up "$(pwd)/peer0/vnet0.conf"
`,
			title: "Positive test",
		},
		{
			input: struct {
				model.VNet
				InterfaceName string
				model.Job
			}{
				VNet:          cleanVNet,
				InterfaceName: "vnet0",
				Job:           model.Job{},
			},
			isError: true,
			title:   "Negative test: Missing network",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderVNet(
				&tt.input.VNet,
				tt.input.InterfaceName,
				&tt.input.Job,
			)

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

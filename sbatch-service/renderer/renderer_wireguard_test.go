package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
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
				InterfaceName: "net0",
			},
			expected: `/usr/bin/cat << 'EOFwireguard' > "$(pwd)/net0.conf"
[Interface]
Address = 10.0.0.1/32
PrivateKey = abc
MTU = 1420
[Peer]
PublicKey = pub
AllowedIPs = 0.0.0.0/0,172.10.0.0/32
Endpoint = 10.0.0.0:30
PresharedKey = sha
PersistentKeepalive = 20
EOFwireguard
/usr/bin/chmod 600 "$(pwd)/net0.conf"
wg-quick up "$(pwd)/net0.conf"
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

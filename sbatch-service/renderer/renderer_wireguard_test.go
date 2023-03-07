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
			expected: `/usr/sbin/ip link add dev net0 type wireguard
/usr/sbin/ip -4 address add 10.0.0.1/32 dev net0
/usr/bin/echo 'abc' > "$(pwd)/wg-privatekey"
/usr/bin/chmod 600 "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 private-key "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 peer 'pub' allowed-ips 0.0.0.0/0,172.10.0.0/32
/usr/bin/wg set net0 peer 'pub' endpoint 10.0.0.0:30
/usr/bin/echo 'sha' > "$(pwd)/wg-preshared-0"
/usr/bin/chmod 600 "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer 'pub' preshared-key "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer 'pub' persistent-keepalive 20
/usr/sbin/ip link set mtu 1420 up dev net0
/usr/sbin/ip -4 route add 0.0.0.0/0 dev net0
/usr/sbin/ip -4 route add 172.10.0.0/32 dev net0
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

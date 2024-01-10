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

var cleanVirtualNetwork = model.VirtualNetwork{
	Name:           "test",
	GatewayAddress: "10.0.0.1/24",
}

func TestRenderVirtualNetwork(t *testing.T) {
	tests := []struct {
		input struct {
			model.VirtualNetwork
			index int
			steps []*model.Step
		}
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: struct {
				model.VirtualNetwork
				index int
				steps []*model.Step
			}{
				index: 0,
				steps: []*model.Step{
					cleanStepWithRun(
						&model.StepRun{
							CustomNetworkInterfaces: []*model.NetworkInterface{
								{
									VNet: &cleanVNet,
								},
								{
									VNet: func() *model.VNet {
										vn := cleanVNet
										vn.Address = "10.0.0.3/24"
										return &vn
									}(),
								},
							},
						},
					),
				},
				VirtualNetwork: cleanVirtualNetwork,
			},
			expected: `nsenter_flags() {
  pid="$1"
  flags="--target=${pid}"
  userns="$(readlink "/proc/${pid}/ns/user")"
  mntns="$(readlink "/proc/${pid}/ns/mnt")"
  netns="$(readlink "/proc/${pid}/ns/net")"

  self_userns="$(readlink /proc/self/ns/user)"
  self_mntns="$(readlink /proc/self/ns/mnt)"
  self_netns="$(readlink /proc/self/ns/net)"

  if [ "${userns}" != "${self_userns}" ]; then
    flags="$flags --preserve-credentials -U"
  fi
  if [ "${mntns}" != "${self_mntns}" ]; then
    flags="$flags -m"
  fi
  if [ "${netns}" != "${self_netns}" ]; then
    flags="$flags -n"
  fi
  /usr/bin/echo "${flags}"
}

wait_for_network_namespace() {
  # Wait that the namespace is ready.
  COUNTER=0
  while [ $COUNTER -lt 40 ]; do
    flags=$(nsenter_flags "$1")
    if /usr/bin/echo "$flags" | grep -qvw -- -n; then
      flags="$flags -n"
    fi
    # shellcheck disable=SC2086
    if /usr/bin/nsenter ${flags} true >/dev/null 2>&1; then
      return 0
    else
      /usr/bin/sleep 0.5
    fi
    COUNTER=$(( COUNTER+1 ))
  done
  exit 1
}

# Find random UDP port
/usr/bin/python -c 'import socket; s=socket.socket(socket.AF_INET, socket.SOCK_DGRAM); s.bind(("", 0)); print(s.getsockname()[1]); s.close()' > "$STORAGE_PATH/vnet0_port"

pk="$(/usr/bin/wg genkey)"
/usr/bin/echo "$pk" > "$STORAGE_PATH/vnet0_pk"
/usr/bin/echo "$(/usr/bin/echo "$pk" | /usr/bin/wg pubkey)" > "$STORAGE_PATH/vnet0_pub"
/usr/bin/echo "$(ip route get 1 | awk 'NR==1 {print $7}')" > "$STORAGE_PATH/vnet0_endpoint"

pk="$(/usr/bin/wg genkey)"
/usr/bin/echo "$pk" > "$STORAGE_PATH/vnet0_peer0_pk"
/usr/bin/echo "$(/usr/bin/echo "$pk" | /usr/bin/wg pubkey)" > "$STORAGE_PATH/vnet0_peer0_pub"

pk="$(/usr/bin/wg genkey)"
/usr/bin/echo "$pk" > "$STORAGE_PATH/vnet0_peer1_pk"
/usr/bin/echo "$(/usr/bin/echo "$pk" | /usr/bin/wg pubkey)" > "$STORAGE_PATH/vnet0_peer1_pub"

/usr/bin/cat << EOF > "$STORAGE_PATH/vnet0.conf"
[Interface]
Address = 10.0.0.1/24
ListenPort = $(/usr/bin/cat "$STORAGE_PATH/vnet0_port")
PrivateKey = $(/usr/bin/cat "$STORAGE_PATH/vnet0_pk")

[Peer]
PublicKey = $(/usr/bin/cat "$STORAGE_PATH/vnet0_peer0_pub")
AllowedIPs = 10.0.0.2

[Peer]
PublicKey = $(/usr/bin/cat "$STORAGE_PATH/vnet0_peer1_pub")
AllowedIPs = 10.0.0.3

EOF

/usr/bin/unshare --map-root-user --net --mount /bin/sh -c '
set -e
umask 0077

nsenter_flags() {
  pid=$1
  flags="--target=${pid}"
  userns="$(readlink "/proc/${pid}/ns/user")"
  mntns="$(readlink "/proc/${pid}/ns/mnt")"
  netns="$(readlink "/proc/${pid}/ns/net")"

  self_userns="$(readlink /proc/self/ns/user)"
  self_mntns="$(readlink /proc/self/ns/mnt)"
  self_netns="$(readlink /proc/self/ns/net)"

  if [ "${userns}" != "${self_userns}" ]; then
    flags="$flags --preserve-credentials -U"
  fi
  if [ "${mntns}" != "${self_mntns}" ]; then
    flags="$flags -m"
  fi
  if [ "${netns}" != "${self_netns}" ]; then
    flags="$flags -n"
  fi
  /usr/bin/echo "${flags}"
}

wait_for_network_device() {
  # Wait that the device appears.
  COUNTER=0
  while [ $COUNTER -lt 40 ]; do
    if /usr/bin/nsenter $(nsenter_flags "$1") ip addr show "$2"; then
      return 0
    else
      /usr/bin/sleep 0.5
    fi
    COUNTER=$(( COUNTER+1 ))
  done
  exit 1
}

wait_for_network_device $$ net0

/usr/bin/wg-quick up "$STORAGE_PATH/vnet0.conf"

/usr/bin/sleep infinity

' &
child=$!

wait_for_network_namespace $child

/usr/bin/pasta --config-net \
  --tcp-ports none \
  --udp-ports "$(cat "$STORAGE_PATH/vnet0_port")" \
  --netns "/proc/$child/ns/net" \
  --userns "/proc/$child/ns/user" \
  --ns-ifname net0 \
  -f &
pasta_pid=$!
cleanup() {
  kill -9 $child $pasta_pid || true
}
trap cleanup EXIT INT TERM
`,
			title: "Positive test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderVirtualNetwork(
				&tt.input.VirtualNetwork,
				tt.input.index,
				tt.input.steps,
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

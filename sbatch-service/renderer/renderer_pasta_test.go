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

func TestRenderPastaNS(t *testing.T) {
	tests := []struct {
		input struct {
			NICs    []*model.NetworkInterface
			DNS     []string
			Command string
			Shell   *string
			Job     *model.Job
		}
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: struct {
				NICs    []*model.NetworkInterface
				DNS     []string
				Command string
				Shell   *string
				Job     *model.Job
			}{
				NICs: []*model.NetworkInterface{
					&cleanWireguardNIC,
				},
				DNS:     []string{"1.1.1.1"},
				Command: "echo 'hello world'",
			},
			expected: `set -e

nsenter_flags() {
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
  echo "${flags}"
}

wait_for_network_namespace() {
  # Wait that the namespace is ready.
  COUNTER=0
  while [ $COUNTER -lt 40 ]; do
    flags=$(nsenter_flags "$1")
    if echo "$flags" | grep -qvw -- -n; then
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

# shellcheck disable=SC2016,SC1078,SC1079
/usr/bin/unshare --map-current-user --net --mount /bin/sh -c '
set -e

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


/usr/bin/cat << '"'"'EOFwireguard'"'"' > "$(pwd)/wg0.conf"
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

/usr/bin/echo "nameserver 1.1.1.1" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf

''echo '"'"'hello world'"'"'' &
child=$!

wait_for_network_namespace $child

/usr/bin/pasta --config-net \
  --tcp-ports none \
  --udp-ports none \
  --netns "/proc/$child/ns/net" \
  --userns "/proc/$child/ns/user" \
  --ns-ifname net0 \
  -f &
pasta_pid=$!

cleanup() {
  kill -9 $child $pasta_pid || true
}
trap cleanup EXIT INT TERM

wait $child
`,
			title: "Positive test with run",
		},
		{
			input: struct {
				NICs    []*model.NetworkInterface
				DNS     []string
				Command string
				Shell   *string
				Job     *model.Job
			}{
				NICs: []*model.NetworkInterface{
					&cleanBoreNIC,
				},
				DNS:     []string{"1.1.1.1"},
				Command: "echo 'hello world'",
			},
			expected: `set -e

nsenter_flags() {
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
  echo "${flags}"
}

wait_for_network_namespace() {
  # Wait that the namespace is ready.
  COUNTER=0
  while [ $COUNTER -lt 40 ]; do
    flags=$(nsenter_flags "$1")
    if echo "$flags" | grep -qvw -- -n; then
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

# shellcheck disable=SC2016,SC1078,SC1079
/usr/bin/unshare --map-current-user --net --mount /bin/sh -c '
set -e

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


/usr/bin/dpsproxy --to.addr address.com:11 --local.addr localhost:22 --secret 'test' -r &
/usr/bin/echo "nameserver 1.1.1.1" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf

''echo '"'"'hello world'"'"'' &
child=$!

wait_for_network_namespace $child

/usr/bin/pasta --config-net \
  --tcp-ports none \
  --udp-ports none \
  --netns "/proc/$child/ns/net" \
  --userns "/proc/$child/ns/user" \
  --ns-ifname net0 \
  -f &
pasta_pid=$!

cleanup() {
  kill -9 $child $pasta_pid || true
}
trap cleanup EXIT INT TERM

wait $child
`,
			title: "Positive test with run bore",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderPastaNS(
				&model.StepRun{
					CustomNetworkInterfaces: tt.input.NICs,
					DNS:                     tt.input.DNS,
					Command:                 tt.input.Command,
					Shell:                   tt.input.Shell,
				},
				tt.input.Job,
				tt.input.Command,
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

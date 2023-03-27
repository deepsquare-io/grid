package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanWireguardNIC = model.NetworkInterface{
	Wireguard: &cleanWireguard,
}

var cleanBoreNIC = model.NetworkInterface{
	Bore: &model.Bore{
		Address:    "address.com",
		Port:       11,
		TargetPort: 22,
	},
}

func TestRenderSlirp4NetNS(t *testing.T) {
	tests := []struct {
		input struct {
			NICs    []*model.NetworkInterface
			DNS     []string
			Command string
			Shell   *string
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
    if nsenter ${flags} true >/dev/null 2>&1; then
      return 0
    else
      /usr/bin/sleep 0.5
    fi
    COUNTER=$(( COUNTER+1 ))
  done
  exit 1
}

# shellcheck disable=SC2016,SC1078,SC1079
/usr/bin/unshare --user --net --mount --map-root-user /bin/sh -c '
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
    if nsenter $(nsenter_flags "$1") ip addr show "$2"; then
      return 0
    else
      /usr/bin/sleep 0.5
    fi
    COUNTER=$(( COUNTER+1 ))
  done
  exit 1
}

wait_for_network_device $$ tap0

/usr/sbin/ip link add dev net0 type wireguard
/usr/sbin/ip -4 address add 10.0.0.1/32 dev net0
/usr/bin/echo '"'"'abc'"'"' > "$(pwd)/wg-privatekey"
/usr/bin/chmod 600 "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 private-key "$(pwd)/wg-privatekey"
/usr/bin/wg set net0 peer '"'"'pub'"'"' allowed-ips 0.0.0.0/0,172.10.0.0/32
/usr/bin/wg set net0 peer '"'"'pub'"'"' endpoint 10.0.0.0:30
/usr/bin/echo '"'"'sha'"'"' > "$(pwd)/wg-preshared-0"
/usr/bin/chmod 600 "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer '"'"'pub'"'"' preshared-key "$(pwd)/wg-preshared-0"
/usr/bin/wg set net0 peer '"'"'pub'"'"' persistent-keepalive 20
/usr/sbin/ip link set mtu 1420 up dev net0
/usr/sbin/ip -4 route add 0.0.0.0/0 dev net0
/usr/sbin/ip -4 route add 172.10.0.0/32 dev net0

/usr/bin/echo "nameserver 1.1.1.1" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf

''echo '"'"'hello world'"'"'' &
child=$!

wait_for_network_namespace $child

/usr/bin/slirp4netns --configure --disable-host-loopback --cidr 169.254.254.0/24 $child tap0 &
slirp_pid=$!

cleanup() {
  kill -9 $child $slirp_pid || true
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
    if nsenter ${flags} true >/dev/null 2>&1; then
      return 0
    else
      /usr/bin/sleep 0.5
    fi
    COUNTER=$(( COUNTER+1 ))
  done
  exit 1
}

# shellcheck disable=SC2016,SC1078,SC1079
/usr/bin/unshare --user --net --mount --map-root-user /bin/sh -c '
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
    if nsenter $(nsenter_flags "$1") ip addr show "$2"; then
      return 0
    else
      /usr/bin/sleep 0.5
    fi
    COUNTER=$(( COUNTER+1 ))
  done
  exit 1
}

wait_for_network_device $$ tap0

/usr/bin/bore -s address.com -p 11 -ls localhost -lp 22 -r &
/usr/bin/echo "nameserver 1.1.1.1" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf

''echo '"'"'hello world'"'"'' &
child=$!

wait_for_network_namespace $child

/usr/bin/slirp4netns --configure --disable-host-loopback --cidr 169.254.254.0/24 $child tap0 &
slirp_pid=$!

cleanup() {
  kill -9 $child $slirp_pid || true
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
			actual, err := renderer.RenderSlirp4NetNS(tt.input.NICs, tt.input.DNS, tt.input.Command, tt.input.Shell)

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

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
/usr/bin/python -c 'import socket; s=socket.socket(socket.AF_INET, socket.SOCK_DGRAM); s.bind(("", 0)); print(s.getsockname()[1]); s.close()' > "$STORAGE_PATH/vnet{{ .Index }}_port"

pk="$(/usr/bin/wg genkey)"
/usr/bin/echo "$pk" > "$STORAGE_PATH/vnet{{ $.Index }}_pk"
/usr/bin/echo "$(/usr/bin/echo "$pk" | /usr/bin/wg pubkey)" > "$STORAGE_PATH/vnet{{ .Index }}_pub"
/usr/bin/echo "$(ip route get 1 | awk 'NR==1 {print $7}')" > "$STORAGE_PATH/vnet{{ .Index }}_endpoint"

{{ range $i, $peer := .Peers -}}
pk="$(/usr/bin/wg genkey)"
/usr/bin/echo "$pk" > "$STORAGE_PATH/vnet{{ $.Index }}_peer{{ $i }}_pk"
/usr/bin/echo "$(/usr/bin/echo "$pk" | /usr/bin/wg pubkey)" > "$STORAGE_PATH/vnet{{ $.Index }}_peer{{ $i }}_pub"

{{ end -}}

/usr/bin/cat << EOF > "$STORAGE_PATH/vnet{{ .Index }}.conf"
[Interface]
Address = {{ .VirtualNetwork.GatewayAddress }}
ListenPort = $(/usr/bin/cat "$STORAGE_PATH/vnet{{ .Index }}_port")
PrivateKey = $(/usr/bin/cat "$STORAGE_PATH/vnet{{ .Index }}_pk")

{{ range $i, $peer := .Peers -}}
[Peer]
PublicKey = $(/usr/bin/cat "$STORAGE_PATH/vnet{{ $.Index }}_peer{{ $i }}_pub")
AllowedIPs = {{ $peer }}

{{ end -}}
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

/usr/bin/wg-quick up "$STORAGE_PATH/vnet{{ $.Index }}.conf"

/usr/bin/sleep infinity

' &
child=$!

wait_for_network_namespace $child

/usr/bin/pasta --config-net \
  --tcp-ports none \
  --udp-ports "$(cat "$STORAGE_PATH/vnet{{ .Index }}_port")" \
  --netns "/proc/$child/ns/net" \
  --userns "/proc/$child/ns/user" \
  --ns-ifname net0 \
  -f &
pasta_pid=$!
cleanup() {
  kill -9 $child $pasta_pid || true
}
trap cleanup EXIT INT TERM

set -e

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
/usr/bin/unshare --user --net --mount --map-root-user {{ if .Shell }}{{ derefStr .Shell }}{{ else }}/bin/sh{{ end }} -c '
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

{{ range $i, $nic := .NICs }}
{{- if $nic.Wireguard -}}
{{ renderWireguard $nic.Wireguard (printf "net%d" $i) | escapeSQuote }}
{{- else if $nic.Bore -}}
/usr/bin/bore -s {{ $nic.Bore.Address }} -p {{ $nic.Bore.Port }} -ls localhost -lp {{ $nic.Bore.TargetPort }} -r &
{{- end -}}
{{- end -}}
{{- range $dns := .DNS }}
/usr/bin/echo "nameserver {{ $dns }}" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
{{- end }}
{{- if gt (len .DNS) 0 }}
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf
{{- end }}

'{{ .Command | squote }} &
child=$!

wait_for_network_namespace $child

/usr/bin/slirp4netns --configure --disable-host-loopback --cidr 169.254.254.0/24 $child tap0 &
slirp_pid=$!

cleanup() {
  kill -9 $child $slirp_pid || true
}
trap cleanup EXIT INT TERM

wait $child

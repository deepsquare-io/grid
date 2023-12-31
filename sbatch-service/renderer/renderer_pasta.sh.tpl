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
/usr/bin/unshare --map-current-user --net --mount{{ if .Run.MapUID }} --map-user={{ .Run.MapUID }}{{ end }}{{ if .Run.MapGID }} --map-group={{ .Run.MapGID }}{{ end }} {{ if .Run.Shell }}{{ derefStr .Run.Shell }}{{ else }}/bin/sh{{ end }} -c '
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

{{ range $i, $nic := .Run.CustomNetworkInterfaces }}
{{- if $nic.Wireguard }}
{{ renderWireguard $nic.Wireguard (printf "wg%d" $i) | escapeSQuote }}
{{- else if $nic.VNet }}
{{ renderVNet $nic.VNet (printf "vnet%d" $i) $.Job | escapeSQuote }}
{{- else if $nic.Bore }}
/usr/bin/dpsproxy --to.addr {{ $nic.Bore.BoreAddress | ignoreNil }}{{ $nic.Bore.Address | ignoreNil }}{{ if $nic.Bore.Port }}:{{ $nic.Bore.Port }}{{ end }} --local.addr localhost:{{ $nic.Bore.TargetPort }}{{ if $nic.Bore.Secret }} --secret {{ derefStr $nic.Bore.Secret | squote }}{{ end }} -r &
{{- end -}}
{{- end -}}
{{- range $dns := .Run.DNS }}
/usr/bin/echo "nameserver {{ $dns }}" > "$(pwd)/resolv.$SLURM_JOB_ID.conf"
{{- end }}
{{- if gt (len .Run.DNS) 0 }}
/usr/bin/mount --bind "$(pwd)/resolv.$SLURM_JOB_ID.conf" /etc/resolv.conf
{{- end }}

'{{ .Command | squote }} &
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

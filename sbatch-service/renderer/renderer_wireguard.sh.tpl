/usr/sbin/ip link add dev {{ .InterfaceName }} type wireguard
{{- range $address := .Wireguard.Address }}
{{- if isCIDRv4 $address }}
/usr/sbin/ip -4 address add {{ $address }} dev {{ $.InterfaceName }}
{{- else if isCIDRv6 $address }}
/usr/sbin/ip -6 address add {{ $address }} dev {{ $.InterfaceName }}
{{- end }}
{{- end }}
/usr/bin/echo {{ .Wireguard.PrivateKey | squote }} > "$(pwd)/wg-privatekey"
/usr/bin/chmod 600 "$(pwd)/wg-privatekey"
/usr/bin/wg set {{ .InterfaceName }} private-key "$(pwd)/wg-privatekey"
{{- range $i,$peer := .Wireguard.Peers }}
/usr/bin/wg set {{ $.InterfaceName }} peer {{ $peer.PublicKey | squote }} allowed-ips {{ $peer.AllowedIPs | join "," }}
{{- if $peer.Endpoint }}
/usr/bin/wg set {{ $.InterfaceName }} peer {{ $peer.PublicKey | squote }} endpoint {{ derefStr $peer.Endpoint }}
{{- end }}
{{- if $peer.PreSharedKey }}
/usr/bin/echo {{ derefStr $peer.PreSharedKey | squote }} > "$(pwd)/wg-preshared-{{ $i }}"
/usr/bin/chmod 600 "$(pwd)/wg-preshared-{{ $i }}"
/usr/bin/wg set {{ $.InterfaceName }} peer {{ $peer.PublicKey | squote }} preshared-key "$(pwd)/wg-preshared-{{ $i }}"
{{- end }}
{{- if $peer.PersistentKeepalive }}
/usr/bin/wg set {{ $.InterfaceName }} peer {{ $peer.PublicKey | squote }} persistent-keepalive {{ derefInt $peer.PersistentKeepalive }}
{{- end }}
{{- end }}
/usr/sbin/ip link set mtu 1420 up dev {{ .InterfaceName }}
{{- range $peer := .Wireguard.Peers }}
{{- range $address := $peer.AllowedIPs }}
{{- if isCIDRv4 $address }}
/usr/sbin/ip -4 route add {{ $address }} dev {{ $.InterfaceName }}
{{- else if isCIDRv6 $address }}
/usr/sbin/ip -6 route add {{ $address }} dev {{ $.InterfaceName }}
{{- end }}
{{- end }}
{{- end }}

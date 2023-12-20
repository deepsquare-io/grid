/usr/bin/cat << 'EOFwireguard' > "$(pwd)/{{ .InterfaceName }}.conf"
[Interface]
Address = {{ .Wireguard.Address | join "," }}
PrivateKey = {{ printf "%s" .Wireguard.PrivateKey | replace "'" "" }}

{{- range $i,$peer := .Wireguard.Peers }}
[Peer]
PublicKey = {{ printf "%s" $peer.PublicKey | replace "'" "" }}
{{- if $peer.AllowedIPs }}
AllowedIPs = {{ $peer.AllowedIPs | join "," }}
{{- end }}
{{- if $peer.Endpoint }}
Endpoint = {{ derefStr $peer.Endpoint }}
{{- end }}
{{- if $peer.PreSharedKey }}
PresharedKey = {{ printf "%s" (derefStr $peer.PreSharedKey) | replace "'" "" }}
{{- end }}
{{- if $peer.PersistentKeepalive }}
PersistentKeepalive = {{ derefInt $peer.PersistentKeepalive }}
{{- end }}
{{- end }}
EOFwireguard
/usr/bin/chmod 600 "$(pwd)/{{ .InterfaceName }}.conf"
wg-quick up "$(pwd)/{{ .InterfaceName }}.conf"

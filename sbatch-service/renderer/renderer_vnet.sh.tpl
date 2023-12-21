/usr/bin/mkdir -p "$(pwd)/peer{{ .PeerIndex }}"
/usr/bin/cat << EOFwireguard > "$(pwd)/peer{{ .PeerIndex }}/{{ .InterfaceName }}.conf"
[Interface]
Address = {{ .VNet.Address }}
PrivateKey = $(cat "$STORAGE_PATH/vnet{{ .VNetIndex }}_peer{{ .PeerIndex }}_pk")

[Peer]
PublicKey = $(cat "$STORAGE_PATH/vnet{{ .VNetIndex }}_pub")
AllowedIPs = {{ .AllowedIPs }}
Endpoint = $(cat "$STORAGE_PATH/vnet{{ .VNetIndex }}_endpoint"):$(cat "$STORAGE_PATH/vnet{{ .VNetIndex }}_port")
PersistentKeepalive = 20
EOFwireguard
/usr/bin/chmod 600 "$(pwd)/peer{{ .PeerIndex }}/{{ .InterfaceName }}.conf"
/usr/bin/wg-quick up "$(pwd)/peer{{ .PeerIndex }}/{{ .InterfaceName }}.conf"

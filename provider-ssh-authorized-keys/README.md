# provider-ssh-authorized-keys

## Build

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o provider-ssh-authorized-keys ./
```

## Install

```sh
# As root
mv provider-ssh-authorized-keys /usr/bin/provider-ssh-authorized-keys
chown root:root /usr/bin/provider-ssh-authorized-keys
```

Wrap the executable with a script:

```sh
cat << EOF > /usr/bin/custom-ssh-authorized-keys
#!/bin/sh
# SSSD
/usr/bin/sss_ssh_authorizedkeys "$1"

# Our authorized keys
/usr/bin/provider-ssh-authorized-keys --supervisor.endpoint localhost:3000
EOF
chown root:root /usr/bin/custom-ssh-authorized-keys
```

SSH Config:

```sh
AuthorizedKeysCommand /usr/bin/custom-ssh-authorized-keys
AuthorizedKeysCommandUser root
```

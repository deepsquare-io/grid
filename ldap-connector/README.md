# LDAP connector

Creates user based on Blockchain events.

## Usage

```
NAME:
   ldap-connector - Create user on job submit.

USAGE:
   ldap-connector [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --avax.endpoint.ws value              Avalanche C-Chain WS endpoint. (default: "wss://testnet.deepsquare.run/ws") [$AVAX_ENDPOINT_WS]
   --jobmanager.smart-contract value     JobManager smart-contract address. (deprecated, if specified, will take over METASCHEDULER_SMART_CONTRACT) [$JOBMANAGER_SMART_CONTRACT]
   --metascheduler.smart-contract value  Metascheduler smart-contract address. (default: "0x3707aB457CF457275b7ec32e203c54df80C299d5") [$METASCHEDULER_SMART_CONTRACT]
   --ldap.url value                      LDAP URL (default: "ldap://example.com:389") [$LDAP_URL]
   --ldap.insecure                       Ignore TLS check. (default: false) [$LDAP_INSECURE]
   --ldap.ca.path value                  LDAP CA path [$LDAP_CA_PATH]
   --ldap.bind.dn value                  LDAP Bind DN (default: "cn=Directory Manager") [$LDAP_BIND_DN]
   --ldap.bind.password value            LDAP Bind password [$LDAP_BIND_PASSWORD]
   --config.path value                   Configuration file path. (default: "config.yaml") [$CONFIG_PATH]
   --help, -h                            show help
```

## Build

```sh
go build -o app ./cmd
```

Or, build and install directly with:

```sh
go install github.com/deepsquare-io/grid/ldap-connector@latest
```

## License

The ldap-connector is [licensed under GPL3](./LICENSE).

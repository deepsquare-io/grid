package main

import (
	"os"

	"github.com/deepsquare-io/the-grid/ldap-connector/config"
	"github.com/deepsquare-io/the-grid/ldap-connector/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/ldap-connector/ldap"
	"github.com/deepsquare-io/the-grid/ldap-connector/logger"
	"github.com/deepsquare-io/the-grid/ldap-connector/watcher"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	ldapURL          string
	ldapInsecure     bool
	ldapCAFile       string
	ldapBindDN       string
	ldapBindPassword string

	ethEndpointWS              string
	metaschedulerSmartContract string

	configPath string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "metascheduler.endpoint.ws",
		Value:       "wss://testnet.deepsquare.run/ws",
		Usage:       "Metascheduler Avalanche C-Chain WS endpoint.",
		Destination: &ethEndpointWS,
		EnvVars:     []string{"METASCHEDULER_ENDPOINT_WS"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       "0x",
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
	&cli.StringFlag{
		Name:        "ldap.url",
		Value:       "ldap://example.com:389",
		Usage:       "LDAP URL",
		Destination: &ldapURL,
		EnvVars:     []string{"LDAP_URL"},
	},
	&cli.BoolFlag{
		Name:        "ldap.insecure",
		Value:       false,
		Usage:       "Ignore TLS check.",
		Destination: &ldapInsecure,
		EnvVars:     []string{"LDAP_INSECURE"},
	},
	&cli.StringFlag{
		Name:        "ldap.ca.path",
		Value:       "",
		Usage:       "LDAP CA path",
		Destination: &ldapCAFile,
		EnvVars:     []string{"LDAP_CA_PATH"},
	},
	&cli.StringFlag{
		Name:        "ldap.bind.dn",
		Value:       "cn=Directory Manager",
		Usage:       "LDAP Bind DN",
		Destination: &ldapBindDN,
		EnvVars:     []string{"LDAP_BIND_DN"},
	},
	&cli.StringFlag{
		Name:        "ldap.bind.password",
		Usage:       "LDAP Bind password",
		Destination: &ldapBindPassword,
		EnvVars:     []string{"LDAP_BIND_PASSWORD"},
	},
	&cli.StringFlag{
		Name:        "config.path",
		Value:       "config.yaml",
		Usage:       "Configuration file path.",
		Destination: &configPath,
		EnvVars:     []string{"CONFIG_PATH"},
	},
}

var app = &cli.App{
	Name:  "ldap-connector",
	Usage: "Create user on job submit.",
	Flags: flags,
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context
		conf, err := config.ParseFile(ldapCAFile)
		if err != nil {
			logger.I.Fatal("config parse failed", zap.Error(err))
		}
		if err := conf.Validate(); err != nil {
			logger.I.Fatal("config validation failed", zap.Error(err))
		}
		ethClientWS, err := ethclient.Dial(ethEndpointWS)
		if err != nil {
			logger.I.Fatal("ethClientWS dial failed", zap.Error(err))
		}
		msWS, err := metascheduler.NewMetaScheduler(common.HexToAddress(metaschedulerSmartContract), ethClientWS)
		if err != nil {
			logger.I.Fatal("metaschedulerWS dial failed", zap.Error(err))
		}
		ldap := ldap.New(
			ldapURL,
			ldapBindDN,
			ldapBindPassword,
			conf,
			ldapInsecure,
			ldapCAFile,
		)

		watcher := watcher.New(
			msWS,
			ldap,
		)

		return watcher.Watch(ctx)
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}

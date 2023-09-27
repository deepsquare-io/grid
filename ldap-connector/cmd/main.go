// Copyright (C) 2023 DeepSquare Asociation
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"os"

	"github.com/deepsquare-io/grid/ldap-connector/config"
	"github.com/deepsquare-io/grid/ldap-connector/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/grid/ldap-connector/ldap"
	"github.com/deepsquare-io/grid/ldap-connector/logger"
	"github.com/deepsquare-io/grid/ldap-connector/watcher"
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

	avaxEndpointWS             string
	jobManagerSmartContract    string
	metaSchedulerSmartContract string

	configPath string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "avax.endpoint.ws",
		Value:       "wss://testnet.deepsquare.run/ws",
		Usage:       "Avalanche C-Chain WS endpoint.",
		Destination: &avaxEndpointWS,
		EnvVars:     []string{"AVAX_ENDPOINT_WS"},
	},
	&cli.StringFlag{
		Name:        "jobmanager.smart-contract",
		Required:    false,
		Usage:       "JobManager smart-contract address. (deprecated, if specified, will take over METASCHEDULER_SMART_CONTRACT)",
		Destination: &jobManagerSmartContract,
		EnvVars:     []string{"JOBMANAGER_SMART_CONTRACT"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       "0x3707aB457CF457275b7ec32e203c54df80C299d5",
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaSchedulerSmartContract,
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
	Name:                 "ldap-connector",
	Usage:                "Create user on job submit.",
	Flags:                flags,
	Suggest:              true,
	EnableBashCompletion: true,
	Copyright: `ldap-connector  Copyright (C) 2023 DeepSquare Association
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.
See the GNU General Public License for more details.`,
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context
		conf, err := config.ParseFile(configPath)
		if err != nil {
			logger.I.Fatal("config parse failed", zap.Error(err))
		}
		if err := conf.Validate(); err != nil {
			logger.I.Fatal("config validation failed", zap.Error(err))
		}
		ethClientWS, err := ethclient.Dial(avaxEndpointWS)
		if err != nil {
			logger.I.Fatal("ethClientWS dial failed", zap.Error(err))
		}

		var contractAddress string
		if jobManagerSmartContract != "" {
			contractAddress = jobManagerSmartContract
		} else {
			contractAddress = metaSchedulerSmartContract
		}
		contract, err := metascheduler.NewMetaScheduler(
			common.HexToAddress(contractAddress),
			ethClientWS,
		)
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
		if err := ldap.HealthCheck(ctx); err != nil {
			return err
		}

		watcher := watcher.New(
			contract,
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

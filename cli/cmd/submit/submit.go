package submit

import (
	"math/big"

	"github.com/deepsquare-io/the-grid/cli/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/pkg/eth"
	"github.com/urfave/cli/v2"
)

var (
	ethEndpoint                string
	ethHexPK                   string
	metaschedulerSmartContract string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "metascheduler.endpoint",
		Value:       "https://testnet.deepsquare.run/rpc",
		Usage:       "Metascheduler RPC endpoint.",
		Destination: &ethEndpoint,
		EnvVars:     []string{"METASCHEDULER_ENDPOINT"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       "0x",
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
	&cli.StringFlag{
		Name:        "eth.private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Required:    true,
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
}

// Container stores the instances for dependency injection.
type Container struct {
	eth *eth.DataSource
}

func Init() *Container {
	ethDataSource := eth.New(
		ethEndpoint,
		ethHexPK,
		metaschedulerSmartContract,
	)

	return &Container{
		eth: ethDataSource,
	}
}

var Command = &cli.Command{
	Name:  "submit",
	Usage: "Submit a job to the DeepSquare Grid.",
	Flags: flags,
	Action: func(ctx *cli.Context) error {
		c := ctx.Context
		container := Init()

		// TODO: Parse
		jobDefinition := metascheduler.JobDefinition{}
		amountLocked := big.NewInt(0)

		return container.eth.RequestNewJob(c, jobDefinition, amountLocked)
	},
}

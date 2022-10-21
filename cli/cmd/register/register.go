package register

import (
	"github.com/deepsquare-io/the-grid/cli/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/deepsquare-io/the-grid/cli/pkg/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	ethEndpoint                string
	ethHexPK                   string
	metaschedulerSmartContract string

	providerDefinition = metascheduler.ProviderDefinition{}
)

func init() {
	viper.AutomaticEnv()
	flags := Command.Flags()
	flags.StringVar(
		&ethEndpoint,
		"metascheduler.endpoint",
		"https://testnet.deepsquare.run/rpc",
		"Metascheduler RPC endpoint. (env: METASCHEDULER_ENDPOINT)",
	)
	if err := viper.BindPFlag("METASCHEDULER_ENDPOINT", flags.Lookup("metascheduler.endpoint")); err != nil {
		logger.I.Fatal("couldn't bind pFlag", zap.Error(err))
	}
	flags.StringVar(
		&metaschedulerSmartContract,
		"metascheduler.smart-contract",
		"",
		"Metascheduler smart-contract address. Must have the prefix 0x. (env: METASCHEDULER_SMART_CONTRACT)",
	)
	// TODO: fix required env
	if err := viper.BindPFlag("METASCHEDULER_SMART_CONTRACT", flags.Lookup("metascheduler.smart-contract")); err != nil {
		logger.I.Fatal("couldn't bind pFlag", zap.Error(err))
	}
	flags.StringVar(
		&ethHexPK,
		"eth.private-key",
		"",
		"An hexadecimal private key for ethereum transactions. (env: ETH_PRIVATE_KEY)",
	)
	if err := viper.BindPFlag("ETH_PRIVATE_KEY", flags.Lookup("eth.private-key")); err != nil {
		logger.I.Fatal("couldn't bind pFlag", zap.Error(err))
	}
	flags.Uint64Var(
		&providerDefinition.CpuPricePerMin,
		"provider.cpu-price-per-min",
		1,
		"CPU Price per minute. 1 cpu minute = 1e6 cred.",
	)
	flags.Uint64Var(
		&providerDefinition.Cpus,
		"provider.cpus",
		1,
		"Allocated CPUs.",
	)
	flags.Uint64Var(
		&providerDefinition.GpuPricePerMin,
		"provider.gpu-price-per-min",
		1,
		"GPU Price per minute. 1 gpu minute = 1e6 cred.",
	)
	flags.Uint64Var(
		&providerDefinition.Gpus,
		"provider.gpus",
		0,
		"Allocated GPUs.",
	)
	flags.Uint64Var(
		&providerDefinition.MemPricePerMin,
		"provider.mem-price-per-min",
		1,
		"Memory Price per minute. 1 Mem minute = 1e6 cred.",
	)
	flags.Uint64Var(
		&providerDefinition.Mem,
		"provider.mem",
		1,
		"Allocated Memory.",
	)
	flags.Uint64Var(
		&providerDefinition.Nodes,
		"provider.nodes",
		1,
		"Allocated Nodes.",
	)
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

var Command = &cobra.Command{
	Use:   "register <provider address>",
	Short: "Register a provider to the DeepSquare Grid",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cmd.Context()

		container := Init()

		tx, err := container.eth.Register(c, common.HexToAddress(args[0]), providerDefinition)
		if err != nil {
			return err
		}

		logger.I.Info("registered provider successfully", zap.String("tx", tx.Hash().String()))

		return nil
	},
}

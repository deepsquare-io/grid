package submit

import (
	"math/big"
	"os"

	"github.com/deepsquare-io/the-grid/cli/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/deepsquare-io/the-grid/cli/pkg/eth"
	"github.com/deepsquare-io/the-grid/cli/pkg/uploader"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	ethEndpoint                string
	ethHexPK                   string
	metaschedulerSmartContract string

	jobDefinition   = metascheduler.JobDefinition{}
	amountLockedStr string
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
		&jobDefinition.GpuPerNode,
		"gpus",
		0,
		"Allocated GPUs per node.",
	)
	flags.Uint64Var(
		&jobDefinition.MemPerNode,
		"mem",
		0,
		"Allocated memory per node (MB).",
	)
	if err := Command.MarkFlagRequired("mem"); err != nil {
		logger.I.Fatal("couldn't mark flag required", zap.Error(err))
	}
	flags.Uint64VarP(
		&jobDefinition.CpuPerTask,
		"cpus-per-task",
		"c",
		0,
		"Allocated CPUs per task.",
	)
	if err := Command.MarkFlagRequired("cpus-per-task"); err != nil {
		logger.I.Fatal("couldn't mark flag required", zap.Error(err))
	}
	flags.Uint64VarP(
		&jobDefinition.Nodes,
		"nodes",
		"N",
		1,
		"Allocated nodes.",
	)
	flags.Uint64VarP(
		&jobDefinition.Ntasks,
		"tasks",
		"n",
		1,
		"Run the same script in parallel if tasks > 1.",
	)
	flags.StringVarP(
		&amountLockedStr,
		"credits",
		"t",
		"",
		"Amount of credits locked for the job, which is equivalent to the time limit.",
	)
	if err := Command.MarkFlagRequired("credits"); err != nil {
		logger.I.Fatal("couldn't mark flag required", zap.Error(err))
	}
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
	Use:   "submit <path to script>",
	Short: "Submit a job to the DeepSquare Grid.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cmd.Context()

		container := Init()

		amountLocked, ok := big.NewInt(0).SetString(amountLockedStr, 10)
		if !ok {
			logger.I.Fatal("couldn't parse value of credits in base 10")
		}

		file, err := os.Open(args[0])
		if err != nil {
			return err
		}
		logger.I.Debug("file handle created", zap.Any("file", file))
		defer func() {
			if err := file.Close(); err != nil {
				logger.I.Fatal("couldn't open file", zap.String("file", args[0]), zap.Error(err))
			}
		}()

		// TODO: implements file hash
		urlResult, err := uploader.UploadFile(file, "job.sh")
		if err != nil {
			logger.I.Fatal("couldn't upload file", zap.String("file", args[0]), zap.Error(err))
		}
		logger.I.Info("hash succeeded", zap.String("hash", urlResult))

		jobDefinition.BatchLocationHash = urlResult

		tx, err := container.eth.RequestNewJob(c, jobDefinition, amountLocked)
		if err != nil {
			return err
		}

		logger.I.Info("requested a new job successfully", zap.String("tx", tx.Hash().String()))

		return nil
	},
}

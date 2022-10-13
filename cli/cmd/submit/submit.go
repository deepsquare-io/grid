package submit

import (
	"bytes"
	"io"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/deepsquare-io/the-grid/cli/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/deepsquare-io/the-grid/cli/pkg/eth"
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
	viper.BindPFlag("METASCHEDULER_ENDPOINT", flags.Lookup("metascheduler.endpoint"))
	flags.StringVar(
		&metaschedulerSmartContract,
		"metascheduler.smart-contract",
		"",
		"Metascheduler smart-contract address. Must have the prefix 0x. (env: METASCHEDULER_SMART_CONTRACT)",
	)
	Command.MarkFlagRequired("metascheduler.smart-contract")
	viper.BindPFlag("METASCHEDULER_SMART_CONTRACT", flags.Lookup("metascheduler.smart-contract"))
	flags.StringVar(
		&ethHexPK,
		"eth.private-key",
		"",
		"An hexadecimal private key for ethereum transactions. (env: ETH_PRIVATE_KEY)",
	)
	Command.MarkFlagRequired("eth.private-key")
	viper.BindPFlag("ETH_PRIVATE_KEY", flags.Lookup("eth.private-key"))
	flags.Uint64Var(
		&jobDefinition.GpuPerNode,
		"res.gpus-per-node",
		0,
		"Allocated GPUs per node.",
	)
	flags.Uint64Var(
		&jobDefinition.MemPerNode,
		"res.mem-per-node",
		0,
		"Allocated memory per node (MB).",
	)
	Command.MarkFlagRequired("res.mem-per-node")
	flags.Uint64Var(
		&jobDefinition.CpuPerTask,
		"res.cpus-per-task",
		0,
		"Allocated CPUs per task.",
	)
	Command.MarkFlagRequired("res.cpus-per-task")
	flags.Uint64Var(
		&jobDefinition.Nodes,
		"res.nodes",
		1,
		"Allocated nodes.",
	)
	flags.Uint64Var(
		&jobDefinition.Ntasks,
		"res.tasks",
		1,
		"Run the same script in parallel if tasks > 1.",
	)
	flags.StringVar(
		&amountLockedStr,
		"res.credit-locked",
		"",
		"Amount of credits locked for the job, which is equivalent to the time limit.",
	)
	Command.MarkFlagRequired("res.credit-locked")
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
	Use:   "submit [path to script]",
	Short: "Submit a job to the DeepSquare Grid.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cmd.Context()

		container := Init()

		amountLocked, ok := big.NewInt(0).SetString(amountLockedStr, 10)
		if !ok {
			logger.I.Fatal("couldn't parse value of res.credit-locked in base 10")
		}

		file, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer func() {
			if err := file.Close(); err != nil {
				logger.I.Fatal("couldn't open file", zap.String("file", args[0]), zap.Error(err))
			}
		}()

		// TODO: implements file hash
		urlResult, err := uploadFile(file, "job.sh")
		if err != nil {
			logger.I.Fatal("couldn't upload file", zap.String("file", args[0]), zap.Error(err))
		}

		jobDefinition.BatchLocationHash = urlResult

		return container.eth.RequestNewJob(c, jobDefinition, amountLocked)
	},
}

func uploadFile(file *os.File, fileName string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("asset", fileName)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	writer.Close()

	res, err := http.Post("https://transfer.sh/", writer.FormDataContentType(), body)
	if err != nil {
		return "", err
	}

	urlResult, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(urlResult), nil
}

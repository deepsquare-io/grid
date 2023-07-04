package submit

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var (
	sbatchEndpoint             string
	ethEndpointRPC             string
	ethHexPK                   string
	metaschedulerSmartContract string

	credits *big.Int
	jobName string
	uses    cli.StringSlice
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "metascheduler.rpc",
		Value:       deepsquare.DefaultRPCEndpoint,
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &ethEndpointRPC,
		EnvVars:     []string{"METASCHEDULER_RPC"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       deepsquare.DefaultMetaSchedulerAddress.Hex(),
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
	&cli.StringFlag{
		Name:        "sbatch.endpoint",
		Value:       deepsquare.DefaultSBatchEndpoint,
		Usage:       "SBatch Service GraphQL endpoint.",
		Destination: &sbatchEndpoint,
		EnvVars:     []string{"SBATCH_ENDPOINT"},
	},
	&cli.StringFlag{
		Name:        "private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Required:    true,
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
	&cli.StringFlag{
		Name:        "job-name",
		Usage:       "The job name.",
		Required:    true,
		Destination: &jobName,
	},
	&cli.StringSliceFlag{
		Name:        "uses",
		Usage:       "Uses flag. Used to filter the clusters. Format: `key=value`",
		Required:    true,
		Destination: &uses,
	},
	&cli.StringFlag{
		Name:  "credits-wei",
		Usage: "Allocated a number of credits. Unit is wei. Is a big int.",
		Action: func(ctx *cli.Context, s string) error {
			c, ok := new(big.Int).SetString(s, 10)
			if !ok {
				return errors.New("failed to parse credits")
			}
			credits = c
			return nil
		},
	},
	&cli.Float64Flag{
		Name:  "credits",
		Usage: "Allocated a number of credits. Unit is 1e18. Is a float and is not precise.",
		Action: func(ctx *cli.Context, val float64) error {
			c := new(big.Float).SetFloat64(val)

			coin := new(big.Float)
			coin.SetInt(big.NewInt(1000000000000000000))

			c.Mul(c, coin)

			result := new(big.Int)
			c.Int(result)
			credits = result
			return nil
		},
	},
}

var Command = cli.Command{
	Name:      "submit",
	Usage:     "Quickly submit a job.",
	ArgsUsage: "<job.yaml>",
	Flags:     flags,
	Action: func(cCtx *cli.Context) error {
		if cCtx.NArg() != 1 {
			return errors.New("missing arguments")
		}
		if credits == nil {
			return errors.New("missing --credits or --credits-wei parameter")
		}
		jobPath := cCtx.Args().First()
		ctx := cCtx.Context
		pk, err := crypto.HexToECDSA(ethHexPK)
		if err != nil {
			return err
		}
		rpcClient, err := rpc.DialOptions(
			ctx,
			ethEndpointRPC,
			rpc.WithHTTPClient(http.DefaultClient),
		)
		if err != nil {
			return err
		}
		defer rpcClient.Close()
		ethClientRPC := ethclient.NewClient(rpcClient)
		chainID, err := ethClientRPC.ChainID(ctx)
		if err != nil {
			return err
		}
		s := sbatch.NewService(http.DefaultClient, sbatchEndpoint)
		scheduler, err := metascheduler.NewJobScheduler(metascheduler.Backend{
			EthereumBackend:      ethClientRPC,
			MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
			ChainID:              chainID,
			UserPrivateKey:       pk,
		}, s)
		if err != nil {
			return err
		}
		dat, err := os.ReadFile(jobPath)
		if err != nil {
			return err
		}
		var job sbatch.Job
		if err := yaml.Unmarshal(dat, &job); err != nil {
			return err
		}

		var jobNameB [32]byte
		copy(jobNameB[:], jobName)

		// Map slices to label
		usesLabels := make([]metaschedulerabi.Label, 0, len(uses.Value()))
		for _, use := range uses.Value() {
			if key, value, ok := strings.Cut(use, "="); ok {
				usesLabels = append(usesLabels, metaschedulerabi.Label{
					Key:   key,
					Value: value,
				})
			}
		}
		jobID, err := scheduler.SubmitJob(ctx, &job, usesLabels, credits, jobNameB)
		if err != nil {
			return err
		}
		fmt.Printf("job %s submitted", hexutil.Encode(jobID[:]))
		return nil
	},
}

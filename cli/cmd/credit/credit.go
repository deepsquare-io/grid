package credit

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"

	"github.com/erikgeiser/promptkit/confirmation"
)

var (
	ethEndpointRPC             string
	ethHexPK                   string
	metaschedulerSmartContract string

	credits    *big.Float
	creditsWei *big.Int
	force      bool
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
		Name:        "private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Required:    true,
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
	&cli.StringFlag{
		Name:  "credits-wei",
		Usage: "Allocated a number of credits. Unit is wei. Is a big int.",
		Action: func(ctx *cli.Context, s string) error {
			c, ok := new(big.Int).SetString(s, 10)
			if !ok {
				return errors.New("failed to parse credits")
			}
			creditsWei = c
			credits = ether.FromWei(c)
			return nil
		},
	},
	&cli.Float64Flag{
		Name:  "credits",
		Usage: "Allocated a number of credits. Unit is 1e18. Is a float and is not precise.",
		Action: func(ctx *cli.Context, val float64) error {
			c := new(big.Float).SetFloat64(val)

			credits = c
			creditsWei = ether.ToWei(credits)

			return nil
		},
	},
	&cli.BoolFlag{
		Name:        "force",
		Usage:       "Do not prompt",
		Destination: &force,
		Aliases:     []string{"f"},
		EnvVars:     []string{"FORCE"},
	},
}

var Command = cli.Command{
	Name:  "credit",
	Usage: "Manage credits.",
	Subcommands: []*cli.Command{
		{
			Name:      "transfer",
			Usage:     "Transfer credits to an another account",
			ArgsUsage: "<0x recipient address>",
			Flags:     flags,
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() != 1 {
					return errors.New("missing arguments")
				}
				if creditsWei == nil {
					return errors.New("missing --credits or --credits-wei parameter")
				}
				ctx := cCtx.Context
				recipient := common.HexToAddress(cCtx.Args().First())
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
				clientset := metascheduler.NewRPCClientSet(metascheduler.Backend{
					EthereumBackend:      ethClientRPC,
					MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
					ChainID:              chainID,
					UserPrivateKey:       pk,
				})

				if !force {
					msg := fmt.Sprintf(
						"Confirm transfer of %s creds (%s wei) to %s?",
						credits.String(),
						creditsWei.String(),
						recipient.Hex(),
					)
					input := confirmation.New(msg, confirmation.No)
					ok, err := input.RunPrompt()
					if err != nil {
						return err
					}
					if !ok {
						fmt.Println("Cancelled.")
						return nil
					}
				}

				if err := clientset.CreditManager().Transfer(ctx, recipient, creditsWei); err != nil {
					return err
				}

				fmt.Println("Done.")
				return nil
			},
		},
	},
}
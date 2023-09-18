package metascheduler

import (
	"context"
	"fmt"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type RPCClientSet struct {
	*Backend
}

// authOpts generate transact options based on the network.
func (c *RPCClientSet) authOpts(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.PendingNonceAt(ctx, c.from())
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.UserPrivateKey, c.ChainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(0)
	auth.GasPrice = gasPrice
	auth.Context = ctx

	return auth, nil
}

func NewRPCClientSet(b Backend) *RPCClientSet {
	return &RPCClientSet{
		Backend: &b,
	}
}

func (c *RPCClientSet) JobScheduler(
	sbatch sbatch.Service,
) types.JobScheduler {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	return &jobScheduler{
		RPCClientSet:  c,
		MetaScheduler: m,
		Service:       sbatch,
	}
}

func (c *RPCClientSet) JobFetcher() types.JobFetcher {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	jobsAddress, err := m.Jobs(&bind.CallOpts{})
	if err != nil {
		panic(fmt.Errorf("failed to fetch JobRepository contract address: %w", err))
	}
	jobs, err := metaschedulerabi.NewIJobRepository(jobsAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate JobRepository: %w", err))
	}

	return &jobFetcher{
		RPCClientSet:   c,
		IJobRepository: jobs,
	}
}

func (c *RPCClientSet) CreditManager() types.CreditManager {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	creditAddress, err := m.Credit(&bind.CallOpts{})
	if err != nil {
		panic(fmt.Errorf("failed to fetch Credit contract address: %w", err))
	}
	ierc20, err := metaschedulerabi.NewIERC20(creditAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate Credit: %w", err))
	}
	return &creditManager{
		RPCClientSet: c,
		IERC20:       ierc20,
	}
}

func (c *RPCClientSet) AllowanceManager() types.AllowanceManager {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	creditAddress, err := m.Credit(&bind.CallOpts{})
	if err != nil {
		panic(fmt.Errorf("failed to fetch CreditManager contract address: %w", err))
	}
	ierc20, err := metaschedulerabi.NewIERC20(creditAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate CreditManager: %w", err))
	}
	return &allowanceManager{
		RPCClientSet: c,
		IERC20:       ierc20,
	}
}

func (c *RPCClientSet) ProviderManager() types.ProviderManager {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	providerManagerAddress, err := m.ProviderManager(&bind.CallOpts{})
	if err != nil {
		panic(fmt.Errorf("failed to fetch CreditManager contract address: %w", err))
	}
	pm, err := metaschedulerabi.NewIProviderManager(providerManagerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate CreditManager: %w", err))
	}
	return &providerManager{
		RPCClientSet:     c,
		IProviderManager: pm,
	}
}

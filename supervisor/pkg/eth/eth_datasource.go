package eth

import (
	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type ClaimJobResponse struct {
	JobID         string
	User          string
	TimeLimit     uint64
	JobDefinition *metascheduler.JobDefinition
}

type DataSource interface {
	// ClaimJob a job.
	//
	// If the queue is not empty, it will claim the job and send it to the SLURM cluster.
	// Else, it will return an error.
	ClaimJob(
		response chan<- ClaimJobResponse,
		done chan<- error,
	) error
}

type dataSource struct {
	client        *ethclient.Client
	metascheduler *metascheduler.MetaScheduler
}

func New(
	rpcEndpoint string,
	metaschedulerAddress string,
) *dataSource {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		logger.I.Fatal("ethclient dial failed", zap.Error(err))
	}

	metasched, err := metascheduler.NewMetaScheduler(common.HexToAddress(metaschedulerAddress), client)
	if err != nil {
		logger.I.Fatal("metascheduler dial failed", zap.Error(err))
	}

	return &dataSource{
		client:        client,
		metascheduler: metasched,
	}
}

func (s *dataSource) ClaimJob(
	response chan<- ClaimJobResponse,
	done chan<- error,
) {
	// TODO: ClaimNextJob

	// TODO: fetch job body
	response <- ClaimJobResponse{}
	done <- nil
}

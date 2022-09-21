package eth

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Service struct {
	client *ethclient.Client
}

func New(address string) *Service {
	client, err := ethclient.Dial(address)
	if err != nil {
		log.Fatal(err)
	}

	return &Service{client}
}

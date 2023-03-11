package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

const INFURA_URL = "https://mainnet.infura.io/v3/"

type EVMClient struct {
	Client *ethclient.Client
}

// NewEVMClient constructs the EVM client
func NewEVMClient(key string) (*EVMClient, error) {

	c := &EVMClient{}

	client, err := ethclient.Dial(INFURA_URL + key)
	if err != nil {
		return nil, fmt.Errorf("can not connect to node: %s", INFURA_URL)
	}

	c.Client = client

	return c, nil

}

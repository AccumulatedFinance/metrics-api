package evm

import (
	"github.com/AccumulatedFinance/metrics-api/binding"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ERC20 struct {
	Address     string
	Symbol      string
	Decimals    int64
	TotalSupply int64
}

// GetERC20 gets ERC20 Token info
func (e *EVMClient) GetERC20TotalSupply(tokenAddress string) (*ERC20, error) {

	token := &ERC20{}

	address := common.HexToAddress(tokenAddress)

	instance, err := binding.NewWrappedToken(address, e.Client)
	if err != nil {
		return nil, err
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	token.Address = tokenAddress
	token.Symbol = symbol
	token.Decimals = int64(decimals)
	token.TotalSupply = totalSupply.Int64()

	return token, nil

}

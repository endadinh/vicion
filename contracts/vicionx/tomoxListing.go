package vicionx

import (
	"github.com/vicion/vicion/accounts/abi/bind"
	"github.com/vicion/vicion/common"
	"github.com/vicion/vicion/contracts/vicionx/contract"
)

type VICIONXListing struct {
	*contract.VICIONXListingSession
	contractBackend bind.ContractBackend
}

func NewMyVICIONXListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*VICIONXListing, error) {
	smartContract, err := contract.NewVICIONXListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &VICIONXListing{
		&contract.VICIONXListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployVICIONXListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *VICIONXListing, error) {
	contractAddr, _, _, err := contract.DeployVICIONXListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyVICIONXListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}

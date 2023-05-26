// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/vicion/vicion/accounts/abi"
	"github.com/vicion/vicion/accounts/abi/bind"
	"github.com/vicion/vicion/common"
	"github.com/vicion/vicion/core/types"
)

// VICIONXListingABI is the input ABI used to generate the binding from.
const VICIONXListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// VICIONXListingBin is the compiled bytecode used for deploying new contracts.
const VICIONXListingBin = `0x608060405234801561001057600080fd5b506102be806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea0000034146101e257600080fd5b6040516068903480156108fc02916000818181858888f1935050505015801561020f573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a723058206d2dc0ce827743c25efa82f99e7830ade39d28e17f4d651573f89e0460a6626a0029`

// DeployVICIONXListing deploys a new Ethereum contract, binding an instance of VICIONXListing to it.
func DeployVICIONXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VICIONXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(VICIONXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VICIONXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VICIONXListing{VICIONXListingCaller: VICIONXListingCaller{contract: contract}, VICIONXListingTransactor: VICIONXListingTransactor{contract: contract}, VICIONXListingFilterer: VICIONXListingFilterer{contract: contract}}, nil
}

// VICIONXListing is an auto generated Go binding around an Ethereum contract.
type VICIONXListing struct {
	VICIONXListingCaller     // Read-only binding to the contract
	VICIONXListingTransactor // Write-only binding to the contract
	VICIONXListingFilterer   // Log filterer for contract events
}

// VICIONXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VICIONXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VICIONXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VICIONXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VICIONXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VICIONXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VICIONXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VICIONXListingSession struct {
	Contract     *VICIONXListing     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VICIONXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VICIONXListingCallerSession struct {
	Contract *VICIONXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VICIONXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VICIONXListingTransactorSession struct {
	Contract     *VICIONXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VICIONXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VICIONXListingRaw struct {
	Contract *VICIONXListing // Generic contract binding to access the raw methods on
}

// VICIONXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VICIONXListingCallerRaw struct {
	Contract *VICIONXListingCaller // Generic read-only contract binding to access the raw methods on
}

// VICIONXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VICIONXListingTransactorRaw struct {
	Contract *VICIONXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVICIONXListing creates a new instance of VICIONXListing, bound to a specific deployed contract.
func NewVICIONXListing(address common.Address, backend bind.ContractBackend) (*VICIONXListing, error) {
	contract, err := bindVICIONXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VICIONXListing{VICIONXListingCaller: VICIONXListingCaller{contract: contract}, VICIONXListingTransactor: VICIONXListingTransactor{contract: contract}, VICIONXListingFilterer: VICIONXListingFilterer{contract: contract}}, nil
}

// NewVICIONXListingCaller creates a new read-only instance of VICIONXListing, bound to a specific deployed contract.
func NewVICIONXListingCaller(address common.Address, caller bind.ContractCaller) (*VICIONXListingCaller, error) {
	contract, err := bindVICIONXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VICIONXListingCaller{contract: contract}, nil
}

// NewVICIONXListingTransactor creates a new write-only instance of VICIONXListing, bound to a specific deployed contract.
func NewVICIONXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*VICIONXListingTransactor, error) {
	contract, err := bindVICIONXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VICIONXListingTransactor{contract: contract}, nil
}

// NewVICIONXListingFilterer creates a new log filterer instance of VICIONXListing, bound to a specific deployed contract.
func NewVICIONXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*VICIONXListingFilterer, error) {
	contract, err := bindVICIONXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VICIONXListingFilterer{contract: contract}, nil
}

// bindVICIONXListing binds a generic wrapper to an already deployed contract.
func bindVICIONXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VICIONXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VICIONXListing *VICIONXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VICIONXListing.Contract.VICIONXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VICIONXListing *VICIONXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VICIONXListing.Contract.VICIONXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VICIONXListing *VICIONXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VICIONXListing.Contract.VICIONXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VICIONXListing *VICIONXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VICIONXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VICIONXListing *VICIONXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VICIONXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VICIONXListing *VICIONXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VICIONXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_VICIONXListing *VICIONXListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VICIONXListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_VICIONXListing *VICIONXListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _VICIONXListing.Contract.GetTokenStatus(&_VICIONXListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_VICIONXListing *VICIONXListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _VICIONXListing.Contract.GetTokenStatus(&_VICIONXListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_VICIONXListing *VICIONXListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _VICIONXListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_VICIONXListing *VICIONXListingSession) Tokens() ([]common.Address, error) {
	return _VICIONXListing.Contract.Tokens(&_VICIONXListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_VICIONXListing *VICIONXListingCallerSession) Tokens() ([]common.Address, error) {
	return _VICIONXListing.Contract.Tokens(&_VICIONXListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_VICIONXListing *VICIONXListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _VICIONXListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_VICIONXListing *VICIONXListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _VICIONXListing.Contract.Apply(&_VICIONXListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_VICIONXListing *VICIONXListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _VICIONXListing.Contract.Apply(&_VICIONXListing.TransactOpts, token)
}

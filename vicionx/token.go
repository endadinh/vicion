package vicionx

import (
	"math/big"
	"strings"

	"github.com/vicion/vicion/contracts/vicionx/contract"
	"github.com/vicion/vicion/log"

	"github.com/vicion/vicion"
	"github.com/vicion/vicion/accounts/abi"
	"github.com/vicion/vicion/accounts/abi/bind/backends"
	"github.com/vicion/vicion/common"
	"github.com/vicion/vicion/consensus"
	"github.com/vicion/vicion/core/state"
)

// GetTokenAbi return token abi
func GetTokenAbi() (*abi.ABI, error) {
	contractABI, err := abi.JSON(strings.NewReader(contract.TRC21ABI))
	if err != nil {
		return nil, err
	}
	return &contractABI, nil
}

// RunContract run smart contract
func RunContract(chain consensus.ChainContext, statedb *state.StateDB, contractAddr common.Address, abi *abi.ABI, method string, args ...interface{}) (interface{}, error) {
	input, err := abi.Pack(method)
	if err != nil {
		return nil, err
	}
	backend := (*backends.SimulatedBackend)(nil)
	fakeCaller := common.HexToAddress("0x0000000000000000000000000000000000000001")
	msg := vicion.CallMsg{To: &contractAddr, Data: input, From: fakeCaller}
	result, err := backend.CallContractWithState(msg, chain, statedb)
	if err != nil {
		return nil, err
	}
	var unpackResult interface{}
	err = abi.Unpack(&unpackResult, method, result)
	if err != nil {
		return nil, err
	}
	return unpackResult, nil
}

func (vicionx *VicionX) GetTokenDecimal(chain consensus.ChainContext, statedb *state.StateDB, tokenAddr common.Address) (*big.Int, error) {
	if tokenDecimal, ok := vicionx.tokenDecimalCache.Get(tokenAddr); ok {
		return tokenDecimal.(*big.Int), nil
	}
	if tokenAddr.String() == common.TomoNativeAddress {
		vicionx.tokenDecimalCache.Add(tokenAddr, common.BasePrice)
		return common.BasePrice, nil
	}
	var decimals uint8
	defer func() {
		log.Debug("GetTokenDecimal from ", "relayerSMC", common.RelayerRegistrationSMC, "tokenAddr", tokenAddr.Hex(), "decimals", decimals)
	}()
	contractABI, err := GetTokenAbi()
	if err != nil {
		return nil, err
	}
	stateCopy := statedb.Copy()
	result, err := RunContract(chain, stateCopy, tokenAddr, contractABI, "decimals")
	if err != nil {
		return nil, err
	}
	decimals = result.(uint8)

	tokenDecimal := new(big.Int).SetUint64(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	vicionx.tokenDecimalCache.Add(tokenAddr, tokenDecimal)
	return tokenDecimal, nil
}

// FIXME: using in unit tests only
func (vicionx *VicionX) SetTokenDecimal(token common.Address, decimal *big.Int) {
	vicionx.tokenDecimalCache.Add(token, decimal)
}

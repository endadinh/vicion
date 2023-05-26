package vm

import (
	"github.com/vicion/vicion/common"
	"github.com/vicion/vicion/log"
	"github.com/vicion/vicion/params"
	"github.com/vicion/vicion/vicionx/tradingstate"
)

const VicionXPriceNumberOfBytesReturn = 32

// vicionxPrice implements a pre-compile contract to get token price in vicionx

type vicionxLastPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}
type vicionxEpochPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}

func (t *vicionxLastPrice) RequiredGas(input []byte) uint64 {
	return params.VicionXPriceGas
}

func (t *vicionxLastPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:])  // 20 bytes from 45-64
		price := t.tradingStateDB.GetLastPrice(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetLastPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), VicionXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, VicionXPriceNumberOfBytesReturn), nil
}

func (t *vicionxLastPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}

func (t *vicionxEpochPrice) RequiredGas(input []byte) uint64 {
	return params.VicionXPriceGas
}

func (t *vicionxEpochPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:])  // 20 bytes from 45-64
		price := t.tradingStateDB.GetMediumPriceBeforeEpoch(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetEpochPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), VicionXPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, VicionXPriceNumberOfBytesReturn), nil
}

func (t *vicionxEpochPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}

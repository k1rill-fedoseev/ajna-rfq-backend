package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

type Order struct {
	Maker         common.Address  `json:"maker,omitempty"`
	Taker         *common.Address `json:"taker,omitempty"`
	Pool          common.Address  `json:"pool,omitempty"`
	Index         int64           `json:"index,omitempty"`
	MakeAmount    decimal.Decimal `json:"makeAmount,omitempty"`
	MinMakeAmount decimal.Decimal `json:"minMakeAmount,omitempty"`
	Expiry        decimal.Decimal `json:"expiry,omitempty"`
	Price         decimal.Decimal `json:"price,omitempty"`
}

type SignedOrder struct {
	Signature hexutil.Bytes `json:"signature"`
	Order
}

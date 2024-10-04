package types

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Order struct {
	LpOrder       bool            `json:"lpOrder,omitempty"`
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

func (order *Order) Validate() error {
	if order.Index <= 0 {
		return errors.New("index must be > 0")
	}
	if order.Index > 7388 {
		return errors.New("index must be <= 7388")
	}
	if order.Price.LessThan(decimal.RequireFromString("100000000000000000")) {
		return errors.New("order price must be >= 0.1")
	}
	if order.Price.GreaterThan(decimal.RequireFromString("1010000000000000000")) {
		return errors.New("order price must be <= 1.01")
	}
	return nil
}

func (order *Order) TakerOrZero() common.Address {
	if order.Taker != nil {
		return *order.Taker
	}
	return common.HexToAddress("0x0000000000000000000000000000000000000000")
}

func (order *Order) Hash(chainId *big.Int, rfqAddress common.Address) (common.Hash, common.Hash, error) {
	data := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Order": []apitypes.Type{
				{Name: "lpOrder", Type: "bool"},
				{Name: "maker", Type: "address"},
				{Name: "taker", Type: "address"},
				{Name: "pool", Type: "address"},
				{Name: "index", Type: "uint256"},
				{Name: "makeAmount", Type: "uint256"},
				{Name: "minMakeAmount", Type: "uint256"},
				{Name: "expiry", Type: "uint256"},
				{Name: "price", Type: "uint256"},
			},
		},
		PrimaryType: "Order",
		Domain: apitypes.TypedDataDomain{
			Name:              "Ajna RFQ",
			Version:           "1",
			ChainId:           (*math.HexOrDecimal256)(chainId),
			VerifyingContract: rfqAddress.String(),
			Salt:              "",
		},
		Message: apitypes.TypedDataMessage{
			"lpOrder":       order.LpOrder,
			"maker":         order.Maker.String(),
			"taker":         order.TakerOrZero().String(),
			"pool":          order.Pool.String(),
			"index":         strconv.FormatInt(order.Index, 10),
			"makeAmount":    order.MakeAmount.String(),
			"minMakeAmount": order.MinMakeAmount.String(),
			"expiry":        order.Expiry.String(),
			"price":         order.Price.String(),
		},
	}
	typedHash, err := data.HashStruct(data.PrimaryType, data.Message)
	if err != nil {
		return common.Hash{}, common.Hash{}, errors.Wrap(err, "can't hash typed data")
	}
	hash, _, err := apitypes.TypedDataAndHash(data)
	if err != nil {
		return common.Hash{}, common.Hash{}, errors.Wrap(err, "can't hash typed data")
	}
	return common.BytesToHash(hash), common.BytesToHash(typedHash), nil
}

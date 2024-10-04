package repo

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"ajna-rfq/internal/types"
)

type OrdersRepo interface {
	ListOrders(filter Filter, limit uint64) ([]*StoredOrder, error)
	Upsert(order StoredOrder) (*StoredOrder, error)
}

type Filter struct {
	ChainId       string           `json:"chainId"`
	LpOrder       bool             `json:"lpOrder"`
	Maker         *common.Address  `json:"maker"`
	Taker         *common.Address  `json:"taker"`
	Active        bool             `json:"active"`
	Pools         []common.Address `json:"pools"`
	CreatedBefore *int64           `json:"createdBefore"`
}

type StoredOrder struct {
	ChainId             string          `json:"chainId"`
	RFQ                 common.Address  `json:"rfq"`
	Hash                common.Hash     `json:"hash"`
	RemainingMakeAmount decimal.Decimal `json:"remainingMakeAmount"`
	ApprovedAmount      decimal.Decimal `json:"approvedAmount"`
	CreatedAt           int64           `json:"createdAt"`
	UpdatedAt           int64           `json:"updatedAt"`
	types.SignedOrder
}

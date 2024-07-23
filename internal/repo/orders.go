package repo

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"ajna-rfq/internal/types"
)

type OrdersRepo interface {
	ListMakeOrders(chainId string, maker common.Address, pools []common.Address, createdBefore *uint64) ([]*StoredOrder, error)
	ListTakeOrders(chainId string, taker *common.Address, pools []common.Address, createdBefore *uint64) ([]*StoredOrder, error)
	ListLastUpdatedActive() ([]*StoredOrder, error)
	Upsert(order *StoredOrder) error
}

type StoredOrder struct {
	ChainId             string          `json:"chainId"`
	RFQ                 common.Address  `json:"rfq"`
	Hash                common.Hash     `json:"hash"`
	RemainingMakeAmount decimal.Decimal `json:"remainingMakeAmount"`
	ApprovedAmount      decimal.Decimal `json:"approvedAmount"`
	UpdatedAt           int64           `json:"updated_at"`
	types.SignedOrder
}

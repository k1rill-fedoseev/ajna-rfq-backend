package service

import (
	"context"
	_ "embed"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"ajna-rfq/internal/config"
	"ajna-rfq/internal/contract"
	"ajna-rfq/internal/repo"
	"ajna-rfq/internal/types"
)

type Service struct {
	repo   repo.OrdersRepo
	chains map[string]*ChainService
}

type ChainService struct {
	chainId    *big.Int
	client     *ethclient.Client
	factory    *contract.FactoryCaller
	rfq        *contract.RFQCaller
	rfqAddress common.Address
	repo       repo.OrdersRepo
}

func NewService(cfg *config.Config, repo repo.OrdersRepo) (*Service, error) {
	chains := make(map[string]*ChainService, len(cfg.Chains))
	for _, chainCfg := range cfg.Chains {
		client, err := ethclient.Dial(chainCfg.RPC)
		if err != nil {
			return nil, errors.Wrapf(err, "can't dial rpc %s", chainCfg.RPC)
		}
		chainId, err := client.ChainID(context.Background())
		if err != nil {
			return nil, errors.Wrapf(err, "can't get chain id %s", chainCfg.RPC)
		}
		chainIdStr := chainId.String()
		if _, ok := chains[chainIdStr]; ok {
			return nil, errors.Errorf("can't use %s, chain id %s already registered", chainCfg.RPC, chainIdStr)
		}
		factory, err := contract.NewFactoryCaller(chainCfg.Factory, client)
		if err != nil {
			return nil, errors.Wrapf(err, "can't create factory caller for %s", chainCfg.Factory)
		}
		rfq, err := contract.NewRFQCaller(chainCfg.RFQ, client)
		if err != nil {
			return nil, errors.Wrapf(err, "can't create rfq caller for %s", chainCfg.RFQ)
		}
		chains[chainIdStr] = &ChainService{
			chainId:    chainId,
			client:     client,
			factory:    factory,
			rfq:        rfq,
			rfqAddress: chainCfg.RFQ,
			repo:       repo,
		}
	}
	return &Service{
		repo:   repo,
		chains: chains,
	}, nil
}

func (s *Service) ChainService(chainId string) (*ChainService, error) {
	if cs, ok := s.chains[chainId]; ok {
		return cs, nil
	}
	return nil, errors.Errorf("chainId %s not configured", chainId)
}

func (cs *ChainService) ListMakeOrders(maker common.Address, pools []common.Address, createdBefore *uint64) ([]*repo.StoredOrder, error) {
	orders, err := cs.repo.ListMakeOrders(cs.chainId.String(), maker, pools, createdBefore)
	if err != nil {
		return nil, errors.Wrap(err, "can't list orders")
	}
	for i, order := range orders {
		if time.Now().Unix()-order.UpdatedAt < 300 {
			continue
		}
		log.Printf("updating order: %s\n", order.Hash)
		newOrder, err := cs.UpdateOrder(order.SignedOrder, false)
		if err != nil {
			log.Printf("can't update order: %v\n", err)
			continue
		}
		orders[i] = newOrder
	}
	return orders, nil
}

func (cs *ChainService) ListTakeOrders(taker *common.Address, pools []common.Address, createdBefore *uint64) ([]*repo.StoredOrder, error) {
	orders, err := cs.repo.ListTakeOrders(cs.chainId.String(), taker, pools, createdBefore)
	if err != nil {
		return nil, errors.Wrap(err, "can't list orders")
	}
	for i, order := range orders {
		newOrder, err := cs.UpdateOrder(order.SignedOrder, false)
		if err != nil {
			log.Printf("can't update order: %v\n", err)
			continue
		}
		orders[i] = newOrder
	}
	return orders, nil
}

func (cs *ChainService) ValidateOrder(order types.SignedOrder) error {
	hash, err := cs.HashOrder(order.Order)
	if err != nil {
		return errors.Wrap(err, "can't hash order")
	}
	valid, err := cs.ValidateOrderSignature(order.Maker, hash, order.Signature)
	if err != nil {
		return errors.Wrap(err, "can't validate order signature")
	}
	if !valid {
		return errors.New("invalid order signature")
	}
	pool, err := contract.NewPoolCaller(order.Pool, cs.client)
	if err != nil {
		return errors.Wrap(err, "can't create pool caller")
	}
	quoteTokenAddress, err := pool.QuoteTokenAddress(nil)
	if err != nil {
		return errors.Wrap(err, "can't call quoteTokenAddress()")
	}
	collateralAddress, err := pool.CollateralAddress(nil)
	if err != nil {
		return errors.Wrap(err, "can't call collateralAddress()")
	}
	subsetHash := common.HexToHash("0x2263c4378b4920f0bef611a3ff22c506afa4745b3319c50b6d704a874990b8b2")
	poolAddress, err := cs.factory.DeployedPools(nil, subsetHash, collateralAddress, quoteTokenAddress)
	if err != nil {
		return errors.Wrap(err, "can't call deployedPools(...)")
	}
	if poolAddress != order.Pool {
		return errors.New("unsupported pool address")
	}
	return nil
}

func (cs *ChainService) CheckApprovedAmount(order types.Order) (bool, *big.Int, error) {
	pool, err := contract.NewPoolCaller(order.Pool, cs.client)
	if err != nil {
		return false, nil, errors.Wrap(err, "can't create pool caller")
	}

	approved, err := pool.ApprovedTransferors(nil, order.Maker, cs.rfqAddress)
	if err != nil {
		return false, nil, errors.Wrap(err, "can't call approvedTransferors(...)")
	}
	if !approved {
		return false, nil, nil
	}
	allowance, err := pool.LpAllowance(nil, big.NewInt(order.Index), cs.rfqAddress, order.Maker)
	if err != nil {
		return false, nil, errors.Wrap(err, "can't call lpAllowance(...)")
	}
	if allowance.Cmp(big.NewInt(0)) == 0 {
		return false, nil, nil
	}
	lenderInfo, err := pool.LenderInfo(nil, big.NewInt(order.Index), order.Maker)
	if err != nil {
		return false, nil, errors.Wrap(err, "can't call lenderInfo(...)")
	}
	if allowance.Cmp(lenderInfo.LpBalance) < 0 {
		return true, allowance, nil
	}
	return true, lenderInfo.LpBalance, nil
}

func (cs *ChainService) ValidateOrderSignature(maker common.Address, hash common.Hash, signature []byte) (bool, error) {
	pk, err := crypto.SigToPub(hash.Bytes(), unpackSignature(signature))
	if err == nil {
		signer := crypto.PubkeyToAddress(*pk)
		if maker == signer {
			return true, nil
		}
	}
	erc1271, err := contract.NewERC1271Caller(maker, cs.client)
	if err != nil {
		return false, errors.Wrap(err, "can't create ERC 1271 caller")
	}
	res, err := erc1271.IsValidSignature(nil, hash, signature)
	if err != nil {
		return false, errors.Wrap(err, "can't call isValidSignature(...)")
	}
	if res == [4]byte{0x16, 0x26, 0xba, 0x7e} {
		return true, nil
	}
	approved, err := cs.rfq.ApprovedOrders(nil, maker, hash)
	if err != nil {
		return false, errors.Wrap(err, "can't check if order approval status")
	}
	return approved, nil
}

func (cs *ChainService) UpdateOrder(order types.SignedOrder, requireApproval bool) (*repo.StoredOrder, error) {
	hash, err := cs.HashOrder(order.Order)
	if err != nil {
		return nil, errors.Wrap(err, "can't hash order")
	}
	rem, err := cs.RemainingAmount(hash, order.Order)
	if err != nil {
		return nil, errors.Wrap(err, "can't get remaining amount")
	}
	approved, approvedAmount, err := cs.CheckApprovedAmount(order.Order)
	if err != nil {
		return nil, errors.Wrap(err, "can't get approved amount")
	}
	if approved && approvedAmount.Cmp(rem) >= 0 {
		approvedAmount = rem
	} else if requireApproval {
		return nil, errors.New("insufficient allowance, full approval is mandatory")
	}
	storedOrder := &repo.StoredOrder{
		ChainId:             cs.chainId.String(),
		RFQ:                 cs.rfqAddress,
		Hash:                hash,
		RemainingMakeAmount: decimal.NewFromBigInt(rem, 0),
		ApprovedAmount:      decimal.NewFromBigInt(approvedAmount, 0),
		SignedOrder:         order,
	}
	err = cs.repo.Upsert(storedOrder)
	if err != nil {
		return nil, errors.Wrap(err, "can't upsert order")
	}
	return storedOrder, nil
}

func (cs *ChainService) RemainingAmount(hash common.Hash, order types.Order) (*big.Int, error) {
	filledAmount, err := cs.rfq.FilledAmounts(nil, order.Maker, hash)
	if err != nil {
		return nil, errors.Wrap(err, "can't call filledAmounts(...)")
	}
	if filledAmount.Cmp(math.MaxBig256) == 0 {
		return big.NewInt(-1), nil
	}
	if filledAmount.Cmp(order.MakeAmount.BigInt()) >= 0 {
		return big.NewInt(0), nil
	}
	return new(big.Int).Sub(order.MakeAmount.BigInt(), filledAmount), nil
}

func (cs *ChainService) HashOrder(order types.Order) (common.Hash, error) {
	taker := "0x0000000000000000000000000000000000000000"
	if order.Taker != nil {
		taker = order.Taker.String()
	}
	data := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "chainId", Type: "uint256"},
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Order": []apitypes.Type{
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
			ChainId:           (*math.HexOrDecimal256)(cs.chainId),
			VerifyingContract: cs.rfqAddress.String(),
			Salt:              "",
		},
		Message: apitypes.TypedDataMessage{
			"maker":         order.Maker.String(),
			"taker":         taker,
			"pool":          order.Pool.String(),
			"index":         strconv.FormatInt(order.Index, 10),
			"makeAmount":    order.MakeAmount.String(),
			"minMakeAmount": order.MinMakeAmount.String(),
			"expiry":        order.Expiry.String(),
			"price":         order.Price.String(),
		},
	}
	hash, _, err := apitypes.TypedDataAndHash(data)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "can't hash typed data")
	}
	return common.BytesToHash(hash), nil
}

func unpackSignature(sig []byte) []byte {
	if len(sig) == 64 {
		unpacked := make([]byte, 65)
		copy(unpacked[1:], sig)
		unpacked[0] = 27 + (sig[32] >> 7)
		unpacked[33] &= 127
		return unpacked
	}
	return sig
}

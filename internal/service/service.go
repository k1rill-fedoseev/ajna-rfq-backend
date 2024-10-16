package service

import (
	"context"
	_ "embed"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
	chainId          *big.Int
	quoteTokens      map[common.Address]common.Address
	quoteTokenScales map[common.Address]*big.Int
	collateralTokens map[common.Address]common.Address
	client           *ethclient.Client
	factory          *contract.FactoryCaller
	rfq              *contract.RFQCaller
	rfqAddress       common.Address
	updateInterval   time.Duration
	repo             repo.OrdersRepo
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
			chainId:          chainId,
			quoteTokens:      make(map[common.Address]common.Address, 10),
			quoteTokenScales: make(map[common.Address]*big.Int, 10),
			collateralTokens: make(map[common.Address]common.Address, 10),
			client:           client,
			factory:          factory,
			rfq:              rfq,
			rfqAddress:       chainCfg.RFQ,
			updateInterval:   time.Duration(chainCfg.UpdateInterval),
			repo:             repo,
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

func (cs *ChainService) ListOrders(filter repo.Filter, limit uint64) ([]*repo.StoredOrder, *int64, error) {
	filter.ChainId = cs.chainId.String()
	res := make([]*repo.StoredOrder, 0, limit)

	if limit == 0 {
		return res, nil, nil
	}

	fetchLimit := limit
	if filter.Active {
		fetchLimit *= 2
	}
	for {
		orders, err := cs.repo.ListOrders(filter, fetchLimit)
		if err != nil {
			return nil, nil, errors.Wrap(err, "can't list orders")
		}
		if len(orders) == 0 {
			return res, nil, nil
		}
		for i, order := range orders {
			newOrder := order
			if time.Now().Add(-cs.updateInterval).Unix() > order.UpdatedAt {
				log.Printf("updating order: %s\n", order.Hash)
				newOrder, err = cs.UpdateOrder(order.SignedOrder, false)
				if err != nil {
					log.Printf("can't update order: %v\n", err)
					newOrder = order
				}
			}
			if !filter.Active || (newOrder.RemainingMakeAmount.IsPositive() && newOrder.ApprovedAmount.IsPositive()) {
				res = append(res, newOrder)
				if len(res) == int(limit) {
					if i < len(orders)-1 || len(orders) == int(fetchLimit) {
						return res, &newOrder.CreatedAt, nil
					}
					return res, nil, nil
				}
			}
			filter.CreatedBefore = &newOrder.CreatedAt
		}
		if len(orders) < int(fetchLimit) {
			return res, nil, nil
		}
	}
}

func (cs *ChainService) ValidateOrder(order types.SignedOrder) error {
	hash, typedHash, err := order.Order.Hash(cs.chainId, cs.rfqAddress)
	if err != nil {
		return errors.Wrap(err, "can't hash order")
	}
	err = order.Order.Validate()
	if err != nil {
		return errors.Wrap(err, "can't validate order")
	}
	valid, err := cs.ValidateOrderSignature(order.Maker, hash, typedHash, order.Signature)
	if err != nil {
		return errors.Wrap(err, "can't validate order signature")
	}
	if !valid {
		return errors.New("invalid order signature")
	}
	quoteAddress, err := cs.QuoteToken(order.Pool)
	if err != nil {
		return err
	}
	collateralAddress, err := cs.CollateralToken(order.Pool)
	if err != nil {
		return err
	}
	subsetHash := common.HexToHash("0x2263c4378b4920f0bef611a3ff22c506afa4745b3319c50b6d704a874990b8b2")
	poolAddress, err := cs.factory.DeployedPools(nil, subsetHash, collateralAddress, quoteAddress)
	if err != nil {
		return errors.Wrap(err, "can't call deployedPools(...)")
	}
	if poolAddress != order.Pool {
		return errors.New("unsupported pool address")
	}
	return nil
}

func (cs *ChainService) QuoteToken(pool common.Address) (common.Address, error) {
	if token, ok := cs.quoteTokens[pool]; ok {
		return token, nil
	}
	caller, err := contract.NewPoolCaller(pool, cs.client)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "can't create pool caller")
	}
	token, err := caller.QuoteTokenAddress(nil)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "can't call quoteTokenAddress()")
	}
	cs.quoteTokens[pool] = token
	return token, nil
}

func (cs *ChainService) QuoteTokenScale(pool common.Address) (*big.Int, error) {
	if scale, ok := cs.quoteTokenScales[pool]; ok {
		return scale, nil
	}
	caller, err := contract.NewPoolCaller(pool, cs.client)
	if err != nil {
		return nil, errors.Wrap(err, "can't create pool caller")
	}
	scale, err := caller.QuoteTokenScale(nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't call quoteTokenScale()")
	}
	cs.quoteTokenScales[pool] = scale
	return scale, nil
}

func (cs *ChainService) CollateralToken(pool common.Address) (common.Address, error) {
	if token, ok := cs.collateralTokens[pool]; ok {
		return token, nil
	}
	caller, err := contract.NewPoolCaller(pool, cs.client)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "can't create pool caller")
	}
	token, err := caller.CollateralAddress(nil)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "can't call collateralAddress()")
	}
	cs.collateralTokens[pool] = token
	return token, nil
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
		return false, big.NewInt(0), nil
	}

	if order.LpOrder {
		index := big.NewInt(order.Index)
		allowance, err := pool.LpAllowance(nil, index, cs.rfqAddress, order.Maker)
		if err != nil {
			return false, nil, errors.Wrap(err, "can't call lpAllowance(...)")
		}
		if allowance.Cmp(big.NewInt(0)) == 0 {
			return false, allowance, nil
		}
		lenderInfo, err := pool.LenderInfo(nil, index, order.Maker)
		if err != nil {
			return false, nil, errors.Wrap(err, "can't call lenderInfo(...)")
		}
		if allowance.Cmp(lenderInfo.LpBalance) < 0 {
			return true, allowance, nil
		}
		return true, lenderInfo.LpBalance, nil
	} else {
		quote, err := cs.QuoteToken(order.Pool)
		if err != nil {
			return false, nil, err
		}
		token, err := contract.NewERC20Caller(quote, cs.client)
		if err != nil {
			return false, nil, errors.Wrap(err, "can't create token caller")
		}
		allowance, err := token.Allowance(nil, order.Maker, cs.rfqAddress)
		if err != nil {
			return false, nil, errors.Wrap(err, "can't call allowance(...)")
		}
		if allowance.Cmp(big.NewInt(0)) == 0 {
			return false, nil, nil
		}
		balance, err := token.BalanceOf(nil, order.Maker)
		if err != nil {
			return false, nil, errors.Wrap(err, "can't call balanceOf(...)")
		}
		if allowance.Cmp(balance) < 0 {
			balance = allowance
		}
		scale, err := cs.QuoteTokenScale(order.Pool)
		if err != nil {
			return false, nil, err
		}
		return true, new(big.Int).Mul(balance, scale), nil
	}
}

func (cs *ChainService) ValidateOrderSignature(maker common.Address, hash common.Hash, typedHash common.Hash, signature []byte) (bool, error) {
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
	approved, err := cs.rfq.ApprovedOrders(nil, maker, typedHash)
	if err != nil {
		return false, errors.Wrap(err, "can't check if order approval status")
	}
	return approved, nil
}

func (cs *ChainService) UpdateOrder(order types.SignedOrder, requireApproval bool) (*repo.StoredOrder, error) {
	_, typedHash, err := order.Order.Hash(cs.chainId, cs.rfqAddress)
	if err != nil {
		return nil, errors.Wrap(err, "can't hash order")
	}
	rem, err := cs.RemainingAmount(typedHash, order.Order)
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
	newOrder, err := cs.repo.Upsert(repo.StoredOrder{
		ChainId:             cs.chainId.String(),
		RFQ:                 cs.rfqAddress,
		Hash:                typedHash,
		RemainingMakeAmount: decimal.NewFromBigInt(rem, 0),
		ApprovedAmount:      decimal.NewFromBigInt(approvedAmount, 0),
		SignedOrder:         order,
	})
	if err != nil {
		return nil, errors.Wrap(err, "can't upsert order")
	}
	return newOrder, nil
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

func unpackSignature(sig []byte) []byte {
	unpacked := make([]byte, 65)
	if len(sig) == 64 {
		copy(unpacked[:64], sig)
		unpacked[32] &= 127
		unpacked[64] = sig[32] >> 7
		return unpacked
	} else {
		copy(unpacked, sig)
		unpacked[64] -= 27
	}
	return unpacked
}

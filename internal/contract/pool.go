// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addCollateral\",\"inputs\":[{\"name\":\"amountToAdd_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"bucketLP_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addQuoteToken\",\"inputs\":[{\"name\":\"amount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"bucketLP_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addedAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveLPTransferors\",\"inputs\":[{\"name\":\"transferors_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvedTransferors\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctionInfo\",\"inputs\":[{\"name\":\"borrower_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"kicker_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bondFactor_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondSize_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kickTime_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"referencePrice_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"neutralPrice_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debtToCollateral_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"head_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"next_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"prev_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"borrowerInfo\",\"inputs\":[{\"name\":\"borrower_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bucketCollateralDust\",\"inputs\":[{\"name\":\"bucketIndex_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"bucketExchangeRate\",\"inputs\":[{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"exchangeRate_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bucketInfo\",\"inputs\":[{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bucketTake\",\"inputs\":[{\"name\":\"borrowerAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositTake_\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnInfo\",\"inputs\":[{\"name\":\"burnEventEpoch_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"collateralScale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"currentBurnEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"debtInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseLPAllowance\",\"inputs\":[{\"name\":\"spender_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"indexes_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amounts_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIndex\",\"inputs\":[{\"name\":\"debt_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositScale\",\"inputs\":[{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositUpToIndex\",\"inputs\":[{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositUtilization\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"drawDebt\",\"inputs\":[{\"name\":\"borrowerAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountToBorrow_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limitIndex_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralToPledge_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emasInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flashFee\",\"inputs\":[{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flashLoan\",\"inputs\":[{\"name\":\"receiver_\",\"type\":\"address\",\"internalType\":\"contractIERC3156FlashBorrower\"},{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseLPAllowance\",\"inputs\":[{\"name\":\"spender_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"indexes_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amounts_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"inflatorInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"rate_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"interestRateInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"kick\",\"inputs\":[{\"name\":\"borrower_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"npLimitIndex_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"kickReserveAuction\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"kickerInfo\",\"inputs\":[{\"name\":\"kicker_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lenderInfo\",\"inputs\":[{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lender_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"lpBalance_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"depositTime_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lenderKick\",\"inputs\":[{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"npLimitIndex_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"loanInfo\",\"inputs\":[{\"name\":\"loanId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"loansInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lpAllowance\",\"inputs\":[{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"spender_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowance_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxFlashLoan\",\"inputs\":[{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"maxLoan_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"moveQuoteToken\",\"inputs\":[{\"name\":\"maxAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromIndex_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toIndex_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fromBucketLP_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toBucketLP_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"movedAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pledgedCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"quoteTokenAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"quoteTokenScale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"removeCollateral\",\"inputs\":[{\"name\":\"maxAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"removedAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redeemedLP_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeQuoteToken\",\"inputs\":[{\"name\":\"maxAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"removedAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redeemedLP_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"repayDebt\",\"inputs\":[{\"name\":\"borrowerAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxQuoteTokenAmountToRepay_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralAmountToPull_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralReceiver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"limitIndex_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountRepaid_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reservesInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeLPAllowance\",\"inputs\":[{\"name\":\"spender_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"indexes_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeLPTransferors\",\"inputs\":[{\"name\":\"transferors_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settle\",\"inputs\":[{\"name\":\"borrowerAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxDepth_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"collateralSettled_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isBorrowerSettled_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stampLoan\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"take\",\"inputs\":[{\"name\":\"borrowerAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callee_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"collateralTaken_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"takeReserves\",\"inputs\":[{\"name\":\"maxAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalAuctionsInPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalT0Debt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalT0DebtInAuction\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferLP\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newOwner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"indexes_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateInterest\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawBonds\",\"inputs\":[{\"name\":\"recipient_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"withdrawnAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddCollateral\",\"inputs\":[{\"name\":\"actor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lpAwarded\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddQuoteToken\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lpAwarded\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lup\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApproveLPTransferors\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"transferors\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionNFTSettle\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionSettle\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondWithdrawn\",\"inputs\":[{\"name\":\"kicker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reciever\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BucketBankruptcy\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"lpForfeited\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BucketTake\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bondChange\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isReward\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BucketTakeLPAwarded\",\"inputs\":[{\"name\":\"taker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"kicker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"lpAwardedTaker\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lpAwardedKicker\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseLPAllowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"indexes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DrawDebt\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountBorrowed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralPledged\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lup\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Flashloan\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseLPAllowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"indexes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InterestUpdateFailure\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Kick\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"debt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KickReserveAuction\",\"inputs\":[{\"name\":\"claimableReservesRemaining\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"auctionPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currentBurnEpoch\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LoanStamped\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MoveQuoteToken\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lpRedeemedFrom\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lpAwardedTo\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lup\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoveCollateral\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lpRedeemed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoveQuoteToken\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lpRedeemed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lup\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RepayDebt\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"quoteRepaid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralPulled\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lup\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReserveAuction\",\"inputs\":[{\"name\":\"claimableReservesRemaining\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"auctionPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currentBurnEpoch\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ResetInterestRate\",\"inputs\":[{\"name\":\"oldRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevokeLPAllowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"indexes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevokeLPTransferors\",\"inputs\":[{\"name\":\"lender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"transferors\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settle\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"settledDebt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Take\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bondChange\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isReward\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferLP\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"lp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateInterestRate\",\"inputs\":[{\"name\":\"oldRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddAboveAuctionPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddAboveAuctionPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountLTMinDebt\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotClearable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotCleared\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotCleared\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotTakeable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionPriceGtBucketPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BorrowerNotSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BorrowerOk\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BorrowerUnderCollateralized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketBankruptcyBlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketIndexOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotMergeToHigherPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DustAmountNotExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FlashloanCallbackFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FlashloanIncorrectBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FlashloanUnavailableForToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientCollateral\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLP\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowancesInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LUPBelowHTP\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LimitIndexExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MoveToSameIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoAuction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoDebt\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoReserves\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoReservesAuction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__DivInputTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__DivOverflow\",\"inputs\":[{\"name\":\"rAbs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__Exp2InputTooBig\",\"inputs\":[{\"name\":\"x\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__FromIntOverflow\",\"inputs\":[{\"name\":\"x\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__FromIntUnderflow\",\"inputs\":[{\"name\":\"x\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__LogInputTooSmall\",\"inputs\":[{\"name\":\"x\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__MulInputTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__MulOverflow\",\"inputs\":[{\"name\":\"rAbs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__SqrtNegativeInput\",\"inputs\":[{\"name\":\"x\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"PRBMathSD59x18__SqrtOverflow\",\"inputs\":[{\"name\":\"x\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"type\":\"error\",\"name\":\"PRBMath__MulDivFixedPointOverflow\",\"inputs\":[{\"name\":\"prod1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PRBMath__MulDivOverflow\",\"inputs\":[{\"name\":\"prod1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PriceBelowLUP\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoveDepositLockedByAuctionDebt\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RemoveDepositLockedByAuctionDebt\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReserveAuctionTooSoon\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransactionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransactionExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToSameOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferorNotApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroDebtToCollateral\",\"inputs\":[]}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// ApprovedTransferors is a free data retrieval call binding the contract method 0xd9606e08.
//
// Solidity: function approvedTransferors(address , address ) view returns(bool)
func (_Pool *PoolCaller) ApprovedTransferors(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "approvedTransferors", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedTransferors is a free data retrieval call binding the contract method 0xd9606e08.
//
// Solidity: function approvedTransferors(address , address ) view returns(bool)
func (_Pool *PoolSession) ApprovedTransferors(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Pool.Contract.ApprovedTransferors(&_Pool.CallOpts, arg0, arg1)
}

// ApprovedTransferors is a free data retrieval call binding the contract method 0xd9606e08.
//
// Solidity: function approvedTransferors(address , address ) view returns(bool)
func (_Pool *PoolCallerSession) ApprovedTransferors(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Pool.Contract.ApprovedTransferors(&_Pool.CallOpts, arg0, arg1)
}

// AuctionInfo is a free data retrieval call binding the contract method 0x0448e51a.
//
// Solidity: function auctionInfo(address borrower_) view returns(address kicker_, uint256 bondFactor_, uint256 bondSize_, uint256 kickTime_, uint256 referencePrice_, uint256 neutralPrice_, uint256 debtToCollateral_, address head_, address next_, address prev_)
func (_Pool *PoolCaller) AuctionInfo(opts *bind.CallOpts, borrower_ common.Address) (struct {
	Kicker           common.Address
	BondFactor       *big.Int
	BondSize         *big.Int
	KickTime         *big.Int
	ReferencePrice   *big.Int
	NeutralPrice     *big.Int
	DebtToCollateral *big.Int
	Head             common.Address
	Next             common.Address
	Prev             common.Address
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "auctionInfo", borrower_)

	outstruct := new(struct {
		Kicker           common.Address
		BondFactor       *big.Int
		BondSize         *big.Int
		KickTime         *big.Int
		ReferencePrice   *big.Int
		NeutralPrice     *big.Int
		DebtToCollateral *big.Int
		Head             common.Address
		Next             common.Address
		Prev             common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Kicker = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BondFactor = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BondSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.KickTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReferencePrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NeutralPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DebtToCollateral = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Head = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.Next = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.Prev = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AuctionInfo is a free data retrieval call binding the contract method 0x0448e51a.
//
// Solidity: function auctionInfo(address borrower_) view returns(address kicker_, uint256 bondFactor_, uint256 bondSize_, uint256 kickTime_, uint256 referencePrice_, uint256 neutralPrice_, uint256 debtToCollateral_, address head_, address next_, address prev_)
func (_Pool *PoolSession) AuctionInfo(borrower_ common.Address) (struct {
	Kicker           common.Address
	BondFactor       *big.Int
	BondSize         *big.Int
	KickTime         *big.Int
	ReferencePrice   *big.Int
	NeutralPrice     *big.Int
	DebtToCollateral *big.Int
	Head             common.Address
	Next             common.Address
	Prev             common.Address
}, error) {
	return _Pool.Contract.AuctionInfo(&_Pool.CallOpts, borrower_)
}

// AuctionInfo is a free data retrieval call binding the contract method 0x0448e51a.
//
// Solidity: function auctionInfo(address borrower_) view returns(address kicker_, uint256 bondFactor_, uint256 bondSize_, uint256 kickTime_, uint256 referencePrice_, uint256 neutralPrice_, uint256 debtToCollateral_, address head_, address next_, address prev_)
func (_Pool *PoolCallerSession) AuctionInfo(borrower_ common.Address) (struct {
	Kicker           common.Address
	BondFactor       *big.Int
	BondSize         *big.Int
	KickTime         *big.Int
	ReferencePrice   *big.Int
	NeutralPrice     *big.Int
	DebtToCollateral *big.Int
	Head             common.Address
	Next             common.Address
	Prev             common.Address
}, error) {
	return _Pool.Contract.AuctionInfo(&_Pool.CallOpts, borrower_)
}

// BorrowerInfo is a free data retrieval call binding the contract method 0xca103d15.
//
// Solidity: function borrowerInfo(address borrower_) view returns(uint256, uint256, uint256)
func (_Pool *PoolCaller) BorrowerInfo(opts *bind.CallOpts, borrower_ common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "borrowerInfo", borrower_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BorrowerInfo is a free data retrieval call binding the contract method 0xca103d15.
//
// Solidity: function borrowerInfo(address borrower_) view returns(uint256, uint256, uint256)
func (_Pool *PoolSession) BorrowerInfo(borrower_ common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.BorrowerInfo(&_Pool.CallOpts, borrower_)
}

// BorrowerInfo is a free data retrieval call binding the contract method 0xca103d15.
//
// Solidity: function borrowerInfo(address borrower_) view returns(uint256, uint256, uint256)
func (_Pool *PoolCallerSession) BorrowerInfo(borrower_ common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.BorrowerInfo(&_Pool.CallOpts, borrower_)
}

// BucketCollateralDust is a free data retrieval call binding the contract method 0x540c1433.
//
// Solidity: function bucketCollateralDust(uint256 bucketIndex_) pure returns(uint256)
func (_Pool *PoolCaller) BucketCollateralDust(opts *bind.CallOpts, bucketIndex_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "bucketCollateralDust", bucketIndex_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BucketCollateralDust is a free data retrieval call binding the contract method 0x540c1433.
//
// Solidity: function bucketCollateralDust(uint256 bucketIndex_) pure returns(uint256)
func (_Pool *PoolSession) BucketCollateralDust(bucketIndex_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.BucketCollateralDust(&_Pool.CallOpts, bucketIndex_)
}

// BucketCollateralDust is a free data retrieval call binding the contract method 0x540c1433.
//
// Solidity: function bucketCollateralDust(uint256 bucketIndex_) pure returns(uint256)
func (_Pool *PoolCallerSession) BucketCollateralDust(bucketIndex_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.BucketCollateralDust(&_Pool.CallOpts, bucketIndex_)
}

// BucketExchangeRate is a free data retrieval call binding the contract method 0x16f8a463.
//
// Solidity: function bucketExchangeRate(uint256 index_) view returns(uint256 exchangeRate_)
func (_Pool *PoolCaller) BucketExchangeRate(opts *bind.CallOpts, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "bucketExchangeRate", index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BucketExchangeRate is a free data retrieval call binding the contract method 0x16f8a463.
//
// Solidity: function bucketExchangeRate(uint256 index_) view returns(uint256 exchangeRate_)
func (_Pool *PoolSession) BucketExchangeRate(index_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.BucketExchangeRate(&_Pool.CallOpts, index_)
}

// BucketExchangeRate is a free data retrieval call binding the contract method 0x16f8a463.
//
// Solidity: function bucketExchangeRate(uint256 index_) view returns(uint256 exchangeRate_)
func (_Pool *PoolCallerSession) BucketExchangeRate(index_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.BucketExchangeRate(&_Pool.CallOpts, index_)
}

// BucketInfo is a free data retrieval call binding the contract method 0xa83de3ec.
//
// Solidity: function bucketInfo(uint256 index_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_Pool *PoolCaller) BucketInfo(opts *bind.CallOpts, index_ *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "bucketInfo", index_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// BucketInfo is a free data retrieval call binding the contract method 0xa83de3ec.
//
// Solidity: function bucketInfo(uint256 index_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_Pool *PoolSession) BucketInfo(index_ *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.BucketInfo(&_Pool.CallOpts, index_)
}

// BucketInfo is a free data retrieval call binding the contract method 0xa83de3ec.
//
// Solidity: function bucketInfo(uint256 index_) view returns(uint256, uint256, uint256, uint256, uint256)
func (_Pool *PoolCallerSession) BucketInfo(index_ *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.BucketInfo(&_Pool.CallOpts, index_)
}

// BurnInfo is a free data retrieval call binding the contract method 0x2c7b2e06.
//
// Solidity: function burnInfo(uint256 burnEventEpoch_) view returns(uint256, uint256, uint256)
func (_Pool *PoolCaller) BurnInfo(opts *bind.CallOpts, burnEventEpoch_ *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "burnInfo", burnEventEpoch_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// BurnInfo is a free data retrieval call binding the contract method 0x2c7b2e06.
//
// Solidity: function burnInfo(uint256 burnEventEpoch_) view returns(uint256, uint256, uint256)
func (_Pool *PoolSession) BurnInfo(burnEventEpoch_ *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.BurnInfo(&_Pool.CallOpts, burnEventEpoch_)
}

// BurnInfo is a free data retrieval call binding the contract method 0x2c7b2e06.
//
// Solidity: function burnInfo(uint256 burnEventEpoch_) view returns(uint256, uint256, uint256)
func (_Pool *PoolCallerSession) BurnInfo(burnEventEpoch_ *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.BurnInfo(&_Pool.CallOpts, burnEventEpoch_)
}

// CollateralAddress is a free data retrieval call binding the contract method 0x48d399e7.
//
// Solidity: function collateralAddress() pure returns(address)
func (_Pool *PoolCaller) CollateralAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "collateralAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralAddress is a free data retrieval call binding the contract method 0x48d399e7.
//
// Solidity: function collateralAddress() pure returns(address)
func (_Pool *PoolSession) CollateralAddress() (common.Address, error) {
	return _Pool.Contract.CollateralAddress(&_Pool.CallOpts)
}

// CollateralAddress is a free data retrieval call binding the contract method 0x48d399e7.
//
// Solidity: function collateralAddress() pure returns(address)
func (_Pool *PoolCallerSession) CollateralAddress() (common.Address, error) {
	return _Pool.Contract.CollateralAddress(&_Pool.CallOpts)
}

// CollateralScale is a free data retrieval call binding the contract method 0xec0bdcfc.
//
// Solidity: function collateralScale() pure returns(uint256)
func (_Pool *PoolCaller) CollateralScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "collateralScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralScale is a free data retrieval call binding the contract method 0xec0bdcfc.
//
// Solidity: function collateralScale() pure returns(uint256)
func (_Pool *PoolSession) CollateralScale() (*big.Int, error) {
	return _Pool.Contract.CollateralScale(&_Pool.CallOpts)
}

// CollateralScale is a free data retrieval call binding the contract method 0xec0bdcfc.
//
// Solidity: function collateralScale() pure returns(uint256)
func (_Pool *PoolCallerSession) CollateralScale() (*big.Int, error) {
	return _Pool.Contract.CollateralScale(&_Pool.CallOpts)
}

// CurrentBurnEpoch is a free data retrieval call binding the contract method 0x4ab1fc36.
//
// Solidity: function currentBurnEpoch() view returns(uint256)
func (_Pool *PoolCaller) CurrentBurnEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "currentBurnEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentBurnEpoch is a free data retrieval call binding the contract method 0x4ab1fc36.
//
// Solidity: function currentBurnEpoch() view returns(uint256)
func (_Pool *PoolSession) CurrentBurnEpoch() (*big.Int, error) {
	return _Pool.Contract.CurrentBurnEpoch(&_Pool.CallOpts)
}

// CurrentBurnEpoch is a free data retrieval call binding the contract method 0x4ab1fc36.
//
// Solidity: function currentBurnEpoch() view returns(uint256)
func (_Pool *PoolCallerSession) CurrentBurnEpoch() (*big.Int, error) {
	return _Pool.Contract.CurrentBurnEpoch(&_Pool.CallOpts)
}

// DebtInfo is a free data retrieval call binding the contract method 0x4d966198.
//
// Solidity: function debtInfo() view returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolCaller) DebtInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "debtInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// DebtInfo is a free data retrieval call binding the contract method 0x4d966198.
//
// Solidity: function debtInfo() view returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolSession) DebtInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.DebtInfo(&_Pool.CallOpts)
}

// DebtInfo is a free data retrieval call binding the contract method 0x4d966198.
//
// Solidity: function debtInfo() view returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolCallerSession) DebtInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.DebtInfo(&_Pool.CallOpts)
}

// DepositIndex is a free data retrieval call binding the contract method 0x329d1a8b.
//
// Solidity: function depositIndex(uint256 debt_) view returns(uint256)
func (_Pool *PoolCaller) DepositIndex(opts *bind.CallOpts, debt_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "depositIndex", debt_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositIndex is a free data retrieval call binding the contract method 0x329d1a8b.
//
// Solidity: function depositIndex(uint256 debt_) view returns(uint256)
func (_Pool *PoolSession) DepositIndex(debt_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.DepositIndex(&_Pool.CallOpts, debt_)
}

// DepositIndex is a free data retrieval call binding the contract method 0x329d1a8b.
//
// Solidity: function depositIndex(uint256 debt_) view returns(uint256)
func (_Pool *PoolCallerSession) DepositIndex(debt_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.DepositIndex(&_Pool.CallOpts, debt_)
}

// DepositScale is a free data retrieval call binding the contract method 0x3ab96ec5.
//
// Solidity: function depositScale(uint256 index_) view returns(uint256)
func (_Pool *PoolCaller) DepositScale(opts *bind.CallOpts, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "depositScale", index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositScale is a free data retrieval call binding the contract method 0x3ab96ec5.
//
// Solidity: function depositScale(uint256 index_) view returns(uint256)
func (_Pool *PoolSession) DepositScale(index_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.DepositScale(&_Pool.CallOpts, index_)
}

// DepositScale is a free data retrieval call binding the contract method 0x3ab96ec5.
//
// Solidity: function depositScale(uint256 index_) view returns(uint256)
func (_Pool *PoolCallerSession) DepositScale(index_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.DepositScale(&_Pool.CallOpts, index_)
}

// DepositSize is a free data retrieval call binding the contract method 0x3fa8fdbb.
//
// Solidity: function depositSize() view returns(uint256)
func (_Pool *PoolCaller) DepositSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "depositSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositSize is a free data retrieval call binding the contract method 0x3fa8fdbb.
//
// Solidity: function depositSize() view returns(uint256)
func (_Pool *PoolSession) DepositSize() (*big.Int, error) {
	return _Pool.Contract.DepositSize(&_Pool.CallOpts)
}

// DepositSize is a free data retrieval call binding the contract method 0x3fa8fdbb.
//
// Solidity: function depositSize() view returns(uint256)
func (_Pool *PoolCallerSession) DepositSize() (*big.Int, error) {
	return _Pool.Contract.DepositSize(&_Pool.CallOpts)
}

// DepositUpToIndex is a free data retrieval call binding the contract method 0xda7951a9.
//
// Solidity: function depositUpToIndex(uint256 index_) view returns(uint256)
func (_Pool *PoolCaller) DepositUpToIndex(opts *bind.CallOpts, index_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "depositUpToIndex", index_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositUpToIndex is a free data retrieval call binding the contract method 0xda7951a9.
//
// Solidity: function depositUpToIndex(uint256 index_) view returns(uint256)
func (_Pool *PoolSession) DepositUpToIndex(index_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.DepositUpToIndex(&_Pool.CallOpts, index_)
}

// DepositUpToIndex is a free data retrieval call binding the contract method 0xda7951a9.
//
// Solidity: function depositUpToIndex(uint256 index_) view returns(uint256)
func (_Pool *PoolCallerSession) DepositUpToIndex(index_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.DepositUpToIndex(&_Pool.CallOpts, index_)
}

// DepositUtilization is a free data retrieval call binding the contract method 0x3a0c8f07.
//
// Solidity: function depositUtilization() view returns(uint256)
func (_Pool *PoolCaller) DepositUtilization(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "depositUtilization")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositUtilization is a free data retrieval call binding the contract method 0x3a0c8f07.
//
// Solidity: function depositUtilization() view returns(uint256)
func (_Pool *PoolSession) DepositUtilization() (*big.Int, error) {
	return _Pool.Contract.DepositUtilization(&_Pool.CallOpts)
}

// DepositUtilization is a free data retrieval call binding the contract method 0x3a0c8f07.
//
// Solidity: function depositUtilization() view returns(uint256)
func (_Pool *PoolCallerSession) DepositUtilization() (*big.Int, error) {
	return _Pool.Contract.DepositUtilization(&_Pool.CallOpts)
}

// EmasInfo is a free data retrieval call binding the contract method 0xe512c061.
//
// Solidity: function emasInfo() view returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolCaller) EmasInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "emasInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// EmasInfo is a free data retrieval call binding the contract method 0xe512c061.
//
// Solidity: function emasInfo() view returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolSession) EmasInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.EmasInfo(&_Pool.CallOpts)
}

// EmasInfo is a free data retrieval call binding the contract method 0xe512c061.
//
// Solidity: function emasInfo() view returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolCallerSession) EmasInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.EmasInfo(&_Pool.CallOpts)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token_, uint256 ) view returns(uint256)
func (_Pool *PoolCaller) FlashFee(opts *bind.CallOpts, token_ common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "flashFee", token_, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token_, uint256 ) view returns(uint256)
func (_Pool *PoolSession) FlashFee(token_ common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Pool.Contract.FlashFee(&_Pool.CallOpts, token_, arg1)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token_, uint256 ) view returns(uint256)
func (_Pool *PoolCallerSession) FlashFee(token_ common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Pool.Contract.FlashFee(&_Pool.CallOpts, token_, arg1)
}

// InflatorInfo is a free data retrieval call binding the contract method 0x063d829f.
//
// Solidity: function inflatorInfo() view returns(uint256, uint256)
func (_Pool *PoolCaller) InflatorInfo(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "inflatorInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// InflatorInfo is a free data retrieval call binding the contract method 0x063d829f.
//
// Solidity: function inflatorInfo() view returns(uint256, uint256)
func (_Pool *PoolSession) InflatorInfo() (*big.Int, *big.Int, error) {
	return _Pool.Contract.InflatorInfo(&_Pool.CallOpts)
}

// InflatorInfo is a free data retrieval call binding the contract method 0x063d829f.
//
// Solidity: function inflatorInfo() view returns(uint256, uint256)
func (_Pool *PoolCallerSession) InflatorInfo() (*big.Int, *big.Int, error) {
	return _Pool.Contract.InflatorInfo(&_Pool.CallOpts)
}

// InterestRateInfo is a free data retrieval call binding the contract method 0x00cdcefb.
//
// Solidity: function interestRateInfo() view returns(uint256, uint256)
func (_Pool *PoolCaller) InterestRateInfo(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "interestRateInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// InterestRateInfo is a free data retrieval call binding the contract method 0x00cdcefb.
//
// Solidity: function interestRateInfo() view returns(uint256, uint256)
func (_Pool *PoolSession) InterestRateInfo() (*big.Int, *big.Int, error) {
	return _Pool.Contract.InterestRateInfo(&_Pool.CallOpts)
}

// InterestRateInfo is a free data retrieval call binding the contract method 0x00cdcefb.
//
// Solidity: function interestRateInfo() view returns(uint256, uint256)
func (_Pool *PoolCallerSession) InterestRateInfo() (*big.Int, *big.Int, error) {
	return _Pool.Contract.InterestRateInfo(&_Pool.CallOpts)
}

// KickerInfo is a free data retrieval call binding the contract method 0x7323f853.
//
// Solidity: function kickerInfo(address kicker_) view returns(uint256, uint256)
func (_Pool *PoolCaller) KickerInfo(opts *bind.CallOpts, kicker_ common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "kickerInfo", kicker_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// KickerInfo is a free data retrieval call binding the contract method 0x7323f853.
//
// Solidity: function kickerInfo(address kicker_) view returns(uint256, uint256)
func (_Pool *PoolSession) KickerInfo(kicker_ common.Address) (*big.Int, *big.Int, error) {
	return _Pool.Contract.KickerInfo(&_Pool.CallOpts, kicker_)
}

// KickerInfo is a free data retrieval call binding the contract method 0x7323f853.
//
// Solidity: function kickerInfo(address kicker_) view returns(uint256, uint256)
func (_Pool *PoolCallerSession) KickerInfo(kicker_ common.Address) (*big.Int, *big.Int, error) {
	return _Pool.Contract.KickerInfo(&_Pool.CallOpts, kicker_)
}

// LenderInfo is a free data retrieval call binding the contract method 0xa749f1a6.
//
// Solidity: function lenderInfo(uint256 index_, address lender_) view returns(uint256 lpBalance_, uint256 depositTime_)
func (_Pool *PoolCaller) LenderInfo(opts *bind.CallOpts, index_ *big.Int, lender_ common.Address) (struct {
	LpBalance   *big.Int
	DepositTime *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "lenderInfo", index_, lender_)

	outstruct := new(struct {
		LpBalance   *big.Int
		DepositTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DepositTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LenderInfo is a free data retrieval call binding the contract method 0xa749f1a6.
//
// Solidity: function lenderInfo(uint256 index_, address lender_) view returns(uint256 lpBalance_, uint256 depositTime_)
func (_Pool *PoolSession) LenderInfo(index_ *big.Int, lender_ common.Address) (struct {
	LpBalance   *big.Int
	DepositTime *big.Int
}, error) {
	return _Pool.Contract.LenderInfo(&_Pool.CallOpts, index_, lender_)
}

// LenderInfo is a free data retrieval call binding the contract method 0xa749f1a6.
//
// Solidity: function lenderInfo(uint256 index_, address lender_) view returns(uint256 lpBalance_, uint256 depositTime_)
func (_Pool *PoolCallerSession) LenderInfo(index_ *big.Int, lender_ common.Address) (struct {
	LpBalance   *big.Int
	DepositTime *big.Int
}, error) {
	return _Pool.Contract.LenderInfo(&_Pool.CallOpts, index_, lender_)
}

// LoanInfo is a free data retrieval call binding the contract method 0x8349d6be.
//
// Solidity: function loanInfo(uint256 loanId_) view returns(address, uint256)
func (_Pool *PoolCaller) LoanInfo(opts *bind.CallOpts, loanId_ *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "loanInfo", loanId_)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LoanInfo is a free data retrieval call binding the contract method 0x8349d6be.
//
// Solidity: function loanInfo(uint256 loanId_) view returns(address, uint256)
func (_Pool *PoolSession) LoanInfo(loanId_ *big.Int) (common.Address, *big.Int, error) {
	return _Pool.Contract.LoanInfo(&_Pool.CallOpts, loanId_)
}

// LoanInfo is a free data retrieval call binding the contract method 0x8349d6be.
//
// Solidity: function loanInfo(uint256 loanId_) view returns(address, uint256)
func (_Pool *PoolCallerSession) LoanInfo(loanId_ *big.Int) (common.Address, *big.Int, error) {
	return _Pool.Contract.LoanInfo(&_Pool.CallOpts, loanId_)
}

// LoansInfo is a free data retrieval call binding the contract method 0x3884cd88.
//
// Solidity: function loansInfo() view returns(address, uint256, uint256)
func (_Pool *PoolCaller) LoansInfo(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "loansInfo")

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// LoansInfo is a free data retrieval call binding the contract method 0x3884cd88.
//
// Solidity: function loansInfo() view returns(address, uint256, uint256)
func (_Pool *PoolSession) LoansInfo() (common.Address, *big.Int, *big.Int, error) {
	return _Pool.Contract.LoansInfo(&_Pool.CallOpts)
}

// LoansInfo is a free data retrieval call binding the contract method 0x3884cd88.
//
// Solidity: function loansInfo() view returns(address, uint256, uint256)
func (_Pool *PoolCallerSession) LoansInfo() (common.Address, *big.Int, *big.Int, error) {
	return _Pool.Contract.LoansInfo(&_Pool.CallOpts)
}

// LpAllowance is a free data retrieval call binding the contract method 0x483cd187.
//
// Solidity: function lpAllowance(uint256 index_, address spender_, address owner_) view returns(uint256 allowance_)
func (_Pool *PoolCaller) LpAllowance(opts *bind.CallOpts, index_ *big.Int, spender_ common.Address, owner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "lpAllowance", index_, spender_, owner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpAllowance is a free data retrieval call binding the contract method 0x483cd187.
//
// Solidity: function lpAllowance(uint256 index_, address spender_, address owner_) view returns(uint256 allowance_)
func (_Pool *PoolSession) LpAllowance(index_ *big.Int, spender_ common.Address, owner_ common.Address) (*big.Int, error) {
	return _Pool.Contract.LpAllowance(&_Pool.CallOpts, index_, spender_, owner_)
}

// LpAllowance is a free data retrieval call binding the contract method 0x483cd187.
//
// Solidity: function lpAllowance(uint256 index_, address spender_, address owner_) view returns(uint256 allowance_)
func (_Pool *PoolCallerSession) LpAllowance(index_ *big.Int, spender_ common.Address, owner_ common.Address) (*big.Int, error) {
	return _Pool.Contract.LpAllowance(&_Pool.CallOpts, index_, spender_, owner_)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token_) view returns(uint256 maxLoan_)
func (_Pool *PoolCaller) MaxFlashLoan(opts *bind.CallOpts, token_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "maxFlashLoan", token_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token_) view returns(uint256 maxLoan_)
func (_Pool *PoolSession) MaxFlashLoan(token_ common.Address) (*big.Int, error) {
	return _Pool.Contract.MaxFlashLoan(&_Pool.CallOpts, token_)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token_) view returns(uint256 maxLoan_)
func (_Pool *PoolCallerSession) MaxFlashLoan(token_ common.Address) (*big.Int, error) {
	return _Pool.Contract.MaxFlashLoan(&_Pool.CallOpts, token_)
}

// PledgedCollateral is a free data retrieval call binding the contract method 0x307ee3b5.
//
// Solidity: function pledgedCollateral() view returns(uint256)
func (_Pool *PoolCaller) PledgedCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "pledgedCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgedCollateral is a free data retrieval call binding the contract method 0x307ee3b5.
//
// Solidity: function pledgedCollateral() view returns(uint256)
func (_Pool *PoolSession) PledgedCollateral() (*big.Int, error) {
	return _Pool.Contract.PledgedCollateral(&_Pool.CallOpts)
}

// PledgedCollateral is a free data retrieval call binding the contract method 0x307ee3b5.
//
// Solidity: function pledgedCollateral() view returns(uint256)
func (_Pool *PoolCallerSession) PledgedCollateral() (*big.Int, error) {
	return _Pool.Contract.PledgedCollateral(&_Pool.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() pure returns(uint8)
func (_Pool *PoolCaller) PoolType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() pure returns(uint8)
func (_Pool *PoolSession) PoolType() (uint8, error) {
	return _Pool.Contract.PoolType(&_Pool.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() pure returns(uint8)
func (_Pool *PoolCallerSession) PoolType() (uint8, error) {
	return _Pool.Contract.PoolType(&_Pool.CallOpts)
}

// QuoteTokenAddress is a free data retrieval call binding the contract method 0xbad34620.
//
// Solidity: function quoteTokenAddress() pure returns(address)
func (_Pool *PoolCaller) QuoteTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "quoteTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteTokenAddress is a free data retrieval call binding the contract method 0xbad34620.
//
// Solidity: function quoteTokenAddress() pure returns(address)
func (_Pool *PoolSession) QuoteTokenAddress() (common.Address, error) {
	return _Pool.Contract.QuoteTokenAddress(&_Pool.CallOpts)
}

// QuoteTokenAddress is a free data retrieval call binding the contract method 0xbad34620.
//
// Solidity: function quoteTokenAddress() pure returns(address)
func (_Pool *PoolCallerSession) QuoteTokenAddress() (common.Address, error) {
	return _Pool.Contract.QuoteTokenAddress(&_Pool.CallOpts)
}

// QuoteTokenScale is a free data retrieval call binding the contract method 0x7b3f8655.
//
// Solidity: function quoteTokenScale() pure returns(uint256)
func (_Pool *PoolCaller) QuoteTokenScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "quoteTokenScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteTokenScale is a free data retrieval call binding the contract method 0x7b3f8655.
//
// Solidity: function quoteTokenScale() pure returns(uint256)
func (_Pool *PoolSession) QuoteTokenScale() (*big.Int, error) {
	return _Pool.Contract.QuoteTokenScale(&_Pool.CallOpts)
}

// QuoteTokenScale is a free data retrieval call binding the contract method 0x7b3f8655.
//
// Solidity: function quoteTokenScale() pure returns(uint256)
func (_Pool *PoolCallerSession) QuoteTokenScale() (*big.Int, error) {
	return _Pool.Contract.QuoteTokenScale(&_Pool.CallOpts)
}

// ReservesInfo is a free data retrieval call binding the contract method 0x5a3b4477.
//
// Solidity: function reservesInfo() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Pool *PoolCaller) ReservesInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "reservesInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// ReservesInfo is a free data retrieval call binding the contract method 0x5a3b4477.
//
// Solidity: function reservesInfo() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Pool *PoolSession) ReservesInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.ReservesInfo(&_Pool.CallOpts)
}

// ReservesInfo is a free data retrieval call binding the contract method 0x5a3b4477.
//
// Solidity: function reservesInfo() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Pool *PoolCallerSession) ReservesInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.ReservesInfo(&_Pool.CallOpts)
}

// TotalAuctionsInPool is a free data retrieval call binding the contract method 0xbcb630d7.
//
// Solidity: function totalAuctionsInPool() view returns(uint256)
func (_Pool *PoolCaller) TotalAuctionsInPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalAuctionsInPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAuctionsInPool is a free data retrieval call binding the contract method 0xbcb630d7.
//
// Solidity: function totalAuctionsInPool() view returns(uint256)
func (_Pool *PoolSession) TotalAuctionsInPool() (*big.Int, error) {
	return _Pool.Contract.TotalAuctionsInPool(&_Pool.CallOpts)
}

// TotalAuctionsInPool is a free data retrieval call binding the contract method 0xbcb630d7.
//
// Solidity: function totalAuctionsInPool() view returns(uint256)
func (_Pool *PoolCallerSession) TotalAuctionsInPool() (*big.Int, error) {
	return _Pool.Contract.TotalAuctionsInPool(&_Pool.CallOpts)
}

// TotalT0Debt is a free data retrieval call binding the contract method 0x5d3637e7.
//
// Solidity: function totalT0Debt() view returns(uint256)
func (_Pool *PoolCaller) TotalT0Debt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalT0Debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalT0Debt is a free data retrieval call binding the contract method 0x5d3637e7.
//
// Solidity: function totalT0Debt() view returns(uint256)
func (_Pool *PoolSession) TotalT0Debt() (*big.Int, error) {
	return _Pool.Contract.TotalT0Debt(&_Pool.CallOpts)
}

// TotalT0Debt is a free data retrieval call binding the contract method 0x5d3637e7.
//
// Solidity: function totalT0Debt() view returns(uint256)
func (_Pool *PoolCallerSession) TotalT0Debt() (*big.Int, error) {
	return _Pool.Contract.TotalT0Debt(&_Pool.CallOpts)
}

// TotalT0DebtInAuction is a free data retrieval call binding the contract method 0x870c764a.
//
// Solidity: function totalT0DebtInAuction() view returns(uint256)
func (_Pool *PoolCaller) TotalT0DebtInAuction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalT0DebtInAuction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalT0DebtInAuction is a free data retrieval call binding the contract method 0x870c764a.
//
// Solidity: function totalT0DebtInAuction() view returns(uint256)
func (_Pool *PoolSession) TotalT0DebtInAuction() (*big.Int, error) {
	return _Pool.Contract.TotalT0DebtInAuction(&_Pool.CallOpts)
}

// TotalT0DebtInAuction is a free data retrieval call binding the contract method 0x870c764a.
//
// Solidity: function totalT0DebtInAuction() view returns(uint256)
func (_Pool *PoolCallerSession) TotalT0DebtInAuction() (*big.Int, error) {
	return _Pool.Contract.TotalT0DebtInAuction(&_Pool.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc861c6e6.
//
// Solidity: function addCollateral(uint256 amountToAdd_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_)
func (_Pool *PoolTransactor) AddCollateral(opts *bind.TransactOpts, amountToAdd_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addCollateral", amountToAdd_, index_, expiry_)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc861c6e6.
//
// Solidity: function addCollateral(uint256 amountToAdd_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_)
func (_Pool *PoolSession) AddCollateral(amountToAdd_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddCollateral(&_Pool.TransactOpts, amountToAdd_, index_, expiry_)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xc861c6e6.
//
// Solidity: function addCollateral(uint256 amountToAdd_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_)
func (_Pool *PoolTransactorSession) AddCollateral(amountToAdd_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddCollateral(&_Pool.TransactOpts, amountToAdd_, index_, expiry_)
}

// AddQuoteToken is a paid mutator transaction binding the contract method 0xf78b0cce.
//
// Solidity: function addQuoteToken(uint256 amount_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_, uint256 addedAmount_)
func (_Pool *PoolTransactor) AddQuoteToken(opts *bind.TransactOpts, amount_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addQuoteToken", amount_, index_, expiry_)
}

// AddQuoteToken is a paid mutator transaction binding the contract method 0xf78b0cce.
//
// Solidity: function addQuoteToken(uint256 amount_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_, uint256 addedAmount_)
func (_Pool *PoolSession) AddQuoteToken(amount_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddQuoteToken(&_Pool.TransactOpts, amount_, index_, expiry_)
}

// AddQuoteToken is a paid mutator transaction binding the contract method 0xf78b0cce.
//
// Solidity: function addQuoteToken(uint256 amount_, uint256 index_, uint256 expiry_) returns(uint256 bucketLP_, uint256 addedAmount_)
func (_Pool *PoolTransactorSession) AddQuoteToken(amount_ *big.Int, index_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddQuoteToken(&_Pool.TransactOpts, amount_, index_, expiry_)
}

// ApproveLPTransferors is a paid mutator transaction binding the contract method 0x7f8baa37.
//
// Solidity: function approveLPTransferors(address[] transferors_) returns()
func (_Pool *PoolTransactor) ApproveLPTransferors(opts *bind.TransactOpts, transferors_ []common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "approveLPTransferors", transferors_)
}

// ApproveLPTransferors is a paid mutator transaction binding the contract method 0x7f8baa37.
//
// Solidity: function approveLPTransferors(address[] transferors_) returns()
func (_Pool *PoolSession) ApproveLPTransferors(transferors_ []common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ApproveLPTransferors(&_Pool.TransactOpts, transferors_)
}

// ApproveLPTransferors is a paid mutator transaction binding the contract method 0x7f8baa37.
//
// Solidity: function approveLPTransferors(address[] transferors_) returns()
func (_Pool *PoolTransactorSession) ApproveLPTransferors(transferors_ []common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ApproveLPTransferors(&_Pool.TransactOpts, transferors_)
}

// BucketTake is a paid mutator transaction binding the contract method 0x0729f62c.
//
// Solidity: function bucketTake(address borrowerAddress_, bool depositTake_, uint256 index_) returns()
func (_Pool *PoolTransactor) BucketTake(opts *bind.TransactOpts, borrowerAddress_ common.Address, depositTake_ bool, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "bucketTake", borrowerAddress_, depositTake_, index_)
}

// BucketTake is a paid mutator transaction binding the contract method 0x0729f62c.
//
// Solidity: function bucketTake(address borrowerAddress_, bool depositTake_, uint256 index_) returns()
func (_Pool *PoolSession) BucketTake(borrowerAddress_ common.Address, depositTake_ bool, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BucketTake(&_Pool.TransactOpts, borrowerAddress_, depositTake_, index_)
}

// BucketTake is a paid mutator transaction binding the contract method 0x0729f62c.
//
// Solidity: function bucketTake(address borrowerAddress_, bool depositTake_, uint256 index_) returns()
func (_Pool *PoolTransactorSession) BucketTake(borrowerAddress_ common.Address, depositTake_ bool, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BucketTake(&_Pool.TransactOpts, borrowerAddress_, depositTake_, index_)
}

// DecreaseLPAllowance is a paid mutator transaction binding the contract method 0x987165ed.
//
// Solidity: function decreaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_Pool *PoolTransactor) DecreaseLPAllowance(opts *bind.TransactOpts, spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "decreaseLPAllowance", spender_, indexes_, amounts_)
}

// DecreaseLPAllowance is a paid mutator transaction binding the contract method 0x987165ed.
//
// Solidity: function decreaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_Pool *PoolSession) DecreaseLPAllowance(spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DecreaseLPAllowance(&_Pool.TransactOpts, spender_, indexes_, amounts_)
}

// DecreaseLPAllowance is a paid mutator transaction binding the contract method 0x987165ed.
//
// Solidity: function decreaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_Pool *PoolTransactorSession) DecreaseLPAllowance(spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DecreaseLPAllowance(&_Pool.TransactOpts, spender_, indexes_, amounts_)
}

// DrawDebt is a paid mutator transaction binding the contract method 0xcfa8ff03.
//
// Solidity: function drawDebt(address borrowerAddress_, uint256 amountToBorrow_, uint256 limitIndex_, uint256 collateralToPledge_) returns()
func (_Pool *PoolTransactor) DrawDebt(opts *bind.TransactOpts, borrowerAddress_ common.Address, amountToBorrow_ *big.Int, limitIndex_ *big.Int, collateralToPledge_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "drawDebt", borrowerAddress_, amountToBorrow_, limitIndex_, collateralToPledge_)
}

// DrawDebt is a paid mutator transaction binding the contract method 0xcfa8ff03.
//
// Solidity: function drawDebt(address borrowerAddress_, uint256 amountToBorrow_, uint256 limitIndex_, uint256 collateralToPledge_) returns()
func (_Pool *PoolSession) DrawDebt(borrowerAddress_ common.Address, amountToBorrow_ *big.Int, limitIndex_ *big.Int, collateralToPledge_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DrawDebt(&_Pool.TransactOpts, borrowerAddress_, amountToBorrow_, limitIndex_, collateralToPledge_)
}

// DrawDebt is a paid mutator transaction binding the contract method 0xcfa8ff03.
//
// Solidity: function drawDebt(address borrowerAddress_, uint256 amountToBorrow_, uint256 limitIndex_, uint256 collateralToPledge_) returns()
func (_Pool *PoolTransactorSession) DrawDebt(borrowerAddress_ common.Address, amountToBorrow_ *big.Int, limitIndex_ *big.Int, collateralToPledge_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DrawDebt(&_Pool.TransactOpts, borrowerAddress_, amountToBorrow_, limitIndex_, collateralToPledge_)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver_, address token_, uint256 amount_, bytes data_) returns(bool)
func (_Pool *PoolTransactor) FlashLoan(opts *bind.TransactOpts, receiver_ common.Address, token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "flashLoan", receiver_, token_, amount_, data_)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver_, address token_, uint256 amount_, bytes data_) returns(bool)
func (_Pool *PoolSession) FlashLoan(receiver_ common.Address, token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Pool.Contract.FlashLoan(&_Pool.TransactOpts, receiver_, token_, amount_, data_)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver_, address token_, uint256 amount_, bytes data_) returns(bool)
func (_Pool *PoolTransactorSession) FlashLoan(receiver_ common.Address, token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Pool.Contract.FlashLoan(&_Pool.TransactOpts, receiver_, token_, amount_, data_)
}

// IncreaseLPAllowance is a paid mutator transaction binding the contract method 0xa918058d.
//
// Solidity: function increaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_Pool *PoolTransactor) IncreaseLPAllowance(opts *bind.TransactOpts, spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "increaseLPAllowance", spender_, indexes_, amounts_)
}

// IncreaseLPAllowance is a paid mutator transaction binding the contract method 0xa918058d.
//
// Solidity: function increaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_Pool *PoolSession) IncreaseLPAllowance(spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.IncreaseLPAllowance(&_Pool.TransactOpts, spender_, indexes_, amounts_)
}

// IncreaseLPAllowance is a paid mutator transaction binding the contract method 0xa918058d.
//
// Solidity: function increaseLPAllowance(address spender_, uint256[] indexes_, uint256[] amounts_) returns()
func (_Pool *PoolTransactorSession) IncreaseLPAllowance(spender_ common.Address, indexes_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.IncreaseLPAllowance(&_Pool.TransactOpts, spender_, indexes_, amounts_)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 rate_) returns()
func (_Pool *PoolTransactor) Initialize(opts *bind.TransactOpts, rate_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initialize", rate_)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 rate_) returns()
func (_Pool *PoolSession) Initialize(rate_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, rate_)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 rate_) returns()
func (_Pool *PoolTransactorSession) Initialize(rate_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, rate_)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address borrower_, uint256 npLimitIndex_) returns()
func (_Pool *PoolTransactor) Kick(opts *bind.TransactOpts, borrower_ common.Address, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "kick", borrower_, npLimitIndex_)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address borrower_, uint256 npLimitIndex_) returns()
func (_Pool *PoolSession) Kick(borrower_ common.Address, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Kick(&_Pool.TransactOpts, borrower_, npLimitIndex_)
}

// Kick is a paid mutator transaction binding the contract method 0x0203d8fb.
//
// Solidity: function kick(address borrower_, uint256 npLimitIndex_) returns()
func (_Pool *PoolTransactorSession) Kick(borrower_ common.Address, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Kick(&_Pool.TransactOpts, borrower_, npLimitIndex_)
}

// KickReserveAuction is a paid mutator transaction binding the contract method 0x5a422b92.
//
// Solidity: function kickReserveAuction() returns()
func (_Pool *PoolTransactor) KickReserveAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "kickReserveAuction")
}

// KickReserveAuction is a paid mutator transaction binding the contract method 0x5a422b92.
//
// Solidity: function kickReserveAuction() returns()
func (_Pool *PoolSession) KickReserveAuction() (*types.Transaction, error) {
	return _Pool.Contract.KickReserveAuction(&_Pool.TransactOpts)
}

// KickReserveAuction is a paid mutator transaction binding the contract method 0x5a422b92.
//
// Solidity: function kickReserveAuction() returns()
func (_Pool *PoolTransactorSession) KickReserveAuction() (*types.Transaction, error) {
	return _Pool.Contract.KickReserveAuction(&_Pool.TransactOpts)
}

// LenderKick is a paid mutator transaction binding the contract method 0xeca48706.
//
// Solidity: function lenderKick(uint256 index_, uint256 npLimitIndex_) returns()
func (_Pool *PoolTransactor) LenderKick(opts *bind.TransactOpts, index_ *big.Int, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "lenderKick", index_, npLimitIndex_)
}

// LenderKick is a paid mutator transaction binding the contract method 0xeca48706.
//
// Solidity: function lenderKick(uint256 index_, uint256 npLimitIndex_) returns()
func (_Pool *PoolSession) LenderKick(index_ *big.Int, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.LenderKick(&_Pool.TransactOpts, index_, npLimitIndex_)
}

// LenderKick is a paid mutator transaction binding the contract method 0xeca48706.
//
// Solidity: function lenderKick(uint256 index_, uint256 npLimitIndex_) returns()
func (_Pool *PoolTransactorSession) LenderKick(index_ *big.Int, npLimitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.LenderKick(&_Pool.TransactOpts, index_, npLimitIndex_)
}

// MoveQuoteToken is a paid mutator transaction binding the contract method 0x332c0e43.
//
// Solidity: function moveQuoteToken(uint256 maxAmount_, uint256 fromIndex_, uint256 toIndex_, uint256 expiry_) returns(uint256 fromBucketLP_, uint256 toBucketLP_, uint256 movedAmount_)
func (_Pool *PoolTransactor) MoveQuoteToken(opts *bind.TransactOpts, maxAmount_ *big.Int, fromIndex_ *big.Int, toIndex_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "moveQuoteToken", maxAmount_, fromIndex_, toIndex_, expiry_)
}

// MoveQuoteToken is a paid mutator transaction binding the contract method 0x332c0e43.
//
// Solidity: function moveQuoteToken(uint256 maxAmount_, uint256 fromIndex_, uint256 toIndex_, uint256 expiry_) returns(uint256 fromBucketLP_, uint256 toBucketLP_, uint256 movedAmount_)
func (_Pool *PoolSession) MoveQuoteToken(maxAmount_ *big.Int, fromIndex_ *big.Int, toIndex_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.MoveQuoteToken(&_Pool.TransactOpts, maxAmount_, fromIndex_, toIndex_, expiry_)
}

// MoveQuoteToken is a paid mutator transaction binding the contract method 0x332c0e43.
//
// Solidity: function moveQuoteToken(uint256 maxAmount_, uint256 fromIndex_, uint256 toIndex_, uint256 expiry_) returns(uint256 fromBucketLP_, uint256 toBucketLP_, uint256 movedAmount_)
func (_Pool *PoolTransactorSession) MoveQuoteToken(maxAmount_ *big.Int, fromIndex_ *big.Int, toIndex_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.MoveQuoteToken(&_Pool.TransactOpts, maxAmount_, fromIndex_, toIndex_, expiry_)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pool *PoolTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pool *PoolSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Pool.Contract.Multicall(&_Pool.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pool *PoolTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Pool.Contract.Multicall(&_Pool.TransactOpts, data)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x6a9b1891.
//
// Solidity: function removeCollateral(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_Pool *PoolTransactor) RemoveCollateral(opts *bind.TransactOpts, maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeCollateral", maxAmount_, index_)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x6a9b1891.
//
// Solidity: function removeCollateral(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_Pool *PoolSession) RemoveCollateral(maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RemoveCollateral(&_Pool.TransactOpts, maxAmount_, index_)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x6a9b1891.
//
// Solidity: function removeCollateral(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_Pool *PoolTransactorSession) RemoveCollateral(maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RemoveCollateral(&_Pool.TransactOpts, maxAmount_, index_)
}

// RemoveQuoteToken is a paid mutator transaction binding the contract method 0xb1f07247.
//
// Solidity: function removeQuoteToken(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_Pool *PoolTransactor) RemoveQuoteToken(opts *bind.TransactOpts, maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeQuoteToken", maxAmount_, index_)
}

// RemoveQuoteToken is a paid mutator transaction binding the contract method 0xb1f07247.
//
// Solidity: function removeQuoteToken(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_Pool *PoolSession) RemoveQuoteToken(maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RemoveQuoteToken(&_Pool.TransactOpts, maxAmount_, index_)
}

// RemoveQuoteToken is a paid mutator transaction binding the contract method 0xb1f07247.
//
// Solidity: function removeQuoteToken(uint256 maxAmount_, uint256 index_) returns(uint256 removedAmount_, uint256 redeemedLP_)
func (_Pool *PoolTransactorSession) RemoveQuoteToken(maxAmount_ *big.Int, index_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RemoveQuoteToken(&_Pool.TransactOpts, maxAmount_, index_)
}

// RepayDebt is a paid mutator transaction binding the contract method 0xa9ff9f77.
//
// Solidity: function repayDebt(address borrowerAddress_, uint256 maxQuoteTokenAmountToRepay_, uint256 collateralAmountToPull_, address collateralReceiver_, uint256 limitIndex_) returns(uint256 amountRepaid_)
func (_Pool *PoolTransactor) RepayDebt(opts *bind.TransactOpts, borrowerAddress_ common.Address, maxQuoteTokenAmountToRepay_ *big.Int, collateralAmountToPull_ *big.Int, collateralReceiver_ common.Address, limitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "repayDebt", borrowerAddress_, maxQuoteTokenAmountToRepay_, collateralAmountToPull_, collateralReceiver_, limitIndex_)
}

// RepayDebt is a paid mutator transaction binding the contract method 0xa9ff9f77.
//
// Solidity: function repayDebt(address borrowerAddress_, uint256 maxQuoteTokenAmountToRepay_, uint256 collateralAmountToPull_, address collateralReceiver_, uint256 limitIndex_) returns(uint256 amountRepaid_)
func (_Pool *PoolSession) RepayDebt(borrowerAddress_ common.Address, maxQuoteTokenAmountToRepay_ *big.Int, collateralAmountToPull_ *big.Int, collateralReceiver_ common.Address, limitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RepayDebt(&_Pool.TransactOpts, borrowerAddress_, maxQuoteTokenAmountToRepay_, collateralAmountToPull_, collateralReceiver_, limitIndex_)
}

// RepayDebt is a paid mutator transaction binding the contract method 0xa9ff9f77.
//
// Solidity: function repayDebt(address borrowerAddress_, uint256 maxQuoteTokenAmountToRepay_, uint256 collateralAmountToPull_, address collateralReceiver_, uint256 limitIndex_) returns(uint256 amountRepaid_)
func (_Pool *PoolTransactorSession) RepayDebt(borrowerAddress_ common.Address, maxQuoteTokenAmountToRepay_ *big.Int, collateralAmountToPull_ *big.Int, collateralReceiver_ common.Address, limitIndex_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RepayDebt(&_Pool.TransactOpts, borrowerAddress_, maxQuoteTokenAmountToRepay_, collateralAmountToPull_, collateralReceiver_, limitIndex_)
}

// RevokeLPAllowance is a paid mutator transaction binding the contract method 0x06e47f26.
//
// Solidity: function revokeLPAllowance(address spender_, uint256[] indexes_) returns()
func (_Pool *PoolTransactor) RevokeLPAllowance(opts *bind.TransactOpts, spender_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "revokeLPAllowance", spender_, indexes_)
}

// RevokeLPAllowance is a paid mutator transaction binding the contract method 0x06e47f26.
//
// Solidity: function revokeLPAllowance(address spender_, uint256[] indexes_) returns()
func (_Pool *PoolSession) RevokeLPAllowance(spender_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RevokeLPAllowance(&_Pool.TransactOpts, spender_, indexes_)
}

// RevokeLPAllowance is a paid mutator transaction binding the contract method 0x06e47f26.
//
// Solidity: function revokeLPAllowance(address spender_, uint256[] indexes_) returns()
func (_Pool *PoolTransactorSession) RevokeLPAllowance(spender_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RevokeLPAllowance(&_Pool.TransactOpts, spender_, indexes_)
}

// RevokeLPTransferors is a paid mutator transaction binding the contract method 0xd39d813f.
//
// Solidity: function revokeLPTransferors(address[] transferors_) returns()
func (_Pool *PoolTransactor) RevokeLPTransferors(opts *bind.TransactOpts, transferors_ []common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "revokeLPTransferors", transferors_)
}

// RevokeLPTransferors is a paid mutator transaction binding the contract method 0xd39d813f.
//
// Solidity: function revokeLPTransferors(address[] transferors_) returns()
func (_Pool *PoolSession) RevokeLPTransferors(transferors_ []common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RevokeLPTransferors(&_Pool.TransactOpts, transferors_)
}

// RevokeLPTransferors is a paid mutator transaction binding the contract method 0xd39d813f.
//
// Solidity: function revokeLPTransferors(address[] transferors_) returns()
func (_Pool *PoolTransactorSession) RevokeLPTransferors(transferors_ []common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RevokeLPTransferors(&_Pool.TransactOpts, transferors_)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address borrowerAddress_, uint256 maxDepth_) returns(uint256 collateralSettled_, bool isBorrowerSettled_)
func (_Pool *PoolTransactor) Settle(opts *bind.TransactOpts, borrowerAddress_ common.Address, maxDepth_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "settle", borrowerAddress_, maxDepth_)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address borrowerAddress_, uint256 maxDepth_) returns(uint256 collateralSettled_, bool isBorrowerSettled_)
func (_Pool *PoolSession) Settle(borrowerAddress_ common.Address, maxDepth_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Settle(&_Pool.TransactOpts, borrowerAddress_, maxDepth_)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address borrowerAddress_, uint256 maxDepth_) returns(uint256 collateralSettled_, bool isBorrowerSettled_)
func (_Pool *PoolTransactorSession) Settle(borrowerAddress_ common.Address, maxDepth_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Settle(&_Pool.TransactOpts, borrowerAddress_, maxDepth_)
}

// StampLoan is a paid mutator transaction binding the contract method 0xce4396d7.
//
// Solidity: function stampLoan() returns()
func (_Pool *PoolTransactor) StampLoan(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "stampLoan")
}

// StampLoan is a paid mutator transaction binding the contract method 0xce4396d7.
//
// Solidity: function stampLoan() returns()
func (_Pool *PoolSession) StampLoan() (*types.Transaction, error) {
	return _Pool.Contract.StampLoan(&_Pool.TransactOpts)
}

// StampLoan is a paid mutator transaction binding the contract method 0xce4396d7.
//
// Solidity: function stampLoan() returns()
func (_Pool *PoolTransactorSession) StampLoan() (*types.Transaction, error) {
	return _Pool.Contract.StampLoan(&_Pool.TransactOpts)
}

// Take is a paid mutator transaction binding the contract method 0x66ae5880.
//
// Solidity: function take(address borrowerAddress_, uint256 maxAmount_, address callee_, bytes data_) returns(uint256 collateralTaken_)
func (_Pool *PoolTransactor) Take(opts *bind.TransactOpts, borrowerAddress_ common.Address, maxAmount_ *big.Int, callee_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "take", borrowerAddress_, maxAmount_, callee_, data_)
}

// Take is a paid mutator transaction binding the contract method 0x66ae5880.
//
// Solidity: function take(address borrowerAddress_, uint256 maxAmount_, address callee_, bytes data_) returns(uint256 collateralTaken_)
func (_Pool *PoolSession) Take(borrowerAddress_ common.Address, maxAmount_ *big.Int, callee_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Pool.Contract.Take(&_Pool.TransactOpts, borrowerAddress_, maxAmount_, callee_, data_)
}

// Take is a paid mutator transaction binding the contract method 0x66ae5880.
//
// Solidity: function take(address borrowerAddress_, uint256 maxAmount_, address callee_, bytes data_) returns(uint256 collateralTaken_)
func (_Pool *PoolTransactorSession) Take(borrowerAddress_ common.Address, maxAmount_ *big.Int, callee_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Pool.Contract.Take(&_Pool.TransactOpts, borrowerAddress_, maxAmount_, callee_, data_)
}

// TakeReserves is a paid mutator transaction binding the contract method 0x42302a9a.
//
// Solidity: function takeReserves(uint256 maxAmount_) returns(uint256 amount_)
func (_Pool *PoolTransactor) TakeReserves(opts *bind.TransactOpts, maxAmount_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "takeReserves", maxAmount_)
}

// TakeReserves is a paid mutator transaction binding the contract method 0x42302a9a.
//
// Solidity: function takeReserves(uint256 maxAmount_) returns(uint256 amount_)
func (_Pool *PoolSession) TakeReserves(maxAmount_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TakeReserves(&_Pool.TransactOpts, maxAmount_)
}

// TakeReserves is a paid mutator transaction binding the contract method 0x42302a9a.
//
// Solidity: function takeReserves(uint256 maxAmount_) returns(uint256 amount_)
func (_Pool *PoolTransactorSession) TakeReserves(maxAmount_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TakeReserves(&_Pool.TransactOpts, maxAmount_)
}

// TransferLP is a paid mutator transaction binding the contract method 0x4efe8af7.
//
// Solidity: function transferLP(address owner_, address newOwner_, uint256[] indexes_) returns()
func (_Pool *PoolTransactor) TransferLP(opts *bind.TransactOpts, owner_ common.Address, newOwner_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferLP", owner_, newOwner_, indexes_)
}

// TransferLP is a paid mutator transaction binding the contract method 0x4efe8af7.
//
// Solidity: function transferLP(address owner_, address newOwner_, uint256[] indexes_) returns()
func (_Pool *PoolSession) TransferLP(owner_ common.Address, newOwner_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferLP(&_Pool.TransactOpts, owner_, newOwner_, indexes_)
}

// TransferLP is a paid mutator transaction binding the contract method 0x4efe8af7.
//
// Solidity: function transferLP(address owner_, address newOwner_, uint256[] indexes_) returns()
func (_Pool *PoolTransactorSession) TransferLP(owner_ common.Address, newOwner_ common.Address, indexes_ []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferLP(&_Pool.TransactOpts, owner_, newOwner_, indexes_)
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_Pool *PoolTransactor) UpdateInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "updateInterest")
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_Pool *PoolSession) UpdateInterest() (*types.Transaction, error) {
	return _Pool.Contract.UpdateInterest(&_Pool.TransactOpts)
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_Pool *PoolTransactorSession) UpdateInterest() (*types.Transaction, error) {
	return _Pool.Contract.UpdateInterest(&_Pool.TransactOpts)
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0xd53e2b1b.
//
// Solidity: function withdrawBonds(address recipient_, uint256 maxAmount_) returns(uint256 withdrawnAmount_)
func (_Pool *PoolTransactor) WithdrawBonds(opts *bind.TransactOpts, recipient_ common.Address, maxAmount_ *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "withdrawBonds", recipient_, maxAmount_)
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0xd53e2b1b.
//
// Solidity: function withdrawBonds(address recipient_, uint256 maxAmount_) returns(uint256 withdrawnAmount_)
func (_Pool *PoolSession) WithdrawBonds(recipient_ common.Address, maxAmount_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.WithdrawBonds(&_Pool.TransactOpts, recipient_, maxAmount_)
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0xd53e2b1b.
//
// Solidity: function withdrawBonds(address recipient_, uint256 maxAmount_) returns(uint256 withdrawnAmount_)
func (_Pool *PoolTransactorSession) WithdrawBonds(recipient_ common.Address, maxAmount_ *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.WithdrawBonds(&_Pool.TransactOpts, recipient_, maxAmount_)
}

// PoolAddCollateralIterator is returned from FilterAddCollateral and is used to iterate over the raw logs and unpacked data for AddCollateral events raised by the Pool contract.
type PoolAddCollateralIterator struct {
	Event *PoolAddCollateral // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolAddCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddCollateral)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolAddCollateral)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolAddCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddCollateral represents a AddCollateral event raised by the Pool contract.
type PoolAddCollateral struct {
	Actor     common.Address
	Index     *big.Int
	Amount    *big.Int
	LpAwarded *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddCollateral is a free log retrieval operation binding the contract event 0xa9387d09ded47dbc173eb751964c0c7b7e0a1165939b958fafc8109337597f94.
//
// Solidity: event AddCollateral(address indexed actor, uint256 indexed index, uint256 amount, uint256 lpAwarded)
func (_Pool *PoolFilterer) FilterAddCollateral(opts *bind.FilterOpts, actor []common.Address, index []*big.Int) (*PoolAddCollateralIterator, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AddCollateral", actorRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddCollateralIterator{contract: _Pool.contract, event: "AddCollateral", logs: logs, sub: sub}, nil
}

// WatchAddCollateral is a free log subscription operation binding the contract event 0xa9387d09ded47dbc173eb751964c0c7b7e0a1165939b958fafc8109337597f94.
//
// Solidity: event AddCollateral(address indexed actor, uint256 indexed index, uint256 amount, uint256 lpAwarded)
func (_Pool *PoolFilterer) WatchAddCollateral(opts *bind.WatchOpts, sink chan<- *PoolAddCollateral, actor []common.Address, index []*big.Int) (event.Subscription, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AddCollateral", actorRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddCollateral)
				if err := _Pool.contract.UnpackLog(event, "AddCollateral", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddCollateral is a log parse operation binding the contract event 0xa9387d09ded47dbc173eb751964c0c7b7e0a1165939b958fafc8109337597f94.
//
// Solidity: event AddCollateral(address indexed actor, uint256 indexed index, uint256 amount, uint256 lpAwarded)
func (_Pool *PoolFilterer) ParseAddCollateral(log types.Log) (*PoolAddCollateral, error) {
	event := new(PoolAddCollateral)
	if err := _Pool.contract.UnpackLog(event, "AddCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddQuoteTokenIterator is returned from FilterAddQuoteToken and is used to iterate over the raw logs and unpacked data for AddQuoteToken events raised by the Pool contract.
type PoolAddQuoteTokenIterator struct {
	Event *PoolAddQuoteToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolAddQuoteTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddQuoteToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolAddQuoteToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolAddQuoteTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddQuoteTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddQuoteToken represents a AddQuoteToken event raised by the Pool contract.
type PoolAddQuoteToken struct {
	Lender    common.Address
	Index     *big.Int
	Amount    *big.Int
	LpAwarded *big.Int
	Lup       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddQuoteToken is a free log retrieval operation binding the contract event 0x8b24a9808cf05d3d8e48ac09e4f649054994a88cfa657b3f4bf340b62137df1e.
//
// Solidity: event AddQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpAwarded, uint256 lup)
func (_Pool *PoolFilterer) FilterAddQuoteToken(opts *bind.FilterOpts, lender []common.Address, index []*big.Int) (*PoolAddQuoteTokenIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AddQuoteToken", lenderRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddQuoteTokenIterator{contract: _Pool.contract, event: "AddQuoteToken", logs: logs, sub: sub}, nil
}

// WatchAddQuoteToken is a free log subscription operation binding the contract event 0x8b24a9808cf05d3d8e48ac09e4f649054994a88cfa657b3f4bf340b62137df1e.
//
// Solidity: event AddQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpAwarded, uint256 lup)
func (_Pool *PoolFilterer) WatchAddQuoteToken(opts *bind.WatchOpts, sink chan<- *PoolAddQuoteToken, lender []common.Address, index []*big.Int) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AddQuoteToken", lenderRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddQuoteToken)
				if err := _Pool.contract.UnpackLog(event, "AddQuoteToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddQuoteToken is a log parse operation binding the contract event 0x8b24a9808cf05d3d8e48ac09e4f649054994a88cfa657b3f4bf340b62137df1e.
//
// Solidity: event AddQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpAwarded, uint256 lup)
func (_Pool *PoolFilterer) ParseAddQuoteToken(log types.Log) (*PoolAddQuoteToken, error) {
	event := new(PoolAddQuoteToken)
	if err := _Pool.contract.UnpackLog(event, "AddQuoteToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolApproveLPTransferorsIterator is returned from FilterApproveLPTransferors and is used to iterate over the raw logs and unpacked data for ApproveLPTransferors events raised by the Pool contract.
type PoolApproveLPTransferorsIterator struct {
	Event *PoolApproveLPTransferors // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolApproveLPTransferorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolApproveLPTransferors)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolApproveLPTransferors)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolApproveLPTransferorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolApproveLPTransferorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolApproveLPTransferors represents a ApproveLPTransferors event raised by the Pool contract.
type PoolApproveLPTransferors struct {
	Lender      common.Address
	Transferors []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterApproveLPTransferors is a free log retrieval operation binding the contract event 0x6dceda33b951e30872202c47c5e3b449112437927dbc475dfaedd3a22889aa54.
//
// Solidity: event ApproveLPTransferors(address indexed lender, address[] transferors)
func (_Pool *PoolFilterer) FilterApproveLPTransferors(opts *bind.FilterOpts, lender []common.Address) (*PoolApproveLPTransferorsIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ApproveLPTransferors", lenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolApproveLPTransferorsIterator{contract: _Pool.contract, event: "ApproveLPTransferors", logs: logs, sub: sub}, nil
}

// WatchApproveLPTransferors is a free log subscription operation binding the contract event 0x6dceda33b951e30872202c47c5e3b449112437927dbc475dfaedd3a22889aa54.
//
// Solidity: event ApproveLPTransferors(address indexed lender, address[] transferors)
func (_Pool *PoolFilterer) WatchApproveLPTransferors(opts *bind.WatchOpts, sink chan<- *PoolApproveLPTransferors, lender []common.Address) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ApproveLPTransferors", lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolApproveLPTransferors)
				if err := _Pool.contract.UnpackLog(event, "ApproveLPTransferors", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproveLPTransferors is a log parse operation binding the contract event 0x6dceda33b951e30872202c47c5e3b449112437927dbc475dfaedd3a22889aa54.
//
// Solidity: event ApproveLPTransferors(address indexed lender, address[] transferors)
func (_Pool *PoolFilterer) ParseApproveLPTransferors(log types.Log) (*PoolApproveLPTransferors, error) {
	event := new(PoolApproveLPTransferors)
	if err := _Pool.contract.UnpackLog(event, "ApproveLPTransferors", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAuctionNFTSettleIterator is returned from FilterAuctionNFTSettle and is used to iterate over the raw logs and unpacked data for AuctionNFTSettle events raised by the Pool contract.
type PoolAuctionNFTSettleIterator struct {
	Event *PoolAuctionNFTSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolAuctionNFTSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAuctionNFTSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolAuctionNFTSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolAuctionNFTSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAuctionNFTSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAuctionNFTSettle represents a AuctionNFTSettle event raised by the Pool contract.
type PoolAuctionNFTSettle struct {
	Borrower   common.Address
	Collateral *big.Int
	Lp         *big.Int
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionNFTSettle is a free log retrieval operation binding the contract event 0xddd6b496b84171d179d9874158b1cbbe422dd482e5523f1b09cb69ebef287841.
//
// Solidity: event AuctionNFTSettle(address indexed borrower, uint256 collateral, uint256 lp, uint256 index)
func (_Pool *PoolFilterer) FilterAuctionNFTSettle(opts *bind.FilterOpts, borrower []common.Address) (*PoolAuctionNFTSettleIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AuctionNFTSettle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolAuctionNFTSettleIterator{contract: _Pool.contract, event: "AuctionNFTSettle", logs: logs, sub: sub}, nil
}

// WatchAuctionNFTSettle is a free log subscription operation binding the contract event 0xddd6b496b84171d179d9874158b1cbbe422dd482e5523f1b09cb69ebef287841.
//
// Solidity: event AuctionNFTSettle(address indexed borrower, uint256 collateral, uint256 lp, uint256 index)
func (_Pool *PoolFilterer) WatchAuctionNFTSettle(opts *bind.WatchOpts, sink chan<- *PoolAuctionNFTSettle, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AuctionNFTSettle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAuctionNFTSettle)
				if err := _Pool.contract.UnpackLog(event, "AuctionNFTSettle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionNFTSettle is a log parse operation binding the contract event 0xddd6b496b84171d179d9874158b1cbbe422dd482e5523f1b09cb69ebef287841.
//
// Solidity: event AuctionNFTSettle(address indexed borrower, uint256 collateral, uint256 lp, uint256 index)
func (_Pool *PoolFilterer) ParseAuctionNFTSettle(log types.Log) (*PoolAuctionNFTSettle, error) {
	event := new(PoolAuctionNFTSettle)
	if err := _Pool.contract.UnpackLog(event, "AuctionNFTSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAuctionSettleIterator is returned from FilterAuctionSettle and is used to iterate over the raw logs and unpacked data for AuctionSettle events raised by the Pool contract.
type PoolAuctionSettleIterator struct {
	Event *PoolAuctionSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolAuctionSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAuctionSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolAuctionSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolAuctionSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAuctionSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAuctionSettle represents a AuctionSettle event raised by the Pool contract.
type PoolAuctionSettle struct {
	Borrower   common.Address
	Collateral *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettle is a free log retrieval operation binding the contract event 0x91a9dcdd01df8b934f14307641e884e0ea6e414bf05fe8daf8c74a28f69b55ee.
//
// Solidity: event AuctionSettle(address indexed borrower, uint256 collateral)
func (_Pool *PoolFilterer) FilterAuctionSettle(opts *bind.FilterOpts, borrower []common.Address) (*PoolAuctionSettleIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AuctionSettle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolAuctionSettleIterator{contract: _Pool.contract, event: "AuctionSettle", logs: logs, sub: sub}, nil
}

// WatchAuctionSettle is a free log subscription operation binding the contract event 0x91a9dcdd01df8b934f14307641e884e0ea6e414bf05fe8daf8c74a28f69b55ee.
//
// Solidity: event AuctionSettle(address indexed borrower, uint256 collateral)
func (_Pool *PoolFilterer) WatchAuctionSettle(opts *bind.WatchOpts, sink chan<- *PoolAuctionSettle, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AuctionSettle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAuctionSettle)
				if err := _Pool.contract.UnpackLog(event, "AuctionSettle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionSettle is a log parse operation binding the contract event 0x91a9dcdd01df8b934f14307641e884e0ea6e414bf05fe8daf8c74a28f69b55ee.
//
// Solidity: event AuctionSettle(address indexed borrower, uint256 collateral)
func (_Pool *PoolFilterer) ParseAuctionSettle(log types.Log) (*PoolAuctionSettle, error) {
	event := new(PoolAuctionSettle)
	if err := _Pool.contract.UnpackLog(event, "AuctionSettle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBondWithdrawnIterator is returned from FilterBondWithdrawn and is used to iterate over the raw logs and unpacked data for BondWithdrawn events raised by the Pool contract.
type PoolBondWithdrawnIterator struct {
	Event *PoolBondWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolBondWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBondWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolBondWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolBondWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBondWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBondWithdrawn represents a BondWithdrawn event raised by the Pool contract.
type PoolBondWithdrawn struct {
	Kicker   common.Address
	Reciever common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBondWithdrawn is a free log retrieval operation binding the contract event 0x1b6622b92ce16ed648b5b93fe47df1cd4c763fdcafe3281bc1dfd5ff7998a94d.
//
// Solidity: event BondWithdrawn(address indexed kicker, address indexed reciever, uint256 amount)
func (_Pool *PoolFilterer) FilterBondWithdrawn(opts *bind.FilterOpts, kicker []common.Address, reciever []common.Address) (*PoolBondWithdrawnIterator, error) {

	var kickerRule []interface{}
	for _, kickerItem := range kicker {
		kickerRule = append(kickerRule, kickerItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "BondWithdrawn", kickerRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return &PoolBondWithdrawnIterator{contract: _Pool.contract, event: "BondWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBondWithdrawn is a free log subscription operation binding the contract event 0x1b6622b92ce16ed648b5b93fe47df1cd4c763fdcafe3281bc1dfd5ff7998a94d.
//
// Solidity: event BondWithdrawn(address indexed kicker, address indexed reciever, uint256 amount)
func (_Pool *PoolFilterer) WatchBondWithdrawn(opts *bind.WatchOpts, sink chan<- *PoolBondWithdrawn, kicker []common.Address, reciever []common.Address) (event.Subscription, error) {

	var kickerRule []interface{}
	for _, kickerItem := range kicker {
		kickerRule = append(kickerRule, kickerItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "BondWithdrawn", kickerRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBondWithdrawn)
				if err := _Pool.contract.UnpackLog(event, "BondWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBondWithdrawn is a log parse operation binding the contract event 0x1b6622b92ce16ed648b5b93fe47df1cd4c763fdcafe3281bc1dfd5ff7998a94d.
//
// Solidity: event BondWithdrawn(address indexed kicker, address indexed reciever, uint256 amount)
func (_Pool *PoolFilterer) ParseBondWithdrawn(log types.Log) (*PoolBondWithdrawn, error) {
	event := new(PoolBondWithdrawn)
	if err := _Pool.contract.UnpackLog(event, "BondWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBucketBankruptcyIterator is returned from FilterBucketBankruptcy and is used to iterate over the raw logs and unpacked data for BucketBankruptcy events raised by the Pool contract.
type PoolBucketBankruptcyIterator struct {
	Event *PoolBucketBankruptcy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolBucketBankruptcyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBucketBankruptcy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolBucketBankruptcy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolBucketBankruptcyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBucketBankruptcyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBucketBankruptcy represents a BucketBankruptcy event raised by the Pool contract.
type PoolBucketBankruptcy struct {
	Index       *big.Int
	LpForfeited *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBucketBankruptcy is a free log retrieval operation binding the contract event 0x30ee43613aaa48d222b158aab1123c5a29d452f8b1a849e5f815dd355923ba85.
//
// Solidity: event BucketBankruptcy(uint256 indexed index, uint256 lpForfeited)
func (_Pool *PoolFilterer) FilterBucketBankruptcy(opts *bind.FilterOpts, index []*big.Int) (*PoolBucketBankruptcyIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "BucketBankruptcy", indexRule)
	if err != nil {
		return nil, err
	}
	return &PoolBucketBankruptcyIterator{contract: _Pool.contract, event: "BucketBankruptcy", logs: logs, sub: sub}, nil
}

// WatchBucketBankruptcy is a free log subscription operation binding the contract event 0x30ee43613aaa48d222b158aab1123c5a29d452f8b1a849e5f815dd355923ba85.
//
// Solidity: event BucketBankruptcy(uint256 indexed index, uint256 lpForfeited)
func (_Pool *PoolFilterer) WatchBucketBankruptcy(opts *bind.WatchOpts, sink chan<- *PoolBucketBankruptcy, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "BucketBankruptcy", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBucketBankruptcy)
				if err := _Pool.contract.UnpackLog(event, "BucketBankruptcy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBucketBankruptcy is a log parse operation binding the contract event 0x30ee43613aaa48d222b158aab1123c5a29d452f8b1a849e5f815dd355923ba85.
//
// Solidity: event BucketBankruptcy(uint256 indexed index, uint256 lpForfeited)
func (_Pool *PoolFilterer) ParseBucketBankruptcy(log types.Log) (*PoolBucketBankruptcy, error) {
	event := new(PoolBucketBankruptcy)
	if err := _Pool.contract.UnpackLog(event, "BucketBankruptcy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBucketTakeIterator is returned from FilterBucketTake and is used to iterate over the raw logs and unpacked data for BucketTake events raised by the Pool contract.
type PoolBucketTakeIterator struct {
	Event *PoolBucketTake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolBucketTakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBucketTake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolBucketTake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolBucketTakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBucketTakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBucketTake represents a BucketTake event raised by the Pool contract.
type PoolBucketTake struct {
	Borrower   common.Address
	Index      *big.Int
	Amount     *big.Int
	Collateral *big.Int
	BondChange *big.Int
	IsReward   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBucketTake is a free log retrieval operation binding the contract event 0xcb6905a59200cc485bfe6d2168392e96a0f204daf9e3937dff19180cb0d796a4.
//
// Solidity: event BucketTake(address indexed borrower, uint256 index, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_Pool *PoolFilterer) FilterBucketTake(opts *bind.FilterOpts, borrower []common.Address) (*PoolBucketTakeIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "BucketTake", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolBucketTakeIterator{contract: _Pool.contract, event: "BucketTake", logs: logs, sub: sub}, nil
}

// WatchBucketTake is a free log subscription operation binding the contract event 0xcb6905a59200cc485bfe6d2168392e96a0f204daf9e3937dff19180cb0d796a4.
//
// Solidity: event BucketTake(address indexed borrower, uint256 index, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_Pool *PoolFilterer) WatchBucketTake(opts *bind.WatchOpts, sink chan<- *PoolBucketTake, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "BucketTake", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBucketTake)
				if err := _Pool.contract.UnpackLog(event, "BucketTake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBucketTake is a log parse operation binding the contract event 0xcb6905a59200cc485bfe6d2168392e96a0f204daf9e3937dff19180cb0d796a4.
//
// Solidity: event BucketTake(address indexed borrower, uint256 index, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_Pool *PoolFilterer) ParseBucketTake(log types.Log) (*PoolBucketTake, error) {
	event := new(PoolBucketTake)
	if err := _Pool.contract.UnpackLog(event, "BucketTake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBucketTakeLPAwardedIterator is returned from FilterBucketTakeLPAwarded and is used to iterate over the raw logs and unpacked data for BucketTakeLPAwarded events raised by the Pool contract.
type PoolBucketTakeLPAwardedIterator struct {
	Event *PoolBucketTakeLPAwarded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolBucketTakeLPAwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBucketTakeLPAwarded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolBucketTakeLPAwarded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolBucketTakeLPAwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBucketTakeLPAwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBucketTakeLPAwarded represents a BucketTakeLPAwarded event raised by the Pool contract.
type PoolBucketTakeLPAwarded struct {
	Taker           common.Address
	Kicker          common.Address
	LpAwardedTaker  *big.Int
	LpAwardedKicker *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBucketTakeLPAwarded is a free log retrieval operation binding the contract event 0xd43d7a2b808648e0814e6795ea3a9b723df6eae4046a7e279a036458f38bc644.
//
// Solidity: event BucketTakeLPAwarded(address indexed taker, address indexed kicker, uint256 lpAwardedTaker, uint256 lpAwardedKicker)
func (_Pool *PoolFilterer) FilterBucketTakeLPAwarded(opts *bind.FilterOpts, taker []common.Address, kicker []common.Address) (*PoolBucketTakeLPAwardedIterator, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var kickerRule []interface{}
	for _, kickerItem := range kicker {
		kickerRule = append(kickerRule, kickerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "BucketTakeLPAwarded", takerRule, kickerRule)
	if err != nil {
		return nil, err
	}
	return &PoolBucketTakeLPAwardedIterator{contract: _Pool.contract, event: "BucketTakeLPAwarded", logs: logs, sub: sub}, nil
}

// WatchBucketTakeLPAwarded is a free log subscription operation binding the contract event 0xd43d7a2b808648e0814e6795ea3a9b723df6eae4046a7e279a036458f38bc644.
//
// Solidity: event BucketTakeLPAwarded(address indexed taker, address indexed kicker, uint256 lpAwardedTaker, uint256 lpAwardedKicker)
func (_Pool *PoolFilterer) WatchBucketTakeLPAwarded(opts *bind.WatchOpts, sink chan<- *PoolBucketTakeLPAwarded, taker []common.Address, kicker []common.Address) (event.Subscription, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var kickerRule []interface{}
	for _, kickerItem := range kicker {
		kickerRule = append(kickerRule, kickerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "BucketTakeLPAwarded", takerRule, kickerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBucketTakeLPAwarded)
				if err := _Pool.contract.UnpackLog(event, "BucketTakeLPAwarded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBucketTakeLPAwarded is a log parse operation binding the contract event 0xd43d7a2b808648e0814e6795ea3a9b723df6eae4046a7e279a036458f38bc644.
//
// Solidity: event BucketTakeLPAwarded(address indexed taker, address indexed kicker, uint256 lpAwardedTaker, uint256 lpAwardedKicker)
func (_Pool *PoolFilterer) ParseBucketTakeLPAwarded(log types.Log) (*PoolBucketTakeLPAwarded, error) {
	event := new(PoolBucketTakeLPAwarded)
	if err := _Pool.contract.UnpackLog(event, "BucketTakeLPAwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDecreaseLPAllowanceIterator is returned from FilterDecreaseLPAllowance and is used to iterate over the raw logs and unpacked data for DecreaseLPAllowance events raised by the Pool contract.
type PoolDecreaseLPAllowanceIterator struct {
	Event *PoolDecreaseLPAllowance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolDecreaseLPAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDecreaseLPAllowance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolDecreaseLPAllowance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolDecreaseLPAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDecreaseLPAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDecreaseLPAllowance represents a DecreaseLPAllowance event raised by the Pool contract.
type PoolDecreaseLPAllowance struct {
	Owner   common.Address
	Spender common.Address
	Indexes []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDecreaseLPAllowance is a free log retrieval operation binding the contract event 0x4a7a52e2fe7e10addaa7f875ecf9ec17563ec12be96c885457061cfc04e05660.
//
// Solidity: event DecreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_Pool *PoolFilterer) FilterDecreaseLPAllowance(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoolDecreaseLPAllowanceIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DecreaseLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolDecreaseLPAllowanceIterator{contract: _Pool.contract, event: "DecreaseLPAllowance", logs: logs, sub: sub}, nil
}

// WatchDecreaseLPAllowance is a free log subscription operation binding the contract event 0x4a7a52e2fe7e10addaa7f875ecf9ec17563ec12be96c885457061cfc04e05660.
//
// Solidity: event DecreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_Pool *PoolFilterer) WatchDecreaseLPAllowance(opts *bind.WatchOpts, sink chan<- *PoolDecreaseLPAllowance, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DecreaseLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDecreaseLPAllowance)
				if err := _Pool.contract.UnpackLog(event, "DecreaseLPAllowance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreaseLPAllowance is a log parse operation binding the contract event 0x4a7a52e2fe7e10addaa7f875ecf9ec17563ec12be96c885457061cfc04e05660.
//
// Solidity: event DecreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_Pool *PoolFilterer) ParseDecreaseLPAllowance(log types.Log) (*PoolDecreaseLPAllowance, error) {
	event := new(PoolDecreaseLPAllowance)
	if err := _Pool.contract.UnpackLog(event, "DecreaseLPAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDrawDebtIterator is returned from FilterDrawDebt and is used to iterate over the raw logs and unpacked data for DrawDebt events raised by the Pool contract.
type PoolDrawDebtIterator struct {
	Event *PoolDrawDebt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolDrawDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDrawDebt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolDrawDebt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolDrawDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDrawDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDrawDebt represents a DrawDebt event raised by the Pool contract.
type PoolDrawDebt struct {
	Borrower          common.Address
	AmountBorrowed    *big.Int
	CollateralPledged *big.Int
	Lup               *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDrawDebt is a free log retrieval operation binding the contract event 0x49a2aab2f4f7ca5c6ba6d413b46a0a09d91d10188fd94b8e23c3225362d12b50.
//
// Solidity: event DrawDebt(address indexed borrower, uint256 amountBorrowed, uint256 collateralPledged, uint256 lup)
func (_Pool *PoolFilterer) FilterDrawDebt(opts *bind.FilterOpts, borrower []common.Address) (*PoolDrawDebtIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DrawDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolDrawDebtIterator{contract: _Pool.contract, event: "DrawDebt", logs: logs, sub: sub}, nil
}

// WatchDrawDebt is a free log subscription operation binding the contract event 0x49a2aab2f4f7ca5c6ba6d413b46a0a09d91d10188fd94b8e23c3225362d12b50.
//
// Solidity: event DrawDebt(address indexed borrower, uint256 amountBorrowed, uint256 collateralPledged, uint256 lup)
func (_Pool *PoolFilterer) WatchDrawDebt(opts *bind.WatchOpts, sink chan<- *PoolDrawDebt, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DrawDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDrawDebt)
				if err := _Pool.contract.UnpackLog(event, "DrawDebt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDrawDebt is a log parse operation binding the contract event 0x49a2aab2f4f7ca5c6ba6d413b46a0a09d91d10188fd94b8e23c3225362d12b50.
//
// Solidity: event DrawDebt(address indexed borrower, uint256 amountBorrowed, uint256 collateralPledged, uint256 lup)
func (_Pool *PoolFilterer) ParseDrawDebt(log types.Log) (*PoolDrawDebt, error) {
	event := new(PoolDrawDebt)
	if err := _Pool.contract.UnpackLog(event, "DrawDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFlashloanIterator is returned from FilterFlashloan and is used to iterate over the raw logs and unpacked data for Flashloan events raised by the Pool contract.
type PoolFlashloanIterator struct {
	Event *PoolFlashloan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolFlashloanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFlashloan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolFlashloan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolFlashloanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFlashloanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFlashloan represents a Flashloan event raised by the Pool contract.
type PoolFlashloan struct {
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFlashloan is a free log retrieval operation binding the contract event 0x6b15284fe89dbd5c436c2e0b06b1bf72e3a0a8e96d1b4a2abd61dfae2d7849a6.
//
// Solidity: event Flashloan(address indexed receiver, address indexed token, uint256 amount)
func (_Pool *PoolFilterer) FilterFlashloan(opts *bind.FilterOpts, receiver []common.Address, token []common.Address) (*PoolFlashloanIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Flashloan", receiverRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolFlashloanIterator{contract: _Pool.contract, event: "Flashloan", logs: logs, sub: sub}, nil
}

// WatchFlashloan is a free log subscription operation binding the contract event 0x6b15284fe89dbd5c436c2e0b06b1bf72e3a0a8e96d1b4a2abd61dfae2d7849a6.
//
// Solidity: event Flashloan(address indexed receiver, address indexed token, uint256 amount)
func (_Pool *PoolFilterer) WatchFlashloan(opts *bind.WatchOpts, sink chan<- *PoolFlashloan, receiver []common.Address, token []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Flashloan", receiverRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFlashloan)
				if err := _Pool.contract.UnpackLog(event, "Flashloan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashloan is a log parse operation binding the contract event 0x6b15284fe89dbd5c436c2e0b06b1bf72e3a0a8e96d1b4a2abd61dfae2d7849a6.
//
// Solidity: event Flashloan(address indexed receiver, address indexed token, uint256 amount)
func (_Pool *PoolFilterer) ParseFlashloan(log types.Log) (*PoolFlashloan, error) {
	event := new(PoolFlashloan)
	if err := _Pool.contract.UnpackLog(event, "Flashloan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolIncreaseLPAllowanceIterator is returned from FilterIncreaseLPAllowance and is used to iterate over the raw logs and unpacked data for IncreaseLPAllowance events raised by the Pool contract.
type PoolIncreaseLPAllowanceIterator struct {
	Event *PoolIncreaseLPAllowance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolIncreaseLPAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolIncreaseLPAllowance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolIncreaseLPAllowance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolIncreaseLPAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolIncreaseLPAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolIncreaseLPAllowance represents a IncreaseLPAllowance event raised by the Pool contract.
type PoolIncreaseLPAllowance struct {
	Owner   common.Address
	Spender common.Address
	Indexes []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncreaseLPAllowance is a free log retrieval operation binding the contract event 0xa9f6ab83637f87ef702b94c10d09430c40cd3d4642d14fc2a07408bde931545f.
//
// Solidity: event IncreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_Pool *PoolFilterer) FilterIncreaseLPAllowance(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoolIncreaseLPAllowanceIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "IncreaseLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolIncreaseLPAllowanceIterator{contract: _Pool.contract, event: "IncreaseLPAllowance", logs: logs, sub: sub}, nil
}

// WatchIncreaseLPAllowance is a free log subscription operation binding the contract event 0xa9f6ab83637f87ef702b94c10d09430c40cd3d4642d14fc2a07408bde931545f.
//
// Solidity: event IncreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_Pool *PoolFilterer) WatchIncreaseLPAllowance(opts *bind.WatchOpts, sink chan<- *PoolIncreaseLPAllowance, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "IncreaseLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolIncreaseLPAllowance)
				if err := _Pool.contract.UnpackLog(event, "IncreaseLPAllowance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseLPAllowance is a log parse operation binding the contract event 0xa9f6ab83637f87ef702b94c10d09430c40cd3d4642d14fc2a07408bde931545f.
//
// Solidity: event IncreaseLPAllowance(address indexed owner, address indexed spender, uint256[] indexes, uint256[] amounts)
func (_Pool *PoolFilterer) ParseIncreaseLPAllowance(log types.Log) (*PoolIncreaseLPAllowance, error) {
	event := new(PoolIncreaseLPAllowance)
	if err := _Pool.contract.UnpackLog(event, "IncreaseLPAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolInterestUpdateFailureIterator is returned from FilterInterestUpdateFailure and is used to iterate over the raw logs and unpacked data for InterestUpdateFailure events raised by the Pool contract.
type PoolInterestUpdateFailureIterator struct {
	Event *PoolInterestUpdateFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolInterestUpdateFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolInterestUpdateFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolInterestUpdateFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolInterestUpdateFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolInterestUpdateFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolInterestUpdateFailure represents a InterestUpdateFailure event raised by the Pool contract.
type PoolInterestUpdateFailure struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInterestUpdateFailure is a free log retrieval operation binding the contract event 0x84da056cd0ff5380ec35a74f131057a96626a24305fa137c235bdbe1b414a396.
//
// Solidity: event InterestUpdateFailure()
func (_Pool *PoolFilterer) FilterInterestUpdateFailure(opts *bind.FilterOpts) (*PoolInterestUpdateFailureIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "InterestUpdateFailure")
	if err != nil {
		return nil, err
	}
	return &PoolInterestUpdateFailureIterator{contract: _Pool.contract, event: "InterestUpdateFailure", logs: logs, sub: sub}, nil
}

// WatchInterestUpdateFailure is a free log subscription operation binding the contract event 0x84da056cd0ff5380ec35a74f131057a96626a24305fa137c235bdbe1b414a396.
//
// Solidity: event InterestUpdateFailure()
func (_Pool *PoolFilterer) WatchInterestUpdateFailure(opts *bind.WatchOpts, sink chan<- *PoolInterestUpdateFailure) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "InterestUpdateFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolInterestUpdateFailure)
				if err := _Pool.contract.UnpackLog(event, "InterestUpdateFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterestUpdateFailure is a log parse operation binding the contract event 0x84da056cd0ff5380ec35a74f131057a96626a24305fa137c235bdbe1b414a396.
//
// Solidity: event InterestUpdateFailure()
func (_Pool *PoolFilterer) ParseInterestUpdateFailure(log types.Log) (*PoolInterestUpdateFailure, error) {
	event := new(PoolInterestUpdateFailure)
	if err := _Pool.contract.UnpackLog(event, "InterestUpdateFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the Pool contract.
type PoolKickIterator struct {
	Event *PoolKick // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolKick)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolKick)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolKick represents a Kick event raised by the Pool contract.
type PoolKick struct {
	Borrower   common.Address
	Debt       *big.Int
	Collateral *big.Int
	Bond       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0x9f9a32e7f0271518f9b1895d0b1f2f4f73ed305e48b0a3782932094f9d00d948.
//
// Solidity: event Kick(address indexed borrower, uint256 debt, uint256 collateral, uint256 bond)
func (_Pool *PoolFilterer) FilterKick(opts *bind.FilterOpts, borrower []common.Address) (*PoolKickIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Kick", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolKickIterator{contract: _Pool.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0x9f9a32e7f0271518f9b1895d0b1f2f4f73ed305e48b0a3782932094f9d00d948.
//
// Solidity: event Kick(address indexed borrower, uint256 debt, uint256 collateral, uint256 bond)
func (_Pool *PoolFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *PoolKick, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Kick", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolKick)
				if err := _Pool.contract.UnpackLog(event, "Kick", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKick is a log parse operation binding the contract event 0x9f9a32e7f0271518f9b1895d0b1f2f4f73ed305e48b0a3782932094f9d00d948.
//
// Solidity: event Kick(address indexed borrower, uint256 debt, uint256 collateral, uint256 bond)
func (_Pool *PoolFilterer) ParseKick(log types.Log) (*PoolKick, error) {
	event := new(PoolKick)
	if err := _Pool.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolKickReserveAuctionIterator is returned from FilterKickReserveAuction and is used to iterate over the raw logs and unpacked data for KickReserveAuction events raised by the Pool contract.
type PoolKickReserveAuctionIterator struct {
	Event *PoolKickReserveAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolKickReserveAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolKickReserveAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolKickReserveAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolKickReserveAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolKickReserveAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolKickReserveAuction represents a KickReserveAuction event raised by the Pool contract.
type PoolKickReserveAuction struct {
	ClaimableReservesRemaining *big.Int
	AuctionPrice               *big.Int
	CurrentBurnEpoch           *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterKickReserveAuction is a free log retrieval operation binding the contract event 0x3dacf6b19b0a84358a76a3338466cd428d1d4e80e53ccfe91b15d9b8021df960.
//
// Solidity: event KickReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_Pool *PoolFilterer) FilterKickReserveAuction(opts *bind.FilterOpts) (*PoolKickReserveAuctionIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "KickReserveAuction")
	if err != nil {
		return nil, err
	}
	return &PoolKickReserveAuctionIterator{contract: _Pool.contract, event: "KickReserveAuction", logs: logs, sub: sub}, nil
}

// WatchKickReserveAuction is a free log subscription operation binding the contract event 0x3dacf6b19b0a84358a76a3338466cd428d1d4e80e53ccfe91b15d9b8021df960.
//
// Solidity: event KickReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_Pool *PoolFilterer) WatchKickReserveAuction(opts *bind.WatchOpts, sink chan<- *PoolKickReserveAuction) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "KickReserveAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolKickReserveAuction)
				if err := _Pool.contract.UnpackLog(event, "KickReserveAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKickReserveAuction is a log parse operation binding the contract event 0x3dacf6b19b0a84358a76a3338466cd428d1d4e80e53ccfe91b15d9b8021df960.
//
// Solidity: event KickReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_Pool *PoolFilterer) ParseKickReserveAuction(log types.Log) (*PoolKickReserveAuction, error) {
	event := new(PoolKickReserveAuction)
	if err := _Pool.contract.UnpackLog(event, "KickReserveAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolLoanStampedIterator is returned from FilterLoanStamped and is used to iterate over the raw logs and unpacked data for LoanStamped events raised by the Pool contract.
type PoolLoanStampedIterator struct {
	Event *PoolLoanStamped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolLoanStampedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolLoanStamped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolLoanStamped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolLoanStampedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolLoanStampedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolLoanStamped represents a LoanStamped event raised by the Pool contract.
type PoolLoanStamped struct {
	Borrower common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLoanStamped is a free log retrieval operation binding the contract event 0x8d6660b4a409414ebe386e9dd200a5c4e75591f0fc98e1272d7ba207d06d4c34.
//
// Solidity: event LoanStamped(address indexed borrower)
func (_Pool *PoolFilterer) FilterLoanStamped(opts *bind.FilterOpts, borrower []common.Address) (*PoolLoanStampedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "LoanStamped", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolLoanStampedIterator{contract: _Pool.contract, event: "LoanStamped", logs: logs, sub: sub}, nil
}

// WatchLoanStamped is a free log subscription operation binding the contract event 0x8d6660b4a409414ebe386e9dd200a5c4e75591f0fc98e1272d7ba207d06d4c34.
//
// Solidity: event LoanStamped(address indexed borrower)
func (_Pool *PoolFilterer) WatchLoanStamped(opts *bind.WatchOpts, sink chan<- *PoolLoanStamped, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "LoanStamped", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolLoanStamped)
				if err := _Pool.contract.UnpackLog(event, "LoanStamped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLoanStamped is a log parse operation binding the contract event 0x8d6660b4a409414ebe386e9dd200a5c4e75591f0fc98e1272d7ba207d06d4c34.
//
// Solidity: event LoanStamped(address indexed borrower)
func (_Pool *PoolFilterer) ParseLoanStamped(log types.Log) (*PoolLoanStamped, error) {
	event := new(PoolLoanStamped)
	if err := _Pool.contract.UnpackLog(event, "LoanStamped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMoveQuoteTokenIterator is returned from FilterMoveQuoteToken and is used to iterate over the raw logs and unpacked data for MoveQuoteToken events raised by the Pool contract.
type PoolMoveQuoteTokenIterator struct {
	Event *PoolMoveQuoteToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolMoveQuoteTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMoveQuoteToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolMoveQuoteToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolMoveQuoteTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMoveQuoteTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMoveQuoteToken represents a MoveQuoteToken event raised by the Pool contract.
type PoolMoveQuoteToken struct {
	Lender         common.Address
	From           *big.Int
	To             *big.Int
	Amount         *big.Int
	LpRedeemedFrom *big.Int
	LpAwardedTo    *big.Int
	Lup            *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMoveQuoteToken is a free log retrieval operation binding the contract event 0x9d7ab6bb30c003ae7d5b583911db0ada7a9e51b0b4ac7ac1bb5e6896e82e4dbe.
//
// Solidity: event MoveQuoteToken(address indexed lender, uint256 indexed from, uint256 indexed to, uint256 amount, uint256 lpRedeemedFrom, uint256 lpAwardedTo, uint256 lup)
func (_Pool *PoolFilterer) FilterMoveQuoteToken(opts *bind.FilterOpts, lender []common.Address, from []*big.Int, to []*big.Int) (*PoolMoveQuoteTokenIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "MoveQuoteToken", lenderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolMoveQuoteTokenIterator{contract: _Pool.contract, event: "MoveQuoteToken", logs: logs, sub: sub}, nil
}

// WatchMoveQuoteToken is a free log subscription operation binding the contract event 0x9d7ab6bb30c003ae7d5b583911db0ada7a9e51b0b4ac7ac1bb5e6896e82e4dbe.
//
// Solidity: event MoveQuoteToken(address indexed lender, uint256 indexed from, uint256 indexed to, uint256 amount, uint256 lpRedeemedFrom, uint256 lpAwardedTo, uint256 lup)
func (_Pool *PoolFilterer) WatchMoveQuoteToken(opts *bind.WatchOpts, sink chan<- *PoolMoveQuoteToken, lender []common.Address, from []*big.Int, to []*big.Int) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "MoveQuoteToken", lenderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMoveQuoteToken)
				if err := _Pool.contract.UnpackLog(event, "MoveQuoteToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMoveQuoteToken is a log parse operation binding the contract event 0x9d7ab6bb30c003ae7d5b583911db0ada7a9e51b0b4ac7ac1bb5e6896e82e4dbe.
//
// Solidity: event MoveQuoteToken(address indexed lender, uint256 indexed from, uint256 indexed to, uint256 amount, uint256 lpRedeemedFrom, uint256 lpAwardedTo, uint256 lup)
func (_Pool *PoolFilterer) ParseMoveQuoteToken(log types.Log) (*PoolMoveQuoteToken, error) {
	event := new(PoolMoveQuoteToken)
	if err := _Pool.contract.UnpackLog(event, "MoveQuoteToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRemoveCollateralIterator is returned from FilterRemoveCollateral and is used to iterate over the raw logs and unpacked data for RemoveCollateral events raised by the Pool contract.
type PoolRemoveCollateralIterator struct {
	Event *PoolRemoveCollateral // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolRemoveCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRemoveCollateral)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolRemoveCollateral)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolRemoveCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRemoveCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRemoveCollateral represents a RemoveCollateral event raised by the Pool contract.
type PoolRemoveCollateral struct {
	Claimer    common.Address
	Index      *big.Int
	Amount     *big.Int
	LpRedeemed *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveCollateral is a free log retrieval operation binding the contract event 0x90895bc82397742e0cea4685e72279103862a03bee6bbe1d71265c7aeb111527.
//
// Solidity: event RemoveCollateral(address indexed claimer, uint256 indexed index, uint256 amount, uint256 lpRedeemed)
func (_Pool *PoolFilterer) FilterRemoveCollateral(opts *bind.FilterOpts, claimer []common.Address, index []*big.Int) (*PoolRemoveCollateralIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RemoveCollateral", claimerRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &PoolRemoveCollateralIterator{contract: _Pool.contract, event: "RemoveCollateral", logs: logs, sub: sub}, nil
}

// WatchRemoveCollateral is a free log subscription operation binding the contract event 0x90895bc82397742e0cea4685e72279103862a03bee6bbe1d71265c7aeb111527.
//
// Solidity: event RemoveCollateral(address indexed claimer, uint256 indexed index, uint256 amount, uint256 lpRedeemed)
func (_Pool *PoolFilterer) WatchRemoveCollateral(opts *bind.WatchOpts, sink chan<- *PoolRemoveCollateral, claimer []common.Address, index []*big.Int) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RemoveCollateral", claimerRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRemoveCollateral)
				if err := _Pool.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveCollateral is a log parse operation binding the contract event 0x90895bc82397742e0cea4685e72279103862a03bee6bbe1d71265c7aeb111527.
//
// Solidity: event RemoveCollateral(address indexed claimer, uint256 indexed index, uint256 amount, uint256 lpRedeemed)
func (_Pool *PoolFilterer) ParseRemoveCollateral(log types.Log) (*PoolRemoveCollateral, error) {
	event := new(PoolRemoveCollateral)
	if err := _Pool.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRemoveQuoteTokenIterator is returned from FilterRemoveQuoteToken and is used to iterate over the raw logs and unpacked data for RemoveQuoteToken events raised by the Pool contract.
type PoolRemoveQuoteTokenIterator struct {
	Event *PoolRemoveQuoteToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolRemoveQuoteTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRemoveQuoteToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolRemoveQuoteToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolRemoveQuoteTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRemoveQuoteTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRemoveQuoteToken represents a RemoveQuoteToken event raised by the Pool contract.
type PoolRemoveQuoteToken struct {
	Lender     common.Address
	Index      *big.Int
	Amount     *big.Int
	LpRedeemed *big.Int
	Lup        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveQuoteToken is a free log retrieval operation binding the contract event 0x0130a7b525bd6b1e72def1ee0b77f3467028a0e958e30174a0c95eb3860fc221.
//
// Solidity: event RemoveQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpRedeemed, uint256 lup)
func (_Pool *PoolFilterer) FilterRemoveQuoteToken(opts *bind.FilterOpts, lender []common.Address, index []*big.Int) (*PoolRemoveQuoteTokenIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RemoveQuoteToken", lenderRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &PoolRemoveQuoteTokenIterator{contract: _Pool.contract, event: "RemoveQuoteToken", logs: logs, sub: sub}, nil
}

// WatchRemoveQuoteToken is a free log subscription operation binding the contract event 0x0130a7b525bd6b1e72def1ee0b77f3467028a0e958e30174a0c95eb3860fc221.
//
// Solidity: event RemoveQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpRedeemed, uint256 lup)
func (_Pool *PoolFilterer) WatchRemoveQuoteToken(opts *bind.WatchOpts, sink chan<- *PoolRemoveQuoteToken, lender []common.Address, index []*big.Int) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RemoveQuoteToken", lenderRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRemoveQuoteToken)
				if err := _Pool.contract.UnpackLog(event, "RemoveQuoteToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveQuoteToken is a log parse operation binding the contract event 0x0130a7b525bd6b1e72def1ee0b77f3467028a0e958e30174a0c95eb3860fc221.
//
// Solidity: event RemoveQuoteToken(address indexed lender, uint256 indexed index, uint256 amount, uint256 lpRedeemed, uint256 lup)
func (_Pool *PoolFilterer) ParseRemoveQuoteToken(log types.Log) (*PoolRemoveQuoteToken, error) {
	event := new(PoolRemoveQuoteToken)
	if err := _Pool.contract.UnpackLog(event, "RemoveQuoteToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRepayDebtIterator is returned from FilterRepayDebt and is used to iterate over the raw logs and unpacked data for RepayDebt events raised by the Pool contract.
type PoolRepayDebtIterator struct {
	Event *PoolRepayDebt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolRepayDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRepayDebt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolRepayDebt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolRepayDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRepayDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRepayDebt represents a RepayDebt event raised by the Pool contract.
type PoolRepayDebt struct {
	Borrower         common.Address
	QuoteRepaid      *big.Int
	CollateralPulled *big.Int
	Lup              *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRepayDebt is a free log retrieval operation binding the contract event 0xef9d6dc34b1e6893b8746b03ac07fd084909654a5cedab240265a8d1bd584dc2.
//
// Solidity: event RepayDebt(address indexed borrower, uint256 quoteRepaid, uint256 collateralPulled, uint256 lup)
func (_Pool *PoolFilterer) FilterRepayDebt(opts *bind.FilterOpts, borrower []common.Address) (*PoolRepayDebtIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RepayDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolRepayDebtIterator{contract: _Pool.contract, event: "RepayDebt", logs: logs, sub: sub}, nil
}

// WatchRepayDebt is a free log subscription operation binding the contract event 0xef9d6dc34b1e6893b8746b03ac07fd084909654a5cedab240265a8d1bd584dc2.
//
// Solidity: event RepayDebt(address indexed borrower, uint256 quoteRepaid, uint256 collateralPulled, uint256 lup)
func (_Pool *PoolFilterer) WatchRepayDebt(opts *bind.WatchOpts, sink chan<- *PoolRepayDebt, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RepayDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRepayDebt)
				if err := _Pool.contract.UnpackLog(event, "RepayDebt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepayDebt is a log parse operation binding the contract event 0xef9d6dc34b1e6893b8746b03ac07fd084909654a5cedab240265a8d1bd584dc2.
//
// Solidity: event RepayDebt(address indexed borrower, uint256 quoteRepaid, uint256 collateralPulled, uint256 lup)
func (_Pool *PoolFilterer) ParseRepayDebt(log types.Log) (*PoolRepayDebt, error) {
	event := new(PoolRepayDebt)
	if err := _Pool.contract.UnpackLog(event, "RepayDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolReserveAuctionIterator is returned from FilterReserveAuction and is used to iterate over the raw logs and unpacked data for ReserveAuction events raised by the Pool contract.
type PoolReserveAuctionIterator struct {
	Event *PoolReserveAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolReserveAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolReserveAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolReserveAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolReserveAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolReserveAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolReserveAuction represents a ReserveAuction event raised by the Pool contract.
type PoolReserveAuction struct {
	ClaimableReservesRemaining *big.Int
	AuctionPrice               *big.Int
	CurrentBurnEpoch           *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterReserveAuction is a free log retrieval operation binding the contract event 0x2115712930a0e5047c8904a9cc557d2b1ba5a21876e04f59249843ce1a74d699.
//
// Solidity: event ReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_Pool *PoolFilterer) FilterReserveAuction(opts *bind.FilterOpts) (*PoolReserveAuctionIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ReserveAuction")
	if err != nil {
		return nil, err
	}
	return &PoolReserveAuctionIterator{contract: _Pool.contract, event: "ReserveAuction", logs: logs, sub: sub}, nil
}

// WatchReserveAuction is a free log subscription operation binding the contract event 0x2115712930a0e5047c8904a9cc557d2b1ba5a21876e04f59249843ce1a74d699.
//
// Solidity: event ReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_Pool *PoolFilterer) WatchReserveAuction(opts *bind.WatchOpts, sink chan<- *PoolReserveAuction) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ReserveAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolReserveAuction)
				if err := _Pool.contract.UnpackLog(event, "ReserveAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveAuction is a log parse operation binding the contract event 0x2115712930a0e5047c8904a9cc557d2b1ba5a21876e04f59249843ce1a74d699.
//
// Solidity: event ReserveAuction(uint256 claimableReservesRemaining, uint256 auctionPrice, uint256 currentBurnEpoch)
func (_Pool *PoolFilterer) ParseReserveAuction(log types.Log) (*PoolReserveAuction, error) {
	event := new(PoolReserveAuction)
	if err := _Pool.contract.UnpackLog(event, "ReserveAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolResetInterestRateIterator is returned from FilterResetInterestRate and is used to iterate over the raw logs and unpacked data for ResetInterestRate events raised by the Pool contract.
type PoolResetInterestRateIterator struct {
	Event *PoolResetInterestRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolResetInterestRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolResetInterestRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolResetInterestRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolResetInterestRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolResetInterestRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolResetInterestRate represents a ResetInterestRate event raised by the Pool contract.
type PoolResetInterestRate struct {
	OldRate *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterResetInterestRate is a free log retrieval operation binding the contract event 0x20ae1d4a2e8d297f3820670c20fc79531e31643d4b201892680e7df3c4ab1599.
//
// Solidity: event ResetInterestRate(uint256 oldRate, uint256 newRate)
func (_Pool *PoolFilterer) FilterResetInterestRate(opts *bind.FilterOpts) (*PoolResetInterestRateIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ResetInterestRate")
	if err != nil {
		return nil, err
	}
	return &PoolResetInterestRateIterator{contract: _Pool.contract, event: "ResetInterestRate", logs: logs, sub: sub}, nil
}

// WatchResetInterestRate is a free log subscription operation binding the contract event 0x20ae1d4a2e8d297f3820670c20fc79531e31643d4b201892680e7df3c4ab1599.
//
// Solidity: event ResetInterestRate(uint256 oldRate, uint256 newRate)
func (_Pool *PoolFilterer) WatchResetInterestRate(opts *bind.WatchOpts, sink chan<- *PoolResetInterestRate) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ResetInterestRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolResetInterestRate)
				if err := _Pool.contract.UnpackLog(event, "ResetInterestRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseResetInterestRate is a log parse operation binding the contract event 0x20ae1d4a2e8d297f3820670c20fc79531e31643d4b201892680e7df3c4ab1599.
//
// Solidity: event ResetInterestRate(uint256 oldRate, uint256 newRate)
func (_Pool *PoolFilterer) ParseResetInterestRate(log types.Log) (*PoolResetInterestRate, error) {
	event := new(PoolResetInterestRate)
	if err := _Pool.contract.UnpackLog(event, "ResetInterestRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRevokeLPAllowanceIterator is returned from FilterRevokeLPAllowance and is used to iterate over the raw logs and unpacked data for RevokeLPAllowance events raised by the Pool contract.
type PoolRevokeLPAllowanceIterator struct {
	Event *PoolRevokeLPAllowance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolRevokeLPAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRevokeLPAllowance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolRevokeLPAllowance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolRevokeLPAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRevokeLPAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRevokeLPAllowance represents a RevokeLPAllowance event raised by the Pool contract.
type PoolRevokeLPAllowance struct {
	Owner   common.Address
	Spender common.Address
	Indexes []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevokeLPAllowance is a free log retrieval operation binding the contract event 0xdf7f78d931b4e040d4598563bab17506dba0aed1a0515c51fd7dbc2a2382afdf.
//
// Solidity: event RevokeLPAllowance(address indexed owner, address indexed spender, uint256[] indexes)
func (_Pool *PoolFilterer) FilterRevokeLPAllowance(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoolRevokeLPAllowanceIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RevokeLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolRevokeLPAllowanceIterator{contract: _Pool.contract, event: "RevokeLPAllowance", logs: logs, sub: sub}, nil
}

// WatchRevokeLPAllowance is a free log subscription operation binding the contract event 0xdf7f78d931b4e040d4598563bab17506dba0aed1a0515c51fd7dbc2a2382afdf.
//
// Solidity: event RevokeLPAllowance(address indexed owner, address indexed spender, uint256[] indexes)
func (_Pool *PoolFilterer) WatchRevokeLPAllowance(opts *bind.WatchOpts, sink chan<- *PoolRevokeLPAllowance, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RevokeLPAllowance", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRevokeLPAllowance)
				if err := _Pool.contract.UnpackLog(event, "RevokeLPAllowance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokeLPAllowance is a log parse operation binding the contract event 0xdf7f78d931b4e040d4598563bab17506dba0aed1a0515c51fd7dbc2a2382afdf.
//
// Solidity: event RevokeLPAllowance(address indexed owner, address indexed spender, uint256[] indexes)
func (_Pool *PoolFilterer) ParseRevokeLPAllowance(log types.Log) (*PoolRevokeLPAllowance, error) {
	event := new(PoolRevokeLPAllowance)
	if err := _Pool.contract.UnpackLog(event, "RevokeLPAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRevokeLPTransferorsIterator is returned from FilterRevokeLPTransferors and is used to iterate over the raw logs and unpacked data for RevokeLPTransferors events raised by the Pool contract.
type PoolRevokeLPTransferorsIterator struct {
	Event *PoolRevokeLPTransferors // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolRevokeLPTransferorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRevokeLPTransferors)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolRevokeLPTransferors)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolRevokeLPTransferorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRevokeLPTransferorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRevokeLPTransferors represents a RevokeLPTransferors event raised by the Pool contract.
type PoolRevokeLPTransferors struct {
	Lender      common.Address
	Transferors []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokeLPTransferors is a free log retrieval operation binding the contract event 0xde63bcd4c57b5d710a16396f80a2797846720923ff52198e806544ccbb720b2b.
//
// Solidity: event RevokeLPTransferors(address indexed lender, address[] transferors)
func (_Pool *PoolFilterer) FilterRevokeLPTransferors(opts *bind.FilterOpts, lender []common.Address) (*PoolRevokeLPTransferorsIterator, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RevokeLPTransferors", lenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolRevokeLPTransferorsIterator{contract: _Pool.contract, event: "RevokeLPTransferors", logs: logs, sub: sub}, nil
}

// WatchRevokeLPTransferors is a free log subscription operation binding the contract event 0xde63bcd4c57b5d710a16396f80a2797846720923ff52198e806544ccbb720b2b.
//
// Solidity: event RevokeLPTransferors(address indexed lender, address[] transferors)
func (_Pool *PoolFilterer) WatchRevokeLPTransferors(opts *bind.WatchOpts, sink chan<- *PoolRevokeLPTransferors, lender []common.Address) (event.Subscription, error) {

	var lenderRule []interface{}
	for _, lenderItem := range lender {
		lenderRule = append(lenderRule, lenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RevokeLPTransferors", lenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRevokeLPTransferors)
				if err := _Pool.contract.UnpackLog(event, "RevokeLPTransferors", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokeLPTransferors is a log parse operation binding the contract event 0xde63bcd4c57b5d710a16396f80a2797846720923ff52198e806544ccbb720b2b.
//
// Solidity: event RevokeLPTransferors(address indexed lender, address[] transferors)
func (_Pool *PoolFilterer) ParseRevokeLPTransferors(log types.Log) (*PoolRevokeLPTransferors, error) {
	event := new(PoolRevokeLPTransferors)
	if err := _Pool.contract.UnpackLog(event, "RevokeLPTransferors", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the Pool contract.
type PoolSettleIterator struct {
	Event *PoolSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSettle represents a Settle event raised by the Pool contract.
type PoolSettle struct {
	Borrower    common.Address
	SettledDebt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed borrower, uint256 settledDebt)
func (_Pool *PoolFilterer) FilterSettle(opts *bind.FilterOpts, borrower []common.Address) (*PoolSettleIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Settle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolSettleIterator{contract: _Pool.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed borrower, uint256 settledDebt)
func (_Pool *PoolFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *PoolSettle, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Settle", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSettle)
				if err := _Pool.contract.UnpackLog(event, "Settle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettle is a log parse operation binding the contract event 0xa3788eedc10ef3ec6982384227c562c6666cf2b6af4f6a583b6d5d0f2ba0d204.
//
// Solidity: event Settle(address indexed borrower, uint256 settledDebt)
func (_Pool *PoolFilterer) ParseSettle(log types.Log) (*PoolSettle, error) {
	event := new(PoolSettle)
	if err := _Pool.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTakeIterator is returned from FilterTake and is used to iterate over the raw logs and unpacked data for Take events raised by the Pool contract.
type PoolTakeIterator struct {
	Event *PoolTake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolTakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolTake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolTakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTake represents a Take event raised by the Pool contract.
type PoolTake struct {
	Borrower   common.Address
	Amount     *big.Int
	Collateral *big.Int
	BondChange *big.Int
	IsReward   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTake is a free log retrieval operation binding the contract event 0x4591b2dfbebff121b3ccd19ae2407e59a9cefd959b35e12d82752cb5bc6eb757.
//
// Solidity: event Take(address indexed borrower, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_Pool *PoolFilterer) FilterTake(opts *bind.FilterOpts, borrower []common.Address) (*PoolTakeIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Take", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &PoolTakeIterator{contract: _Pool.contract, event: "Take", logs: logs, sub: sub}, nil
}

// WatchTake is a free log subscription operation binding the contract event 0x4591b2dfbebff121b3ccd19ae2407e59a9cefd959b35e12d82752cb5bc6eb757.
//
// Solidity: event Take(address indexed borrower, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_Pool *PoolFilterer) WatchTake(opts *bind.WatchOpts, sink chan<- *PoolTake, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Take", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTake)
				if err := _Pool.contract.UnpackLog(event, "Take", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTake is a log parse operation binding the contract event 0x4591b2dfbebff121b3ccd19ae2407e59a9cefd959b35e12d82752cb5bc6eb757.
//
// Solidity: event Take(address indexed borrower, uint256 amount, uint256 collateral, uint256 bondChange, bool isReward)
func (_Pool *PoolFilterer) ParseTake(log types.Log) (*PoolTake, error) {
	event := new(PoolTake)
	if err := _Pool.contract.UnpackLog(event, "Take", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTransferLPIterator is returned from FilterTransferLP and is used to iterate over the raw logs and unpacked data for TransferLP events raised by the Pool contract.
type PoolTransferLPIterator struct {
	Event *PoolTransferLP // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolTransferLPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTransferLP)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolTransferLP)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolTransferLPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTransferLPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTransferLP represents a TransferLP event raised by the Pool contract.
type PoolTransferLP struct {
	Owner    common.Address
	NewOwner common.Address
	Indexes  []*big.Int
	Lp       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferLP is a free log retrieval operation binding the contract event 0x98b14a8359f3da2f702dde7a51271b67ea43d27b8712e860408b8bf8dd0eb682.
//
// Solidity: event TransferLP(address owner, address newOwner, uint256[] indexes, uint256 lp)
func (_Pool *PoolFilterer) FilterTransferLP(opts *bind.FilterOpts) (*PoolTransferLPIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "TransferLP")
	if err != nil {
		return nil, err
	}
	return &PoolTransferLPIterator{contract: _Pool.contract, event: "TransferLP", logs: logs, sub: sub}, nil
}

// WatchTransferLP is a free log subscription operation binding the contract event 0x98b14a8359f3da2f702dde7a51271b67ea43d27b8712e860408b8bf8dd0eb682.
//
// Solidity: event TransferLP(address owner, address newOwner, uint256[] indexes, uint256 lp)
func (_Pool *PoolFilterer) WatchTransferLP(opts *bind.WatchOpts, sink chan<- *PoolTransferLP) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "TransferLP")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTransferLP)
				if err := _Pool.contract.UnpackLog(event, "TransferLP", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferLP is a log parse operation binding the contract event 0x98b14a8359f3da2f702dde7a51271b67ea43d27b8712e860408b8bf8dd0eb682.
//
// Solidity: event TransferLP(address owner, address newOwner, uint256[] indexes, uint256 lp)
func (_Pool *PoolFilterer) ParseTransferLP(log types.Log) (*PoolTransferLP, error) {
	event := new(PoolTransferLP)
	if err := _Pool.contract.UnpackLog(event, "TransferLP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUpdateInterestRateIterator is returned from FilterUpdateInterestRate and is used to iterate over the raw logs and unpacked data for UpdateInterestRate events raised by the Pool contract.
type PoolUpdateInterestRateIterator struct {
	Event *PoolUpdateInterestRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolUpdateInterestRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUpdateInterestRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolUpdateInterestRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolUpdateInterestRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUpdateInterestRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUpdateInterestRate represents a UpdateInterestRate event raised by the Pool contract.
type PoolUpdateInterestRate struct {
	OldRate *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateInterestRate is a free log retrieval operation binding the contract event 0x2463616ef8e6f9bddf00e4964b853ad9050f87cd3c73985d2ee6b6c8a8336991.
//
// Solidity: event UpdateInterestRate(uint256 oldRate, uint256 newRate)
func (_Pool *PoolFilterer) FilterUpdateInterestRate(opts *bind.FilterOpts) (*PoolUpdateInterestRateIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "UpdateInterestRate")
	if err != nil {
		return nil, err
	}
	return &PoolUpdateInterestRateIterator{contract: _Pool.contract, event: "UpdateInterestRate", logs: logs, sub: sub}, nil
}

// WatchUpdateInterestRate is a free log subscription operation binding the contract event 0x2463616ef8e6f9bddf00e4964b853ad9050f87cd3c73985d2ee6b6c8a8336991.
//
// Solidity: event UpdateInterestRate(uint256 oldRate, uint256 newRate)
func (_Pool *PoolFilterer) WatchUpdateInterestRate(opts *bind.WatchOpts, sink chan<- *PoolUpdateInterestRate) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "UpdateInterestRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUpdateInterestRate)
				if err := _Pool.contract.UnpackLog(event, "UpdateInterestRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateInterestRate is a log parse operation binding the contract event 0x2463616ef8e6f9bddf00e4964b853ad9050f87cd3c73985d2ee6b6c8a8336991.
//
// Solidity: event UpdateInterestRate(uint256 oldRate, uint256 newRate)
func (_Pool *PoolFilterer) ParseUpdateInterestRate(log types.Log) (*PoolUpdateInterestRate, error) {
	event := new(PoolUpdateInterestRate)
	if err := _Pool.contract.UnpackLog(event, "UpdateInterestRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// IAjnaRFQOrder is an auto generated low-level Go binding around an user-defined struct.
type IAjnaRFQOrder struct {
	Maker         common.Address
	Taker         common.Address
	Pool          common.Address
	Index         *big.Int
	MakeAmount    *big.Int
	MinMakeAmount *big.Int
	Expiry        *big.Int
	Price         *big.Int
}

// RFQMetaData contains all meta data concerning the RFQ contract.
var RFQMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ORDER_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveOrder\",\"inputs\":[{\"name\":\"order_\",\"type\":\"tuple\",\"internalType\":\"structIAjnaRFQ.Order\",\"components\":[{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minMakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveOrder\",\"inputs\":[{\"name\":\"hash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvedOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOrder\",\"inputs\":[{\"name\":\"hash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fillOrder\",\"inputs\":[{\"name\":\"order_\",\"type\":\"tuple\",\"internalType\":\"structIAjnaRFQ.Order\",\"components\":[{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minMakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature_\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"to_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fillQuoteAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minFillQuoteAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"filledAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structIAjnaRFQ.Order\",\"components\":[{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minMakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFee\",\"inputs\":[{\"name\":\"fee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FillAmountTooLowForMaker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FillAmountTooLowForTaker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FillExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLPBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMsgValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MissingLP\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderAlreadyFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderCancelled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// RFQABI is the input ABI used to generate the binding from.
// Deprecated: Use RFQMetaData.ABI instead.
var RFQABI = RFQMetaData.ABI

// RFQ is an auto generated Go binding around an Ethereum contract.
type RFQ struct {
	RFQCaller     // Read-only binding to the contract
	RFQTransactor // Write-only binding to the contract
	RFQFilterer   // Log filterer for contract events
}

// RFQCaller is an auto generated read-only Go binding around an Ethereum contract.
type RFQCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RFQTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RFQTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RFQFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RFQFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RFQSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RFQSession struct {
	Contract     *RFQ              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RFQCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RFQCallerSession struct {
	Contract *RFQCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RFQTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RFQTransactorSession struct {
	Contract     *RFQTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RFQRaw is an auto generated low-level Go binding around an Ethereum contract.
type RFQRaw struct {
	Contract *RFQ // Generic contract binding to access the raw methods on
}

// RFQCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RFQCallerRaw struct {
	Contract *RFQCaller // Generic read-only contract binding to access the raw methods on
}

// RFQTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RFQTransactorRaw struct {
	Contract *RFQTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRFQ creates a new instance of RFQ, bound to a specific deployed contract.
func NewRFQ(address common.Address, backend bind.ContractBackend) (*RFQ, error) {
	contract, err := bindRFQ(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RFQ{RFQCaller: RFQCaller{contract: contract}, RFQTransactor: RFQTransactor{contract: contract}, RFQFilterer: RFQFilterer{contract: contract}}, nil
}

// NewRFQCaller creates a new read-only instance of RFQ, bound to a specific deployed contract.
func NewRFQCaller(address common.Address, caller bind.ContractCaller) (*RFQCaller, error) {
	contract, err := bindRFQ(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RFQCaller{contract: contract}, nil
}

// NewRFQTransactor creates a new write-only instance of RFQ, bound to a specific deployed contract.
func NewRFQTransactor(address common.Address, transactor bind.ContractTransactor) (*RFQTransactor, error) {
	contract, err := bindRFQ(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RFQTransactor{contract: contract}, nil
}

// NewRFQFilterer creates a new log filterer instance of RFQ, bound to a specific deployed contract.
func NewRFQFilterer(address common.Address, filterer bind.ContractFilterer) (*RFQFilterer, error) {
	contract, err := bindRFQ(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RFQFilterer{contract: contract}, nil
}

// bindRFQ binds a generic wrapper to an already deployed contract.
func bindRFQ(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RFQMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RFQ *RFQRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RFQ.Contract.RFQCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RFQ *RFQRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RFQ.Contract.RFQTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RFQ *RFQRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RFQ.Contract.RFQTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RFQ *RFQCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RFQ.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RFQ *RFQTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RFQ.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RFQ *RFQTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RFQ.Contract.contract.Transact(opts, method, params...)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_RFQ *RFQCaller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RFQ.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_RFQ *RFQSession) ORDERTYPEHASH() ([32]byte, error) {
	return _RFQ.Contract.ORDERTYPEHASH(&_RFQ.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_RFQ *RFQCallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _RFQ.Contract.ORDERTYPEHASH(&_RFQ.CallOpts)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe5cf3498.
//
// Solidity: function approvedOrders(address , bytes32 ) view returns(bool)
func (_RFQ *RFQCaller) ApprovedOrders(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _RFQ.contract.Call(opts, &out, "approvedOrders", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe5cf3498.
//
// Solidity: function approvedOrders(address , bytes32 ) view returns(bool)
func (_RFQ *RFQSession) ApprovedOrders(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RFQ.Contract.ApprovedOrders(&_RFQ.CallOpts, arg0, arg1)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe5cf3498.
//
// Solidity: function approvedOrders(address , bytes32 ) view returns(bool)
func (_RFQ *RFQCallerSession) ApprovedOrders(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RFQ.Contract.ApprovedOrders(&_RFQ.CallOpts, arg0, arg1)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_RFQ *RFQCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _RFQ.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_RFQ *RFQSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _RFQ.Contract.Eip712Domain(&_RFQ.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_RFQ *RFQCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _RFQ.Contract.Eip712Domain(&_RFQ.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_RFQ *RFQCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RFQ.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_RFQ *RFQSession) Fee() (*big.Int, error) {
	return _RFQ.Contract.Fee(&_RFQ.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_RFQ *RFQCallerSession) Fee() (*big.Int, error) {
	return _RFQ.Contract.Fee(&_RFQ.CallOpts)
}

// FilledAmounts is a free data retrieval call binding the contract method 0xb0de1f61.
//
// Solidity: function filledAmounts(address , bytes32 ) view returns(uint256)
func (_RFQ *RFQCaller) FilledAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RFQ.contract.Call(opts, &out, "filledAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilledAmounts is a free data retrieval call binding the contract method 0xb0de1f61.
//
// Solidity: function filledAmounts(address , bytes32 ) view returns(uint256)
func (_RFQ *RFQSession) FilledAmounts(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _RFQ.Contract.FilledAmounts(&_RFQ.CallOpts, arg0, arg1)
}

// FilledAmounts is a free data retrieval call binding the contract method 0xb0de1f61.
//
// Solidity: function filledAmounts(address , bytes32 ) view returns(uint256)
func (_RFQ *RFQCallerSession) FilledAmounts(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _RFQ.Contract.FilledAmounts(&_RFQ.CallOpts, arg0, arg1)
}

// HashOrder is a free data retrieval call binding the contract method 0x38053926.
//
// Solidity: function hashOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32)
func (_RFQ *RFQCaller) HashOrder(opts *bind.CallOpts, order IAjnaRFQOrder) ([32]byte, error) {
	var out []interface{}
	err := _RFQ.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x38053926.
//
// Solidity: function hashOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32)
func (_RFQ *RFQSession) HashOrder(order IAjnaRFQOrder) ([32]byte, error) {
	return _RFQ.Contract.HashOrder(&_RFQ.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x38053926.
//
// Solidity: function hashOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order) pure returns(bytes32)
func (_RFQ *RFQCallerSession) HashOrder(order IAjnaRFQOrder) ([32]byte, error) {
	return _RFQ.Contract.HashOrder(&_RFQ.CallOpts, order)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RFQ *RFQCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RFQ.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RFQ *RFQSession) Owner() (common.Address, error) {
	return _RFQ.Contract.Owner(&_RFQ.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RFQ *RFQCallerSession) Owner() (common.Address, error) {
	return _RFQ.Contract.Owner(&_RFQ.CallOpts)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x3285d5cb.
//
// Solidity: function approveOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order_) returns(bytes32 hash)
func (_RFQ *RFQTransactor) ApproveOrder(opts *bind.TransactOpts, order_ IAjnaRFQOrder) (*types.Transaction, error) {
	return _RFQ.contract.Transact(opts, "approveOrder", order_)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x3285d5cb.
//
// Solidity: function approveOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order_) returns(bytes32 hash)
func (_RFQ *RFQSession) ApproveOrder(order_ IAjnaRFQOrder) (*types.Transaction, error) {
	return _RFQ.Contract.ApproveOrder(&_RFQ.TransactOpts, order_)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x3285d5cb.
//
// Solidity: function approveOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order_) returns(bytes32 hash)
func (_RFQ *RFQTransactorSession) ApproveOrder(order_ IAjnaRFQOrder) (*types.Transaction, error) {
	return _RFQ.Contract.ApproveOrder(&_RFQ.TransactOpts, order_)
}

// ApproveOrder0 is a paid mutator transaction binding the contract method 0xf115b1c1.
//
// Solidity: function approveOrder(bytes32 hash_) returns()
func (_RFQ *RFQTransactor) ApproveOrder0(opts *bind.TransactOpts, hash_ [32]byte) (*types.Transaction, error) {
	return _RFQ.contract.Transact(opts, "approveOrder0", hash_)
}

// ApproveOrder0 is a paid mutator transaction binding the contract method 0xf115b1c1.
//
// Solidity: function approveOrder(bytes32 hash_) returns()
func (_RFQ *RFQSession) ApproveOrder0(hash_ [32]byte) (*types.Transaction, error) {
	return _RFQ.Contract.ApproveOrder0(&_RFQ.TransactOpts, hash_)
}

// ApproveOrder0 is a paid mutator transaction binding the contract method 0xf115b1c1.
//
// Solidity: function approveOrder(bytes32 hash_) returns()
func (_RFQ *RFQTransactorSession) ApproveOrder0(hash_ [32]byte) (*types.Transaction, error) {
	return _RFQ.Contract.ApproveOrder0(&_RFQ.TransactOpts, hash_)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 hash_) returns()
func (_RFQ *RFQTransactor) CancelOrder(opts *bind.TransactOpts, hash_ [32]byte) (*types.Transaction, error) {
	return _RFQ.contract.Transact(opts, "cancelOrder", hash_)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 hash_) returns()
func (_RFQ *RFQSession) CancelOrder(hash_ [32]byte) (*types.Transaction, error) {
	return _RFQ.Contract.CancelOrder(&_RFQ.TransactOpts, hash_)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 hash_) returns()
func (_RFQ *RFQTransactorSession) CancelOrder(hash_ [32]byte) (*types.Transaction, error) {
	return _RFQ.Contract.CancelOrder(&_RFQ.TransactOpts, hash_)
}

// FillOrder is a paid mutator transaction binding the contract method 0xdff1507c.
//
// Solidity: function fillOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order_, bytes signature_, address to_, uint256 fillQuoteAmount_, uint256 minFillQuoteAmount_, uint256 expiry_) payable returns(uint256, uint256)
func (_RFQ *RFQTransactor) FillOrder(opts *bind.TransactOpts, order_ IAjnaRFQOrder, signature_ []byte, to_ common.Address, fillQuoteAmount_ *big.Int, minFillQuoteAmount_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _RFQ.contract.Transact(opts, "fillOrder", order_, signature_, to_, fillQuoteAmount_, minFillQuoteAmount_, expiry_)
}

// FillOrder is a paid mutator transaction binding the contract method 0xdff1507c.
//
// Solidity: function fillOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order_, bytes signature_, address to_, uint256 fillQuoteAmount_, uint256 minFillQuoteAmount_, uint256 expiry_) payable returns(uint256, uint256)
func (_RFQ *RFQSession) FillOrder(order_ IAjnaRFQOrder, signature_ []byte, to_ common.Address, fillQuoteAmount_ *big.Int, minFillQuoteAmount_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _RFQ.Contract.FillOrder(&_RFQ.TransactOpts, order_, signature_, to_, fillQuoteAmount_, minFillQuoteAmount_, expiry_)
}

// FillOrder is a paid mutator transaction binding the contract method 0xdff1507c.
//
// Solidity: function fillOrder((address,address,address,uint256,uint256,uint256,uint256,uint256) order_, bytes signature_, address to_, uint256 fillQuoteAmount_, uint256 minFillQuoteAmount_, uint256 expiry_) payable returns(uint256, uint256)
func (_RFQ *RFQTransactorSession) FillOrder(order_ IAjnaRFQOrder, signature_ []byte, to_ common.Address, fillQuoteAmount_ *big.Int, minFillQuoteAmount_ *big.Int, expiry_ *big.Int) (*types.Transaction, error) {
	return _RFQ.Contract.FillOrder(&_RFQ.TransactOpts, order_, signature_, to_, fillQuoteAmount_, minFillQuoteAmount_, expiry_)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RFQ *RFQTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _RFQ.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RFQ *RFQSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _RFQ.Contract.Multicall(&_RFQ.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RFQ *RFQTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _RFQ.Contract.Multicall(&_RFQ.TransactOpts, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RFQ *RFQTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RFQ.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RFQ *RFQSession) RenounceOwnership() (*types.Transaction, error) {
	return _RFQ.Contract.RenounceOwnership(&_RFQ.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RFQ *RFQTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RFQ.Contract.RenounceOwnership(&_RFQ.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RFQ *RFQTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RFQ.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RFQ *RFQSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RFQ.Contract.TransferOwnership(&_RFQ.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RFQ *RFQTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RFQ.Contract.TransferOwnership(&_RFQ.TransactOpts, newOwner)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 fee_) returns()
func (_RFQ *RFQTransactor) UpdateFee(opts *bind.TransactOpts, fee_ *big.Int) (*types.Transaction, error) {
	return _RFQ.contract.Transact(opts, "updateFee", fee_)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 fee_) returns()
func (_RFQ *RFQSession) UpdateFee(fee_ *big.Int) (*types.Transaction, error) {
	return _RFQ.Contract.UpdateFee(&_RFQ.TransactOpts, fee_)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 fee_) returns()
func (_RFQ *RFQTransactorSession) UpdateFee(fee_ *big.Int) (*types.Transaction, error) {
	return _RFQ.Contract.UpdateFee(&_RFQ.TransactOpts, fee_)
}

// RFQEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the RFQ contract.
type RFQEIP712DomainChangedIterator struct {
	Event *RFQEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *RFQEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RFQEIP712DomainChanged)
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
		it.Event = new(RFQEIP712DomainChanged)
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
func (it *RFQEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RFQEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RFQEIP712DomainChanged represents a EIP712DomainChanged event raised by the RFQ contract.
type RFQEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_RFQ *RFQFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*RFQEIP712DomainChangedIterator, error) {

	logs, sub, err := _RFQ.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &RFQEIP712DomainChangedIterator{contract: _RFQ.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_RFQ *RFQFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *RFQEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _RFQ.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RFQEIP712DomainChanged)
				if err := _RFQ.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_RFQ *RFQFilterer) ParseEIP712DomainChanged(log types.Log) (*RFQEIP712DomainChanged, error) {
	event := new(RFQEIP712DomainChanged)
	if err := _RFQ.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RFQOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RFQ contract.
type RFQOwnershipTransferredIterator struct {
	Event *RFQOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RFQOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RFQOwnershipTransferred)
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
		it.Event = new(RFQOwnershipTransferred)
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
func (it *RFQOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RFQOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RFQOwnershipTransferred represents a OwnershipTransferred event raised by the RFQ contract.
type RFQOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RFQ *RFQFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RFQOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RFQ.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RFQOwnershipTransferredIterator{contract: _RFQ.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RFQ *RFQFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RFQOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RFQ.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RFQOwnershipTransferred)
				if err := _RFQ.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RFQ *RFQFilterer) ParseOwnershipTransferred(log types.Log) (*RFQOwnershipTransferred, error) {
	event := new(RFQOwnershipTransferred)
	if err := _RFQ.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

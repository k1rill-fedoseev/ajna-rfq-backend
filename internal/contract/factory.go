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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ajna_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CreateFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DecimalsNotCompliant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeployQuoteCollateralSameToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeployWithZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolInterestRateInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20_NON_SUBSET_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ajna\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"interestRate_\",\"type\":\"uint256\"}],\"name\":\"deployPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployedPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deployedPoolsList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployedPoolsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfDeployedPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"contractERC20Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// ERC20NONSUBSETHASH is a free data retrieval call binding the contract method 0xc38dc7fc.
//
// Solidity: function ERC20_NON_SUBSET_HASH() view returns(bytes32)
func (_Factory *FactoryCaller) ERC20NONSUBSETHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "ERC20_NON_SUBSET_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20NONSUBSETHASH is a free data retrieval call binding the contract method 0xc38dc7fc.
//
// Solidity: function ERC20_NON_SUBSET_HASH() view returns(bytes32)
func (_Factory *FactorySession) ERC20NONSUBSETHASH() ([32]byte, error) {
	return _Factory.Contract.ERC20NONSUBSETHASH(&_Factory.CallOpts)
}

// ERC20NONSUBSETHASH is a free data retrieval call binding the contract method 0xc38dc7fc.
//
// Solidity: function ERC20_NON_SUBSET_HASH() view returns(bytes32)
func (_Factory *FactoryCallerSession) ERC20NONSUBSETHASH() ([32]byte, error) {
	return _Factory.Contract.ERC20NONSUBSETHASH(&_Factory.CallOpts)
}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_Factory *FactoryCaller) MAXRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "MAX_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_Factory *FactorySession) MAXRATE() (*big.Int, error) {
	return _Factory.Contract.MAXRATE(&_Factory.CallOpts)
}

// MAXRATE is a free data retrieval call binding the contract method 0xc24dbebd.
//
// Solidity: function MAX_RATE() view returns(uint256)
func (_Factory *FactoryCallerSession) MAXRATE() (*big.Int, error) {
	return _Factory.Contract.MAXRATE(&_Factory.CallOpts)
}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_Factory *FactoryCaller) MINRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "MIN_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_Factory *FactorySession) MINRATE() (*big.Int, error) {
	return _Factory.Contract.MINRATE(&_Factory.CallOpts)
}

// MINRATE is a free data retrieval call binding the contract method 0xd819bfef.
//
// Solidity: function MIN_RATE() view returns(uint256)
func (_Factory *FactoryCallerSession) MINRATE() (*big.Int, error) {
	return _Factory.Contract.MINRATE(&_Factory.CallOpts)
}

// Ajna is a free data retrieval call binding the contract method 0xbb6da0dd.
//
// Solidity: function ajna() view returns(address)
func (_Factory *FactoryCaller) Ajna(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "ajna")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ajna is a free data retrieval call binding the contract method 0xbb6da0dd.
//
// Solidity: function ajna() view returns(address)
func (_Factory *FactorySession) Ajna() (common.Address, error) {
	return _Factory.Contract.Ajna(&_Factory.CallOpts)
}

// Ajna is a free data retrieval call binding the contract method 0xbb6da0dd.
//
// Solidity: function ajna() view returns(address)
func (_Factory *FactoryCallerSession) Ajna() (common.Address, error) {
	return _Factory.Contract.Ajna(&_Factory.CallOpts)
}

// DeployedPools is a free data retrieval call binding the contract method 0x7f165b0b.
//
// Solidity: function deployedPools(bytes32 , address , address ) view returns(address)
func (_Factory *FactoryCaller) DeployedPools(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "deployedPools", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedPools is a free data retrieval call binding the contract method 0x7f165b0b.
//
// Solidity: function deployedPools(bytes32 , address , address ) view returns(address)
func (_Factory *FactorySession) DeployedPools(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	return _Factory.Contract.DeployedPools(&_Factory.CallOpts, arg0, arg1, arg2)
}

// DeployedPools is a free data retrieval call binding the contract method 0x7f165b0b.
//
// Solidity: function deployedPools(bytes32 , address , address ) view returns(address)
func (_Factory *FactoryCallerSession) DeployedPools(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	return _Factory.Contract.DeployedPools(&_Factory.CallOpts, arg0, arg1, arg2)
}

// DeployedPoolsList is a free data retrieval call binding the contract method 0xa387245c.
//
// Solidity: function deployedPoolsList(uint256 ) view returns(address)
func (_Factory *FactoryCaller) DeployedPoolsList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "deployedPoolsList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedPoolsList is a free data retrieval call binding the contract method 0xa387245c.
//
// Solidity: function deployedPoolsList(uint256 ) view returns(address)
func (_Factory *FactorySession) DeployedPoolsList(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.DeployedPoolsList(&_Factory.CallOpts, arg0)
}

// DeployedPoolsList is a free data retrieval call binding the contract method 0xa387245c.
//
// Solidity: function deployedPoolsList(uint256 ) view returns(address)
func (_Factory *FactoryCallerSession) DeployedPoolsList(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.DeployedPoolsList(&_Factory.CallOpts, arg0)
}

// GetDeployedPoolsList is a free data retrieval call binding the contract method 0x2b6983af.
//
// Solidity: function getDeployedPoolsList() view returns(address[])
func (_Factory *FactoryCaller) GetDeployedPoolsList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getDeployedPoolsList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeployedPoolsList is a free data retrieval call binding the contract method 0x2b6983af.
//
// Solidity: function getDeployedPoolsList() view returns(address[])
func (_Factory *FactorySession) GetDeployedPoolsList() ([]common.Address, error) {
	return _Factory.Contract.GetDeployedPoolsList(&_Factory.CallOpts)
}

// GetDeployedPoolsList is a free data retrieval call binding the contract method 0x2b6983af.
//
// Solidity: function getDeployedPoolsList() view returns(address[])
func (_Factory *FactoryCallerSession) GetDeployedPoolsList() ([]common.Address, error) {
	return _Factory.Contract.GetDeployedPoolsList(&_Factory.CallOpts)
}

// GetNumberOfDeployedPools is a free data retrieval call binding the contract method 0xb3d4cfa4.
//
// Solidity: function getNumberOfDeployedPools() view returns(uint256)
func (_Factory *FactoryCaller) GetNumberOfDeployedPools(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getNumberOfDeployedPools")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfDeployedPools is a free data retrieval call binding the contract method 0xb3d4cfa4.
//
// Solidity: function getNumberOfDeployedPools() view returns(uint256)
func (_Factory *FactorySession) GetNumberOfDeployedPools() (*big.Int, error) {
	return _Factory.Contract.GetNumberOfDeployedPools(&_Factory.CallOpts)
}

// GetNumberOfDeployedPools is a free data retrieval call binding the contract method 0xb3d4cfa4.
//
// Solidity: function getNumberOfDeployedPools() view returns(uint256)
func (_Factory *FactoryCallerSession) GetNumberOfDeployedPools() (*big.Int, error) {
	return _Factory.Contract.GetNumberOfDeployedPools(&_Factory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Factory *FactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Factory *FactorySession) Implementation() (common.Address, error) {
	return _Factory.Contract.Implementation(&_Factory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Factory *FactoryCallerSession) Implementation() (common.Address, error) {
	return _Factory.Contract.Implementation(&_Factory.CallOpts)
}

// DeployPool is a paid mutator transaction binding the contract method 0xa3232bf3.
//
// Solidity: function deployPool(address collateral_, address quote_, uint256 interestRate_) returns(address pool_)
func (_Factory *FactoryTransactor) DeployPool(opts *bind.TransactOpts, collateral_ common.Address, quote_ common.Address, interestRate_ *big.Int) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "deployPool", collateral_, quote_, interestRate_)
}

// DeployPool is a paid mutator transaction binding the contract method 0xa3232bf3.
//
// Solidity: function deployPool(address collateral_, address quote_, uint256 interestRate_) returns(address pool_)
func (_Factory *FactorySession) DeployPool(collateral_ common.Address, quote_ common.Address, interestRate_ *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.DeployPool(&_Factory.TransactOpts, collateral_, quote_, interestRate_)
}

// DeployPool is a paid mutator transaction binding the contract method 0xa3232bf3.
//
// Solidity: function deployPool(address collateral_, address quote_, uint256 interestRate_) returns(address pool_)
func (_Factory *FactoryTransactorSession) DeployPool(collateral_ common.Address, quote_ common.Address, interestRate_ *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.DeployPool(&_Factory.TransactOpts, collateral_, quote_, interestRate_)
}

// FactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Factory contract.
type FactoryPoolCreatedIterator struct {
	Event *FactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *FactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryPoolCreated)
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
		it.Event = new(FactoryPoolCreated)
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
func (it *FactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryPoolCreated represents a PoolCreated event raised by the Factory contract.
type FactoryPoolCreated struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address pool_)
func (_Factory *FactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts) (*FactoryPoolCreatedIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return &FactoryPoolCreatedIterator{contract: _Factory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address pool_)
func (_Factory *FactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *FactoryPoolCreated) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryPoolCreated)
				if err := _Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc.
//
// Solidity: event PoolCreated(address pool_)
func (_Factory *FactoryFilterer) ParsePoolCreated(log types.Log) (*FactoryPoolCreated, error) {
	event := new(FactoryPoolCreated)
	if err := _Factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

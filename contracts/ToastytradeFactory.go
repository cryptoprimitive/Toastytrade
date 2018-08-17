// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ToastytradeFactoryABI is the input ABI used to generate the binding from.
const ToastytradeFactoryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"notifyServiceProvider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notifyServiceFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"autoreleaseInterval\",\"type\":\"uint256\"},{\"name\":\"_details\",\"type\":\"string\"}],\"name\":\"newToastytrade\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toastytradeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"details\",\"type\":\"string\"}],\"name\":\"ToastytradeCreated\",\"type\":\"event\"}]"

// ToastytradeFactory is an auto generated Go binding around an Ethereum contract.
type ToastytradeFactory struct {
	ToastytradeFactoryCaller     // Read-only binding to the contract
	ToastytradeFactoryTransactor // Write-only binding to the contract
	ToastytradeFactoryFilterer   // Log filterer for contract events
}

// ToastytradeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ToastytradeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToastytradeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ToastytradeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToastytradeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ToastytradeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToastytradeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ToastytradeFactorySession struct {
	Contract     *ToastytradeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ToastytradeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ToastytradeFactoryCallerSession struct {
	Contract *ToastytradeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ToastytradeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ToastytradeFactoryTransactorSession struct {
	Contract     *ToastytradeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ToastytradeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ToastytradeFactoryRaw struct {
	Contract *ToastytradeFactory // Generic contract binding to access the raw methods on
}

// ToastytradeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ToastytradeFactoryCallerRaw struct {
	Contract *ToastytradeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ToastytradeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ToastytradeFactoryTransactorRaw struct {
	Contract *ToastytradeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToastytradeFactory creates a new instance of ToastytradeFactory, bound to a specific deployed contract.
func NewToastytradeFactory(address common.Address, backend bind.ContractBackend) (*ToastytradeFactory, error) {
	contract, err := bindToastytradeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ToastytradeFactory{ToastytradeFactoryCaller: ToastytradeFactoryCaller{contract: contract}, ToastytradeFactoryTransactor: ToastytradeFactoryTransactor{contract: contract}, ToastytradeFactoryFilterer: ToastytradeFactoryFilterer{contract: contract}}, nil
}

// NewToastytradeFactoryCaller creates a new read-only instance of ToastytradeFactory, bound to a specific deployed contract.
func NewToastytradeFactoryCaller(address common.Address, caller bind.ContractCaller) (*ToastytradeFactoryCaller, error) {
	contract, err := bindToastytradeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ToastytradeFactoryCaller{contract: contract}, nil
}

// NewToastytradeFactoryTransactor creates a new write-only instance of ToastytradeFactory, bound to a specific deployed contract.
func NewToastytradeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ToastytradeFactoryTransactor, error) {
	contract, err := bindToastytradeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ToastytradeFactoryTransactor{contract: contract}, nil
}

// NewToastytradeFactoryFilterer creates a new log filterer instance of ToastytradeFactory, bound to a specific deployed contract.
func NewToastytradeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ToastytradeFactoryFilterer, error) {
	contract, err := bindToastytradeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ToastytradeFactoryFilterer{contract: contract}, nil
}

// bindToastytradeFactory binds a generic wrapper to an already deployed contract.
func bindToastytradeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ToastytradeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ToastytradeFactory *ToastytradeFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ToastytradeFactory.Contract.ToastytradeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ToastytradeFactory *ToastytradeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ToastytradeFactory.Contract.ToastytradeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ToastytradeFactory *ToastytradeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ToastytradeFactory.Contract.ToastytradeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ToastytradeFactory *ToastytradeFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ToastytradeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ToastytradeFactory *ToastytradeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ToastytradeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ToastytradeFactory *ToastytradeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ToastytradeFactory.Contract.contract.Transact(opts, method, params...)
}

// NotifyServiceFee is a free data retrieval call binding the contract method 0x5ee27777.
//
// Solidity: function notifyServiceFee() constant returns(uint256)
func (_ToastytradeFactory *ToastytradeFactoryCaller) NotifyServiceFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ToastytradeFactory.contract.Call(opts, out, "notifyServiceFee")
	return *ret0, err
}

// NotifyServiceFee is a free data retrieval call binding the contract method 0x5ee27777.
//
// Solidity: function notifyServiceFee() constant returns(uint256)
func (_ToastytradeFactory *ToastytradeFactorySession) NotifyServiceFee() (*big.Int, error) {
	return _ToastytradeFactory.Contract.NotifyServiceFee(&_ToastytradeFactory.CallOpts)
}

// NotifyServiceFee is a free data retrieval call binding the contract method 0x5ee27777.
//
// Solidity: function notifyServiceFee() constant returns(uint256)
func (_ToastytradeFactory *ToastytradeFactoryCallerSession) NotifyServiceFee() (*big.Int, error) {
	return _ToastytradeFactory.Contract.NotifyServiceFee(&_ToastytradeFactory.CallOpts)
}

// NotifyServiceProvider is a free data retrieval call binding the contract method 0x22ce8e18.
//
// Solidity: function notifyServiceProvider() constant returns(address)
func (_ToastytradeFactory *ToastytradeFactoryCaller) NotifyServiceProvider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ToastytradeFactory.contract.Call(opts, out, "notifyServiceProvider")
	return *ret0, err
}

// NotifyServiceProvider is a free data retrieval call binding the contract method 0x22ce8e18.
//
// Solidity: function notifyServiceProvider() constant returns(address)
func (_ToastytradeFactory *ToastytradeFactorySession) NotifyServiceProvider() (common.Address, error) {
	return _ToastytradeFactory.Contract.NotifyServiceProvider(&_ToastytradeFactory.CallOpts)
}

// NotifyServiceProvider is a free data retrieval call binding the contract method 0x22ce8e18.
//
// Solidity: function notifyServiceProvider() constant returns(address)
func (_ToastytradeFactory *ToastytradeFactoryCallerSession) NotifyServiceProvider() (common.Address, error) {
	return _ToastytradeFactory.Contract.NotifyServiceProvider(&_ToastytradeFactory.CallOpts)
}

// NewToastytrade is a paid mutator transaction binding the contract method 0xe10c2957.
//
// Solidity: function newToastytrade(seller address, autoreleaseInterval uint256, _details string) returns()
func (_ToastytradeFactory *ToastytradeFactoryTransactor) NewToastytrade(opts *bind.TransactOpts, seller common.Address, autoreleaseInterval *big.Int, _details string) (*types.Transaction, error) {
	return _ToastytradeFactory.contract.Transact(opts, "newToastytrade", seller, autoreleaseInterval, _details)
}

// NewToastytrade is a paid mutator transaction binding the contract method 0xe10c2957.
//
// Solidity: function newToastytrade(seller address, autoreleaseInterval uint256, _details string) returns()
func (_ToastytradeFactory *ToastytradeFactorySession) NewToastytrade(seller common.Address, autoreleaseInterval *big.Int, _details string) (*types.Transaction, error) {
	return _ToastytradeFactory.Contract.NewToastytrade(&_ToastytradeFactory.TransactOpts, seller, autoreleaseInterval, _details)
}

// NewToastytrade is a paid mutator transaction binding the contract method 0xe10c2957.
//
// Solidity: function newToastytrade(seller address, autoreleaseInterval uint256, _details string) returns()
func (_ToastytradeFactory *ToastytradeFactoryTransactorSession) NewToastytrade(seller common.Address, autoreleaseInterval *big.Int, _details string) (*types.Transaction, error) {
	return _ToastytradeFactory.Contract.NewToastytrade(&_ToastytradeFactory.TransactOpts, seller, autoreleaseInterval, _details)
}

// ToastytradeFactoryToastytradeCreatedIterator is returned from FilterToastytradeCreated and is used to iterate over the raw logs and unpacked data for ToastytradeCreated events raised by the ToastytradeFactory contract.
type ToastytradeFactoryToastytradeCreatedIterator struct {
	Event *ToastytradeFactoryToastytradeCreated // Event containing the contract specifics and raw log

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
func (it *ToastytradeFactoryToastytradeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ToastytradeFactoryToastytradeCreated)
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
		it.Event = new(ToastytradeFactoryToastytradeCreated)
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
func (it *ToastytradeFactoryToastytradeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ToastytradeFactoryToastytradeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ToastytradeFactoryToastytradeCreated represents a ToastytradeCreated event raised by the ToastytradeFactory contract.
type ToastytradeFactoryToastytradeCreated struct {
	ToastytradeAddress common.Address
	Details            string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterToastytradeCreated is a free log retrieval operation binding the contract event 0xa81bcfe83c120fa0c03db992f6be954fd2205edafded521ab1059a0318ac1465.
//
// Solidity: e ToastytradeCreated(toastytradeAddress address, details string)
func (_ToastytradeFactory *ToastytradeFactoryFilterer) FilterToastytradeCreated(opts *bind.FilterOpts) (*ToastytradeFactoryToastytradeCreatedIterator, error) {

	logs, sub, err := _ToastytradeFactory.contract.FilterLogs(opts, "ToastytradeCreated")
	if err != nil {
		return nil, err
	}
	return &ToastytradeFactoryToastytradeCreatedIterator{contract: _ToastytradeFactory.contract, event: "ToastytradeCreated", logs: logs, sub: sub}, nil
}

// WatchToastytradeCreated is a free log subscription operation binding the contract event 0xa81bcfe83c120fa0c03db992f6be954fd2205edafded521ab1059a0318ac1465.
//
// Solidity: e ToastytradeCreated(toastytradeAddress address, details string)
func (_ToastytradeFactory *ToastytradeFactoryFilterer) WatchToastytradeCreated(opts *bind.WatchOpts, sink chan<- *ToastytradeFactoryToastytradeCreated) (event.Subscription, error) {

	logs, sub, err := _ToastytradeFactory.contract.WatchLogs(opts, "ToastytradeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ToastytradeFactoryToastytradeCreated)
				if err := _ToastytradeFactory.contract.UnpackLog(event, "ToastytradeCreated", log); err != nil {
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

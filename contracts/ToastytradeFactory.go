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

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"notifyServiceProvider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notifyServiceFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"autoreleaseInterval\",\"type\":\"uint256\"},{\"name\":\"_details\",\"type\":\"string\"}],\"name\":\"newToastytrade\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toastytradeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"details\",\"type\":\"string\"}],\"name\":\"ToastytradeCreated\",\"type\":\"event\"}]"

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// NotifyServiceFee is a free data retrieval call binding the contract method 0x5ee27777.
//
// Solidity: function notifyServiceFee() constant returns(uint256)
func (_Main *MainCaller) NotifyServiceFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "notifyServiceFee")
	return *ret0, err
}

// NotifyServiceFee is a free data retrieval call binding the contract method 0x5ee27777.
//
// Solidity: function notifyServiceFee() constant returns(uint256)
func (_Main *MainSession) NotifyServiceFee() (*big.Int, error) {
	return _Main.Contract.NotifyServiceFee(&_Main.CallOpts)
}

// NotifyServiceFee is a free data retrieval call binding the contract method 0x5ee27777.
//
// Solidity: function notifyServiceFee() constant returns(uint256)
func (_Main *MainCallerSession) NotifyServiceFee() (*big.Int, error) {
	return _Main.Contract.NotifyServiceFee(&_Main.CallOpts)
}

// NotifyServiceProvider is a free data retrieval call binding the contract method 0x22ce8e18.
//
// Solidity: function notifyServiceProvider() constant returns(address)
func (_Main *MainCaller) NotifyServiceProvider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "notifyServiceProvider")
	return *ret0, err
}

// NotifyServiceProvider is a free data retrieval call binding the contract method 0x22ce8e18.
//
// Solidity: function notifyServiceProvider() constant returns(address)
func (_Main *MainSession) NotifyServiceProvider() (common.Address, error) {
	return _Main.Contract.NotifyServiceProvider(&_Main.CallOpts)
}

// NotifyServiceProvider is a free data retrieval call binding the contract method 0x22ce8e18.
//
// Solidity: function notifyServiceProvider() constant returns(address)
func (_Main *MainCallerSession) NotifyServiceProvider() (common.Address, error) {
	return _Main.Contract.NotifyServiceProvider(&_Main.CallOpts)
}

// NewToastytrade is a paid mutator transaction binding the contract method 0xe10c2957.
//
// Solidity: function newToastytrade(seller address, autoreleaseInterval uint256, _details string) returns()
func (_Main *MainTransactor) NewToastytrade(opts *bind.TransactOpts, seller common.Address, autoreleaseInterval *big.Int, _details string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "newToastytrade", seller, autoreleaseInterval, _details)
}

// NewToastytrade is a paid mutator transaction binding the contract method 0xe10c2957.
//
// Solidity: function newToastytrade(seller address, autoreleaseInterval uint256, _details string) returns()
func (_Main *MainSession) NewToastytrade(seller common.Address, autoreleaseInterval *big.Int, _details string) (*types.Transaction, error) {
	return _Main.Contract.NewToastytrade(&_Main.TransactOpts, seller, autoreleaseInterval, _details)
}

// NewToastytrade is a paid mutator transaction binding the contract method 0xe10c2957.
//
// Solidity: function newToastytrade(seller address, autoreleaseInterval uint256, _details string) returns()
func (_Main *MainTransactorSession) NewToastytrade(seller common.Address, autoreleaseInterval *big.Int, _details string) (*types.Transaction, error) {
	return _Main.Contract.NewToastytrade(&_Main.TransactOpts, seller, autoreleaseInterval, _details)
}

// MainToastytradeCreatedIterator is returned from FilterToastytradeCreated and is used to iterate over the raw logs and unpacked data for ToastytradeCreated events raised by the Main contract.
type MainToastytradeCreatedIterator struct {
	Event *MainToastytradeCreated // Event containing the contract specifics and raw log

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
func (it *MainToastytradeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainToastytradeCreated)
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
		it.Event = new(MainToastytradeCreated)
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
func (it *MainToastytradeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainToastytradeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainToastytradeCreated represents a ToastytradeCreated event raised by the Main contract.
type MainToastytradeCreated struct {
	ToastytradeAddress common.Address
	Details            string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterToastytradeCreated is a free log retrieval operation binding the contract event 0xa81bcfe83c120fa0c03db992f6be954fd2205edafded521ab1059a0318ac1465.
//
// Solidity: e ToastytradeCreated(toastytradeAddress address, details string)
func (_Main *MainFilterer) FilterToastytradeCreated(opts *bind.FilterOpts) (*MainToastytradeCreatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ToastytradeCreated")
	if err != nil {
		return nil, err
	}
	return &MainToastytradeCreatedIterator{contract: _Main.contract, event: "ToastytradeCreated", logs: logs, sub: sub}, nil
}

// WatchToastytradeCreated is a free log subscription operation binding the contract event 0xa81bcfe83c120fa0c03db992f6be954fd2205edafded521ab1059a0318ac1465.
//
// Solidity: e ToastytradeCreated(toastytradeAddress address, details string)
func (_Main *MainFilterer) WatchToastytradeCreated(opts *bind.WatchOpts, sink chan<- *MainToastytradeCreated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ToastytradeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainToastytradeCreated)
				if err := _Main.contract.UnpackLog(event, "ToastytradeCreated", log); err != nil {
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

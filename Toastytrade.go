// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ToastytradeABI is the input ABI used to generate the binding from.
const ToastytradeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"details\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BP\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"autoreleaseInterval\",\"type\":\"uint256\"},{\"name\":\"_details\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// Toastytrade is an auto generated Go binding around an Ethereum contract.
type Toastytrade struct {
	ToastytradeCaller     // Read-only binding to the contract
	ToastytradeTransactor // Write-only binding to the contract
	ToastytradeFilterer   // Log filterer for contract events
}

// ToastytradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ToastytradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToastytradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ToastytradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToastytradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ToastytradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToastytradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ToastytradeSession struct {
	Contract     *Toastytrade      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ToastytradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ToastytradeCallerSession struct {
	Contract *ToastytradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ToastytradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ToastytradeTransactorSession struct {
	Contract     *ToastytradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ToastytradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ToastytradeRaw struct {
	Contract *Toastytrade // Generic contract binding to access the raw methods on
}

// ToastytradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ToastytradeCallerRaw struct {
	Contract *ToastytradeCaller // Generic read-only contract binding to access the raw methods on
}

// ToastytradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ToastytradeTransactorRaw struct {
	Contract *ToastytradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToastytrade creates a new instance of Toastytrade, bound to a specific deployed contract.
func NewToastytrade(address common.Address, backend bind.ContractBackend) (*Toastytrade, error) {
	contract, err := bindToastytrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Toastytrade{ToastytradeCaller: ToastytradeCaller{contract: contract}, ToastytradeTransactor: ToastytradeTransactor{contract: contract}, ToastytradeFilterer: ToastytradeFilterer{contract: contract}}, nil
}

// NewToastytradeCaller creates a new read-only instance of Toastytrade, bound to a specific deployed contract.
func NewToastytradeCaller(address common.Address, caller bind.ContractCaller) (*ToastytradeCaller, error) {
	contract, err := bindToastytrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ToastytradeCaller{contract: contract}, nil
}

// NewToastytradeTransactor creates a new write-only instance of Toastytrade, bound to a specific deployed contract.
func NewToastytradeTransactor(address common.Address, transactor bind.ContractTransactor) (*ToastytradeTransactor, error) {
	contract, err := bindToastytrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ToastytradeTransactor{contract: contract}, nil
}

// NewToastytradeFilterer creates a new log filterer instance of Toastytrade, bound to a specific deployed contract.
func NewToastytradeFilterer(address common.Address, filterer bind.ContractFilterer) (*ToastytradeFilterer, error) {
	contract, err := bindToastytrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ToastytradeFilterer{contract: contract}, nil
}

// bindToastytrade binds a generic wrapper to an already deployed contract.
func bindToastytrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ToastytradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Toastytrade *ToastytradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Toastytrade.Contract.ToastytradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Toastytrade *ToastytradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Toastytrade.Contract.ToastytradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Toastytrade *ToastytradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Toastytrade.Contract.ToastytradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Toastytrade *ToastytradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Toastytrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Toastytrade *ToastytradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Toastytrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Toastytrade *ToastytradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Toastytrade.Contract.contract.Transact(opts, method, params...)
}

// BP is a free data retrieval call binding the contract method 0x9cfdbd5e.
//
// Solidity: function BP() constant returns(address)
func (_Toastytrade *ToastytradeCaller) BP(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Toastytrade.contract.Call(opts, out, "BP")
	return *ret0, err
}

// BP is a free data retrieval call binding the contract method 0x9cfdbd5e.
//
// Solidity: function BP() constant returns(address)
func (_Toastytrade *ToastytradeSession) BP() (common.Address, error) {
	return _Toastytrade.Contract.BP(&_Toastytrade.CallOpts)
}

// BP is a free data retrieval call binding the contract method 0x9cfdbd5e.
//
// Solidity: function BP() constant returns(address)
func (_Toastytrade *ToastytradeCallerSession) BP() (common.Address, error) {
	return _Toastytrade.Contract.BP(&_Toastytrade.CallOpts)
}

// Details is a free data retrieval call binding the contract method 0x565974d3.
//
// Solidity: function details() constant returns(string)
func (_Toastytrade *ToastytradeCaller) Details(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Toastytrade.contract.Call(opts, out, "details")
	return *ret0, err
}

// Details is a free data retrieval call binding the contract method 0x565974d3.
//
// Solidity: function details() constant returns(string)
func (_Toastytrade *ToastytradeSession) Details() (string, error) {
	return _Toastytrade.Contract.Details(&_Toastytrade.CallOpts)
}

// Details is a free data retrieval call binding the contract method 0x565974d3.
//
// Solidity: function details() constant returns(string)
func (_Toastytrade *ToastytradeCallerSession) Details() (string, error) {
	return _Toastytrade.Contract.Details(&_Toastytrade.CallOpts)
}

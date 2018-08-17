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
const MainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"payer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"commit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"title\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"worker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoreleaseInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"statement\",\"type\":\"string\"}],\"name\":\"logPayerStatement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFullState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addFunds\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"recoverFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoreleaseTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountReleased\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"startClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"triggerClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"statement\",\"type\":\"string\"}],\"name\":\"logWorkerStatement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountDeposited\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"payerIsOpening\",\"type\":\"bool\"},{\"name\":\"creator\",\"type\":\"address\"},{\"name\":\"_commitThreshold\",\"type\":\"uint256\"},{\"name\":\"_autoreleaseInterval\",\"type\":\"uint256\"},{\"name\":\"_title\",\"type\":\"string\"},{\"name\":\"initialStatement\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"payerOpened\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commitThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"autoreleaseInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"title\",\"type\":\"string\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"statement\",\"type\":\"string\"}],\"name\":\"PayerStatement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"statement\",\"type\":\"string\"}],\"name\":\"WorkerStatement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"committer\",\"type\":\"address\"}],\"name\":\"Committed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ClaimStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ClaimCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ClaimTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unclosed\",\"type\":\"event\"}]"

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

// AmountBurned is a free data retrieval call binding the contract method 0x5290d773.
//
// Solidity: function amountBurned() constant returns(uint256)
func (_Main *MainCaller) AmountBurned(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "amountBurned")
	return *ret0, err
}

// AmountBurned is a free data retrieval call binding the contract method 0x5290d773.
//
// Solidity: function amountBurned() constant returns(uint256)
func (_Main *MainSession) AmountBurned() (*big.Int, error) {
	return _Main.Contract.AmountBurned(&_Main.CallOpts)
}

// AmountBurned is a free data retrieval call binding the contract method 0x5290d773.
//
// Solidity: function amountBurned() constant returns(uint256)
func (_Main *MainCallerSession) AmountBurned() (*big.Int, error) {
	return _Main.Contract.AmountBurned(&_Main.CallOpts)
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_Main *MainCaller) AmountDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "amountDeposited")
	return *ret0, err
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_Main *MainSession) AmountDeposited() (*big.Int, error) {
	return _Main.Contract.AmountDeposited(&_Main.CallOpts)
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_Main *MainCallerSession) AmountDeposited() (*big.Int, error) {
	return _Main.Contract.AmountDeposited(&_Main.CallOpts)
}

// AmountReleased is a free data retrieval call binding the contract method 0xdc7454dd.
//
// Solidity: function amountReleased() constant returns(uint256)
func (_Main *MainCaller) AmountReleased(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "amountReleased")
	return *ret0, err
}

// AmountReleased is a free data retrieval call binding the contract method 0xdc7454dd.
//
// Solidity: function amountReleased() constant returns(uint256)
func (_Main *MainSession) AmountReleased() (*big.Int, error) {
	return _Main.Contract.AmountReleased(&_Main.CallOpts)
}

// AmountReleased is a free data retrieval call binding the contract method 0xdc7454dd.
//
// Solidity: function amountReleased() constant returns(uint256)
func (_Main *MainCallerSession) AmountReleased() (*big.Int, error) {
	return _Main.Contract.AmountReleased(&_Main.CallOpts)
}

// AutoreleaseInterval is a free data retrieval call binding the contract method 0x67aff919.
//
// Solidity: function autoreleaseInterval() constant returns(uint256)
func (_Main *MainCaller) AutoreleaseInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "autoreleaseInterval")
	return *ret0, err
}

// AutoreleaseInterval is a free data retrieval call binding the contract method 0x67aff919.
//
// Solidity: function autoreleaseInterval() constant returns(uint256)
func (_Main *MainSession) AutoreleaseInterval() (*big.Int, error) {
	return _Main.Contract.AutoreleaseInterval(&_Main.CallOpts)
}

// AutoreleaseInterval is a free data retrieval call binding the contract method 0x67aff919.
//
// Solidity: function autoreleaseInterval() constant returns(uint256)
func (_Main *MainCallerSession) AutoreleaseInterval() (*big.Int, error) {
	return _Main.Contract.AutoreleaseInterval(&_Main.CallOpts)
}

// AutoreleaseTime is a free data retrieval call binding the contract method 0xbc308233.
//
// Solidity: function autoreleaseTime() constant returns(uint256)
func (_Main *MainCaller) AutoreleaseTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "autoreleaseTime")
	return *ret0, err
}

// AutoreleaseTime is a free data retrieval call binding the contract method 0xbc308233.
//
// Solidity: function autoreleaseTime() constant returns(uint256)
func (_Main *MainSession) AutoreleaseTime() (*big.Int, error) {
	return _Main.Contract.AutoreleaseTime(&_Main.CallOpts)
}

// AutoreleaseTime is a free data retrieval call binding the contract method 0xbc308233.
//
// Solidity: function autoreleaseTime() constant returns(uint256)
func (_Main *MainCallerSession) AutoreleaseTime() (*big.Int, error) {
	return _Main.Contract.AutoreleaseTime(&_Main.CallOpts)
}

// CommitThreshold is a free data retrieval call binding the contract method 0xec1e74a7.
//
// Solidity: function commitThreshold() constant returns(uint256)
func (_Main *MainCaller) CommitThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "commitThreshold")
	return *ret0, err
}

// CommitThreshold is a free data retrieval call binding the contract method 0xec1e74a7.
//
// Solidity: function commitThreshold() constant returns(uint256)
func (_Main *MainSession) CommitThreshold() (*big.Int, error) {
	return _Main.Contract.CommitThreshold(&_Main.CallOpts)
}

// CommitThreshold is a free data retrieval call binding the contract method 0xec1e74a7.
//
// Solidity: function commitThreshold() constant returns(uint256)
func (_Main *MainCallerSession) CommitThreshold() (*big.Int, error) {
	return _Main.Contract.CommitThreshold(&_Main.CallOpts)
}

// GetFullState is a free data retrieval call binding the contract method 0x972161f7.
//
// Solidity: function getFullState() constant returns(uint8, address, address, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Main *MainCaller) GetFullState(opts *bind.CallOpts) (uint8, common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0  = new(uint8)
		ret1  = new(common.Address)
		ret2  = new(common.Address)
		ret3  = new(string)
		ret4  = new(*big.Int)
		ret5  = new(*big.Int)
		ret6  = new(*big.Int)
		ret7  = new(*big.Int)
		ret8  = new(*big.Int)
		ret9  = new(*big.Int)
		ret10 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
		ret9,
		ret10,
	}
	err := _Main.contract.Call(opts, out, "getFullState")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, *ret9, *ret10, err
}

// GetFullState is a free data retrieval call binding the contract method 0x972161f7.
//
// Solidity: function getFullState() constant returns(uint8, address, address, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Main *MainSession) GetFullState() (uint8, common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Main.Contract.GetFullState(&_Main.CallOpts)
}

// GetFullState is a free data retrieval call binding the contract method 0x972161f7.
//
// Solidity: function getFullState() constant returns(uint8, address, address, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Main *MainCallerSession) GetFullState() (uint8, common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Main.Contract.GetFullState(&_Main.CallOpts)
}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() constant returns(address)
func (_Main *MainCaller) Payer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "payer")
	return *ret0, err
}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() constant returns(address)
func (_Main *MainSession) Payer() (common.Address, error) {
	return _Main.Contract.Payer(&_Main.CallOpts)
}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() constant returns(address)
func (_Main *MainCallerSession) Payer() (common.Address, error) {
	return _Main.Contract.Payer(&_Main.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main *MainCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main *MainSession) State() (uint8, error) {
	return _Main.Contract.State(&_Main.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main *MainCallerSession) State() (uint8, error) {
	return _Main.Contract.State(&_Main.CallOpts)
}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() constant returns(string)
func (_Main *MainCaller) Title(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "title")
	return *ret0, err
}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() constant returns(string)
func (_Main *MainSession) Title() (string, error) {
	return _Main.Contract.Title(&_Main.CallOpts)
}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() constant returns(string)
func (_Main *MainCallerSession) Title() (string, error) {
	return _Main.Contract.Title(&_Main.CallOpts)
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() constant returns(address)
func (_Main *MainCaller) Worker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "worker")
	return *ret0, err
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() constant returns(address)
func (_Main *MainSession) Worker() (common.Address, error) {
	return _Main.Contract.Worker(&_Main.CallOpts)
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() constant returns(address)
func (_Main *MainCallerSession) Worker() (common.Address, error) {
	return _Main.Contract.Worker(&_Main.CallOpts)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_Main *MainTransactor) AddFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addFunds")
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_Main *MainSession) AddFunds() (*types.Transaction, error) {
	return _Main.Contract.AddFunds(&_Main.TransactOpts)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_Main *MainTransactorSession) AddFunds() (*types.Transaction, error) {
	return _Main.Contract.AddFunds(&_Main.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_Main *MainTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_Main *MainSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Burn(&_Main.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_Main *MainTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Burn(&_Main.TransactOpts, amount)
}

// CancelClaim is a paid mutator transaction binding the contract method 0x2bd7cc0f.
//
// Solidity: function cancelClaim() returns()
func (_Main *MainTransactor) CancelClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "cancelClaim")
}

// CancelClaim is a paid mutator transaction binding the contract method 0x2bd7cc0f.
//
// Solidity: function cancelClaim() returns()
func (_Main *MainSession) CancelClaim() (*types.Transaction, error) {
	return _Main.Contract.CancelClaim(&_Main.TransactOpts)
}

// CancelClaim is a paid mutator transaction binding the contract method 0x2bd7cc0f.
//
// Solidity: function cancelClaim() returns()
func (_Main *MainTransactorSession) CancelClaim() (*types.Transaction, error) {
	return _Main.Contract.CancelClaim(&_Main.TransactOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_Main *MainTransactor) Commit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "commit")
}

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_Main *MainSession) Commit() (*types.Transaction, error) {
	return _Main.Contract.Commit(&_Main.TransactOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_Main *MainTransactorSession) Commit() (*types.Transaction, error) {
	return _Main.Contract.Commit(&_Main.TransactOpts)
}

// LogPayerStatement is a paid mutator transaction binding the contract method 0x7345da39.
//
// Solidity: function logPayerStatement(statement string) returns()
func (_Main *MainTransactor) LogPayerStatement(opts *bind.TransactOpts, statement string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "logPayerStatement", statement)
}

// LogPayerStatement is a paid mutator transaction binding the contract method 0x7345da39.
//
// Solidity: function logPayerStatement(statement string) returns()
func (_Main *MainSession) LogPayerStatement(statement string) (*types.Transaction, error) {
	return _Main.Contract.LogPayerStatement(&_Main.TransactOpts, statement)
}

// LogPayerStatement is a paid mutator transaction binding the contract method 0x7345da39.
//
// Solidity: function logPayerStatement(statement string) returns()
func (_Main *MainTransactorSession) LogPayerStatement(statement string) (*types.Transaction, error) {
	return _Main.Contract.LogPayerStatement(&_Main.TransactOpts, statement)
}

// LogWorkerStatement is a paid mutator transaction binding the contract method 0xf3c74496.
//
// Solidity: function logWorkerStatement(statement string) returns()
func (_Main *MainTransactor) LogWorkerStatement(opts *bind.TransactOpts, statement string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "logWorkerStatement", statement)
}

// LogWorkerStatement is a paid mutator transaction binding the contract method 0xf3c74496.
//
// Solidity: function logWorkerStatement(statement string) returns()
func (_Main *MainSession) LogWorkerStatement(statement string) (*types.Transaction, error) {
	return _Main.Contract.LogWorkerStatement(&_Main.TransactOpts, statement)
}

// LogWorkerStatement is a paid mutator transaction binding the contract method 0xf3c74496.
//
// Solidity: function logWorkerStatement(statement string) returns()
func (_Main *MainTransactorSession) LogWorkerStatement(statement string) (*types.Transaction, error) {
	return _Main.Contract.LogWorkerStatement(&_Main.TransactOpts, statement)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Main *MainTransactor) RecoverFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "recoverFunds")
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Main *MainSession) RecoverFunds() (*types.Transaction, error) {
	return _Main.Contract.RecoverFunds(&_Main.TransactOpts)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_Main *MainTransactorSession) RecoverFunds() (*types.Transaction, error) {
	return _Main.Contract.RecoverFunds(&_Main.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(amount uint256) returns()
func (_Main *MainTransactor) Release(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "release", amount)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(amount uint256) returns()
func (_Main *MainSession) Release(amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Release(&_Main.TransactOpts, amount)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(amount uint256) returns()
func (_Main *MainTransactorSession) Release(amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.Release(&_Main.TransactOpts, amount)
}

// StartClaim is a paid mutator transaction binding the contract method 0xecbfc077.
//
// Solidity: function startClaim() returns()
func (_Main *MainTransactor) StartClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "startClaim")
}

// StartClaim is a paid mutator transaction binding the contract method 0xecbfc077.
//
// Solidity: function startClaim() returns()
func (_Main *MainSession) StartClaim() (*types.Transaction, error) {
	return _Main.Contract.StartClaim(&_Main.TransactOpts)
}

// StartClaim is a paid mutator transaction binding the contract method 0xecbfc077.
//
// Solidity: function startClaim() returns()
func (_Main *MainTransactorSession) StartClaim() (*types.Transaction, error) {
	return _Main.Contract.StartClaim(&_Main.TransactOpts)
}

// TriggerClaim is a paid mutator transaction binding the contract method 0xef21fe5a.
//
// Solidity: function triggerClaim() returns()
func (_Main *MainTransactor) TriggerClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "triggerClaim")
}

// TriggerClaim is a paid mutator transaction binding the contract method 0xef21fe5a.
//
// Solidity: function triggerClaim() returns()
func (_Main *MainSession) TriggerClaim() (*types.Transaction, error) {
	return _Main.Contract.TriggerClaim(&_Main.TransactOpts)
}

// TriggerClaim is a paid mutator transaction binding the contract method 0xef21fe5a.
//
// Solidity: function triggerClaim() returns()
func (_Main *MainTransactorSession) TriggerClaim() (*types.Transaction, error) {
	return _Main.Contract.TriggerClaim(&_Main.TransactOpts)
}

// MainClaimCanceledIterator is returned from FilterClaimCanceled and is used to iterate over the raw logs and unpacked data for ClaimCanceled events raised by the Main contract.
type MainClaimCanceledIterator struct {
	Event *MainClaimCanceled // Event containing the contract specifics and raw log

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
func (it *MainClaimCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainClaimCanceled)
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
		it.Event = new(MainClaimCanceled)
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
func (it *MainClaimCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainClaimCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainClaimCanceled represents a ClaimCanceled event raised by the Main contract.
type MainClaimCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimCanceled is a free log retrieval operation binding the contract event 0x599fe6791b76067426490176d39fb8ef37ca6315dc7228225bfc1c589df0555b.
//
// Solidity: e ClaimCanceled()
func (_Main *MainFilterer) FilterClaimCanceled(opts *bind.FilterOpts) (*MainClaimCanceledIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ClaimCanceled")
	if err != nil {
		return nil, err
	}
	return &MainClaimCanceledIterator{contract: _Main.contract, event: "ClaimCanceled", logs: logs, sub: sub}, nil
}

// WatchClaimCanceled is a free log subscription operation binding the contract event 0x599fe6791b76067426490176d39fb8ef37ca6315dc7228225bfc1c589df0555b.
//
// Solidity: e ClaimCanceled()
func (_Main *MainFilterer) WatchClaimCanceled(opts *bind.WatchOpts, sink chan<- *MainClaimCanceled) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ClaimCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainClaimCanceled)
				if err := _Main.contract.UnpackLog(event, "ClaimCanceled", log); err != nil {
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

// MainClaimStartedIterator is returned from FilterClaimStarted and is used to iterate over the raw logs and unpacked data for ClaimStarted events raised by the Main contract.
type MainClaimStartedIterator struct {
	Event *MainClaimStarted // Event containing the contract specifics and raw log

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
func (it *MainClaimStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainClaimStarted)
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
		it.Event = new(MainClaimStarted)
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
func (it *MainClaimStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainClaimStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainClaimStarted represents a ClaimStarted event raised by the Main contract.
type MainClaimStarted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimStarted is a free log retrieval operation binding the contract event 0xc2e00d7be4f68ed88532b9273d506a7aae151286bb6babdb81735dfb224b21a6.
//
// Solidity: e ClaimStarted()
func (_Main *MainFilterer) FilterClaimStarted(opts *bind.FilterOpts) (*MainClaimStartedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ClaimStarted")
	if err != nil {
		return nil, err
	}
	return &MainClaimStartedIterator{contract: _Main.contract, event: "ClaimStarted", logs: logs, sub: sub}, nil
}

// WatchClaimStarted is a free log subscription operation binding the contract event 0xc2e00d7be4f68ed88532b9273d506a7aae151286bb6babdb81735dfb224b21a6.
//
// Solidity: e ClaimStarted()
func (_Main *MainFilterer) WatchClaimStarted(opts *bind.WatchOpts, sink chan<- *MainClaimStarted) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ClaimStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainClaimStarted)
				if err := _Main.contract.UnpackLog(event, "ClaimStarted", log); err != nil {
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

// MainClaimTriggeredIterator is returned from FilterClaimTriggered and is used to iterate over the raw logs and unpacked data for ClaimTriggered events raised by the Main contract.
type MainClaimTriggeredIterator struct {
	Event *MainClaimTriggered // Event containing the contract specifics and raw log

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
func (it *MainClaimTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainClaimTriggered)
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
		it.Event = new(MainClaimTriggered)
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
func (it *MainClaimTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainClaimTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainClaimTriggered represents a ClaimTriggered event raised by the Main contract.
type MainClaimTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimTriggered is a free log retrieval operation binding the contract event 0x0d41024acbae97d7b0daeab1487048fcaedc58b0d754949c140f70f880c628e9.
//
// Solidity: e ClaimTriggered()
func (_Main *MainFilterer) FilterClaimTriggered(opts *bind.FilterOpts) (*MainClaimTriggeredIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ClaimTriggered")
	if err != nil {
		return nil, err
	}
	return &MainClaimTriggeredIterator{contract: _Main.contract, event: "ClaimTriggered", logs: logs, sub: sub}, nil
}

// WatchClaimTriggered is a free log subscription operation binding the contract event 0x0d41024acbae97d7b0daeab1487048fcaedc58b0d754949c140f70f880c628e9.
//
// Solidity: e ClaimTriggered()
func (_Main *MainFilterer) WatchClaimTriggered(opts *bind.WatchOpts, sink chan<- *MainClaimTriggered) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ClaimTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainClaimTriggered)
				if err := _Main.contract.UnpackLog(event, "ClaimTriggered", log); err != nil {
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

// MainClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the Main contract.
type MainClosedIterator struct {
	Event *MainClosed // Event containing the contract specifics and raw log

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
func (it *MainClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainClosed)
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
		it.Event = new(MainClosed)
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
func (it *MainClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainClosed represents a Closed event raised by the Main contract.
type MainClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: e Closed()
func (_Main *MainFilterer) FilterClosed(opts *bind.FilterOpts) (*MainClosedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &MainClosedIterator{contract: _Main.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: e Closed()
func (_Main *MainFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *MainClosed) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainClosed)
				if err := _Main.contract.UnpackLog(event, "Closed", log); err != nil {
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

// MainCommittedIterator is returned from FilterCommitted and is used to iterate over the raw logs and unpacked data for Committed events raised by the Main contract.
type MainCommittedIterator struct {
	Event *MainCommitted // Event containing the contract specifics and raw log

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
func (it *MainCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainCommitted)
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
		it.Event = new(MainCommitted)
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
func (it *MainCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainCommitted represents a Committed event raised by the Main contract.
type MainCommitted struct {
	Committer common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommitted is a free log retrieval operation binding the contract event 0x385d85909904c479680cfb49104dd25dd686a79a13b842e5ab5f1fab8fa0fb2a.
//
// Solidity: e Committed(committer address)
func (_Main *MainFilterer) FilterCommitted(opts *bind.FilterOpts) (*MainCommittedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Committed")
	if err != nil {
		return nil, err
	}
	return &MainCommittedIterator{contract: _Main.contract, event: "Committed", logs: logs, sub: sub}, nil
}

// WatchCommitted is a free log subscription operation binding the contract event 0x385d85909904c479680cfb49104dd25dd686a79a13b842e5ab5f1fab8fa0fb2a.
//
// Solidity: e Committed(committer address)
func (_Main *MainFilterer) WatchCommitted(opts *bind.WatchOpts, sink chan<- *MainCommitted) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Committed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainCommitted)
				if err := _Main.contract.UnpackLog(event, "Committed", log); err != nil {
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

// MainCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Main contract.
type MainCreatedIterator struct {
	Event *MainCreated // Event containing the contract specifics and raw log

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
func (it *MainCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainCreated)
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
		it.Event = new(MainCreated)
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
func (it *MainCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainCreated represents a Created event raised by the Main contract.
type MainCreated struct {
	ContractAddress     common.Address
	PayerOpened         bool
	Creator             common.Address
	CommitThreshold     *big.Int
	AutoreleaseInterval *big.Int
	Title               string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x7948a17e5fb02dd2f672a909a6ae3292d179707215209444a747ffe9fc50d418.
//
// Solidity: e Created(contractAddress indexed address, payerOpened bool, creator address, commitThreshold uint256, autoreleaseInterval uint256, title string)
func (_Main *MainFilterer) FilterCreated(opts *bind.FilterOpts, contractAddress []common.Address) (*MainCreatedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Created", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &MainCreatedIterator{contract: _Main.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x7948a17e5fb02dd2f672a909a6ae3292d179707215209444a747ffe9fc50d418.
//
// Solidity: e Created(contractAddress indexed address, payerOpened bool, creator address, commitThreshold uint256, autoreleaseInterval uint256, title string)
func (_Main *MainFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *MainCreated, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Created", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainCreated)
				if err := _Main.contract.UnpackLog(event, "Created", log); err != nil {
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

// MainFundsAddedIterator is returned from FilterFundsAdded and is used to iterate over the raw logs and unpacked data for FundsAdded events raised by the Main contract.
type MainFundsAddedIterator struct {
	Event *MainFundsAdded // Event containing the contract specifics and raw log

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
func (it *MainFundsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainFundsAdded)
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
		it.Event = new(MainFundsAdded)
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
func (it *MainFundsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainFundsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainFundsAdded represents a FundsAdded event raised by the Main contract.
type MainFundsAdded struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsAdded is a free log retrieval operation binding the contract event 0x8fe10ae416f22f5e5220b0018a6c1d4ff534d6aa3a471f2a20cb7747fe63e5b9.
//
// Solidity: e FundsAdded(from address, amount uint256)
func (_Main *MainFilterer) FilterFundsAdded(opts *bind.FilterOpts) (*MainFundsAddedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "FundsAdded")
	if err != nil {
		return nil, err
	}
	return &MainFundsAddedIterator{contract: _Main.contract, event: "FundsAdded", logs: logs, sub: sub}, nil
}

// WatchFundsAdded is a free log subscription operation binding the contract event 0x8fe10ae416f22f5e5220b0018a6c1d4ff534d6aa3a471f2a20cb7747fe63e5b9.
//
// Solidity: e FundsAdded(from address, amount uint256)
func (_Main *MainFilterer) WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *MainFundsAdded) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "FundsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainFundsAdded)
				if err := _Main.contract.UnpackLog(event, "FundsAdded", log); err != nil {
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

// MainFundsBurnedIterator is returned from FilterFundsBurned and is used to iterate over the raw logs and unpacked data for FundsBurned events raised by the Main contract.
type MainFundsBurnedIterator struct {
	Event *MainFundsBurned // Event containing the contract specifics and raw log

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
func (it *MainFundsBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainFundsBurned)
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
		it.Event = new(MainFundsBurned)
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
func (it *MainFundsBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainFundsBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainFundsBurned represents a FundsBurned event raised by the Main contract.
type MainFundsBurned struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsBurned is a free log retrieval operation binding the contract event 0xe2a0d56d128408deff6c63b30ce69c78024280bc67a251ee2bb096dc08ff1c1e.
//
// Solidity: e FundsBurned(amount uint256)
func (_Main *MainFilterer) FilterFundsBurned(opts *bind.FilterOpts) (*MainFundsBurnedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "FundsBurned")
	if err != nil {
		return nil, err
	}
	return &MainFundsBurnedIterator{contract: _Main.contract, event: "FundsBurned", logs: logs, sub: sub}, nil
}

// WatchFundsBurned is a free log subscription operation binding the contract event 0xe2a0d56d128408deff6c63b30ce69c78024280bc67a251ee2bb096dc08ff1c1e.
//
// Solidity: e FundsBurned(amount uint256)
func (_Main *MainFilterer) WatchFundsBurned(opts *bind.WatchOpts, sink chan<- *MainFundsBurned) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "FundsBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainFundsBurned)
				if err := _Main.contract.UnpackLog(event, "FundsBurned", log); err != nil {
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

// MainFundsRecoveredIterator is returned from FilterFundsRecovered and is used to iterate over the raw logs and unpacked data for FundsRecovered events raised by the Main contract.
type MainFundsRecoveredIterator struct {
	Event *MainFundsRecovered // Event containing the contract specifics and raw log

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
func (it *MainFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainFundsRecovered)
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
		it.Event = new(MainFundsRecovered)
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
func (it *MainFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainFundsRecovered represents a FundsRecovered event raised by the Main contract.
type MainFundsRecovered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFundsRecovered is a free log retrieval operation binding the contract event 0x8bc5aab0b8d1d51bcc031c58eb657027aac7eaa971cc1038d29846400ca22fc5.
//
// Solidity: e FundsRecovered()
func (_Main *MainFilterer) FilterFundsRecovered(opts *bind.FilterOpts) (*MainFundsRecoveredIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &MainFundsRecoveredIterator{contract: _Main.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

// WatchFundsRecovered is a free log subscription operation binding the contract event 0x8bc5aab0b8d1d51bcc031c58eb657027aac7eaa971cc1038d29846400ca22fc5.
//
// Solidity: e FundsRecovered()
func (_Main *MainFilterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *MainFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainFundsRecovered)
				if err := _Main.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
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

// MainFundsReleasedIterator is returned from FilterFundsReleased and is used to iterate over the raw logs and unpacked data for FundsReleased events raised by the Main contract.
type MainFundsReleasedIterator struct {
	Event *MainFundsReleased // Event containing the contract specifics and raw log

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
func (it *MainFundsReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainFundsReleased)
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
		it.Event = new(MainFundsReleased)
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
func (it *MainFundsReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainFundsReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainFundsReleased represents a FundsReleased event raised by the Main contract.
type MainFundsReleased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsReleased is a free log retrieval operation binding the contract event 0x952b264c8e0a06cddb4bbaa6d6af1d565145329fd95bbe72cb2b53942b2dc966.
//
// Solidity: e FundsReleased(amount uint256)
func (_Main *MainFilterer) FilterFundsReleased(opts *bind.FilterOpts) (*MainFundsReleasedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "FundsReleased")
	if err != nil {
		return nil, err
	}
	return &MainFundsReleasedIterator{contract: _Main.contract, event: "FundsReleased", logs: logs, sub: sub}, nil
}

// WatchFundsReleased is a free log subscription operation binding the contract event 0x952b264c8e0a06cddb4bbaa6d6af1d565145329fd95bbe72cb2b53942b2dc966.
//
// Solidity: e FundsReleased(amount uint256)
func (_Main *MainFilterer) WatchFundsReleased(opts *bind.WatchOpts, sink chan<- *MainFundsReleased) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "FundsReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainFundsReleased)
				if err := _Main.contract.UnpackLog(event, "FundsReleased", log); err != nil {
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

// MainPayerStatementIterator is returned from FilterPayerStatement and is used to iterate over the raw logs and unpacked data for PayerStatement events raised by the Main contract.
type MainPayerStatementIterator struct {
	Event *MainPayerStatement // Event containing the contract specifics and raw log

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
func (it *MainPayerStatementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainPayerStatement)
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
		it.Event = new(MainPayerStatement)
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
func (it *MainPayerStatementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainPayerStatementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainPayerStatement represents a PayerStatement event raised by the Main contract.
type MainPayerStatement struct {
	Statement string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayerStatement is a free log retrieval operation binding the contract event 0x21dce665866130bddd42cadae51db6d5093826abb5e5309d67ab8589c7e92694.
//
// Solidity: e PayerStatement(statement string)
func (_Main *MainFilterer) FilterPayerStatement(opts *bind.FilterOpts) (*MainPayerStatementIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "PayerStatement")
	if err != nil {
		return nil, err
	}
	return &MainPayerStatementIterator{contract: _Main.contract, event: "PayerStatement", logs: logs, sub: sub}, nil
}

// WatchPayerStatement is a free log subscription operation binding the contract event 0x21dce665866130bddd42cadae51db6d5093826abb5e5309d67ab8589c7e92694.
//
// Solidity: e PayerStatement(statement string)
func (_Main *MainFilterer) WatchPayerStatement(opts *bind.WatchOpts, sink chan<- *MainPayerStatement) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "PayerStatement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainPayerStatement)
				if err := _Main.contract.UnpackLog(event, "PayerStatement", log); err != nil {
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

// MainUnclosedIterator is returned from FilterUnclosed and is used to iterate over the raw logs and unpacked data for Unclosed events raised by the Main contract.
type MainUnclosedIterator struct {
	Event *MainUnclosed // Event containing the contract specifics and raw log

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
func (it *MainUnclosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainUnclosed)
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
		it.Event = new(MainUnclosed)
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
func (it *MainUnclosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainUnclosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainUnclosed represents a Unclosed event raised by the Main contract.
type MainUnclosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnclosed is a free log retrieval operation binding the contract event 0x295a49ca32ac44ceb5c58aec886eeaf13b1a9cadee420af4c63ed7f1bc75b75b.
//
// Solidity: e Unclosed()
func (_Main *MainFilterer) FilterUnclosed(opts *bind.FilterOpts) (*MainUnclosedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Unclosed")
	if err != nil {
		return nil, err
	}
	return &MainUnclosedIterator{contract: _Main.contract, event: "Unclosed", logs: logs, sub: sub}, nil
}

// WatchUnclosed is a free log subscription operation binding the contract event 0x295a49ca32ac44ceb5c58aec886eeaf13b1a9cadee420af4c63ed7f1bc75b75b.
//
// Solidity: e Unclosed()
func (_Main *MainFilterer) WatchUnclosed(opts *bind.WatchOpts, sink chan<- *MainUnclosed) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Unclosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainUnclosed)
				if err := _Main.contract.UnpackLog(event, "Unclosed", log); err != nil {
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

// MainWorkerStatementIterator is returned from FilterWorkerStatement and is used to iterate over the raw logs and unpacked data for WorkerStatement events raised by the Main contract.
type MainWorkerStatementIterator struct {
	Event *MainWorkerStatement // Event containing the contract specifics and raw log

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
func (it *MainWorkerStatementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainWorkerStatement)
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
		it.Event = new(MainWorkerStatement)
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
func (it *MainWorkerStatementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainWorkerStatementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainWorkerStatement represents a WorkerStatement event raised by the Main contract.
type MainWorkerStatement struct {
	Statement string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWorkerStatement is a free log retrieval operation binding the contract event 0x337c87ca7e10f4ba0201da47ad3a16b990a1198718c55f51688d80da2a35cb75.
//
// Solidity: e WorkerStatement(statement string)
func (_Main *MainFilterer) FilterWorkerStatement(opts *bind.FilterOpts) (*MainWorkerStatementIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "WorkerStatement")
	if err != nil {
		return nil, err
	}
	return &MainWorkerStatementIterator{contract: _Main.contract, event: "WorkerStatement", logs: logs, sub: sub}, nil
}

// WatchWorkerStatement is a free log subscription operation binding the contract event 0x337c87ca7e10f4ba0201da47ad3a16b990a1198718c55f51688d80da2a35cb75.
//
// Solidity: e WorkerStatement(statement string)
func (_Main *MainFilterer) WatchWorkerStatement(opts *bind.WatchOpts, sink chan<- *MainWorkerStatement) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "WorkerStatement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainWorkerStatement)
				if err := _Main.contract.UnpackLog(event, "WorkerStatement", log); err != nil {
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

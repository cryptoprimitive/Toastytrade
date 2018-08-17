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

// BurnablePaymentABI is the input ABI used to generate the binding from.
const BurnablePaymentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"payer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"delayAutorelease\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"commit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"title\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"worker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoreleaseInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"statement\",\"type\":\"string\"}],\"name\":\"logPayerStatement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"triggerAutorelease\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFullState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addFunds\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"recoverFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoreleaseTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountReleased\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"statement\",\"type\":\"string\"}],\"name\":\"logWorkerStatement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountDeposited\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"payerIsOpening\",\"type\":\"bool\"},{\"name\":\"creator\",\"type\":\"address\"},{\"name\":\"_commitThreshold\",\"type\":\"uint256\"},{\"name\":\"_autoreleaseInterval\",\"type\":\"uint256\"},{\"name\":\"_title\",\"type\":\"string\"},{\"name\":\"initialStatement\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"payerOpened\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commitThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"autoreleaseInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"title\",\"type\":\"string\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"statement\",\"type\":\"string\"}],\"name\":\"PayerStatement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"statement\",\"type\":\"string\"}],\"name\":\"WorkerStatement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"committer\",\"type\":\"address\"}],\"name\":\"Committed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unclosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AutoreleaseDelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AutoreleaseTriggered\",\"type\":\"event\"}]"

// BurnablePayment is an auto generated Go binding around an Ethereum contract.
type BurnablePayment struct {
	BurnablePaymentCaller     // Read-only binding to the contract
	BurnablePaymentTransactor // Write-only binding to the contract
	BurnablePaymentFilterer   // Log filterer for contract events
}

// BurnablePaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnablePaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnablePaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnablePaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnablePaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnablePaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnablePaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnablePaymentSession struct {
	Contract     *BurnablePayment  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnablePaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnablePaymentCallerSession struct {
	Contract *BurnablePaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BurnablePaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnablePaymentTransactorSession struct {
	Contract     *BurnablePaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BurnablePaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnablePaymentRaw struct {
	Contract *BurnablePayment // Generic contract binding to access the raw methods on
}

// BurnablePaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnablePaymentCallerRaw struct {
	Contract *BurnablePaymentCaller // Generic read-only contract binding to access the raw methods on
}

// BurnablePaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnablePaymentTransactorRaw struct {
	Contract *BurnablePaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnablePayment creates a new instance of BurnablePayment, bound to a specific deployed contract.
func NewBurnablePayment(address common.Address, backend bind.ContractBackend) (*BurnablePayment, error) {
	contract, err := bindBurnablePayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnablePayment{BurnablePaymentCaller: BurnablePaymentCaller{contract: contract}, BurnablePaymentTransactor: BurnablePaymentTransactor{contract: contract}, BurnablePaymentFilterer: BurnablePaymentFilterer{contract: contract}}, nil
}

// NewBurnablePaymentCaller creates a new read-only instance of BurnablePayment, bound to a specific deployed contract.
func NewBurnablePaymentCaller(address common.Address, caller bind.ContractCaller) (*BurnablePaymentCaller, error) {
	contract, err := bindBurnablePayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentCaller{contract: contract}, nil
}

// NewBurnablePaymentTransactor creates a new write-only instance of BurnablePayment, bound to a specific deployed contract.
func NewBurnablePaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnablePaymentTransactor, error) {
	contract, err := bindBurnablePayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentTransactor{contract: contract}, nil
}

// NewBurnablePaymentFilterer creates a new log filterer instance of BurnablePayment, bound to a specific deployed contract.
func NewBurnablePaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnablePaymentFilterer, error) {
	contract, err := bindBurnablePayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentFilterer{contract: contract}, nil
}

// bindBurnablePayment binds a generic wrapper to an already deployed contract.
func bindBurnablePayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnablePaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnablePayment *BurnablePaymentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnablePayment.Contract.BurnablePaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnablePayment *BurnablePaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnablePayment.Contract.BurnablePaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnablePayment *BurnablePaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnablePayment.Contract.BurnablePaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnablePayment *BurnablePaymentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnablePayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnablePayment *BurnablePaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnablePayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnablePayment *BurnablePaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnablePayment.Contract.contract.Transact(opts, method, params...)
}

// AmountBurned is a free data retrieval call binding the contract method 0x5290d773.
//
// Solidity: function amountBurned() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCaller) AmountBurned(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "amountBurned")
	return *ret0, err
}

// AmountBurned is a free data retrieval call binding the contract method 0x5290d773.
//
// Solidity: function amountBurned() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentSession) AmountBurned() (*big.Int, error) {
	return _BurnablePayment.Contract.AmountBurned(&_BurnablePayment.CallOpts)
}

// AmountBurned is a free data retrieval call binding the contract method 0x5290d773.
//
// Solidity: function amountBurned() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCallerSession) AmountBurned() (*big.Int, error) {
	return _BurnablePayment.Contract.AmountBurned(&_BurnablePayment.CallOpts)
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCaller) AmountDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "amountDeposited")
	return *ret0, err
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentSession) AmountDeposited() (*big.Int, error) {
	return _BurnablePayment.Contract.AmountDeposited(&_BurnablePayment.CallOpts)
}

// AmountDeposited is a free data retrieval call binding the contract method 0xf86ccd41.
//
// Solidity: function amountDeposited() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCallerSession) AmountDeposited() (*big.Int, error) {
	return _BurnablePayment.Contract.AmountDeposited(&_BurnablePayment.CallOpts)
}

// AmountReleased is a free data retrieval call binding the contract method 0xdc7454dd.
//
// Solidity: function amountReleased() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCaller) AmountReleased(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "amountReleased")
	return *ret0, err
}

// AmountReleased is a free data retrieval call binding the contract method 0xdc7454dd.
//
// Solidity: function amountReleased() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentSession) AmountReleased() (*big.Int, error) {
	return _BurnablePayment.Contract.AmountReleased(&_BurnablePayment.CallOpts)
}

// AmountReleased is a free data retrieval call binding the contract method 0xdc7454dd.
//
// Solidity: function amountReleased() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCallerSession) AmountReleased() (*big.Int, error) {
	return _BurnablePayment.Contract.AmountReleased(&_BurnablePayment.CallOpts)
}

// AutoreleaseInterval is a free data retrieval call binding the contract method 0x67aff919.
//
// Solidity: function autoreleaseInterval() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCaller) AutoreleaseInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "autoreleaseInterval")
	return *ret0, err
}

// AutoreleaseInterval is a free data retrieval call binding the contract method 0x67aff919.
//
// Solidity: function autoreleaseInterval() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentSession) AutoreleaseInterval() (*big.Int, error) {
	return _BurnablePayment.Contract.AutoreleaseInterval(&_BurnablePayment.CallOpts)
}

// AutoreleaseInterval is a free data retrieval call binding the contract method 0x67aff919.
//
// Solidity: function autoreleaseInterval() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCallerSession) AutoreleaseInterval() (*big.Int, error) {
	return _BurnablePayment.Contract.AutoreleaseInterval(&_BurnablePayment.CallOpts)
}

// AutoreleaseTime is a free data retrieval call binding the contract method 0xbc308233.
//
// Solidity: function autoreleaseTime() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCaller) AutoreleaseTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "autoreleaseTime")
	return *ret0, err
}

// AutoreleaseTime is a free data retrieval call binding the contract method 0xbc308233.
//
// Solidity: function autoreleaseTime() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentSession) AutoreleaseTime() (*big.Int, error) {
	return _BurnablePayment.Contract.AutoreleaseTime(&_BurnablePayment.CallOpts)
}

// AutoreleaseTime is a free data retrieval call binding the contract method 0xbc308233.
//
// Solidity: function autoreleaseTime() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCallerSession) AutoreleaseTime() (*big.Int, error) {
	return _BurnablePayment.Contract.AutoreleaseTime(&_BurnablePayment.CallOpts)
}

// CommitThreshold is a free data retrieval call binding the contract method 0xec1e74a7.
//
// Solidity: function commitThreshold() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCaller) CommitThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "commitThreshold")
	return *ret0, err
}

// CommitThreshold is a free data retrieval call binding the contract method 0xec1e74a7.
//
// Solidity: function commitThreshold() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentSession) CommitThreshold() (*big.Int, error) {
	return _BurnablePayment.Contract.CommitThreshold(&_BurnablePayment.CallOpts)
}

// CommitThreshold is a free data retrieval call binding the contract method 0xec1e74a7.
//
// Solidity: function commitThreshold() constant returns(uint256)
func (_BurnablePayment *BurnablePaymentCallerSession) CommitThreshold() (*big.Int, error) {
	return _BurnablePayment.Contract.CommitThreshold(&_BurnablePayment.CallOpts)
}

// GetFullState is a free data retrieval call binding the contract method 0x972161f7.
//
// Solidity: function getFullState() constant returns(uint8, address, address, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_BurnablePayment *BurnablePaymentCaller) GetFullState(opts *bind.CallOpts) (uint8, common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
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
	err := _BurnablePayment.contract.Call(opts, out, "getFullState")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, *ret9, *ret10, err
}

// GetFullState is a free data retrieval call binding the contract method 0x972161f7.
//
// Solidity: function getFullState() constant returns(uint8, address, address, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_BurnablePayment *BurnablePaymentSession) GetFullState() (uint8, common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BurnablePayment.Contract.GetFullState(&_BurnablePayment.CallOpts)
}

// GetFullState is a free data retrieval call binding the contract method 0x972161f7.
//
// Solidity: function getFullState() constant returns(uint8, address, address, string, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_BurnablePayment *BurnablePaymentCallerSession) GetFullState() (uint8, common.Address, common.Address, string, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BurnablePayment.Contract.GetFullState(&_BurnablePayment.CallOpts)
}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() constant returns(address)
func (_BurnablePayment *BurnablePaymentCaller) Payer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "payer")
	return *ret0, err
}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() constant returns(address)
func (_BurnablePayment *BurnablePaymentSession) Payer() (common.Address, error) {
	return _BurnablePayment.Contract.Payer(&_BurnablePayment.CallOpts)
}

// Payer is a free data retrieval call binding the contract method 0x123119cd.
//
// Solidity: function payer() constant returns(address)
func (_BurnablePayment *BurnablePaymentCallerSession) Payer() (common.Address, error) {
	return _BurnablePayment.Contract.Payer(&_BurnablePayment.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BurnablePayment *BurnablePaymentCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BurnablePayment *BurnablePaymentSession) State() (uint8, error) {
	return _BurnablePayment.Contract.State(&_BurnablePayment.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BurnablePayment *BurnablePaymentCallerSession) State() (uint8, error) {
	return _BurnablePayment.Contract.State(&_BurnablePayment.CallOpts)
}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() constant returns(string)
func (_BurnablePayment *BurnablePaymentCaller) Title(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "title")
	return *ret0, err
}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() constant returns(string)
func (_BurnablePayment *BurnablePaymentSession) Title() (string, error) {
	return _BurnablePayment.Contract.Title(&_BurnablePayment.CallOpts)
}

// Title is a free data retrieval call binding the contract method 0x4a79d50c.
//
// Solidity: function title() constant returns(string)
func (_BurnablePayment *BurnablePaymentCallerSession) Title() (string, error) {
	return _BurnablePayment.Contract.Title(&_BurnablePayment.CallOpts)
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() constant returns(address)
func (_BurnablePayment *BurnablePaymentCaller) Worker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnablePayment.contract.Call(opts, out, "worker")
	return *ret0, err
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() constant returns(address)
func (_BurnablePayment *BurnablePaymentSession) Worker() (common.Address, error) {
	return _BurnablePayment.Contract.Worker(&_BurnablePayment.CallOpts)
}

// Worker is a free data retrieval call binding the contract method 0x4d547ada.
//
// Solidity: function worker() constant returns(address)
func (_BurnablePayment *BurnablePaymentCallerSession) Worker() (common.Address, error) {
	return _BurnablePayment.Contract.Worker(&_BurnablePayment.CallOpts)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_BurnablePayment *BurnablePaymentTransactor) AddFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "addFunds")
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_BurnablePayment *BurnablePaymentSession) AddFunds() (*types.Transaction, error) {
	return _BurnablePayment.Contract.AddFunds(&_BurnablePayment.TransactOpts)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) AddFunds() (*types.Transaction, error) {
	return _BurnablePayment.Contract.AddFunds(&_BurnablePayment.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_BurnablePayment *BurnablePaymentTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_BurnablePayment *BurnablePaymentSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnablePayment.Contract.Burn(&_BurnablePayment.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnablePayment.Contract.Burn(&_BurnablePayment.TransactOpts, amount)
}

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_BurnablePayment *BurnablePaymentTransactor) Commit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "commit")
}

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_BurnablePayment *BurnablePaymentSession) Commit() (*types.Transaction, error) {
	return _BurnablePayment.Contract.Commit(&_BurnablePayment.TransactOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) Commit() (*types.Transaction, error) {
	return _BurnablePayment.Contract.Commit(&_BurnablePayment.TransactOpts)
}

// DelayAutorelease is a paid mutator transaction binding the contract method 0x127b0901.
//
// Solidity: function delayAutorelease() returns()
func (_BurnablePayment *BurnablePaymentTransactor) DelayAutorelease(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "delayAutorelease")
}

// DelayAutorelease is a paid mutator transaction binding the contract method 0x127b0901.
//
// Solidity: function delayAutorelease() returns()
func (_BurnablePayment *BurnablePaymentSession) DelayAutorelease() (*types.Transaction, error) {
	return _BurnablePayment.Contract.DelayAutorelease(&_BurnablePayment.TransactOpts)
}

// DelayAutorelease is a paid mutator transaction binding the contract method 0x127b0901.
//
// Solidity: function delayAutorelease() returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) DelayAutorelease() (*types.Transaction, error) {
	return _BurnablePayment.Contract.DelayAutorelease(&_BurnablePayment.TransactOpts)
}

// LogPayerStatement is a paid mutator transaction binding the contract method 0x7345da39.
//
// Solidity: function logPayerStatement(statement string) returns()
func (_BurnablePayment *BurnablePaymentTransactor) LogPayerStatement(opts *bind.TransactOpts, statement string) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "logPayerStatement", statement)
}

// LogPayerStatement is a paid mutator transaction binding the contract method 0x7345da39.
//
// Solidity: function logPayerStatement(statement string) returns()
func (_BurnablePayment *BurnablePaymentSession) LogPayerStatement(statement string) (*types.Transaction, error) {
	return _BurnablePayment.Contract.LogPayerStatement(&_BurnablePayment.TransactOpts, statement)
}

// LogPayerStatement is a paid mutator transaction binding the contract method 0x7345da39.
//
// Solidity: function logPayerStatement(statement string) returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) LogPayerStatement(statement string) (*types.Transaction, error) {
	return _BurnablePayment.Contract.LogPayerStatement(&_BurnablePayment.TransactOpts, statement)
}

// LogWorkerStatement is a paid mutator transaction binding the contract method 0xf3c74496.
//
// Solidity: function logWorkerStatement(statement string) returns()
func (_BurnablePayment *BurnablePaymentTransactor) LogWorkerStatement(opts *bind.TransactOpts, statement string) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "logWorkerStatement", statement)
}

// LogWorkerStatement is a paid mutator transaction binding the contract method 0xf3c74496.
//
// Solidity: function logWorkerStatement(statement string) returns()
func (_BurnablePayment *BurnablePaymentSession) LogWorkerStatement(statement string) (*types.Transaction, error) {
	return _BurnablePayment.Contract.LogWorkerStatement(&_BurnablePayment.TransactOpts, statement)
}

// LogWorkerStatement is a paid mutator transaction binding the contract method 0xf3c74496.
//
// Solidity: function logWorkerStatement(statement string) returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) LogWorkerStatement(statement string) (*types.Transaction, error) {
	return _BurnablePayment.Contract.LogWorkerStatement(&_BurnablePayment.TransactOpts, statement)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_BurnablePayment *BurnablePaymentTransactor) RecoverFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "recoverFunds")
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_BurnablePayment *BurnablePaymentSession) RecoverFunds() (*types.Transaction, error) {
	return _BurnablePayment.Contract.RecoverFunds(&_BurnablePayment.TransactOpts)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xb79550be.
//
// Solidity: function recoverFunds() returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) RecoverFunds() (*types.Transaction, error) {
	return _BurnablePayment.Contract.RecoverFunds(&_BurnablePayment.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(amount uint256) returns()
func (_BurnablePayment *BurnablePaymentTransactor) Release(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "release", amount)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(amount uint256) returns()
func (_BurnablePayment *BurnablePaymentSession) Release(amount *big.Int) (*types.Transaction, error) {
	return _BurnablePayment.Contract.Release(&_BurnablePayment.TransactOpts, amount)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(amount uint256) returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) Release(amount *big.Int) (*types.Transaction, error) {
	return _BurnablePayment.Contract.Release(&_BurnablePayment.TransactOpts, amount)
}

// TriggerAutorelease is a paid mutator transaction binding the contract method 0x7d6ad4cd.
//
// Solidity: function triggerAutorelease() returns()
func (_BurnablePayment *BurnablePaymentTransactor) TriggerAutorelease(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnablePayment.contract.Transact(opts, "triggerAutorelease")
}

// TriggerAutorelease is a paid mutator transaction binding the contract method 0x7d6ad4cd.
//
// Solidity: function triggerAutorelease() returns()
func (_BurnablePayment *BurnablePaymentSession) TriggerAutorelease() (*types.Transaction, error) {
	return _BurnablePayment.Contract.TriggerAutorelease(&_BurnablePayment.TransactOpts)
}

// TriggerAutorelease is a paid mutator transaction binding the contract method 0x7d6ad4cd.
//
// Solidity: function triggerAutorelease() returns()
func (_BurnablePayment *BurnablePaymentTransactorSession) TriggerAutorelease() (*types.Transaction, error) {
	return _BurnablePayment.Contract.TriggerAutorelease(&_BurnablePayment.TransactOpts)
}

// BurnablePaymentAutoreleaseDelayedIterator is returned from FilterAutoreleaseDelayed and is used to iterate over the raw logs and unpacked data for AutoreleaseDelayed events raised by the BurnablePayment contract.
type BurnablePaymentAutoreleaseDelayedIterator struct {
	Event *BurnablePaymentAutoreleaseDelayed // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentAutoreleaseDelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentAutoreleaseDelayed)
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
		it.Event = new(BurnablePaymentAutoreleaseDelayed)
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
func (it *BurnablePaymentAutoreleaseDelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentAutoreleaseDelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentAutoreleaseDelayed represents a AutoreleaseDelayed event raised by the BurnablePayment contract.
type BurnablePaymentAutoreleaseDelayed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAutoreleaseDelayed is a free log retrieval operation binding the contract event 0x94306ecb3b8f13e878988b316e670b3a84f5b22fb40a6d534096390fdfd050ff.
//
// Solidity: e AutoreleaseDelayed()
func (_BurnablePayment *BurnablePaymentFilterer) FilterAutoreleaseDelayed(opts *bind.FilterOpts) (*BurnablePaymentAutoreleaseDelayedIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "AutoreleaseDelayed")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentAutoreleaseDelayedIterator{contract: _BurnablePayment.contract, event: "AutoreleaseDelayed", logs: logs, sub: sub}, nil
}

// WatchAutoreleaseDelayed is a free log subscription operation binding the contract event 0x94306ecb3b8f13e878988b316e670b3a84f5b22fb40a6d534096390fdfd050ff.
//
// Solidity: e AutoreleaseDelayed()
func (_BurnablePayment *BurnablePaymentFilterer) WatchAutoreleaseDelayed(opts *bind.WatchOpts, sink chan<- *BurnablePaymentAutoreleaseDelayed) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "AutoreleaseDelayed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentAutoreleaseDelayed)
				if err := _BurnablePayment.contract.UnpackLog(event, "AutoreleaseDelayed", log); err != nil {
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

// BurnablePaymentAutoreleaseTriggeredIterator is returned from FilterAutoreleaseTriggered and is used to iterate over the raw logs and unpacked data for AutoreleaseTriggered events raised by the BurnablePayment contract.
type BurnablePaymentAutoreleaseTriggeredIterator struct {
	Event *BurnablePaymentAutoreleaseTriggered // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentAutoreleaseTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentAutoreleaseTriggered)
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
		it.Event = new(BurnablePaymentAutoreleaseTriggered)
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
func (it *BurnablePaymentAutoreleaseTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentAutoreleaseTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentAutoreleaseTriggered represents a AutoreleaseTriggered event raised by the BurnablePayment contract.
type BurnablePaymentAutoreleaseTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAutoreleaseTriggered is a free log retrieval operation binding the contract event 0xac0ca7228365809cfba9bd7a1549620db6e2c9a4176dae1d163edb0ade5d8606.
//
// Solidity: e AutoreleaseTriggered()
func (_BurnablePayment *BurnablePaymentFilterer) FilterAutoreleaseTriggered(opts *bind.FilterOpts) (*BurnablePaymentAutoreleaseTriggeredIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "AutoreleaseTriggered")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentAutoreleaseTriggeredIterator{contract: _BurnablePayment.contract, event: "AutoreleaseTriggered", logs: logs, sub: sub}, nil
}

// WatchAutoreleaseTriggered is a free log subscription operation binding the contract event 0xac0ca7228365809cfba9bd7a1549620db6e2c9a4176dae1d163edb0ade5d8606.
//
// Solidity: e AutoreleaseTriggered()
func (_BurnablePayment *BurnablePaymentFilterer) WatchAutoreleaseTriggered(opts *bind.WatchOpts, sink chan<- *BurnablePaymentAutoreleaseTriggered) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "AutoreleaseTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentAutoreleaseTriggered)
				if err := _BurnablePayment.contract.UnpackLog(event, "AutoreleaseTriggered", log); err != nil {
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

// BurnablePaymentClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the BurnablePayment contract.
type BurnablePaymentClosedIterator struct {
	Event *BurnablePaymentClosed // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentClosed)
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
		it.Event = new(BurnablePaymentClosed)
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
func (it *BurnablePaymentClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentClosed represents a Closed event raised by the BurnablePayment contract.
type BurnablePaymentClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: e Closed()
func (_BurnablePayment *BurnablePaymentFilterer) FilterClosed(opts *bind.FilterOpts) (*BurnablePaymentClosedIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentClosedIterator{contract: _BurnablePayment.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: e Closed()
func (_BurnablePayment *BurnablePaymentFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *BurnablePaymentClosed) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentClosed)
				if err := _BurnablePayment.contract.UnpackLog(event, "Closed", log); err != nil {
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

// BurnablePaymentCommittedIterator is returned from FilterCommitted and is used to iterate over the raw logs and unpacked data for Committed events raised by the BurnablePayment contract.
type BurnablePaymentCommittedIterator struct {
	Event *BurnablePaymentCommitted // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentCommitted)
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
		it.Event = new(BurnablePaymentCommitted)
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
func (it *BurnablePaymentCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentCommitted represents a Committed event raised by the BurnablePayment contract.
type BurnablePaymentCommitted struct {
	Committer common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommitted is a free log retrieval operation binding the contract event 0x385d85909904c479680cfb49104dd25dd686a79a13b842e5ab5f1fab8fa0fb2a.
//
// Solidity: e Committed(committer address)
func (_BurnablePayment *BurnablePaymentFilterer) FilterCommitted(opts *bind.FilterOpts) (*BurnablePaymentCommittedIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "Committed")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentCommittedIterator{contract: _BurnablePayment.contract, event: "Committed", logs: logs, sub: sub}, nil
}

// WatchCommitted is a free log subscription operation binding the contract event 0x385d85909904c479680cfb49104dd25dd686a79a13b842e5ab5f1fab8fa0fb2a.
//
// Solidity: e Committed(committer address)
func (_BurnablePayment *BurnablePaymentFilterer) WatchCommitted(opts *bind.WatchOpts, sink chan<- *BurnablePaymentCommitted) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "Committed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentCommitted)
				if err := _BurnablePayment.contract.UnpackLog(event, "Committed", log); err != nil {
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

// BurnablePaymentCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the BurnablePayment contract.
type BurnablePaymentCreatedIterator struct {
	Event *BurnablePaymentCreated // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentCreated)
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
		it.Event = new(BurnablePaymentCreated)
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
func (it *BurnablePaymentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentCreated represents a Created event raised by the BurnablePayment contract.
type BurnablePaymentCreated struct {
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
func (_BurnablePayment *BurnablePaymentFilterer) FilterCreated(opts *bind.FilterOpts, contractAddress []common.Address) (*BurnablePaymentCreatedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "Created", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentCreatedIterator{contract: _BurnablePayment.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x7948a17e5fb02dd2f672a909a6ae3292d179707215209444a747ffe9fc50d418.
//
// Solidity: e Created(contractAddress indexed address, payerOpened bool, creator address, commitThreshold uint256, autoreleaseInterval uint256, title string)
func (_BurnablePayment *BurnablePaymentFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *BurnablePaymentCreated, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "Created", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentCreated)
				if err := _BurnablePayment.contract.UnpackLog(event, "Created", log); err != nil {
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

// BurnablePaymentFundsAddedIterator is returned from FilterFundsAdded and is used to iterate over the raw logs and unpacked data for FundsAdded events raised by the BurnablePayment contract.
type BurnablePaymentFundsAddedIterator struct {
	Event *BurnablePaymentFundsAdded // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentFundsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentFundsAdded)
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
		it.Event = new(BurnablePaymentFundsAdded)
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
func (it *BurnablePaymentFundsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentFundsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentFundsAdded represents a FundsAdded event raised by the BurnablePayment contract.
type BurnablePaymentFundsAdded struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsAdded is a free log retrieval operation binding the contract event 0x8fe10ae416f22f5e5220b0018a6c1d4ff534d6aa3a471f2a20cb7747fe63e5b9.
//
// Solidity: e FundsAdded(from address, amount uint256)
func (_BurnablePayment *BurnablePaymentFilterer) FilterFundsAdded(opts *bind.FilterOpts) (*BurnablePaymentFundsAddedIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "FundsAdded")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentFundsAddedIterator{contract: _BurnablePayment.contract, event: "FundsAdded", logs: logs, sub: sub}, nil
}

// WatchFundsAdded is a free log subscription operation binding the contract event 0x8fe10ae416f22f5e5220b0018a6c1d4ff534d6aa3a471f2a20cb7747fe63e5b9.
//
// Solidity: e FundsAdded(from address, amount uint256)
func (_BurnablePayment *BurnablePaymentFilterer) WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *BurnablePaymentFundsAdded) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "FundsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentFundsAdded)
				if err := _BurnablePayment.contract.UnpackLog(event, "FundsAdded", log); err != nil {
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

// BurnablePaymentFundsBurnedIterator is returned from FilterFundsBurned and is used to iterate over the raw logs and unpacked data for FundsBurned events raised by the BurnablePayment contract.
type BurnablePaymentFundsBurnedIterator struct {
	Event *BurnablePaymentFundsBurned // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentFundsBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentFundsBurned)
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
		it.Event = new(BurnablePaymentFundsBurned)
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
func (it *BurnablePaymentFundsBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentFundsBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentFundsBurned represents a FundsBurned event raised by the BurnablePayment contract.
type BurnablePaymentFundsBurned struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsBurned is a free log retrieval operation binding the contract event 0xe2a0d56d128408deff6c63b30ce69c78024280bc67a251ee2bb096dc08ff1c1e.
//
// Solidity: e FundsBurned(amount uint256)
func (_BurnablePayment *BurnablePaymentFilterer) FilterFundsBurned(opts *bind.FilterOpts) (*BurnablePaymentFundsBurnedIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "FundsBurned")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentFundsBurnedIterator{contract: _BurnablePayment.contract, event: "FundsBurned", logs: logs, sub: sub}, nil
}

// WatchFundsBurned is a free log subscription operation binding the contract event 0xe2a0d56d128408deff6c63b30ce69c78024280bc67a251ee2bb096dc08ff1c1e.
//
// Solidity: e FundsBurned(amount uint256)
func (_BurnablePayment *BurnablePaymentFilterer) WatchFundsBurned(opts *bind.WatchOpts, sink chan<- *BurnablePaymentFundsBurned) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "FundsBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentFundsBurned)
				if err := _BurnablePayment.contract.UnpackLog(event, "FundsBurned", log); err != nil {
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

// BurnablePaymentFundsRecoveredIterator is returned from FilterFundsRecovered and is used to iterate over the raw logs and unpacked data for FundsRecovered events raised by the BurnablePayment contract.
type BurnablePaymentFundsRecoveredIterator struct {
	Event *BurnablePaymentFundsRecovered // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentFundsRecovered)
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
		it.Event = new(BurnablePaymentFundsRecovered)
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
func (it *BurnablePaymentFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentFundsRecovered represents a FundsRecovered event raised by the BurnablePayment contract.
type BurnablePaymentFundsRecovered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFundsRecovered is a free log retrieval operation binding the contract event 0x8bc5aab0b8d1d51bcc031c58eb657027aac7eaa971cc1038d29846400ca22fc5.
//
// Solidity: e FundsRecovered()
func (_BurnablePayment *BurnablePaymentFilterer) FilterFundsRecovered(opts *bind.FilterOpts) (*BurnablePaymentFundsRecoveredIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentFundsRecoveredIterator{contract: _BurnablePayment.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

// WatchFundsRecovered is a free log subscription operation binding the contract event 0x8bc5aab0b8d1d51bcc031c58eb657027aac7eaa971cc1038d29846400ca22fc5.
//
// Solidity: e FundsRecovered()
func (_BurnablePayment *BurnablePaymentFilterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *BurnablePaymentFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentFundsRecovered)
				if err := _BurnablePayment.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
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

// BurnablePaymentFundsReleasedIterator is returned from FilterFundsReleased and is used to iterate over the raw logs and unpacked data for FundsReleased events raised by the BurnablePayment contract.
type BurnablePaymentFundsReleasedIterator struct {
	Event *BurnablePaymentFundsReleased // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentFundsReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentFundsReleased)
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
		it.Event = new(BurnablePaymentFundsReleased)
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
func (it *BurnablePaymentFundsReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentFundsReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentFundsReleased represents a FundsReleased event raised by the BurnablePayment contract.
type BurnablePaymentFundsReleased struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsReleased is a free log retrieval operation binding the contract event 0x952b264c8e0a06cddb4bbaa6d6af1d565145329fd95bbe72cb2b53942b2dc966.
//
// Solidity: e FundsReleased(amount uint256)
func (_BurnablePayment *BurnablePaymentFilterer) FilterFundsReleased(opts *bind.FilterOpts) (*BurnablePaymentFundsReleasedIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "FundsReleased")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentFundsReleasedIterator{contract: _BurnablePayment.contract, event: "FundsReleased", logs: logs, sub: sub}, nil
}

// WatchFundsReleased is a free log subscription operation binding the contract event 0x952b264c8e0a06cddb4bbaa6d6af1d565145329fd95bbe72cb2b53942b2dc966.
//
// Solidity: e FundsReleased(amount uint256)
func (_BurnablePayment *BurnablePaymentFilterer) WatchFundsReleased(opts *bind.WatchOpts, sink chan<- *BurnablePaymentFundsReleased) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "FundsReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentFundsReleased)
				if err := _BurnablePayment.contract.UnpackLog(event, "FundsReleased", log); err != nil {
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

// BurnablePaymentPayerStatementIterator is returned from FilterPayerStatement and is used to iterate over the raw logs and unpacked data for PayerStatement events raised by the BurnablePayment contract.
type BurnablePaymentPayerStatementIterator struct {
	Event *BurnablePaymentPayerStatement // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentPayerStatementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentPayerStatement)
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
		it.Event = new(BurnablePaymentPayerStatement)
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
func (it *BurnablePaymentPayerStatementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentPayerStatementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentPayerStatement represents a PayerStatement event raised by the BurnablePayment contract.
type BurnablePaymentPayerStatement struct {
	Statement string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayerStatement is a free log retrieval operation binding the contract event 0x21dce665866130bddd42cadae51db6d5093826abb5e5309d67ab8589c7e92694.
//
// Solidity: e PayerStatement(statement string)
func (_BurnablePayment *BurnablePaymentFilterer) FilterPayerStatement(opts *bind.FilterOpts) (*BurnablePaymentPayerStatementIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "PayerStatement")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentPayerStatementIterator{contract: _BurnablePayment.contract, event: "PayerStatement", logs: logs, sub: sub}, nil
}

// WatchPayerStatement is a free log subscription operation binding the contract event 0x21dce665866130bddd42cadae51db6d5093826abb5e5309d67ab8589c7e92694.
//
// Solidity: e PayerStatement(statement string)
func (_BurnablePayment *BurnablePaymentFilterer) WatchPayerStatement(opts *bind.WatchOpts, sink chan<- *BurnablePaymentPayerStatement) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "PayerStatement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentPayerStatement)
				if err := _BurnablePayment.contract.UnpackLog(event, "PayerStatement", log); err != nil {
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

// BurnablePaymentUnclosedIterator is returned from FilterUnclosed and is used to iterate over the raw logs and unpacked data for Unclosed events raised by the BurnablePayment contract.
type BurnablePaymentUnclosedIterator struct {
	Event *BurnablePaymentUnclosed // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentUnclosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentUnclosed)
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
		it.Event = new(BurnablePaymentUnclosed)
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
func (it *BurnablePaymentUnclosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentUnclosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentUnclosed represents a Unclosed event raised by the BurnablePayment contract.
type BurnablePaymentUnclosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnclosed is a free log retrieval operation binding the contract event 0x295a49ca32ac44ceb5c58aec886eeaf13b1a9cadee420af4c63ed7f1bc75b75b.
//
// Solidity: e Unclosed()
func (_BurnablePayment *BurnablePaymentFilterer) FilterUnclosed(opts *bind.FilterOpts) (*BurnablePaymentUnclosedIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "Unclosed")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentUnclosedIterator{contract: _BurnablePayment.contract, event: "Unclosed", logs: logs, sub: sub}, nil
}

// WatchUnclosed is a free log subscription operation binding the contract event 0x295a49ca32ac44ceb5c58aec886eeaf13b1a9cadee420af4c63ed7f1bc75b75b.
//
// Solidity: e Unclosed()
func (_BurnablePayment *BurnablePaymentFilterer) WatchUnclosed(opts *bind.WatchOpts, sink chan<- *BurnablePaymentUnclosed) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "Unclosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentUnclosed)
				if err := _BurnablePayment.contract.UnpackLog(event, "Unclosed", log); err != nil {
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

// BurnablePaymentWorkerStatementIterator is returned from FilterWorkerStatement and is used to iterate over the raw logs and unpacked data for WorkerStatement events raised by the BurnablePayment contract.
type BurnablePaymentWorkerStatementIterator struct {
	Event *BurnablePaymentWorkerStatement // Event containing the contract specifics and raw log

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
func (it *BurnablePaymentWorkerStatementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnablePaymentWorkerStatement)
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
		it.Event = new(BurnablePaymentWorkerStatement)
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
func (it *BurnablePaymentWorkerStatementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnablePaymentWorkerStatementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnablePaymentWorkerStatement represents a WorkerStatement event raised by the BurnablePayment contract.
type BurnablePaymentWorkerStatement struct {
	Statement string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWorkerStatement is a free log retrieval operation binding the contract event 0x337c87ca7e10f4ba0201da47ad3a16b990a1198718c55f51688d80da2a35cb75.
//
// Solidity: e WorkerStatement(statement string)
func (_BurnablePayment *BurnablePaymentFilterer) FilterWorkerStatement(opts *bind.FilterOpts) (*BurnablePaymentWorkerStatementIterator, error) {

	logs, sub, err := _BurnablePayment.contract.FilterLogs(opts, "WorkerStatement")
	if err != nil {
		return nil, err
	}
	return &BurnablePaymentWorkerStatementIterator{contract: _BurnablePayment.contract, event: "WorkerStatement", logs: logs, sub: sub}, nil
}

// WatchWorkerStatement is a free log subscription operation binding the contract event 0x337c87ca7e10f4ba0201da47ad3a16b990a1198718c55f51688d80da2a35cb75.
//
// Solidity: e WorkerStatement(statement string)
func (_BurnablePayment *BurnablePaymentFilterer) WatchWorkerStatement(opts *bind.WatchOpts, sink chan<- *BurnablePaymentWorkerStatement) (event.Subscription, error) {

	logs, sub, err := _BurnablePayment.contract.WatchLogs(opts, "WorkerStatement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnablePaymentWorkerStatement)
				if err := _BurnablePayment.contract.UnpackLog(event, "WorkerStatement", log); err != nil {
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tc

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

// TcMetaData contains all meta data concerning the Tc contract.
var TcMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nullifierHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIELD_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_HISTORY_SIZE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRootIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denomination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filledSubtrees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIHasher\",\"name\":\"_hasher\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasher\",\"outputs\":[{\"internalType\":\"contractIHasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"isKnownRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nullifierHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_nullifierHash\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_refund\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nullifierHash\",\"type\":\"bytes32\"}],\"name\":\"isSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_nullifierHashes\",\"type\":\"bytes32[]\"}],\"name\":\"isSpentArray\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"spent\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TcABI is the input ABI used to generate the binding from.
// Deprecated: Use TcMetaData.ABI instead.
var TcABI = TcMetaData.ABI

// Tc is an auto generated Go binding around an Ethereum contract.
type Tc struct {
	TcCaller     // Read-only binding to the contract
	TcTransactor // Write-only binding to the contract
	TcFilterer   // Log filterer for contract events
}

// TcCaller is an auto generated read-only Go binding around an Ethereum contract.
type TcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TcSession struct {
	Contract     *Tc               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TcCallerSession struct {
	Contract *TcCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TcTransactorSession struct {
	Contract     *TcTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TcRaw is an auto generated low-level Go binding around an Ethereum contract.
type TcRaw struct {
	Contract *Tc // Generic contract binding to access the raw methods on
}

// TcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TcCallerRaw struct {
	Contract *TcCaller // Generic read-only contract binding to access the raw methods on
}

// TcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TcTransactorRaw struct {
	Contract *TcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTc creates a new instance of Tc, bound to a specific deployed contract.
func NewTc(address common.Address, backend bind.ContractBackend) (*Tc, error) {
	contract, err := bindTc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tc{TcCaller: TcCaller{contract: contract}, TcTransactor: TcTransactor{contract: contract}, TcFilterer: TcFilterer{contract: contract}}, nil
}

// NewTcCaller creates a new read-only instance of Tc, bound to a specific deployed contract.
func NewTcCaller(address common.Address, caller bind.ContractCaller) (*TcCaller, error) {
	contract, err := bindTc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TcCaller{contract: contract}, nil
}

// NewTcTransactor creates a new write-only instance of Tc, bound to a specific deployed contract.
func NewTcTransactor(address common.Address, transactor bind.ContractTransactor) (*TcTransactor, error) {
	contract, err := bindTc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TcTransactor{contract: contract}, nil
}

// NewTcFilterer creates a new log filterer instance of Tc, bound to a specific deployed contract.
func NewTcFilterer(address common.Address, filterer bind.ContractFilterer) (*TcFilterer, error) {
	contract, err := bindTc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TcFilterer{contract: contract}, nil
}

// bindTc binds a generic wrapper to an already deployed contract.
func bindTc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tc *TcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tc.Contract.TcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tc *TcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tc.Contract.TcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tc *TcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tc.Contract.TcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tc *TcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tc *TcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tc *TcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tc.Contract.contract.Transact(opts, method, params...)
}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_Tc *TcCaller) FIELDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "FIELD_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_Tc *TcSession) FIELDSIZE() (*big.Int, error) {
	return _Tc.Contract.FIELDSIZE(&_Tc.CallOpts)
}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_Tc *TcCallerSession) FIELDSIZE() (*big.Int, error) {
	return _Tc.Contract.FIELDSIZE(&_Tc.CallOpts)
}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_Tc *TcCaller) ROOTHISTORYSIZE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "ROOT_HISTORY_SIZE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_Tc *TcSession) ROOTHISTORYSIZE() (uint32, error) {
	return _Tc.Contract.ROOTHISTORYSIZE(&_Tc.CallOpts)
}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_Tc *TcCallerSession) ROOTHISTORYSIZE() (uint32, error) {
	return _Tc.Contract.ROOTHISTORYSIZE(&_Tc.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Tc *TcCaller) ZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Tc *TcSession) ZEROVALUE() (*big.Int, error) {
	return _Tc.Contract.ZEROVALUE(&_Tc.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Tc *TcCallerSession) ZEROVALUE() (*big.Int, error) {
	return _Tc.Contract.ZEROVALUE(&_Tc.CallOpts)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Tc *TcCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Tc *TcSession) Commitments(arg0 [32]byte) (bool, error) {
	return _Tc.Contract.Commitments(&_Tc.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Tc *TcCallerSession) Commitments(arg0 [32]byte) (bool, error) {
	return _Tc.Contract.Commitments(&_Tc.CallOpts, arg0)
}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_Tc *TcCaller) CurrentRootIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "currentRootIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_Tc *TcSession) CurrentRootIndex() (uint32, error) {
	return _Tc.Contract.CurrentRootIndex(&_Tc.CallOpts)
}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_Tc *TcCallerSession) CurrentRootIndex() (uint32, error) {
	return _Tc.Contract.CurrentRootIndex(&_Tc.CallOpts)
}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_Tc *TcCaller) Denomination(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "denomination")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_Tc *TcSession) Denomination() (*big.Int, error) {
	return _Tc.Contract.Denomination(&_Tc.CallOpts)
}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_Tc *TcCallerSession) Denomination() (*big.Int, error) {
	return _Tc.Contract.Denomination(&_Tc.CallOpts)
}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_Tc *TcCaller) FilledSubtrees(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "filledSubtrees", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_Tc *TcSession) FilledSubtrees(arg0 *big.Int) ([32]byte, error) {
	return _Tc.Contract.FilledSubtrees(&_Tc.CallOpts, arg0)
}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_Tc *TcCallerSession) FilledSubtrees(arg0 *big.Int) ([32]byte, error) {
	return _Tc.Contract.FilledSubtrees(&_Tc.CallOpts, arg0)
}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_Tc *TcCaller) GetLastRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "getLastRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_Tc *TcSession) GetLastRoot() ([32]byte, error) {
	return _Tc.Contract.GetLastRoot(&_Tc.CallOpts)
}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_Tc *TcCallerSession) GetLastRoot() ([32]byte, error) {
	return _Tc.Contract.GetLastRoot(&_Tc.CallOpts)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Tc *TcCaller) HashLeftRight(opts *bind.CallOpts, _hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "hashLeftRight", _hasher, _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Tc *TcSession) HashLeftRight(_hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Tc.Contract.HashLeftRight(&_Tc.CallOpts, _hasher, _left, _right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Tc *TcCallerSession) HashLeftRight(_hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Tc.Contract.HashLeftRight(&_Tc.CallOpts, _hasher, _left, _right)
}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_Tc *TcCaller) Hasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "hasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_Tc *TcSession) Hasher() (common.Address, error) {
	return _Tc.Contract.Hasher(&_Tc.CallOpts)
}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_Tc *TcCallerSession) Hasher() (common.Address, error) {
	return _Tc.Contract.Hasher(&_Tc.CallOpts)
}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_Tc *TcCaller) IsKnownRoot(opts *bind.CallOpts, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "isKnownRoot", _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_Tc *TcSession) IsKnownRoot(_root [32]byte) (bool, error) {
	return _Tc.Contract.IsKnownRoot(&_Tc.CallOpts, _root)
}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_Tc *TcCallerSession) IsKnownRoot(_root [32]byte) (bool, error) {
	return _Tc.Contract.IsKnownRoot(&_Tc.CallOpts, _root)
}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_Tc *TcCaller) IsSpent(opts *bind.CallOpts, _nullifierHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "isSpent", _nullifierHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_Tc *TcSession) IsSpent(_nullifierHash [32]byte) (bool, error) {
	return _Tc.Contract.IsSpent(&_Tc.CallOpts, _nullifierHash)
}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_Tc *TcCallerSession) IsSpent(_nullifierHash [32]byte) (bool, error) {
	return _Tc.Contract.IsSpent(&_Tc.CallOpts, _nullifierHash)
}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_Tc *TcCaller) IsSpentArray(opts *bind.CallOpts, _nullifierHashes [][32]byte) ([]bool, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "isSpentArray", _nullifierHashes)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_Tc *TcSession) IsSpentArray(_nullifierHashes [][32]byte) ([]bool, error) {
	return _Tc.Contract.IsSpentArray(&_Tc.CallOpts, _nullifierHashes)
}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_Tc *TcCallerSession) IsSpentArray(_nullifierHashes [][32]byte) ([]bool, error) {
	return _Tc.Contract.IsSpentArray(&_Tc.CallOpts, _nullifierHashes)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_Tc *TcCaller) Levels(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "levels")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_Tc *TcSession) Levels() (uint32, error) {
	return _Tc.Contract.Levels(&_Tc.CallOpts)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_Tc *TcCallerSession) Levels() (uint32, error) {
	return _Tc.Contract.Levels(&_Tc.CallOpts)
}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_Tc *TcCaller) NextIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "nextIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_Tc *TcSession) NextIndex() (uint32, error) {
	return _Tc.Contract.NextIndex(&_Tc.CallOpts)
}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_Tc *TcCallerSession) NextIndex() (uint32, error) {
	return _Tc.Contract.NextIndex(&_Tc.CallOpts)
}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_Tc *TcCaller) NullifierHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "nullifierHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_Tc *TcSession) NullifierHashes(arg0 [32]byte) (bool, error) {
	return _Tc.Contract.NullifierHashes(&_Tc.CallOpts, arg0)
}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_Tc *TcCallerSession) NullifierHashes(arg0 [32]byte) (bool, error) {
	return _Tc.Contract.NullifierHashes(&_Tc.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_Tc *TcCaller) Roots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_Tc *TcSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _Tc.Contract.Roots(&_Tc.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_Tc *TcCallerSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _Tc.Contract.Roots(&_Tc.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Tc *TcCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Tc *TcSession) Verifier() (common.Address, error) {
	return _Tc.Contract.Verifier(&_Tc.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Tc *TcCallerSession) Verifier() (common.Address, error) {
	return _Tc.Contract.Verifier(&_Tc.CallOpts)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_Tc *TcCaller) Zeros(opts *bind.CallOpts, i *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Tc.contract.Call(opts, &out, "zeros", i)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_Tc *TcSession) Zeros(i *big.Int) ([32]byte, error) {
	return _Tc.Contract.Zeros(&_Tc.CallOpts, i)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_Tc *TcCallerSession) Zeros(i *big.Int) ([32]byte, error) {
	return _Tc.Contract.Zeros(&_Tc.CallOpts, i)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_Tc *TcTransactor) Deposit(opts *bind.TransactOpts, _commitment [32]byte) (*types.Transaction, error) {
	return _Tc.contract.Transact(opts, "deposit", _commitment)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_Tc *TcSession) Deposit(_commitment [32]byte) (*types.Transaction, error) {
	return _Tc.Contract.Deposit(&_Tc.TransactOpts, _commitment)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_Tc *TcTransactorSession) Deposit(_commitment [32]byte) (*types.Transaction, error) {
	return _Tc.Contract.Deposit(&_Tc.TransactOpts, _commitment)
}

// Withdraw is a paid mutator transaction binding the contract method 0x21a0adb6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient, address _relayer, uint256 _fee, uint256 _refund) payable returns()
func (_Tc *TcTransactor) Withdraw(opts *bind.TransactOpts, _proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address, _relayer common.Address, _fee *big.Int, _refund *big.Int) (*types.Transaction, error) {
	return _Tc.contract.Transact(opts, "withdraw", _proof, _root, _nullifierHash, _recipient, _relayer, _fee, _refund)
}

// Withdraw is a paid mutator transaction binding the contract method 0x21a0adb6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient, address _relayer, uint256 _fee, uint256 _refund) payable returns()
func (_Tc *TcSession) Withdraw(_proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address, _relayer common.Address, _fee *big.Int, _refund *big.Int) (*types.Transaction, error) {
	return _Tc.Contract.Withdraw(&_Tc.TransactOpts, _proof, _root, _nullifierHash, _recipient, _relayer, _fee, _refund)
}

// Withdraw is a paid mutator transaction binding the contract method 0x21a0adb6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient, address _relayer, uint256 _fee, uint256 _refund) payable returns()
func (_Tc *TcTransactorSession) Withdraw(_proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address, _relayer common.Address, _fee *big.Int, _refund *big.Int) (*types.Transaction, error) {
	return _Tc.Contract.Withdraw(&_Tc.TransactOpts, _proof, _root, _nullifierHash, _recipient, _relayer, _fee, _refund)
}

// TcDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Tc contract.
type TcDepositIterator struct {
	Event *TcDeposit // Event containing the contract specifics and raw log

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
func (it *TcDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TcDeposit)
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
		it.Event = new(TcDeposit)
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
func (it *TcDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TcDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TcDeposit represents a Deposit event raised by the Tc contract.
type TcDeposit struct {
	Commitment [32]byte
	LeafIndex  uint32
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xa945e51eec50ab98c161376f0db4cf2aeba3ec92755fe2fcd388bdbbb80ff196.
//
// Solidity: event Deposit(bytes32 indexed commitment, uint32 leafIndex, uint256 timestamp)
func (_Tc *TcFilterer) FilterDeposit(opts *bind.FilterOpts, commitment [][32]byte) (*TcDepositIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _Tc.contract.FilterLogs(opts, "Deposit", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &TcDepositIterator{contract: _Tc.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xa945e51eec50ab98c161376f0db4cf2aeba3ec92755fe2fcd388bdbbb80ff196.
//
// Solidity: event Deposit(bytes32 indexed commitment, uint32 leafIndex, uint256 timestamp)
func (_Tc *TcFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TcDeposit, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _Tc.contract.WatchLogs(opts, "Deposit", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TcDeposit)
				if err := _Tc.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xa945e51eec50ab98c161376f0db4cf2aeba3ec92755fe2fcd388bdbbb80ff196.
//
// Solidity: event Deposit(bytes32 indexed commitment, uint32 leafIndex, uint256 timestamp)
func (_Tc *TcFilterer) ParseDeposit(log types.Log) (*TcDeposit, error) {
	event := new(TcDeposit)
	if err := _Tc.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TcWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Tc contract.
type TcWithdrawalIterator struct {
	Event *TcWithdrawal // Event containing the contract specifics and raw log

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
func (it *TcWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TcWithdrawal)
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
		it.Event = new(TcWithdrawal)
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
func (it *TcWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TcWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TcWithdrawal represents a Withdrawal event raised by the Tc contract.
type TcWithdrawal struct {
	To            common.Address
	NullifierHash [32]byte
	Relayer       common.Address
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xe9e508bad6d4c3227e881ca19068f099da81b5164dd6d62b2eaf1e8bc6c34931.
//
// Solidity: event Withdrawal(address to, bytes32 nullifierHash, address indexed relayer, uint256 fee)
func (_Tc *TcFilterer) FilterWithdrawal(opts *bind.FilterOpts, relayer []common.Address) (*TcWithdrawalIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Tc.contract.FilterLogs(opts, "Withdrawal", relayerRule)
	if err != nil {
		return nil, err
	}
	return &TcWithdrawalIterator{contract: _Tc.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xe9e508bad6d4c3227e881ca19068f099da81b5164dd6d62b2eaf1e8bc6c34931.
//
// Solidity: event Withdrawal(address to, bytes32 nullifierHash, address indexed relayer, uint256 fee)
func (_Tc *TcFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *TcWithdrawal, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Tc.contract.WatchLogs(opts, "Withdrawal", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TcWithdrawal)
				if err := _Tc.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xe9e508bad6d4c3227e881ca19068f099da81b5164dd6d62b2eaf1e8bc6c34931.
//
// Solidity: event Withdrawal(address to, bytes32 nullifierHash, address indexed relayer, uint256 fee)
func (_Tc *TcFilterer) ParseWithdrawal(log types.Log) (*TcWithdrawal, error) {
	event := new(TcWithdrawal)
	if err := _Tc.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

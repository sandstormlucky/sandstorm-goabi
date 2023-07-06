// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethtc

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

// EthtcMetaData contains all meta data concerning the Ethtc contract.
var EthtcMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nullifierHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIELD_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_HISTORY_SIZE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRootIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denomination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filledSubtrees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIHasher\",\"name\":\"_hasher\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasher\",\"outputs\":[{\"internalType\":\"contractIHasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"isKnownRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nullifierHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_nullifierHash\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_refund\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nullifierHash\",\"type\":\"bytes32\"}],\"name\":\"isSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_nullifierHashes\",\"type\":\"bytes32[]\"}],\"name\":\"isSpentArray\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"spent\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EthtcABI is the input ABI used to generate the binding from.
// Deprecated: Use EthtcMetaData.ABI instead.
var EthtcABI = EthtcMetaData.ABI

// Ethtc is an auto generated Go binding around an Ethereum contract.
type Ethtc struct {
	EthtcCaller     // Read-only binding to the contract
	EthtcTransactor // Write-only binding to the contract
	EthtcFilterer   // Log filterer for contract events
}

// EthtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthtcSession struct {
	Contract     *Ethtc            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthtcCallerSession struct {
	Contract *EthtcCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthtcTransactorSession struct {
	Contract     *EthtcTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthtcRaw struct {
	Contract *Ethtc // Generic contract binding to access the raw methods on
}

// EthtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthtcCallerRaw struct {
	Contract *EthtcCaller // Generic read-only contract binding to access the raw methods on
}

// EthtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthtcTransactorRaw struct {
	Contract *EthtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthtc creates a new instance of Ethtc, bound to a specific deployed contract.
func NewEthtc(address common.Address, backend bind.ContractBackend) (*Ethtc, error) {
	contract, err := bindEthtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethtc{EthtcCaller: EthtcCaller{contract: contract}, EthtcTransactor: EthtcTransactor{contract: contract}, EthtcFilterer: EthtcFilterer{contract: contract}}, nil
}

// NewEthtcCaller creates a new read-only instance of Ethtc, bound to a specific deployed contract.
func NewEthtcCaller(address common.Address, caller bind.ContractCaller) (*EthtcCaller, error) {
	contract, err := bindEthtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthtcCaller{contract: contract}, nil
}

// NewEthtcTransactor creates a new write-only instance of Ethtc, bound to a specific deployed contract.
func NewEthtcTransactor(address common.Address, transactor bind.ContractTransactor) (*EthtcTransactor, error) {
	contract, err := bindEthtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthtcTransactor{contract: contract}, nil
}

// NewEthtcFilterer creates a new log filterer instance of Ethtc, bound to a specific deployed contract.
func NewEthtcFilterer(address common.Address, filterer bind.ContractFilterer) (*EthtcFilterer, error) {
	contract, err := bindEthtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthtcFilterer{contract: contract}, nil
}

// bindEthtc binds a generic wrapper to an already deployed contract.
func bindEthtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethtc *EthtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethtc.Contract.EthtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethtc *EthtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethtc.Contract.EthtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethtc *EthtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethtc.Contract.EthtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethtc *EthtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethtc *EthtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethtc *EthtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethtc.Contract.contract.Transact(opts, method, params...)
}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_Ethtc *EthtcCaller) FIELDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "FIELD_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_Ethtc *EthtcSession) FIELDSIZE() (*big.Int, error) {
	return _Ethtc.Contract.FIELDSIZE(&_Ethtc.CallOpts)
}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_Ethtc *EthtcCallerSession) FIELDSIZE() (*big.Int, error) {
	return _Ethtc.Contract.FIELDSIZE(&_Ethtc.CallOpts)
}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_Ethtc *EthtcCaller) ROOTHISTORYSIZE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "ROOT_HISTORY_SIZE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_Ethtc *EthtcSession) ROOTHISTORYSIZE() (uint32, error) {
	return _Ethtc.Contract.ROOTHISTORYSIZE(&_Ethtc.CallOpts)
}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_Ethtc *EthtcCallerSession) ROOTHISTORYSIZE() (uint32, error) {
	return _Ethtc.Contract.ROOTHISTORYSIZE(&_Ethtc.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Ethtc *EthtcCaller) ZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Ethtc *EthtcSession) ZEROVALUE() (*big.Int, error) {
	return _Ethtc.Contract.ZEROVALUE(&_Ethtc.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Ethtc *EthtcCallerSession) ZEROVALUE() (*big.Int, error) {
	return _Ethtc.Contract.ZEROVALUE(&_Ethtc.CallOpts)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Ethtc *EthtcCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Ethtc *EthtcSession) Commitments(arg0 [32]byte) (bool, error) {
	return _Ethtc.Contract.Commitments(&_Ethtc.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Ethtc *EthtcCallerSession) Commitments(arg0 [32]byte) (bool, error) {
	return _Ethtc.Contract.Commitments(&_Ethtc.CallOpts, arg0)
}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_Ethtc *EthtcCaller) CurrentRootIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "currentRootIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_Ethtc *EthtcSession) CurrentRootIndex() (uint32, error) {
	return _Ethtc.Contract.CurrentRootIndex(&_Ethtc.CallOpts)
}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_Ethtc *EthtcCallerSession) CurrentRootIndex() (uint32, error) {
	return _Ethtc.Contract.CurrentRootIndex(&_Ethtc.CallOpts)
}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_Ethtc *EthtcCaller) Denomination(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "denomination")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_Ethtc *EthtcSession) Denomination() (*big.Int, error) {
	return _Ethtc.Contract.Denomination(&_Ethtc.CallOpts)
}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_Ethtc *EthtcCallerSession) Denomination() (*big.Int, error) {
	return _Ethtc.Contract.Denomination(&_Ethtc.CallOpts)
}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_Ethtc *EthtcCaller) FilledSubtrees(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "filledSubtrees", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_Ethtc *EthtcSession) FilledSubtrees(arg0 *big.Int) ([32]byte, error) {
	return _Ethtc.Contract.FilledSubtrees(&_Ethtc.CallOpts, arg0)
}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_Ethtc *EthtcCallerSession) FilledSubtrees(arg0 *big.Int) ([32]byte, error) {
	return _Ethtc.Contract.FilledSubtrees(&_Ethtc.CallOpts, arg0)
}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_Ethtc *EthtcCaller) GetLastRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "getLastRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_Ethtc *EthtcSession) GetLastRoot() ([32]byte, error) {
	return _Ethtc.Contract.GetLastRoot(&_Ethtc.CallOpts)
}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_Ethtc *EthtcCallerSession) GetLastRoot() ([32]byte, error) {
	return _Ethtc.Contract.GetLastRoot(&_Ethtc.CallOpts)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Ethtc *EthtcCaller) HashLeftRight(opts *bind.CallOpts, _hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "hashLeftRight", _hasher, _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Ethtc *EthtcSession) HashLeftRight(_hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Ethtc.Contract.HashLeftRight(&_Ethtc.CallOpts, _hasher, _left, _right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Ethtc *EthtcCallerSession) HashLeftRight(_hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Ethtc.Contract.HashLeftRight(&_Ethtc.CallOpts, _hasher, _left, _right)
}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_Ethtc *EthtcCaller) Hasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "hasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_Ethtc *EthtcSession) Hasher() (common.Address, error) {
	return _Ethtc.Contract.Hasher(&_Ethtc.CallOpts)
}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_Ethtc *EthtcCallerSession) Hasher() (common.Address, error) {
	return _Ethtc.Contract.Hasher(&_Ethtc.CallOpts)
}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_Ethtc *EthtcCaller) IsKnownRoot(opts *bind.CallOpts, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "isKnownRoot", _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_Ethtc *EthtcSession) IsKnownRoot(_root [32]byte) (bool, error) {
	return _Ethtc.Contract.IsKnownRoot(&_Ethtc.CallOpts, _root)
}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_Ethtc *EthtcCallerSession) IsKnownRoot(_root [32]byte) (bool, error) {
	return _Ethtc.Contract.IsKnownRoot(&_Ethtc.CallOpts, _root)
}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_Ethtc *EthtcCaller) IsSpent(opts *bind.CallOpts, _nullifierHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "isSpent", _nullifierHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_Ethtc *EthtcSession) IsSpent(_nullifierHash [32]byte) (bool, error) {
	return _Ethtc.Contract.IsSpent(&_Ethtc.CallOpts, _nullifierHash)
}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_Ethtc *EthtcCallerSession) IsSpent(_nullifierHash [32]byte) (bool, error) {
	return _Ethtc.Contract.IsSpent(&_Ethtc.CallOpts, _nullifierHash)
}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_Ethtc *EthtcCaller) IsSpentArray(opts *bind.CallOpts, _nullifierHashes [][32]byte) ([]bool, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "isSpentArray", _nullifierHashes)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_Ethtc *EthtcSession) IsSpentArray(_nullifierHashes [][32]byte) ([]bool, error) {
	return _Ethtc.Contract.IsSpentArray(&_Ethtc.CallOpts, _nullifierHashes)
}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_Ethtc *EthtcCallerSession) IsSpentArray(_nullifierHashes [][32]byte) ([]bool, error) {
	return _Ethtc.Contract.IsSpentArray(&_Ethtc.CallOpts, _nullifierHashes)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_Ethtc *EthtcCaller) Levels(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "levels")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_Ethtc *EthtcSession) Levels() (uint32, error) {
	return _Ethtc.Contract.Levels(&_Ethtc.CallOpts)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_Ethtc *EthtcCallerSession) Levels() (uint32, error) {
	return _Ethtc.Contract.Levels(&_Ethtc.CallOpts)
}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_Ethtc *EthtcCaller) NextIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "nextIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_Ethtc *EthtcSession) NextIndex() (uint32, error) {
	return _Ethtc.Contract.NextIndex(&_Ethtc.CallOpts)
}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_Ethtc *EthtcCallerSession) NextIndex() (uint32, error) {
	return _Ethtc.Contract.NextIndex(&_Ethtc.CallOpts)
}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_Ethtc *EthtcCaller) NullifierHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "nullifierHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_Ethtc *EthtcSession) NullifierHashes(arg0 [32]byte) (bool, error) {
	return _Ethtc.Contract.NullifierHashes(&_Ethtc.CallOpts, arg0)
}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_Ethtc *EthtcCallerSession) NullifierHashes(arg0 [32]byte) (bool, error) {
	return _Ethtc.Contract.NullifierHashes(&_Ethtc.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_Ethtc *EthtcCaller) Roots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_Ethtc *EthtcSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _Ethtc.Contract.Roots(&_Ethtc.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_Ethtc *EthtcCallerSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _Ethtc.Contract.Roots(&_Ethtc.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Ethtc *EthtcCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Ethtc *EthtcSession) Verifier() (common.Address, error) {
	return _Ethtc.Contract.Verifier(&_Ethtc.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Ethtc *EthtcCallerSession) Verifier() (common.Address, error) {
	return _Ethtc.Contract.Verifier(&_Ethtc.CallOpts)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_Ethtc *EthtcCaller) Zeros(opts *bind.CallOpts, i *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Ethtc.contract.Call(opts, &out, "zeros", i)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_Ethtc *EthtcSession) Zeros(i *big.Int) ([32]byte, error) {
	return _Ethtc.Contract.Zeros(&_Ethtc.CallOpts, i)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_Ethtc *EthtcCallerSession) Zeros(i *big.Int) ([32]byte, error) {
	return _Ethtc.Contract.Zeros(&_Ethtc.CallOpts, i)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_Ethtc *EthtcTransactor) Deposit(opts *bind.TransactOpts, _commitment [32]byte) (*types.Transaction, error) {
	return _Ethtc.contract.Transact(opts, "deposit", _commitment)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_Ethtc *EthtcSession) Deposit(_commitment [32]byte) (*types.Transaction, error) {
	return _Ethtc.Contract.Deposit(&_Ethtc.TransactOpts, _commitment)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_Ethtc *EthtcTransactorSession) Deposit(_commitment [32]byte) (*types.Transaction, error) {
	return _Ethtc.Contract.Deposit(&_Ethtc.TransactOpts, _commitment)
}

// Withdraw is a paid mutator transaction binding the contract method 0x21a0adb6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient, address _relayer, uint256 _fee, uint256 _refund) payable returns()
func (_Ethtc *EthtcTransactor) Withdraw(opts *bind.TransactOpts, _proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address, _relayer common.Address, _fee *big.Int, _refund *big.Int) (*types.Transaction, error) {
	return _Ethtc.contract.Transact(opts, "withdraw", _proof, _root, _nullifierHash, _recipient, _relayer, _fee, _refund)
}

// Withdraw is a paid mutator transaction binding the contract method 0x21a0adb6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient, address _relayer, uint256 _fee, uint256 _refund) payable returns()
func (_Ethtc *EthtcSession) Withdraw(_proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address, _relayer common.Address, _fee *big.Int, _refund *big.Int) (*types.Transaction, error) {
	return _Ethtc.Contract.Withdraw(&_Ethtc.TransactOpts, _proof, _root, _nullifierHash, _recipient, _relayer, _fee, _refund)
}

// Withdraw is a paid mutator transaction binding the contract method 0x21a0adb6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient, address _relayer, uint256 _fee, uint256 _refund) payable returns()
func (_Ethtc *EthtcTransactorSession) Withdraw(_proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address, _relayer common.Address, _fee *big.Int, _refund *big.Int) (*types.Transaction, error) {
	return _Ethtc.Contract.Withdraw(&_Ethtc.TransactOpts, _proof, _root, _nullifierHash, _recipient, _relayer, _fee, _refund)
}

// EthtcDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Ethtc contract.
type EthtcDepositIterator struct {
	Event *EthtcDeposit // Event containing the contract specifics and raw log

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
func (it *EthtcDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthtcDeposit)
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
		it.Event = new(EthtcDeposit)
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
func (it *EthtcDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthtcDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthtcDeposit represents a Deposit event raised by the Ethtc contract.
type EthtcDeposit struct {
	Commitment [32]byte
	LeafIndex  uint32
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xa945e51eec50ab98c161376f0db4cf2aeba3ec92755fe2fcd388bdbbb80ff196.
//
// Solidity: event Deposit(bytes32 indexed commitment, uint32 leafIndex, uint256 timestamp)
func (_Ethtc *EthtcFilterer) FilterDeposit(opts *bind.FilterOpts, commitment [][32]byte) (*EthtcDepositIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _Ethtc.contract.FilterLogs(opts, "Deposit", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &EthtcDepositIterator{contract: _Ethtc.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xa945e51eec50ab98c161376f0db4cf2aeba3ec92755fe2fcd388bdbbb80ff196.
//
// Solidity: event Deposit(bytes32 indexed commitment, uint32 leafIndex, uint256 timestamp)
func (_Ethtc *EthtcFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EthtcDeposit, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _Ethtc.contract.WatchLogs(opts, "Deposit", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthtcDeposit)
				if err := _Ethtc.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Ethtc *EthtcFilterer) ParseDeposit(log types.Log) (*EthtcDeposit, error) {
	event := new(EthtcDeposit)
	if err := _Ethtc.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthtcWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Ethtc contract.
type EthtcWithdrawalIterator struct {
	Event *EthtcWithdrawal // Event containing the contract specifics and raw log

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
func (it *EthtcWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthtcWithdrawal)
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
		it.Event = new(EthtcWithdrawal)
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
func (it *EthtcWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthtcWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthtcWithdrawal represents a Withdrawal event raised by the Ethtc contract.
type EthtcWithdrawal struct {
	To            common.Address
	NullifierHash [32]byte
	Relayer       common.Address
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xe9e508bad6d4c3227e881ca19068f099da81b5164dd6d62b2eaf1e8bc6c34931.
//
// Solidity: event Withdrawal(address to, bytes32 nullifierHash, address indexed relayer, uint256 fee)
func (_Ethtc *EthtcFilterer) FilterWithdrawal(opts *bind.FilterOpts, relayer []common.Address) (*EthtcWithdrawalIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Ethtc.contract.FilterLogs(opts, "Withdrawal", relayerRule)
	if err != nil {
		return nil, err
	}
	return &EthtcWithdrawalIterator{contract: _Ethtc.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xe9e508bad6d4c3227e881ca19068f099da81b5164dd6d62b2eaf1e8bc6c34931.
//
// Solidity: event Withdrawal(address to, bytes32 nullifierHash, address indexed relayer, uint256 fee)
func (_Ethtc *EthtcFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *EthtcWithdrawal, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Ethtc.contract.WatchLogs(opts, "Withdrawal", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthtcWithdrawal)
				if err := _Ethtc.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_Ethtc *EthtcFilterer) ParseWithdrawal(log types.Log) (*EthtcWithdrawal, error) {
	event := new(EthtcWithdrawal)
	if err := _Ethtc.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

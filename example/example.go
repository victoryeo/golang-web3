// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package example

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
)

// ExampleMetaData contains all meta data concerning the Example contract.
var ExampleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"Hello\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"a\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"b\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exEncode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"hello\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExampleABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleMetaData.ABI instead.
var ExampleABI = ExampleMetaData.ABI

// Example is an auto generated Go binding around an Ethereum contract.
type Example struct {
	ExampleCaller     // Read-only binding to the contract
	ExampleTransactor // Write-only binding to the contract
	ExampleFilterer   // Log filterer for contract events
}

// ExampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleSession struct {
	Contract     *Example          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleCallerSession struct {
	Contract *ExampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ExampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleTransactorSession struct {
	Contract     *ExampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleRaw struct {
	Contract *Example // Generic contract binding to access the raw methods on
}

// ExampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleCallerRaw struct {
	Contract *ExampleCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleTransactorRaw struct {
	Contract *ExampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExample creates a new instance of Example, bound to a specific deployed contract.
func NewExample(address common.Address, backend bind.ContractBackend) (*Example, error) {
	contract, err := bindExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Example{ExampleCaller: ExampleCaller{contract: contract}, ExampleTransactor: ExampleTransactor{contract: contract}, ExampleFilterer: ExampleFilterer{contract: contract}}, nil
}

// NewExampleCaller creates a new read-only instance of Example, bound to a specific deployed contract.
func NewExampleCaller(address common.Address, caller bind.ContractCaller) (*ExampleCaller, error) {
	contract, err := bindExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleCaller{contract: contract}, nil
}

// NewExampleTransactor creates a new write-only instance of Example, bound to a specific deployed contract.
func NewExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleTransactor, error) {
	contract, err := bindExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleTransactor{contract: contract}, nil
}

// NewExampleFilterer creates a new log filterer instance of Example, bound to a specific deployed contract.
func NewExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleFilterer, error) {
	contract, err := bindExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleFilterer{contract: contract}, nil
}

// bindExample binds a generic wrapper to an already deployed contract.
func bindExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExampleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Example *ExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Example.Contract.ExampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Example *ExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Example.Contract.ExampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Example *ExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Example.Contract.ExampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Example *ExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Example.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Example *ExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Example.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Example *ExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Example.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(string)
func (_Example *ExampleCaller) A(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Example.contract.Call(opts, &out, "a")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(string)
func (_Example *ExampleSession) A() (string, error) {
	return _Example.Contract.A(&_Example.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() view returns(string)
func (_Example *ExampleCallerSession) A() (string, error) {
	return _Example.Contract.A(&_Example.CallOpts)
}

// B is a free data retrieval call binding the contract method 0x4df7e3d0.
//
// Solidity: function b() view returns(string)
func (_Example *ExampleCaller) B(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Example.contract.Call(opts, &out, "b")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// B is a free data retrieval call binding the contract method 0x4df7e3d0.
//
// Solidity: function b() view returns(string)
func (_Example *ExampleSession) B() (string, error) {
	return _Example.Contract.B(&_Example.CallOpts)
}

// B is a free data retrieval call binding the contract method 0x4df7e3d0.
//
// Solidity: function b() view returns(string)
func (_Example *ExampleCallerSession) B() (string, error) {
	return _Example.Contract.B(&_Example.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Example *ExampleCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Example.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Example *ExampleSession) Count() (*big.Int, error) {
	return _Example.Contract.Count(&_Example.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Example *ExampleCallerSession) Count() (*big.Int, error) {
	return _Example.Contract.Count(&_Example.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Example *ExampleCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Example.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Example *ExampleSession) Name() (string, error) {
	return _Example.Contract.Name(&_Example.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Example *ExampleCallerSession) Name() (string, error) {
	return _Example.Contract.Name(&_Example.CallOpts)
}

// ExEncode is a paid mutator transaction binding the contract method 0xeccb8553.
//
// Solidity: function exEncode() returns()
func (_Example *ExampleTransactor) ExEncode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Example.contract.Transact(opts, "exEncode")
}

// ExEncode is a paid mutator transaction binding the contract method 0xeccb8553.
//
// Solidity: function exEncode() returns()
func (_Example *ExampleSession) ExEncode() (*types.Transaction, error) {
	return _Example.Contract.ExEncode(&_Example.TransactOpts)
}

// ExEncode is a paid mutator transaction binding the contract method 0xeccb8553.
//
// Solidity: function exEncode() returns()
func (_Example *ExampleTransactorSession) ExEncode() (*types.Transaction, error) {
	return _Example.Contract.ExEncode(&_Example.TransactOpts)
}

// Hello is a paid mutator transaction binding the contract method 0xa777d0dc.
//
// Solidity: function hello(string _name) returns()
func (_Example *ExampleTransactor) Hello(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Example.contract.Transact(opts, "hello", _name)
}

// Hello is a paid mutator transaction binding the contract method 0xa777d0dc.
//
// Solidity: function hello(string _name) returns()
func (_Example *ExampleSession) Hello(_name string) (*types.Transaction, error) {
	return _Example.Contract.Hello(&_Example.TransactOpts, _name)
}

// Hello is a paid mutator transaction binding the contract method 0xa777d0dc.
//
// Solidity: function hello(string _name) returns()
func (_Example *ExampleTransactorSession) Hello(_name string) (*types.Transaction, error) {
	return _Example.Contract.Hello(&_Example.TransactOpts, _name)
}

// ExampleHelloIterator is returned from FilterHello and is used to iterate over the raw logs and unpacked data for Hello events raised by the Example contract.
type ExampleHelloIterator struct {
	Event *ExampleHello // Event containing the contract specifics and raw log

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
func (it *ExampleHelloIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleHello)
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
		it.Event = new(ExampleHello)
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
func (it *ExampleHelloIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleHelloIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleHello represents a Hello event raised by the Example contract.
type ExampleHello struct {
	Name  string
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHello is a free log retrieval operation binding the contract event 0x410c78c042da0f8716528df5d22e9a41f525de79edb2377e7a1355924b80e8a0.
//
// Solidity: event Hello(string name, uint256 count)
func (_Example *ExampleFilterer) FilterHello(opts *bind.FilterOpts) (*ExampleHelloIterator, error) {

	logs, sub, err := _Example.contract.FilterLogs(opts, "Hello")
	if err != nil {
		return nil, err
	}
	return &ExampleHelloIterator{contract: _Example.contract, event: "Hello", logs: logs, sub: sub}, nil
}

// WatchHello is a free log subscription operation binding the contract event 0x410c78c042da0f8716528df5d22e9a41f525de79edb2377e7a1355924b80e8a0.
//
// Solidity: event Hello(string name, uint256 count)
func (_Example *ExampleFilterer) WatchHello(opts *bind.WatchOpts, sink chan<- *ExampleHello) (event.Subscription, error) {

	logs, sub, err := _Example.contract.WatchLogs(opts, "Hello")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleHello)
				if err := _Example.contract.UnpackLog(event, "Hello", log); err != nil {
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

// ParseHello is a log parse operation binding the contract event 0x410c78c042da0f8716528df5d22e9a41f525de79edb2377e7a1355924b80e8a0.
//
// Solidity: event Hello(string name, uint256 count)
func (_Example *ExampleFilterer) ParseHello(log types.Log) (*ExampleHello, error) {
	event := new(ExampleHello)
	if err := _Example.contract.UnpackLog(event, "Hello", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

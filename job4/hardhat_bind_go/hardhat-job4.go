// 代码由 abigen 自动生成 - 请勿手动编辑。
// 本文件是自动生成的合约绑定，任何手动修改都可能在重新生成时丢失。

package hardhat_bind_go

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// 强制引用这些导入，避免它们未被使用时报错。
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
	_ = time.Tick
	_ = context.Background
)

// HardhatBindGoMetaData 包含 HardhatBindGo 合约的全部元数据。
var HardhatBindGoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CountChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103708061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80632baeceb71461004e578063a87d942c14610058578063d09de08a14610076578063d826f88f14610080575b5f5ffd5b61005661008a565b005b610060610136565b60405161006d9190610216565b60405180910390f35b61007e61013e565b005b6100886101a7565b005b5f5f54116100cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100c490610289565b60405180910390fd5b60015f5f8282546100de91906102d4565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f5460405161012c9190610216565b60405180910390a2565b5f5f54905090565b60015f5f82825461014f9190610307565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f5460405161019d9190610216565b60405180910390a2565b5f5f819055503373ffffffffffffffffffffffffffffffffffffffff167fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f546040516101f49190610216565b60405180910390a2565b5f819050919050565b610210816101fe565b82525050565b5f6020820190506102295f830184610207565b92915050565b5f82825260208201905092915050565b7f636f756e74206973207a65726f000000000000000000000000000000000000005f82015250565b5f610273600d8361022f565b915061027e8261023f565b602082019050919050565b5f6020820190508181035f8301526102a081610267565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6102de826101fe565b91506102e9836101fe565b9250828203905081811115610301576103006102a7565b5b92915050565b5f610311826101fe565b915061031c836101fe565b9250828201905080821115610334576103336102a7565b5b9291505056fea264697066735822122077a0372b02ecac0ebe37ce7b69d1da925e3cf3078ddc1eddc7c023f9bcd0cbfb64736f6c63430008240033",
}

// HardhatBindGoABI 是用于生成本绑定的输入 ABI。
// 已废弃：请改用 HardhatBindGoMetaData.ABI。
var HardhatBindGoABI = HardhatBindGoMetaData.ABI

// HardhatBindGoBin 是用于部署新合约的编译字节码。
// 已废弃：请改用 HardhatBindGoMetaData.Bin。
var HardhatBindGoBin = HardhatBindGoMetaData.Bin

// DeployHardhatBindGo 部署一个新的以太坊合约，并把一个 HardhatBindGo 实例与之绑定。
func DeployHardhatBindGo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HardhatBindGo, error) {
	parsed, err := HardhatBindGoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HardhatBindGoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HardhatBindGo{HardhatBindGoCaller: HardhatBindGoCaller{contract: contract}, HardhatBindGoTransactor: HardhatBindGoTransactor{contract: contract}, HardhatBindGoFilterer: HardhatBindGoFilterer{contract: contract}}, nil
}

// HardhatBindGo 是围绕以太坊合约的自动生成 Go 绑定。
type HardhatBindGo struct {
	HardhatBindGoCaller     // Read-only binding to the contract
	HardhatBindGoTransactor // Write-only binding to the contract
	HardhatBindGoFilterer   // Log filterer for contract events
}

// HardhatBindGoCaller 是围绕以太坊合约的自动生成只读 Go 绑定。
type HardhatBindGoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HardhatBindGoTransactor 是围绕以太坊合约的自动生成只写 Go 绑定。
type HardhatBindGoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HardhatBindGoFilterer 是针对以太坊合约事件的自动生成日志过滤 Go 绑定。
type HardhatBindGoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HardhatBindGoSession 是围绕以太坊合约的自动生成 Go 绑定，
// 并预先设置了调用与交易选项。
type HardhatBindGoSession struct {
	Contract     *HardhatBindGo    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HardhatBindGoCallerSession 是围绕以太坊合约的自动生成只读 Go 绑定，
// 并预先设置了调用选项。
type HardhatBindGoCallerSession struct {
	Contract *HardhatBindGoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HardhatBindGoTransactorSession 是围绕以太坊合约的自动生成只写 Go 绑定，
// 并预先设置了交易选项。
type HardhatBindGoTransactorSession struct {
	Contract     *HardhatBindGoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HardhatBindGoRaw 是围绕以太坊合约的自动生成底层 Go 绑定。
type HardhatBindGoRaw struct {
	Contract *HardhatBindGo // Generic contract binding to access the raw methods on
}

// HardhatBindGoCallerRaw 是围绕以太坊合约的自动生成底层只读 Go 绑定。
type HardhatBindGoCallerRaw struct {
	Contract *HardhatBindGoCaller // Generic read-only contract binding to access the raw methods on
}

// HardhatBindGoTransactorRaw 是围绕以太坊合约的自动生成底层只写 Go 绑定。
type HardhatBindGoTransactorRaw struct {
	Contract *HardhatBindGoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHardhatBindGo 创建一个新的 HardhatBindGo 实例，并将其绑定到指定的已部署合约上。
func NewHardhatBindGo(address common.Address, backend bind.ContractBackend) (*HardhatBindGo, error) {
	contract, err := bindHardhatBindGo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HardhatBindGo{HardhatBindGoCaller: HardhatBindGoCaller{contract: contract}, HardhatBindGoTransactor: HardhatBindGoTransactor{contract: contract}, HardhatBindGoFilterer: HardhatBindGoFilterer{contract: contract}}, nil
}

// NewHardhatBindGoCaller 创建一个新的只读 HardhatBindGo 实例，并将其绑定到指定的已部署合约上。
func NewHardhatBindGoCaller(address common.Address, caller bind.ContractCaller) (*HardhatBindGoCaller, error) {
	contract, err := bindHardhatBindGo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HardhatBindGoCaller{contract: contract}, nil
}

// NewHardhatBindGoTransactor 创建一个新的只写 HardhatBindGo 实例，并将其绑定到指定的已部署合约上。
func NewHardhatBindGoTransactor(address common.Address, transactor bind.ContractTransactor) (*HardhatBindGoTransactor, error) {
	contract, err := bindHardhatBindGo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HardhatBindGoTransactor{contract: contract}, nil
}

// NewHardhatBindGoFilterer 创建一个新的日志过滤 HardhatBindGo 实例，并将其绑定到指定的已部署合约上。
func NewHardhatBindGoFilterer(address common.Address, filterer bind.ContractFilterer) (*HardhatBindGoFilterer, error) {
	contract, err := bindHardhatBindGo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HardhatBindGoFilterer{contract: contract}, nil
}

// bindHardhatBindGo 将通用包装器绑定到一个已部署的合约上。
func bindHardhatBindGo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HardhatBindGoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call 以 params 作为输入参数调用合约的（只读）方法，
// 并把输出写入 result。对于简单返回值，结果是单个字段；
// 匿名多返回值时是接口切片，命名返回值时则是结构体。
func (_HardhatBindGo *HardhatBindGoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HardhatBindGo.Contract.HardhatBindGoCaller.contract.Call(opts, result, method, params...)
}

// Transfer 发起一笔普通交易向合约转账，如果合约定义了默认方法，
// 则会调用该默认方法。
func (_HardhatBindGo *HardhatBindGoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HardhatBindGo.Contract.HardhatBindGoTransactor.contract.Transfer(opts)
}

// Transact 以 params 作为输入参数调用会改变状态的（付费）合约方法。
func (_HardhatBindGo *HardhatBindGoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HardhatBindGo.Contract.HardhatBindGoTransactor.contract.Transact(opts, method, params...)
}

// Call 以 params 作为输入参数调用合约的（只读）方法，
// 并把输出写入 result。对于简单返回值，结果是单个字段；
// 匿名多返回值时是接口切片，命名返回值时则是结构体。
func (_HardhatBindGo *HardhatBindGoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HardhatBindGo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer 发起一笔普通交易向合约转账，如果合约定义了默认方法，
// 则会调用该默认方法。
func (_HardhatBindGo *HardhatBindGoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HardhatBindGo.Contract.contract.Transfer(opts)
}

// Transact 以 params 作为输入参数调用会改变状态的（付费）合约方法。
func (_HardhatBindGo *HardhatBindGoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HardhatBindGo.Contract.contract.Transact(opts, method, params...)
}

// GetCount 是对合约方法 0xa87d942c 的免费数据读取调用绑定。
//
// Solidity: function getCount() view returns(uint256)
func (_HardhatBindGo *HardhatBindGoCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HardhatBindGo.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount 是对合约方法 0xa87d942c 的免费数据读取调用绑定。
//
// Solidity: function getCount() view returns(uint256)
func (_HardhatBindGo *HardhatBindGoSession) GetCount() (*big.Int, error) {
	return _HardhatBindGo.Contract.GetCount(&_HardhatBindGo.CallOpts)
}

// GetCount 是对合约方法 0xa87d942c 的免费数据读取调用绑定。
//
// Solidity: function getCount() view returns(uint256)
func (_HardhatBindGo *HardhatBindGoCallerSession) GetCount() (*big.Int, error) {
	return _HardhatBindGo.Contract.GetCount(&_HardhatBindGo.CallOpts)
}

// Decrement 是对合约方法 0x2baeceb7 的付费状态修改交易绑定。
//
// Solidity: function decrement() returns()
func (_HardhatBindGo *HardhatBindGoTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HardhatBindGo.contract.Transact(opts, "decrement")
}

// Decrement 是对合约方法 0x2baeceb7 的付费状态修改交易绑定。
//
// Solidity: function decrement() returns()
func (_HardhatBindGo *HardhatBindGoSession) Decrement() (*types.Transaction, error) {
	return _HardhatBindGo.Contract.Decrement(&_HardhatBindGo.TransactOpts)
}

// Decrement 是对合约方法 0x2baeceb7 的付费状态修改交易绑定。
//
// Solidity: function decrement() returns()
func (_HardhatBindGo *HardhatBindGoTransactorSession) Decrement() (*types.Transaction, error) {
	return _HardhatBindGo.Contract.Decrement(&_HardhatBindGo.TransactOpts)
}

// Increment 是对合约方法 0xd09de08a 的付费状态修改交易绑定。
//
// Solidity: function increment() returns()
func (_HardhatBindGo *HardhatBindGoTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HardhatBindGo.contract.Transact(opts, "increment")
}

// Increment 是对合约方法 0xd09de08a 的付费状态修改交易绑定。
//
// Solidity: function increment() returns()
func (_HardhatBindGo *HardhatBindGoSession) Increment() (*types.Transaction, error) {
	return _HardhatBindGo.Contract.Increment(&_HardhatBindGo.TransactOpts)
}

// Increment 是对合约方法 0xd09de08a 的付费状态修改交易绑定。
//
// Solidity: function increment() returns()
func (_HardhatBindGo *HardhatBindGoTransactorSession) Increment() (*types.Transaction, error) {
	return _HardhatBindGo.Contract.Increment(&_HardhatBindGo.TransactOpts)
}

// Reset 是对合约方法 0xd826f88f 的付费状态修改交易绑定。
//
// Solidity: function reset() returns()
func (_HardhatBindGo *HardhatBindGoTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HardhatBindGo.contract.Transact(opts, "reset")
}

// Reset 是对合约方法 0xd826f88f 的付费状态修改交易绑定。
//
// Solidity: function reset() returns()
func (_HardhatBindGo *HardhatBindGoSession) Reset() (*types.Transaction, error) {
	return _HardhatBindGo.Contract.Reset(&_HardhatBindGo.TransactOpts)
}

// Reset 是对合约方法 0xd826f88f 的付费状态修改交易绑定。
//
// Solidity: function reset() returns()
func (_HardhatBindGo *HardhatBindGoTransactorSession) Reset() (*types.Transaction, error) {
	return _HardhatBindGo.Contract.Reset(&_HardhatBindGo.TransactOpts)
}

// HardhatBindGoCountChangedIterator 由 FilterCountChanged 返回，用于遍历 HardhatBindGo 合约发出的 CountChanged 事件的原始日志和解包后的数据。
type HardhatBindGoCountChangedIterator struct {
	Event *HardhatBindGoCountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next 把迭代器推进到下一个事件，并返回是否还有更多事件。
// 当发生获取或解析错误时同样返回 false，
// 具体的失败原因可以通过 Error() 查询。
func (it *HardhatBindGoCountChangedIterator) Next() bool {
	// 如果迭代器已失败，则停止迭代。
	if it.fail != nil {
		return false
	}
	// 如果迭代器已完成，就直接交付当前可用的日志。
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HardhatBindGoCountChanged)
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
	// 迭代仍在进行中，等待数据或错误事件。
	select {
	case log := <-it.logs:
		it.Event = new(HardhatBindGoCountChanged)
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

// Error 返回过滤过程中出现的任何获取或解析错误。
func (it *HardhatBindGoCountChangedIterator) Error() error {
	return it.fail
}

// Close 终止迭代过程，释放所有尚待处理的
// 底层资源。
func (it *HardhatBindGoCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HardhatBindGoCountChanged 表示由 HardhatBindGo 合约发出的 CountChanged 事件。
type HardhatBindGoCountChanged struct {
	NewCount *big.Int
	Caller   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountChanged 是对合约事件 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82 的免费日志检索操作绑定。
//
// Solidity: event CountChanged(uint256 newCount, address indexed caller)
func (_HardhatBindGo *HardhatBindGoFilterer) FilterCountChanged(opts *bind.FilterOpts, caller []common.Address) (*HardhatBindGoCountChangedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _HardhatBindGo.contract.FilterLogs(opts, "CountChanged", callerRule)
	if err != nil {
		return nil, err
	}
	return &HardhatBindGoCountChangedIterator{contract: _HardhatBindGo.contract, event: "CountChanged", logs: logs, sub: sub}, nil
}

// WatchCountChanged 是对合约事件 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82 的免费日志订阅操作绑定。
//
// Solidity: event CountChanged(uint256 newCount, address indexed caller)
func (_HardhatBindGo *HardhatBindGoFilterer) WatchCountChanged(opts *bind.WatchOpts, sink chan<- *HardhatBindGoCountChanged, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _HardhatBindGo.contract.WatchLogs(opts, "CountChanged", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// 收到新日志，解析该事件并转发给用户。
				event := new(HardhatBindGoCountChanged)
				if err := _HardhatBindGo.contract.UnpackLog(event, "CountChanged", log); err != nil {
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

// ParseCountChanged 是对合约事件 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82 的日志解析操作绑定。
//
// Solidity: event CountChanged(uint256 newCount, address indexed caller)
func (_HardhatBindGo *HardhatBindGoFilterer) ParseCountChanged(log types.Log) (*HardhatBindGoCountChanged, error) {
	event := new(HardhatBindGoCountChanged)
	if err := _HardhatBindGo.contract.UnpackLog(event, "CountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

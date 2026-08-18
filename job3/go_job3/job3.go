// 代码自动生成 - 请勿手动编辑。
// 此文件是自动生成的绑定代码，任何手动修改都将丢失。

package go_job3

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

// 引用导入以避免未使用时的编译错误。
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

// GoJob3MetaData 包含 GoJob3 合约的所有元数据。
var GoJob3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CountChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103708061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80632baeceb71461004e578063a87d942c14610058578063d09de08a14610076578063d826f88f14610080575b5f5ffd5b61005661008a565b005b610060610136565b60405161006d9190610216565b60405180910390f35b61007e61013e565b005b6100886101a7565b005b5f5f54116100cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100c490610289565b60405180910390fd5b60015f5f8282546100de91906102d4565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f5460405161012c9190610216565b60405180910390a2565b5f5f54905090565b60015f5f82825461014f9190610307565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f5460405161019d9190610216565b60405180910390a2565b5f5f819055503373ffffffffffffffffffffffffffffffffffffffff167fb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a825f546040516101f49190610216565b60405180910390a2565b5f819050919050565b610210816101fe565b82525050565b5f6020820190506102295f830184610207565b92915050565b5f82825260208201905092915050565b7f636f756e74206973207a65726f000000000000000000000000000000000000005f82015250565b5f610273600d8361022f565b915061027e8261023f565b602082019050919050565b5f6020820190508181035f8301526102a081610267565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6102de826101fe565b91506102e9836101fe565b9250828203905081811115610301576103006102a7565b5b92915050565b5f610311826101fe565b915061031c836101fe565b9250828201905080821115610334576103336102a7565b5b9291505056fea264697066735822122090dc92e31fe1b8bedf8ac3c6886d300d7a8260385dea4ae34d63c900bdca0da164736f6c63430008240033",
}

// GoJob3ABI 是用于生成绑定的输入 ABI。
// 已废弃：请使用 GoJob3MetaData.ABI 代替。
var GoJob3ABI = GoJob3MetaData.ABI

// GoJob3Bin 是用于部署新合约的编译字节码。
// 已废弃：请使用 GoJob3MetaData.Bin 代替。
var GoJob3Bin = GoJob3MetaData.Bin

// DeployGoJob3 部署一个新的以太坊合约，并将 GoJob3 实例绑定到它。
func DeployGoJob3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GoJob3, error) {
	parsed, err := GoJob3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GoJob3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GoJob3{GoJob3Caller: GoJob3Caller{contract: contract}, GoJob3Transactor: GoJob3Transactor{contract: contract}, GoJob3Filterer: GoJob3Filterer{contract: contract}}, nil
}

// GoJob3 是围绕以太坊合约自动生成的 Go 绑定。
type GoJob3 struct {
	GoJob3Caller     // 合约的只读绑定
	GoJob3Transactor // 合约的只写绑定
	GoJob3Filterer   // 合约事件的日志过滤器
}

// GoJob3Caller 是围绕以太坊合约自动生成的只读 Go 绑定。
type GoJob3Caller struct {
	contract *bind.BoundContract // 用于底层调用的通用合约包装器
}

// GoJob3Transactor 是围绕以太坊合约自动生成的只写 Go 绑定。
type GoJob3Transactor struct {
	contract *bind.BoundContract // 用于底层调用的通用合约包装器
}

// GoJob3Filterer 是围绕以太坊合约事件自动生成的日志过滤 Go 绑定。
type GoJob3Filterer struct {
	contract *bind.BoundContract // 用于底层调用的通用合约包装器
}

// GoJob3Session 是围绕以太坊合约自动生成的 Go 绑定，
// 并预设了调用和交易选项。
type GoJob3Session struct {
	Contract     *GoJob3           // 用于设置会话的通用合约绑定
	CallOpts     bind.CallOpts     // 本次会话使用的调用选项
	TransactOpts bind.TransactOpts // 本次会话使用的交易授权选项
}

// GoJob3CallerSession 是围绕以太坊合约自动生成的只读 Go 绑定，
// 并预设了调用选项。
type GoJob3CallerSession struct {
	Contract *GoJob3Caller // 用于设置会话的通用合约调用绑定
	CallOpts bind.CallOpts // 本次会话使用的调用选项
}

// GoJob3TransactorSession 是围绕以太坊合约自动生成的只写 Go 绑定，
// 并预设了交易选项。
type GoJob3TransactorSession struct {
	Contract     *GoJob3Transactor // 用于设置会话的通用合约交易绑定
	TransactOpts bind.TransactOpts // 本次会话使用的交易授权选项
}

// GoJob3Raw 是围绕以太坊合约自动生成的底层 Go 绑定。
type GoJob3Raw struct {
	Contract *GoJob3 // 用于访问原始方法的通用合约绑定
}

// GoJob3CallerRaw 是围绕以太坊合约自动生成的底层只读 Go 绑定。
type GoJob3CallerRaw struct {
	Contract *GoJob3Caller // 用于访问原始方法的只读合约绑定
}

// GoJob3TransactorRaw 是围绕以太坊合约自动生成的底层只写 Go 绑定。
type GoJob3TransactorRaw struct {
	Contract *GoJob3Transactor // 用于访问原始方法的只写合约绑定
}

// NewGoJob3 创建一个新的 GoJob3 实例，绑定到特定已部署的合约。
func NewGoJob3(address common.Address, backend bind.ContractBackend) (*GoJob3, error) {
	contract, err := bindGoJob3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoJob3{GoJob3Caller: GoJob3Caller{contract: contract}, GoJob3Transactor: GoJob3Transactor{contract: contract}, GoJob3Filterer: GoJob3Filterer{contract: contract}}, nil
}

// NewGoJob3Caller 创建一个新的只读 GoJob3 实例，绑定到特定已部署的合约。
func NewGoJob3Caller(address common.Address, caller bind.ContractCaller) (*GoJob3Caller, error) {
	contract, err := bindGoJob3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoJob3Caller{contract: contract}, nil
}

// NewGoJob3Transactor 创建一个新的只写 GoJob3 实例，绑定到特定已部署的合约。
func NewGoJob3Transactor(address common.Address, transactor bind.ContractTransactor) (*GoJob3Transactor, error) {
	contract, err := bindGoJob3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoJob3Transactor{contract: contract}, nil
}

// NewGoJob3Filterer 创建一个新的 GoJob3 日志过滤器实例，绑定到特定已部署的合约。
func NewGoJob3Filterer(address common.Address, filterer bind.ContractFilterer) (*GoJob3Filterer, error) {
	contract, err := bindGoJob3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoJob3Filterer{contract: contract}, nil
}

// bindGoJob3 将通用包装器绑定到已部署的合约。
func bindGoJob3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GoJob3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call 调用合约的（常量）方法，将 params 作为输入值，
// 并将输出设置到 result 中。结果类型可能是简单返回的单个字段、
// 匿名返回的接口切片，或命名返回的结构体。
func (_GoJob3 *GoJob3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoJob3.Contract.GoJob3Caller.contract.Call(opts, result, method, params...)
}

// Transfer 发起一个普通交易，将资金转移到合约，
// 如果合约有默认方法则调用它。
func (_GoJob3 *GoJob3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoJob3.Contract.GoJob3Transactor.contract.Transfer(opts)
}

// Transact 以 params 作为输入值调用合约的（付费）方法。
func (_GoJob3 *GoJob3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoJob3.Contract.GoJob3Transactor.contract.Transact(opts, method, params...)
}

// Call 调用合约的（常量）方法，将 params 作为输入值，
// 并将输出设置到 result 中。结果类型可能是简单返回的单个字段、
// 匿名返回的接口切片，或命名返回的结构体。
func (_GoJob3 *GoJob3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoJob3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer 发起一个普通交易，将资金转移到合约，
// 如果合约有默认方法则调用它。
func (_GoJob3 *GoJob3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoJob3.Contract.contract.Transfer(opts)
}

// Transact 以 params 作为输入值调用合约的（付费）方法。
func (_GoJob3 *GoJob3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoJob3.Contract.contract.Transact(opts, method, params...)
}

// GetCount 是绑定合约方法 0xa87d942c 的免费数据查询调用。
//
// Solidity: function getCount() view returns(uint256)
func (_GoJob3 *GoJob3Caller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GoJob3.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount 是绑定合约方法 0xa87d942c 的免费数据查询调用。
//
// Solidity: function getCount() view returns(uint256)
func (_GoJob3 *GoJob3Session) GetCount() (*big.Int, error) {
	return _GoJob3.Contract.GetCount(&_GoJob3.CallOpts)
}

// GetCount 是绑定合约方法 0xa87d942c 的免费数据查询调用。
//
// Solidity: function getCount() view returns(uint256)
func (_GoJob3 *GoJob3CallerSession) GetCount() (*big.Int, error) {
	return _GoJob3.Contract.GetCount(&_GoJob3.CallOpts)
}

// Decrement 是绑定合约方法 0x2baeceb7 的付费变更交易。
//
// Solidity: function decrement() returns()
func (_GoJob3 *GoJob3Transactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoJob3.contract.Transact(opts, "decrement")
}

// Decrement 是绑定合约方法 0x2baeceb7 的付费变更交易。
//
// Solidity: function decrement() returns()
func (_GoJob3 *GoJob3Session) Decrement() (*types.Transaction, error) {
	return _GoJob3.Contract.Decrement(&_GoJob3.TransactOpts)
}

// Decrement 是绑定合约方法 0x2baeceb7 的付费变更交易。
//
// Solidity: function decrement() returns()
func (_GoJob3 *GoJob3TransactorSession) Decrement() (*types.Transaction, error) {
	return _GoJob3.Contract.Decrement(&_GoJob3.TransactOpts)
}

// Increment 是绑定合约方法 0xd09de08a 的付费变更交易。
//
// Solidity: function increment() returns()
func (_GoJob3 *GoJob3Transactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoJob3.contract.Transact(opts, "increment")
}

// Increment 是绑定合约方法 0xd09de08a 的付费变更交易。
//
// Solidity: function increment() returns()
func (_GoJob3 *GoJob3Session) Increment() (*types.Transaction, error) {
	return _GoJob3.Contract.Increment(&_GoJob3.TransactOpts)
}

// Increment 是绑定合约方法 0xd09de08a 的付费变更交易。
//
// Solidity: function increment() returns()
func (_GoJob3 *GoJob3TransactorSession) Increment() (*types.Transaction, error) {
	return _GoJob3.Contract.Increment(&_GoJob3.TransactOpts)
}

// Reset 是绑定合约方法 0xd826f88f 的付费变更交易。
//
// Solidity: function reset() returns()
func (_GoJob3 *GoJob3Transactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoJob3.contract.Transact(opts, "reset")
}

// Reset 是绑定合约方法 0xd826f88f 的付费变更交易。
//
// Solidity: function reset() returns()
func (_GoJob3 *GoJob3Session) Reset() (*types.Transaction, error) {
	return _GoJob3.Contract.Reset(&_GoJob3.TransactOpts)
}

// Reset 是绑定合约方法 0xd826f88f 的付费变更交易。
//
// Solidity: function reset() returns()
func (_GoJob3 *GoJob3TransactorSession) Reset() (*types.Transaction, error) {
	return _GoJob3.Contract.Reset(&_GoJob3.TransactOpts)
}

// GoJob3CountChangedIterator 由 FilterCountChanged 返回，用于遍历 GoJob3 合约
// 触发的 CountChanged 事件的原始日志和已解包数据。
type GoJob3CountChangedIterator struct {
	Event *GoJob3CountChanged // 包含合约详情和原始日志的事件

	contract *bind.BoundContract // 用于解包事件数据的通用合约
	event    string              // 用于解包事件数据的事件名称

	logs chan types.Log        // 接收已找到的合约事件的日志通道
	sub  ethereum.Subscription // 用于错误、完成和终止的订阅
	done bool                  // 订阅是否已完成日志传递
	fail error                 // 发生的错误，用于停止迭代
}

// Next 将迭代器推进到下一个事件，返回是否还有更多事件。
// 如果发生检索或解析错误，则返回 false，
// 可通过 Error() 查询具体的失败原因。
func (it *GoJob3CountChangedIterator) Next() bool {
	// 如果迭代器出错，停止迭代
	if it.fail != nil {
		return false
	}
	// 如果迭代器已完成，直接传递可用的数据
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GoJob3CountChanged)
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
	// 迭代器仍在进行中，等待数据或错误事件
	select {
	case log := <-it.logs:
		it.Event = new(GoJob3CountChanged)
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

// Error 返回过滤过程中发生的任何检索或解析错误。
func (it *GoJob3CountChangedIterator) Error() error {
	return it.fail
}

// Close 终止迭代过程，释放所有待处理的底层资源。
func (it *GoJob3CountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GoJob3CountChanged 表示 GoJob3 合约触发的 CountChanged 事件。
type GoJob3CountChanged struct {
	NewCount *big.Int
	Caller   common.Address
	Raw      types.Log // 区块链特定的上下文信息
}

// FilterCountChanged 是绑定合约事件 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82 的免费日志检索操作。
//
// Solidity: event CountChanged(uint256 newCount, address indexed caller)
func (_GoJob3 *GoJob3Filterer) FilterCountChanged(opts *bind.FilterOpts, caller []common.Address) (*GoJob3CountChangedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _GoJob3.contract.FilterLogs(opts, "CountChanged", callerRule)
	if err != nil {
		return nil, err
	}
	return &GoJob3CountChangedIterator{contract: _GoJob3.contract, event: "CountChanged", logs: logs, sub: sub}, nil
}

// WatchCountChanged 是绑定合约事件 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82 的免费日志订阅操作。
//
// Solidity: event CountChanged(uint256 newCount, address indexed caller)
func (_GoJob3 *GoJob3Filterer) WatchCountChanged(opts *bind.WatchOpts, sink chan<- *GoJob3CountChanged, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _GoJob3.contract.WatchLogs(opts, "CountChanged", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// 收到新日志，解析事件并转发给用户
				event := new(GoJob3CountChanged)
				if err := _GoJob3.contract.UnpackLog(event, "CountChanged", log); err != nil {
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

// ParseCountChanged 是绑定合约事件 0xb0be3ede0b8207b0d3c8899f9b86dc42673423e1532e9aa4e0c4c03580cb7a82 的日志解析操作。
//
// Solidity: event CountChanged(uint256 newCount, address indexed caller)
func (_GoJob3 *GoJob3Filterer) ParseCountChanged(log types.Log) (*GoJob3CountChanged, error) {
	event := new(GoJob3CountChanged)
	if err := _GoJob3.contract.UnpackLog(event, "CountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

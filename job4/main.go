package main

// Job4 CLI 用于部署和调用 Job4.sol 对应的计数器合约。
//
// 与 job3 不同，job4 同时包含两套由不同工具链产物生成的 abigen 绑定：
//   - foundry_bind_go/: 用 Foundry(forge) 导出的 ABI/BIN 生成，合约类型为 FoundryBindGo；
//   - hardhat_bind_go/: 用 Hardhat 编译产物的 ABI/BIN 生成，合约类型为 HardhatBindGo。
//
// 两套 ABI 完全一致（getCount/increment/decrement/reset 与 CountChanged 事件），
// 因此 main.go 面向统一的 counter 接口编程，通过 -bind 参数选择实际使用的绑定：
//   - deploy 模式会调用所选绑定的 DeployXxx 部署新合约，并打印部署后的合约地址。
//   - getCount 是只读调用，不需要私钥；increment/decrement/reset/deploy 都会发送交易，必须指定私钥。
//   - CountChanged 模式通过所选绑定的 ParseCountChanged 解析交易回执中的事件日志。
//
// 常用命令:
//
//	go run main.go -bind foundry -mode deploy -private-key <PRIVATE_KEY>
//	go run main.go -bind hardhat  -mode deploy -private-key <PRIVATE_KEY>
//	go run main.go -bind foundry -mode getCount -contract <JOB4_CONTRACT_ADDRESS>
//	go run main.go -bind hardhat  -mode increment -contract <JOB4_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
//	go run main.go -bind hardhat  -mode increment -contract 0x73511669fd4dE447feD18BB79bAFeAC93aB7F31f -private-key 0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e
//	go run main.go -bind foundry -mode decrement -contract <JOB4_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
//	go run main.go -bind hardhat  -mode reset -contract <JOB4_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
//	go run main.go -bind foundry -mode CountChanged -contract <JOB4_CONTRACT_ADDRESS> -tx <TX_HASH>
//
// 可选环境变量:
//
//	ETH_RPC_URL=http://127.0.0.1:8545
//	PRIVATE_KEY=<PRIVATE_KEY>
//	CONTRACT_ADDRESS=<JOB4_CONTRACT_ADDRESS>
//	TX_HASH=<TX_HASH>

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	foundrybind "job4/foundry_bind_go"
	hardhatbind "job4/hardhat_bind_go"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// countChangedEvent 是 CountChanged 事件在两种绑定中的统一视图，
// 让上层逻辑不必区分 FoundryBindGoCountChanged 和 HardhatBindGoCountChanged 的类型差异。
type countChangedEvent struct {
	NewCount *big.Int
	Caller   common.Address
}

// counter 抽象两套生成绑定暴露的相同方法集，业务逻辑只面向该接口编程。
type counter interface {
	GetCount(opts *bind.CallOpts) (*big.Int, error)
	Increment(opts *bind.TransactOpts) (*types.Transaction, error)
	Decrement(opts *bind.TransactOpts) (*types.Transaction, error)
	Reset(opts *bind.TransactOpts) (*types.Transaction, error)
	ParseCountChanged(lg types.Log) (*countChangedEvent, error)
}

// binding 描述一套生成绑定：如何部署新合约，以及如何把已部署地址包装成 counter。
type binding struct {
	label      string // 绑定名称，仅用于输出展示
	deploy     func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, error)
	newCounter func(common.Address, bind.ContractBackend) (counter, error)
}

// foundryCounter 复用 FoundryBindGo 生成的读写方法，仅适配 ParseCountChanged 的返回类型。
type foundryCounter struct{ *foundrybind.FoundryBindGo }

func (c foundryCounter) ParseCountChanged(lg types.Log) (*countChangedEvent, error) {
	event, err := c.FoundryBindGo.ParseCountChanged(lg)
	if err != nil {
		return nil, err
	}
	return &countChangedEvent{NewCount: event.NewCount, Caller: event.Caller}, nil
}

// hardhatCounter 复用 HardhatBindGo 生成的读写方法，仅适配 ParseCountChanged 的返回类型。
type hardhatCounter struct{ *hardhatbind.HardhatBindGo }

func (c hardhatCounter) ParseCountChanged(lg types.Log) (*countChangedEvent, error) {
	event, err := c.HardhatBindGo.ParseCountChanged(lg)
	if err != nil {
		return nil, err
	}
	return &countChangedEvent{NewCount: event.NewCount, Caller: event.Caller}, nil
}

// bindings 登记全部可用的绑定，key 即 -bind 参数的取值。
var bindings = map[string]binding{
	"foundry": {
		label: "foundry",
		deploy: func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
			addr, tx, _, err := foundrybind.DeployFoundryBindGo(auth, backend)
			return addr, tx, err
		},
		newCounter: func(addr common.Address, backend bind.ContractBackend) (counter, error) {
			c, err := foundrybind.NewFoundryBindGo(addr, backend)
			if err != nil {
				return nil, err
			}
			return foundryCounter{FoundryBindGo: c}, nil
		},
	},
	"hardhat": {
		label: "hardhat",
		deploy: func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
			addr, tx, _, err := hardhatbind.DeployHardhatBindGo(auth, backend)
			return addr, tx, err
		},
		newCounter: func(addr common.Address, backend bind.ContractBackend) (counter, error) {
			c, err := hardhatbind.NewHardhatBindGo(addr, backend)
			if err != nil {
				return nil, err
			}
			return hardhatCounter{HardhatBindGo: c}, nil
		},
	},
}

func main() {
	mode, bindName, contract, txHash, privateKeyHex := parseFlags()
	selected := selectBinding(bindName)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, env("ETH_RPC_URL", "http://127.0.0.1:8545"))
	if err != nil {
		log.Fatalf("连接节点失败: %v", err)
	}
	defer client.Close()

	switch mode {
	case "deploy":
		require("private-key / PRIVATE_KEY", privateKeyHex)
		deploy(ctx, client, selected, privateKeyHex)
	case "getCount":
		require("contract / CONTRACT_ADDRESS", contract)
		getCount(ctx, client, selected, contract)
	case "increment", "decrement", "reset":
		require("contract / CONTRACT_ADDRESS", contract)
		require("private-key / PRIVATE_KEY", privateKeyHex)
		write(ctx, client, selected, contract, mode, privateKeyHex)
	case "CountChanged":
		require("contract / CONTRACT_ADDRESS", contract)
		require("tx / TX_HASH", txHash)
		parseCountChanged(ctx, client, selected, contract, txHash)
	default:
		log.Fatalf("未知 mode: %s（可选: deploy, getCount, increment, decrement, reset, CountChanged）", mode)
	}
}

// parseFlags 统一读取命令行参数和环境变量。
// 命令行参数优先；如果未传 -bind、-contract、-tx 或 -private-key，则回退到对应环境变量。
func parseFlags() (mode, bindName, contract, txHash, privateKeyHex string) {
	modeFlag := flag.String("mode", "getCount", "操作: deploy | getCount | increment | decrement | reset | CountChanged")
	bindFlag := flag.String("bind", "foundry", "使用哪套生成绑定: foundry | hardhat")
	contractFlag := flag.String("contract", "", "Job4 计数器合约地址（也可用 CONTRACT_ADDRESS）")
	txHashFlag := flag.String("tx", "", "交易哈希，用于解析 CountChanged（也可用 TX_HASH）")
	privateKeyFlag := flag.String("private-key", "", "交易签名私钥（也可用 PRIVATE_KEY）")
	flag.Parse()

	return *modeFlag,
		*bindFlag,
		first(*contractFlag, os.Getenv("CONTRACT_ADDRESS")),
		first(*txHashFlag, os.Getenv("TX_HASH")),
		first(*privateKeyFlag, os.Getenv("PRIVATE_KEY"))
}

// selectBinding 校验 -bind 参数并返回对应的绑定集合。
func selectBinding(name string) binding {
	selected, ok := bindings[name]
	if !ok {
		log.Fatalf("未知 bind: %s（可选: foundry, hardhat）", name)
	}
	return selected
}

// deploy 使用所选绑定的 DeployXxx 部署新合约。
// 部署交易被挖出后，receipt.ContractAddress 就是后续 getCount/write/CountChanged 要传入的地址。
func deploy(ctx context.Context, client *ethclient.Client, b binding, privateKeyHex string) {
	auth, chainID := transactOpts(ctx, client, privateKeyHex)
	addr, tx, err := b.deploy(auth, client)
	if err != nil {
		log.Fatalf("部署 Job4 合约失败: %v", err)
	}

	printTx("deploy", b, auth.From, addr, chainID, tx)
	receipt := waitMined(ctx, client, tx, "部署")
	printReceipt(receipt)
	fmt.Printf("合约地址: %s\n", receipt.ContractAddress.Hex())
}

// getCount 是只读链上调用，只需要 RPC 和合约地址，不需要账户私钥。
func getCount(ctx context.Context, client *ethclient.Client, b binding, contractHex string) {
	addr, contract := bindCounter(b, contractHex, client)
	count, err := contract.GetCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatalf("调用 getCount 失败: %v", err)
	}

	fmt.Println("========== getCount ==========")
	fmt.Printf("绑定:     %s\n", b.label)
	fmt.Printf("合约:     %s\n", addr.Hex())
	fmt.Printf("当前计数: %s\n", count.String())
	fmt.Println("==============================")
}

// write 处理所有会改变链上状态的方法。
// 这里用一个 switch 把 mode 映射到统一 counter 接口上的 Increment/Decrement/Reset。
func write(ctx context.Context, client *ethclient.Client, b binding, contractHex, method, privateKeyHex string) {
	addr, contract := bindCounter(b, contractHex, client)
	auth, chainID := transactOpts(ctx, client, privateKeyHex)

	var tx *types.Transaction
	var err error
	switch method {
	case "increment":
		tx, err = contract.Increment(auth)
	case "decrement":
		tx, err = contract.Decrement(auth)
	case "reset":
		tx, err = contract.Reset(auth)
	}
	if err != nil {
		log.Fatalf("发送 %s 交易失败（合约未部署或 decrement 时 count=0）: %v", method, err)
	}

	printTx(method, b, auth.From, addr, chainID, tx)
	receipt := waitMined(ctx, client, tx, method)
	printReceipt(receipt)
	printCountChanged(contract, receipt.Logs)
}

// parseCountChanged 按交易哈希读取回执，然后只解析其中属于该合约的 CountChanged 日志。
func parseCountChanged(ctx context.Context, client *ethclient.Client, b binding, contractHex, txHashHex string) {
	addr, contract := bindCounter(b, contractHex, client)
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHashHex))
	if err != nil {
		log.Fatalf("查询交易回执失败: %v", err)
	}

	fmt.Println("========== CountChanged ==========")
	fmt.Printf("绑定:     %s\n", b.label)
	fmt.Printf("合约:     %s\n", addr.Hex())
	fmt.Printf("交易哈希: %s\n", common.HexToHash(txHashHex).Hex())
	printReceipt(receipt)
	printCountChanged(contract, receipt.Logs)
	fmt.Println("==================================")
}

// bindCounter 通过所选绑定把一个已部署合约地址包装成统一的 counter 接口。
// 绑定本身不会发交易，也不会检查该地址是否真的有合约代码；真正的错误会在调用方法时出现。
func bindCounter(b binding, contractHex string, client *ethclient.Client) (common.Address, counter) {
	addr := common.HexToAddress(contractHex)
	contract, err := b.newCounter(addr, client)
	if err != nil {
		log.Fatalf("绑定 Job4 合约失败: %v", err)
	}
	return addr, contract
}

// transactOpts 创建部署和写交易共用的签名配置。
// 私钥可通过 -private-key 或 PRIVATE_KEY 指定；只读查询不需要创建这个配置。
func transactOpts(ctx context.Context, client *ethclient.Client, privateKeyHex string) (*bind.TransactOpts, *big.Int) {
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatalf("获取 chainID 失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey(privateKeyHex), chainID)
	if err != nil {
		log.Fatalf("创建交易授权失败: %v", err)
	}
	auth.Context = ctx
	return auth, chainID
}

// privateKey 解析十六进制私钥，支持带或不带 0x 前缀。
func privateKey(privateKeyHex string) *ecdsa.PrivateKey {
	key, err := crypto.HexToECDSA(trim0x(privateKeyHex))
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}
	return key
}

// printCountChanged 使用所选绑定的 ParseCountChanged 解码事件。
// 非 CountChanged 日志会解析失败并被跳过，因此同一笔交易里有其它合约日志也不会影响结果。
func printCountChanged(contract counter, logs []*types.Log) {
	found := false
	for _, lg := range logs {
		event, err := contract.ParseCountChanged(*lg)
		if err != nil {
			continue
		}
		found = true
		fmt.Printf("事件: CountChanged(newCount=%s, caller=%s)\n", event.NewCount, event.Caller.Hex())
	}
	if !found {
		fmt.Println("未找到 CountChanged 事件")
	}
}

// waitMined 等待交易被打包；超时由 main 中创建的 ctx 控制。
func waitMined(ctx context.Context, client *ethclient.Client, tx *types.Transaction, action string) *types.Receipt {
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("等待 %s 回执失败: %v", action, err)
	}
	return receipt
}

func printTx(action string, b binding, from, contract common.Address, chainID *big.Int, tx *types.Transaction) {
	fmt.Printf("========== %s(%s) 已发送 ==========\n", action, b.label)
	fmt.Printf("绑定:     %s\n", b.label)
	fmt.Printf("调用方:   %s\n", from.Hex())
	fmt.Printf("合约:     %s\n", contract.Hex())
	fmt.Printf("ChainID:  %s\n", chainID)
	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())
	fmt.Println("==============================")
}

func printReceipt(receipt *types.Receipt) {
	fmt.Printf("交易状态: %d (1=成功)\n", receipt.Status)
	fmt.Printf("区块号:   %d\n", receipt.BlockNumber.Uint64())
}

func require(name, value string) {
	if value == "" {
		log.Fatalf("请设置 %s", name)
	}
}

func env(key, fallback string) string {
	return first(os.Getenv(key), fallback)
}

func first(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func trim0x(s string) string {
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		return s[2:]
	}
	return s
}

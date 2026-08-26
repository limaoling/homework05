package main

// Job3 CLI 负责部署和调用 Job3.sol 对应的计数器合约。
//
// 设计要点:
//   - main.go 不再手写 ABI 编码/解码，而是直接使用 go_job3/job3.go 生成的 Go 绑定。
//   - deploy 模式会调用 DeployGoJob3 部署新合约，并打印部署后的合约地址。
//   - getCount 是只读调用，不需要私钥；increment/decrement/reset/deploy 都会发送交易，必须指定私钥。
//   - CountChanged 模式通过已部署合约绑定的 ParseCountChanged 解析交易回执中的事件日志。
//
// 常用命令:
//
//	go run main.go -mode deploy -private-key <PRIVATE_KEY>
//	go run main.go -mode getCount -contract <JOB3_CONTRACT_ADDRESS>
//	go run main.go -mode increment -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
//	go run main.go -mode decrement -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
//	go run main.go -mode reset -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
//	go run main.go -mode CountChanged -contract <JOB3_CONTRACT_ADDRESS> -tx <TX_HASH>
//
// 可选环境变量:
//
//	ETH_RPC_URL=http://127.0.0.1:8545
//	PRIVATE_KEY=<PRIVATE_KEY>
//	CONTRACT_ADDRESS=<JOB3_CONTRACT_ADDRESS>
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

	gojob3 "job3/go_job3"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	mode, contract, txHash, privateKeyHex := parseFlags()

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
		deploy(ctx, client, privateKeyHex)
	case "getCount":
		require("contract / CONTRACT_ADDRESS", contract)
		getCount(ctx, client, contract)
	case "increment", "decrement", "reset":
		require("contract / CONTRACT_ADDRESS", contract)
		require("private-key / PRIVATE_KEY", privateKeyHex)
		write(ctx, client, contract, mode, privateKeyHex)
	case "CountChanged":
		require("contract / CONTRACT_ADDRESS", contract)
		require("tx / TX_HASH", txHash)
		parseCountChanged(ctx, client, contract, txHash)
	default:
		log.Fatalf("未知 mode: %s（可选: deploy, getCount, increment, decrement, reset, CountChanged）", mode)
	}
}

// parseFlags 统一读取命令行参数和环境变量。
// 命令行参数优先；如果未传 -contract、-tx 或 -private-key，则回退到对应环境变量。
func parseFlags() (mode, contract, txHash, privateKeyHex string) {
	modeFlag := flag.String("mode", "getCount", "操作: deploy | getCount | increment | decrement | reset | CountChanged")
	contractFlag := flag.String("contract", "", "Job3 合约地址（也可用 CONTRACT_ADDRESS）")
	txHashFlag := flag.String("tx", "", "交易哈希，用于解析 CountChanged（也可用 TX_HASH）")
	privateKeyFlag := flag.String("private-key", "", "交易签名私钥（也可用 PRIVATE_KEY）")
	flag.Parse()

	return *modeFlag,
		first(*contractFlag, os.Getenv("CONTRACT_ADDRESS")),
		first(*txHashFlag, os.Getenv("TX_HASH")),
		first(*privateKeyFlag, os.Getenv("PRIVATE_KEY"))
}

// deploy 使用 abigen 生成的 DeployGoJob3 部署新合约。
// 部署交易被挖出后，receipt.ContractAddress 就是后续 getCount/write/CountChanged 要传入的地址。
func deploy(ctx context.Context, client *ethclient.Client, privateKeyHex string) {
	auth, chainID := transactOpts(ctx, client, privateKeyHex)
	addr, tx, _, err := gojob3.DeployGoJob3(auth, client)
	if err != nil {
		log.Fatalf("部署 Job3 合约失败: %v", err)
	}

	printTx("deploy", auth.From, addr, chainID, tx)
	receipt := waitMined(ctx, client, tx, "部署")
	printReceipt(receipt)
	fmt.Printf("合约地址: %s\n", receipt.ContractAddress.Hex())
}

// getCount 是只读链上调用，只需要 RPC 和合约地址，不需要账户私钥。
func getCount(ctx context.Context, client *ethclient.Client, contractHex string) {
	addr, contract := bindJob3(client, contractHex)
	count, err := contract.GetCount(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Fatalf("调用 getCount 失败: %v", err)
	}

	fmt.Println("========== getCount ==========")
	fmt.Printf("合约:     %s\n", addr.Hex())
	fmt.Printf("当前计数: %s\n", count.String())
	fmt.Println("==============================")
}

// write 处理所有会改变链上状态的方法。
// 这里用一个 switch 把 mode 映射到生成绑定里的 Increment/Decrement/Reset 方法。
func write(ctx context.Context, client *ethclient.Client, contractHex, method, privateKeyHex string) {
	addr, contract := bindJob3(client, contractHex)
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

	printTx(method, auth.From, addr, chainID, tx)
	receipt := waitMined(ctx, client, tx, method)
	printReceipt(receipt)
	printCountChanged(contract, receipt.Logs)
}

// parseCountChanged 按交易哈希读取回执，然后只解析其中属于 Job3 的 CountChanged 日志。
func parseCountChanged(ctx context.Context, client *ethclient.Client, contractHex, txHashHex string) {
	addr, contract := bindJob3(client, contractHex)
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHashHex))
	if err != nil {
		log.Fatalf("查询交易回执失败: %v", err)
	}

	fmt.Println("========== CountChanged ==========")
	fmt.Printf("合约:     %s\n", addr.Hex())
	fmt.Printf("交易哈希: %s\n", common.HexToHash(txHashHex).Hex())
	printReceipt(receipt)
	printCountChanged(contract, receipt.Logs)
	fmt.Println("==================================")
}

// bindJob3 把一个已部署合约地址包装成 Go 对象。
// 绑定本身不会发交易，也不会检查该地址是否真的有合约代码；真正的错误会在调用方法时出现。
func bindJob3(client *ethclient.Client, contractHex string) (common.Address, *gojob3.GoJob3) {
	addr := common.HexToAddress(contractHex)
	contract, err := gojob3.NewGoJob3(addr, client)
	if err != nil {
		log.Fatalf("绑定 Job3 合约失败: %v", err)
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

// printCountChanged 使用生成绑定的 ParseCountChanged 解码事件。
// 非 CountChanged 日志会解析失败并被跳过，因此同一笔交易里有其它合约日志也不会影响结果。
func printCountChanged(contract *gojob3.GoJob3, logs []*types.Log) {
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

func printTx(action string, from, contract common.Address, chainID *big.Int, tx *types.Transaction) {
	fmt.Printf("========== %s 已发送 ==========\n", action)
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

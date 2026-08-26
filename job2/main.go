package main

// Package main 提供了一个用于与名为 'Job2' 的智能合约进行交互的 CLI 工具。
// 它支持读取合约的计数、通过交易修改计数，以及解析合约发出的事件。
//
// 执行命令示例:
//
//	# 读取计数器
//	go run main.go -contract <JOB2_CONTRACT_ADDRESS> -mode getCount
//	go run main.go -contract 0x5FbDB2315678afecb367f032d93F642f64180aa3 -mode getCount
//
//	# 发送交易: increment / decrement / reset
//	go run main.go -contract <JOB2_CONTRACT_ADDRESS> -mode increment
//	go run main.go -contract 0x5FbDB2315678afecb367f032d93F642f64180aa3 -mode increment

//
//	# 解析 CountChanged 事件
//	go run main.go -contract <JOB2_CONTRACT_ADDRESS> -mode CountChanged -tx <TX_HASH>
//	go run main.go -contract 0x5FbDB2315678afecb367f032d93F642f64180aa3 -mode CountChanged -tx 0x379f09bd200a526964e7f9773b2cf278bb3fa2974750425d423b9d7404c96ca8
//
// 可选环境变量:
//
//	ETH_RPC_URL=http://127.0.0.1:8545
//	PRIVATE_KEY=<PRIVATE_KEY>
//	CONTRACT_ADDRESS=<JOB2_CONTRACT_ADDRESS>
//	TX_HASH=<TX_HASH>

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 解析操作模式和合约地址的 CLI 标志。
	mode := flag.String("mode", "getCount", "操作: getCount | increment | decrement | reset | CountChanged")
	contractHex := flag.String("contract", "", "Job2 合约地址（也可用环境变量 CONTRACT_ADDRESS）")
	txHashHex := flag.String("tx", "", "交易哈希，用于解析 CountChanged（也可用 TX_HASH）")
	flag.Parse()

	if *contractHex == "" {
		*contractHex = os.Getenv("CONTRACT_ADDRESS")
	}
	if *txHashHex == "" {
		*txHashHex = os.Getenv("TX_HASH")
	}

	// 验证输入。
	if *mode == "" {
		log.Fatalf("操作类型不能为空")
	}

	// 通过 RPC 连接以太坊节点。
	rpcURL := os.Getenv("ETH_RPC_URL")
	if rpcURL == "" {
		rpcURL = "http://127.0.0.1:8545"
	}

	// 设置用于节点交互的超时上下文。
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("连接节点失败: %v", err)
	}
	defer client.Close()

	// 加载并解析合约 ABI。
	//abiBytes, err := os.ReadFile("./foundry-abi/foundry-job2.abi.json")
	abiBytes, err := os.ReadFile("./hardhat3-abi/hardhat3-job2.abi.json")
	if err != nil {
		log.Fatalf("读取 ABI 失败: %v", err)
	}
	parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Fatalf("解析 ABI 失败: %v", err)
	}

	// 根据请求的模式进行路由。
	switch *mode {
	case "getCount":
		require("contract / CONTRACT_ADDRESS", *contractHex)
		handleGetCount(ctx, client, parsedABI, *contractHex)
	case "increment", "decrement", "reset":
		require("contract / CONTRACT_ADDRESS", *contractHex)
		handleWrite(ctx, client, parsedABI, *contractHex, *mode)
	case "CountChanged":
		require("tx / TX_HASH", *txHashHex)
		handleCountChanged(ctx, client, parsedABI, *txHashHex)
	default:
		log.Fatalf("未知 mode: %s（可选: getCount, increment, decrement, reset, CountChanged）", *mode)
	}
}

// handleGetCount 调用合约的 'getCount' 函数以读取当前计数器的值。
func handleGetCount(ctx context.Context, client *ethclient.Client, parsedABI abi.ABI, contractHex string) {
	contract := common.HexToAddress(contractHex)
	data, err := parsedABI.Pack("getCount")
	if err != nil {
		log.Fatalf("编码 getCount 失败: %v", err)
	}

	// 对合约执行只读调用。
	out, err := client.CallContract(ctx, ethereum.CallMsg{To: &contract, Data: data}, nil)
	if err != nil {
		log.Fatalf("调用 getCount 失败: %v", err)
	}

	// 根据 ABI 解包输出数据。
	results, err := parsedABI.Unpack("getCount", out)
	if err != nil {
		log.Fatalf("解码 getCount 返回值失败: %v", err)
	}
	count := results[0].(*big.Int)

	fmt.Println("========== getCount ==========")
	fmt.Printf("合约:     %s\n", contract.Hex())
	fmt.Printf("当前计数: %s\n", count.String())
	fmt.Println("==============================")
}

// handleWrite 准备并提交一笔交易来调用修改状态的方法（increment/decrement/reset）。
func handleWrite(ctx context.Context, client *ethclient.Client, parsedABI abi.ABI, contractHex, method string) {
	privateKey, fromAddress := loadAccount()
	contract := common.HexToAddress(contractHex)

	// 对方法调用数据进行编码。
	data, err := parsedABI.Pack(method)
	if err != nil {
		log.Fatalf("编码 %s 失败: %v", method, err)
	}

	// 收集必要的交易参数。
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatalf("获取 nonce 失败: %v", err)
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("获取 gasPrice 失败: %v", err)
	}
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatalf("获取 chainID 失败: %v", err)
	}

	// 估算交易的 gas 限制。
	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From: fromAddress,
		To:   &contract,
		Data: data,
	})
	if err != nil {
		log.Fatalf("估算 gas 失败（合约未部署或 decrement 时 count=0）: %v", err)
	}

	// 创建、签名并发送交易。
	tx := types.NewTransaction(nonce, contract, big.NewInt(0), gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("签名交易失败: %v", err)
	}
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		log.Fatalf("发送交易失败: %v", err)
	}

	fmt.Printf("========== %s 已发送 ==========\n", method)
	fmt.Printf("调用方:   %s\n", fromAddress.Hex())
	fmt.Printf("合约:     %s\n", contract.Hex())
	fmt.Printf("ChainID:  %s\n", chainID.String())
	fmt.Printf("Nonce:    %d\n", nonce)
	fmt.Printf("GasLimit: %d\n", gasLimit)
	fmt.Printf("GasPrice: %s wei\n", gasPrice.String())
	fmt.Printf("交易哈希: %s\n", signedTx.Hash().Hex())
	fmt.Println("==============================")

	// 等待交易被挖掘并显示回执状态。
	receipt, err := waitForReceipt(signedTx.Hash(), client)
	if err != nil {
		log.Fatalf("等待回执失败: %v", err)
	}
	fmt.Printf("交易状态: %d (1=成功)\n", receipt.Status)
	fmt.Printf("区块号:   %d\n", receipt.BlockNumber.Uint64())
	printCountChangedLogs(parsedABI, receipt.Logs)
}

// handleCountChanged 根据哈希获取交易回执并解析任何 'CountChanged' 事件。
func handleCountChanged(ctx context.Context, client *ethclient.Client, parsedABI abi.ABI, txHashHex string) {
	receipt, err := client.TransactionReceipt(ctx, common.HexToHash(txHashHex))
	if err != nil {
		log.Fatalf("查询交易回执失败: %v", err)
	}

	fmt.Println("========== CountChanged ==========")
	fmt.Printf("交易哈希: %s\n", common.HexToHash(txHashHex).Hex())
	fmt.Printf("交易状态: %d\n", receipt.Status)
	printCountChangedLogs(parsedABI, receipt.Logs)
	fmt.Println("==================================")
}

// printCountChangedLogs 遍历交易日志以查找并解码 'CountChanged' 事件。
func printCountChangedLogs(parsedABI abi.ABI, logs []*types.Log) {
	event := parsedABI.Events["CountChanged"]
	found := false
	for _, lg := range logs {
		// 仅处理匹配 CountChanged 事件签名的日志。
		if len(lg.Topics) == 0 || lg.Topics[0] != event.ID {
			continue
		}
		found = true

		// 解码事件数据。
		vals, err := parsedABI.Unpack("CountChanged", lg.Data)
		if err != nil {
			log.Fatalf("解码 CountChanged Data 失败: %v", err)
		}
		newCount := vals[0].(*big.Int)

		// 从 topics 中解析调用者地址。
		caller := common.Address{}
		if len(lg.Topics) > 1 {
			caller = common.BytesToAddress(lg.Topics[1].Bytes())
		}
		fmt.Printf("事件: CountChanged(newCount=%s, caller=%s)\n", newCount.String(), caller.Hex())
	}
	if !found {
		fmt.Println("未找到 CountChanged 事件")
	}
}

// waitForReceipt 轮询区块链，直到交易回执可用。
func waitForReceipt(txHash common.Hash, client *ethclient.Client) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
		}
	}
}

// loadAccount 从环境变量加载私钥并推导出对应的公用地址。
func loadAccount() (*ecdsa.PrivateKey, common.Address) {
	// Anvil 默认账户 #0；可用 PRIVATE_KEY 覆盖
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		privateKeyHex = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	}
	privateKey, err := crypto.HexToECDSA(trim0x(privateKeyHex))
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}
	pub, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法转换公钥")
	}
	return privateKey, crypto.PubkeyToAddress(*pub)
}

// require 验证提供的值是否为空，用于 CLI 输入。
func require(name, value string) {
	if value == "" {
		log.Fatalf("请设置 %s", name)
	}
}

// trim0x 从十六进制字符串中移除可选的 "0x" 或 "0X" 前缀。
func trim0x(s string) string {
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		return s[2:]
	}
	return s
}

package main

// Package main 演示了如何使用 go-ethereum 库与以太坊区块链进行交互。
// 它包括查询区块信息和在网络上执行基本以太币转账的示例。

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 从环境变量加载 RPC URL，如果未设置，则回退到本地节点。
	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	if rpcURL == "" {
		rpcURL = "http://127.0.0.1:8545"
	}
	ctx := context.Background()

	// 建立与以太坊节点的连接。
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("连接 Sepolia 失败: %v", err)
	}
	defer client.Close()

	// 1. 查询区块: 查询并显示区块的详细信息。
	queryBlock(ctx, client)

	// 2. 发送交易: 构造并广播一笔以太币转账。
	// 需要配置 PRIVATE_KEY 和 TO_ADDRESS。
	sendTransaction(ctx, client)
}

// queryBlock 查询指定区块号的区块信息（未设置 BLOCK_NUMBER 则查最新区块）
func queryBlock(ctx context.Context, client *ethclient.Client) {
	var blockNumber *big.Int
	if n := os.Getenv("BLOCK_NUMBER"); n != "" {
		bn, ok := new(big.Int).SetString(n, 10)
		if !ok {
			log.Fatalf("无效的 BLOCK_NUMBER: %s", n)
		}
		blockNumber = bn
	}

	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Fatalf("查询区块失败: %v", err)
	}

	fmt.Println("========== 区块信息 ==========")
	fmt.Printf("区块号:     %d\n", block.NumberU64())
	fmt.Printf("区块哈希:   %s\n", block.Hash().Hex())
	fmt.Printf("父哈希:     %s\n", block.ParentHash().Hex())
	fmt.Printf("时间戳:     %d\n", block.Time())
	fmt.Printf("交易数量:   %d\n", len(block.Transactions()))
	fmt.Printf("矿工地址:   %s\n", block.Coinbase().Hex())
	fmt.Printf("Gas Limit:  %d\n", block.GasLimit())
	fmt.Printf("Gas Used:   %d\n", block.GasUsed())
	fmt.Printf("Difficulty: %s\n", block.Difficulty().String())
	fmt.Printf("区块大小:   %d bytes\n", block.Size())
	fmt.Println("==============================")
}

// sendTransaction 构造并发送一笔 Sepolia 以太币转账交易
func sendTransaction(ctx context.Context, client *ethclient.Client) {
	//privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKeyHex := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	//toAddressHex := os.Getenv("TO_ADDRESS")
	toAddressHex := "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"

	privateKey, err := crypto.HexToECDSA(trim0x(privateKeyHex))
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法转换公钥")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatalf("获取 nonce 失败: %v", err)
	}

	// 默认转账 0.001 ETH；可通过 VALUE_WEI 覆盖（单位：wei）
	value := big.NewInt(1e15) // 0.001 ether = 10^15 wei
	if v := os.Getenv("VALUE_WEI"); v != "" {
		wei, ok := new(big.Int).SetString(v, 10)
		if !ok {
			log.Fatalf("无效的 VALUE_WEI: %s", v)
		}
		value = wei
	}

	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("获取 gasPrice 失败: %v", err)
	}

	toAddress := common.HexToAddress(toAddressHex)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatalf("获取 chainID 失败: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("签名交易失败: %v", err)
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatalf("发送交易失败: %v", err)
	}

	fmt.Println("========== 交易已发送 ==========")
	fmt.Printf("发送方:   %s\n", fromAddress.Hex())
	fmt.Printf("接收方:   %s\n", toAddress.Hex())
	fmt.Printf("金额:     %s wei\n", value.String())
	fmt.Printf("ChainID:  %s\n", chainID.String())
	fmt.Printf("Nonce:    %d\n", nonce)
	fmt.Printf("GasPrice: %s wei\n", gasPrice.String())
	fmt.Printf("交易哈希: %s\n", signedTx.Hash().Hex())
	fmt.Println("================================")
}

// trim0x 去掉十六进制字符串开头的 "0x"/"0X" 前缀。
// crypto.HexToECDSA 等部分 API 不接受带前缀的输入。
func trim0x(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}

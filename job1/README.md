# job1: 基础区块链交互

演示如何使用 `go-ethereum` 库与以太坊网络进行最基础的交互：**查询区块信息**和**发送 ETH 转账交易**。不涉及智能合约。

## 功能

### 1. 查询区块

通过 `ethclient.BlockByNumber` 查询指定区块号（未设置 `BLOCK_NUMBER` 时查最新区块），打印：

- 区块号、区块哈希、父哈希
- 时间戳、交易数量、矿工地址
- Gas Limit / Gas Used / Difficulty / 区块大小

### 2. 发送转账交易

完整走一遍链上交易的流程：

1. `crypto.HexToECDSA` 解析私钥，并从私钥推导出发送方地址；
2. `PendingNonceAt` 获取账户 nonce；
3. `SuggestGasPrice` 获取建议 gas 价格，构造 `types.NewTransaction`（固定 gasLimit 21000，标准转账消耗）；
4. `NetworkID` 拿到 chainID，用 EIP-155 签名器签名；
5. `SendTransaction` 广播，打印交易哈希等信息。

## 环境变量

| 变量 | 说明 | 默认值 |
| --- | --- | --- |
| `SEPOLIA_RPC_URL` | 以太坊节点 RPC 地址 | `http://127.0.0.1:8545` |
| `BLOCK_NUMBER` | 要查询的区块号（十进制），不设置则查最新区块 | 最新区块 |
| `PRIVATE_KEY` | 发送方私钥 | 见下方说明 |
| `TO_ADDRESS` | 接收方地址 | 见下方说明 |
| `VALUE_WEI` | 转账金额，单位 wei | `1000000000000000`（0.001 ETH） |

> **注意**：当前 `main.go` 中私钥和收款地址被硬编码为 Anvil/Hardhat 默认测试账户（Account #0 → Account #1），环境变量 `PRIVATE_KEY` / `TO_ADDRESS` 不生效。连本地开发链时可直接运行；如需自定义账户，请把这两行改回读取环境变量的写法。

## 运行说明

1. 准备一个可用节点，任选其一：
   - 本地链：`npx hardhat node` 或 `anvil`（默认 RPC 即 `http://127.0.0.1:8545`）
   - 测试网：设置 `SEPOLIA_RPC_URL` 为 Sepolia 节点地址
2. 运行：

```bash
go run main.go
```

## 输出示例

```text
========== 区块信息 ==========
区块号:     5
区块哈希:   0x...
父哈希:     0x...
时间戳:     1756000000
交易数量:   1
矿工地址:   0x0000000000000000000000000000000000000000
Gas Limit:  30000000
Gas Used:   21000
Difficulty: 0
区块大小:   721 bytes
==============================
========== 交易已发送 ==========
发送方:   0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
接收方:   0x70997970C51812dc3A010C7d01b50e0d17dc79C8
金额:     1000000000000000 wei
ChainID:  31337
Nonce:    0
GasPrice: 2000000000 wei
交易哈希: 0x...
================================
```

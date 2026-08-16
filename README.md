# 以太坊交互练习项目

本仓库包含两个基于 Go 语言和以太坊生态的练习项目。

## 项目结构

```text
/
├── job1/              # 项目 1：基础区块链交互
└── job2/              # 项目 2：智能合约交互
```

## job1: 基础区块链交互

此项目演示了如何使用 `go-ethereum` 库与以太坊网络进行基础交互。

### 功能
- 查询指定或最新区块的详细信息（哈希、父哈希、时间戳、交易数等）。
- 构造并发送一笔简单的以太币（ETH）转账交易。

### 运行说明
1. 进入 `job1` 目录。
2. 设置环境变量：
   - `SEPOLIA_RPC_URL` (可选，默认为 `http://127.0.0.1:8545`)
   - `BLOCK_NUMBER` (查询特定区块)
   - `PRIVATE_KEY` / `TO_ADDRESS` / `VALUE_WEI` (转账配置)
3. 运行：`go run main.go`

---

## job2: 智能合约交互

此项目演示了如何通过 Go 代码部署（或交互）智能合约。

### 核心组成
- `job2-contract/`: 使用 Foundry 开发的 Solidity 智能合约 (`Job2.sol`)。
- `abi/job2.abi.json`: 合约的 ABI 定义。
- `main.go`: Go 交互脚本。

### 功能
- `getCount`: 只读调用合约中的 `getCount()` 方法查询当前计数。
- `increment`/`decrement`/`reset`: 发送交易修改合约状态。
- `CountChanged`: 查询特定交易的事件。

### 运行说明
1. 进入 `job2` 目录。
2. 确保已正确配置 `job2-contract` 并通过 foundry 编译得到 ABI。
3. 运行交互脚本：`go run main.go -mode <操作> -contract <合约地址>`
   - 可选模式: `getCount`, `increment`, `decrement`, `reset`, `CountChanged`

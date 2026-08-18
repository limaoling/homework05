# 以太坊交互练习项目

本仓库包含三个基于 Go 语言和以太坊生态的练习项目。

## 项目结构

```text
/
├── job1/              # 项目 1：基础区块链交互
├── job2/              # 项目 2：手动 ABI 智能合约交互
└── job3/              # 项目 3：使用 abigen Go 绑定交互智能合约
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

此项目演示了如何通过 Go 代码和 ABI 与智能合约交互。

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

---

## job3: 使用 Go 绑定交互智能合约

此项目演示了如何使用 `abigen` 生成的 Go 绑定代码部署和调用 `Job3` 计数器合约。

### 核心组成
- `job3.sol`: Solidity 计数器合约源码。
- `go_job3/job3.go`: 由 ABI 和 bytecode 生成的 Go 合约绑定。
- `job3_sol_Job3.abi` / `job3_sol_Job3.bin`: 合约 ABI 和字节码文件。
- `main.go`: 使用生成绑定进行部署、查询、写交易和事件解析的 CLI。

### 功能
- `deploy`: 部署新的 `Job3` 合约。
- `getCount`: 只读调用 `getCount()` 查询当前计数。
- `increment` / `decrement` / `reset`: 发送交易修改计数器状态。
- `CountChanged`: 按交易哈希解析 `CountChanged` 事件。

### 参数和环境变量
- `-mode`: 操作模式，可选 `deploy`, `getCount`, `increment`, `decrement`, `reset`, `CountChanged`。
- `-contract`: 已部署的 `Job3` 合约地址，也可用 `CONTRACT_ADDRESS`。
- `-tx`: 交易哈希，用于 `CountChanged` 模式，也可用 `TX_HASH`。
- `-private-key`: 部署和写交易使用的私钥，也可用 `PRIVATE_KEY`。
- `ETH_RPC_URL`: RPC 地址，默认 `http://127.0.0.1:8545`。

### 运行说明
1. 进入 `job3` 目录。
2. 启动本地链或准备好可用 RPC，例如 Anvil 默认 RPC：`http://127.0.0.1:8545`。
3. 部署合约，必须指定私钥：

```bash
go run main.go -mode deploy -private-key <PRIVATE_KEY>
```

也可以通过环境变量指定私钥：

```bash
PRIVATE_KEY=<PRIVATE_KEY> go run main.go -mode deploy
```

部署成功后，程序会打印 `合约地址`。后续调用需要使用这个地址。

### 常用命令

```bash
# 查询当前计数，不需要私钥
go run main.go -mode getCount -contract <JOB3_CONTRACT_ADDRESS>

# 递增计数，需要私钥签名交易
go run main.go -mode increment -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>

# 递减计数，需要私钥签名交易；当 count 为 0 时会失败
go run main.go -mode decrement -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>

# 重置计数，需要私钥签名交易
go run main.go -mode reset -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>

# 解析某笔交易中的 CountChanged 事件
go run main.go -mode CountChanged -contract <JOB3_CONTRACT_ADDRESS> -tx <TX_HASH>
```

也可以用环境变量减少重复参数：

```bash
export ETH_RPC_URL=http://127.0.0.1:8545
export PRIVATE_KEY=<PRIVATE_KEY>
export CONTRACT_ADDRESS=<JOB3_CONTRACT_ADDRESS>

go run main.go -mode getCount
go run main.go -mode increment
go run main.go -mode reset
```

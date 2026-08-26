# job2: 手动 ABI 智能合约交互

用三种方式实现并交互同一个 `Job2` 计数器合约（递增 / 递减 / 重置，触发 `CountChanged` 事件），重点演示 **Go 端如何脱离代码生成工具、直接解析 ABI 完成调用与事件解析**。

## 项目结构

```text
job2/
├── foundry-contract/    Foundry 版合约项目（forge/cast/anvil），见其 README.md
├── foundry-abi/         Foundry 导出的 ABI（foundry-Job2.abi.json）
├── hardhat3-contract/   Hardhat 3 版合约项目（mocha + ethers + Ignition），见其 README.md
├── hardhat3-abi/        Hardhat 导出的 ABI（hardhat3-Job2.abi.json）
└── main.go              Go CLI：手动解析 ABI 与合约交互
```

两个合约目录是同一个 `Job2.sol` 的两种工具链实现，任选其一部署即可；`main.go` 通过读取对应目录导出的 ABI JSON 文件工作。

## Go 客户端（main.go）

### 功能与模式

| `-mode` | 说明 | 是否需要私钥 |
| --- | --- | --- |
| `getCount` | 只读调用 `getCount()` 查询当前计数 | 否 |
| `increment` / `decrement` / `reset` | 发交易修改计数器状态 | 是 |
| `CountChanged` | 按交易哈希查询回执，解析事件 | 否 |

### 参数与环境变量

| 参数 / 变量 | 说明 | 默认值 |
| --- | --- | --- |
| `-mode` | 操作模式 | `getCount` |
| `-contract` / `CONTRACT_ADDRESS` | Job2 合约地址（写操作必填） | 无 |
| `-tx` / `TX_HASH` | 交易哈希（`CountChanged` 模式必填） | 无 |
| `ETH_RPC_URL` | 节点 RPC 地址 | `http://127.0.0.1:8545` |
| `PRIVATE_KEY` | 签名私钥 | Anvil/Hardhat 默认账户 #0 |

> 注意：ABI 路径硬编码在 `main.go:84`，当前读取 `./hardhat3-abi/hardhat3-job2.abi.json`；如需改用 Foundry 导出的 ABI，切换为注释中的 `foundry-abi` 路径即可。

## 运行步骤

1. 启动本地链（`anvil` 或 `npx hardhat node`），并用任一合约项目部署 `Job2`，拿到合约地址；
2. 在 `job2` 目录下运行：

```bash
# 查询当前计数
go run main.go -mode getCount -contract <JOB2_CONTRACT_ADDRESS>

# 发送交易修改状态（默认使用 Anvil 测试账户签名）
go run main.go -mode increment -contract <JOB2_CONTRACT_ADDRESS>
go run main.go -mode decrement -contract <JOB2_CONTRACT_ADDRESS>
go run main.go -mode reset -contract <JOB2_CONTRACT_ADDRESS>

# 按交易哈希解析 CountChanged 事件
go run main.go -mode CountChanged -contract <JOB2_CONTRACT_ADDRESS> -tx <TX_HASH>
```

也可以用环境变量减少重复参数：

```bash
export ETH_RPC_URL=http://127.0.0.1:8545
export CONTRACT_ADDRESS=<JOB2_CONTRACT_ADDRESS>

go run main.go -mode getCount
go run main.go -mode increment
```

## 实现要点

- **方法调用**：`abi.JSON` 解析 ABI 文件后，用 `Pack(method)` 编码 calldata，`CallContract` 做只读调用并用 `Unpack` 解码返回值；
- **写交易**：手动获取 nonce、gasPrice、chainID，`EstimateGas` 估算 gas，构造 `types.NewTransaction` 后经 EIP-155 签名广播，再轮询等待回执；
- **事件解析**：遍历回执中的 Logs，用 `Topics[0] == event.ID` 过滤出 `CountChanged`，从 `Data` 解码 `newCount`、从 `Topics[1]` 还原 indexed 的 `caller` 地址。

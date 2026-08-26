# 以太坊交互练习项目

本仓库包含四个基于 Go 语言和以太坊生态的递进式练习项目：从基础的区块查询与转账，到手动解析 ABI、使用 abigen 生成绑定，再到多工具链产物的双绑定实践。

## 项目结构

```text
/
├── job1/   # 项目 1：基础区块链交互（查区块、ETH 转账）
├── job2/   # 项目 2：手动解析 ABI 与智能合约交互
├── job3/   # 项目 3：使用 abigen 生成的 Go 绑定交互智能合约
└── job4/   # 项目 4：同一合约的两套工具链产物生成双绑定，统一接口切换调用
```

> 每个 `job*` 目录都是一个独立的 Go module（`module job1` ~ `module job4`），请进入对应目录后再运行 `go run main.go`。

## 环境准备

| 工具 | 用途 | 安装 |
| --- | --- | --- |
| Go ≥ 1.25 | 运行各项目的 CLI | https://go.dev/dl/ |
| Foundry（forge / cast / anvil） | 编译、测试合约，启动本地链，导出 ABI | https://book.getfoundry.sh/getting-started/installation |
| Node.js + npm | Hardhat 3 合约项目 | https://nodejs.org/ |
| abigen（go-ethereum 自带工具） | 由 ABI/BIN 生成 Go 绑定（job3/job4） | `go install github.com/ethereum/go-ethereum/cmd/abigen@latest` |

本地链任选其一（默认 RPC 均为 `http://127.0.0.1:8545`，自带测试账户）：

```bash
anvil                    # Foundry 本地节点
npx hardhat node         # Hardhat 常驻节点（在任一 hardhat3-contract 目录下）
```

---

## job1: 基础区块链交互

演示如何使用 `go-ethereum` 库与以太坊网络进行基础交互，不涉及智能合约。

- **查询区块**：按 `BLOCK_NUMBER` 查询指定区块（缺省查最新），打印哈希、时间戳、交易数等信息；
- **发送转账**：构造签名交易完成一笔 ETH 转账。

配置项：`SEPOLIA_RPC_URL`、`BLOCK_NUMBER`、`PRIVATE_KEY`、`TO_ADDRESS`、`VALUE_WEI`。

详见 [job1/README.md](job1/README.md)。

## job2: 手动 ABI 智能合约交互

用 **Foundry** 和 **Hardhat 3** 两条工具链实现同一个 `Job2` 计数器合约，Go 端脱离代码生成工具，直接用 `abi.JSON` 解析 ABI 完成调用与事件解析——重点理解 ABI 编码原理。

- `foundry-contract/` + `foundry-abi/`：Foundry 版合约及其导出 ABI；
- `hardhat3-contract/` + `hardhat3-abi/`：Hardhat 3 版合约及其导出 ABI；
- `main.go`：读取 ABI JSON，手动 pack 参数 / unpack 返回值和事件。

模式：`getCount` / `increment` / `decrement` / `reset` / `CountChanged`。

详见 [job2/README.md](job2/README.md)。

## job3: 使用 abigen Go 绑定交互智能合约

用 `solc` 编译 `Job3` 计数器合约得到 `.abi` / `.bin`，再用 `abigen` 生成强类型 Go 绑定（`go_job3/`）。像调用普通 Go 方法一样操作合约：部署（`DeployGoJob3`）、读写（`GetCount` / `Increment` …）、事件解析（`ParseCountChanged`）。

模式：`deploy` / `getCount` / `increment` / `decrement` / `reset` / `CountChanged`。

详见 [job3/README.md](job3/README.md)。

## job4: 同一合约的双 abigen 绑定

用 Foundry 和 Hardhat 3 分别编译内容完全相同的 `Job4` 计数器合约，各自导出 ABI/BIN 后生成**两套** abigen 绑定（`foundry_bind_go/` 的 `FoundryBindGo`、`hardhat_bind_go/` 的 `HardhatBindGo`）。两套 ABI 一致，因此 `main.go` 面向统一的 `counter` 接口编程，通过 `-bind foundry|hardhat` 切换实际使用的绑定。

模式：`deploy` / `getCount` / `increment` / `decrement` / `reset` / `CountChanged`。

详见 [job4/README.md](job4/README.md)。

---

## 四个项目对比

| | job1 | job2 | job3 | job4 |
| --- | --- | --- | --- | --- |
| 定位 | 基础链上交互入门 | 手动解析 ABI 与合约交互 | abigen 生成绑定的工程实践 | 多工具链产物生成双绑定 |
| 是否涉及智能合约 | 否 | 是（Job2 计数器） | 是（Job3 计数器） | 是（Job4 计数器） |
| 合约开发工具链 | – | Foundry + Hardhat 3（双实现） | solc 直接编译 | Foundry + Hardhat 3（同源码双编译） |
| 绑定方式 | – | 无绑定，`abi.JSON` 手动编解码 | 1 套 abigen 绑定 | 2 套 abigen 绑定，运行时 `-bind` 切换 |
| 合约调用方式 | 不调用合约，直接发转账交易 | 手工对齐 ABI 顺序 pack/unpack | 强类型方法调用 | 统一 `counter` 接口 + 适配器收敛两套绑定差异 |
| 合约部署 | 不涉及 | forge script / Ignition 部署 | CLI 内 `DeployGoJob3` | CLI 内按所选绑定 `DeployXxx` |
| 事件解析 | 无事件 | 按 topic/data 手动解码 | 绑定的 Parse 方法自动解码 | 统一接口的 Parse 方法自动解码 |
| 类型安全 | – | 弱：参数靠 `[]interface{}` | 强：编译期类型检查 | 强：编译期类型检查 |
| 适用场景 | 理解区块、交易、签名等底层概念 | 理解 ABI 编码原理 | 生产环境推荐的工程化做法 | 多工具链产物复用与接口抽象 |

简单来说：**job1** 打基础（区块与转账），**job2** 深入原理（手动 ABI 编解码），**job3** 走工程实践（abigen 自动生成类型安全的绑定），**job4** 进一步演练"多来源产物 + 接口抽象"的组合方式。

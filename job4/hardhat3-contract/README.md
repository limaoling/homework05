# Job4 合约项目（Hardhat 3 + Mocha + Ethers）

基于 Hardhat 3 的示例项目，使用 `mocha` 编写 TypeScript 测试、`ethers` 进行链上交互。合约为一个简单的计数器 `Job4`（递增、递减、重置，并触发 `CountChanged` 事件），与 `../foundry-contract/src/Job4.sol` 内容完全相同。本目录的编译产物（ABI/BIN）用于生成上级目录 `../hardhat_bind_go/` 中的 abigen Go 绑定。

## 项目结构

```text
contracts/Job4.sol              Solidity 合约源码
contracts/Job4.t.sol            Solidity 单元测试
test/Job4.ts                    TypeScript 集成测试（mocha + ethers）
scripts/job4-tx.ts              独立脚本：部署并依次调用 increment/decrement/reset
ignition/modules/Job4.ts        Hardhat Ignition 部署模块
artifacts/                      编译产物（含 ABI 与 bytecode）
contracts_Job4_sol_Job4.abi     导出的合约 ABI（供 abigen 使用）
contracts_Job4_sol_Job4.bin     导出的合约字节码（供 abigen 使用）
hardhat.config.ts               配置：Solidity 版本与网络
```

## 合约接口

| 方法 | 说明 |
| --- | --- |
| `getCount()` view | 返回当前计数 |
| `increment()` | 计数 +1，触发 `CountChanged` |
| `decrement()` | 计数 -1，触发 `CountChanged`；计数为 0 时回退（`count is zero`） |
| `reset()` | 计数归零，触发 `CountChanged` |

事件：`CountChanged(uint256 newCount, address indexed caller)`。

## 使用

### 安装依赖

```shell
npm install
```

### 运行测试

```shell
npx hardhat test              # 全部测试（Solidity + TypeScript）
npx hardhat test solidity     # 仅 Solidity 测试
npx hardhat test mocha        # 仅 TypeScript 测试
```

### 部署

本地常驻节点上用 Ignition 部署：

```shell
# 终端 1：启动常驻节点（默认 RPC http://127.0.0.1:8545，自带测试账户）
npx hardhat node

# 终端 2：部署
npx hardhat ignition deploy ./ignition/modules/Job4.ts --network localhost
```

也可以用内置的模拟网络直接跑交互脚本（脚本会自行部署一个新合约）：

```shell
npx hardhat run scripts/job4-tx.ts --network hardhatOp
```

部署到 Sepolia 需要一个有资金的账户。配置中使用了 `SEPOLIA_RPC_URL` / `SEPOLIA_PRIVATE_KEY` 配置变量，可通过 keystore 设置：

```shell
npx hardhat keystore set SEPOLIA_PRIVATE_KEY
npx hardhat keystore set SEPOLIA_RPC_URL
npx hardhat ignition deploy --network sepolia ignition/modules/Job4.ts
```

### 交互脚本

`scripts/job4-tx.ts` 会部署合约并依次调用 `increment` / `decrement` / `reset`，打印每次的计数值：

```shell
npx hardhat run scripts/job4-tx.ts --network hardhatOp
```

### 导出 ABI 与字节码

从编译产物中提取 ABI 和部署字节码（文件名沿用 solc 的 `<源码文件>_<合约名>` 规则），供 `abigen` 生成 Go 绑定（即上级目录的 `hardhat_bind_go/`，生成命令见其 README）：

```shell
jq -c '.abi'      artifacts/Job4.sol/Job4.json > contracts_Job4_sol_Job4.abi
jq -r '.bytecode' artifacts/Job4.sol/Job4.json > contracts_Job4_sol_Job4.bin
```

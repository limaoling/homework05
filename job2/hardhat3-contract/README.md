# Job2 合约项目（Hardhat 3 + Mocha + Ethers）

基于 Hardhat 3 的示例项目，使用 `mocha` 编写 TypeScript 测试、`ethers` 进行链上交互。合约为一个简单的计数器 `Job2`（递增、递减、重置，并触发 `CountChanged` 事件）。

## 项目结构

```
contracts/        Solidity 源码（Job2.sol）及 Solidity 单元测试（*.t.sol）
test/             TypeScript 集成测试（mocha + ethers）
scripts/          独立脚本（npx hardhat run 执行），如 job2-tx.ts
ignition/         Hardhat Ignition 部署模块
hardhat.config.ts
```

## 使用

### 运行测试

```shell
npx hardhat test              # 全部测试
npx hardhat test solidity     # 仅 Solidity 测试
npx hardhat test mocha        # 仅 TypeScript 测试
```

### 部署

本地网络部署：

```shell
# 先启动常驻节点：npx hardhat node
npx hardhat ignition deploy ignition/modules/Job2.ts --network localhost
```

部署到 Sepolia 需要一个有资金的账户。配置中使用了 `SEPOLIA_PRIVATE_KEY` 配置变量，可通过 keystore 设置：

```shell
npx hardhat keystore set SEPOLIA_PRIVATE_KEY
npx hardhat ignition deploy --network sepolia ignition/modules/Job2.ts
```

### 交互脚本

部署合约并依次调用 `increment` / `decrement` / `reset`，打印每次的计数值：

```shell
npx hardhat run scripts/job2-tx.ts --network localhost
```

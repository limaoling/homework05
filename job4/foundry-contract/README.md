# Job4 合约项目（Foundry）

使用 Foundry 工具链开发的 `Job4` 计数器合约（递增、递减、重置，并触发 `CountChanged` 事件）。本目录的编译产物（ABI/BIN）用于生成上级目录 `../foundry_bind_go/` 中的 abigen Go 绑定。

## 项目结构

```text
src/Job4.sol              合约源码
script/DeployJob4.s.sol   部署脚本
test/Job4.t.sol           Solidity 测试
out/                      forge build 产物（含 ABI）
broadcast/                forge script 部署记录
Job4.abi / Job4.bin       导出的合约 ABI 与字节码（供 abigen 使用）
```

> Foundry 由四部分组成：**Forge**（测试框架）、**Cast**（链上交互瑞士军刀）、**Anvil**（本地节点）、**Chisel**（Solidity REPL）。文档见 https://book.getfoundry.sh/

## 合约接口

| 方法 | 说明 |
| --- | --- |
| `getCount()` view | 返回当前计数 |
| `increment()` | 计数 +1，触发 `CountChanged` |
| `decrement()` | 计数 -1，触发 `CountChanged`；计数为 0 时回退（`count is zero`） |
| `reset()` | 计数归零，触发 `CountChanged` |

事件：`CountChanged(uint256 newCount, address indexed caller)`。

## 使用

### 编译

```shell
forge build
```

### 测试 / 格式化 / Gas 快照

```shell
forge test
forge fmt
forge snapshot
```

### 启动本地节点并部署

```shell
# 终端 1：启动本地节点
anvil

# 终端 2：用 Anvil 默认账户 #0 部署
forge script script/DeployJob4.s.sol:DeployJob4 \
  --rpc-url http://127.0.0.1:8545 \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --broadcast
```

部署成功后，终端会打印合约地址；部署记录保存在 `broadcast/` 目录下。

### 用 Cast 与合约交互

```shell
# 查询计数
cast call <JOB4_ADDRESS> "getCount()(uint256)" --rpc-url http://127.0.0.1:8545

# 递增
cast send <JOB4_ADDRESS> "increment()" --private-key <PRIVATE_KEY> --rpc-url http://127.0.0.1:8545
```

### 导出 ABI 与字节码

从编译产物中提取 ABI 和部署字节码，供 `abigen` 生成 Go 绑定（即上级目录的 `foundry_bind_go/`，生成命令见其 README）：

```shell
forge inspect Job4 abi      > Job4.abi
forge inspect Job4 bytecode > Job4.bin
```

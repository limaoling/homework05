# job3: 使用 abigen Go 绑定交互智能合约

使用 `abigen` 生成的强类型 Go 绑定，完成 `Job3` 计数器合约的**部署、调用与事件解析**。与 job2 的区别在于：不再手写 ABI 编解码，而是像调用普通 Go 方法一样操作合约。

## 项目结构

```text
job3/
├── Job3.sol               Solidity 计数器合约源码
├── Job3_sol_Job3.abi      合约 ABI
├── Job3_sol_Job3.bin      合约字节码
├── go_job3/job3.go        abigen 生成的 Go 合约绑定
└── main.go                CLI：基于生成绑定部署和调用合约
```

### `.abi` / `.bin` 文件的由来

`Job3_sol_Job3.abi` 和 `Job3_sol_Job3.bin` 是用 `solc` 编译 `Job3.sol` 得到的（文件名规则为 `<源码文件>_<合约名>.abi/.bin`）：

```bash
# 需要先安装 solc（如 brew install solidity）
solc --abi --bin Job3.sol -o . --overwrite
```

- `--abi`：输出合约的应用二进制接口，abigen 据此生成方法/事件签名；
- `--bin`：输出的 EVM 字节码，用于 `deploy` 模式部署新合约；
- `-o .`：输出到当前目录；`--overwrite`：覆盖旧文件。

### 生成 Go 绑定

拿到 `.abi` / `.bin` 后，用 `abigen` 生成 `go_job3/job3.go`（合约或 ABI 变更后需重新执行）：

```bash
abigen --abi Job3_sol_Job3.abi \
       --bin Job3_sol_Job3.bin \
       --pkg  go_job3 \
       --type GoJob3 \
       --out  go_job3/job3.go
```

## 功能与模式

| `-mode` | 说明 | 需要私钥 |
| --- | --- | --- |
| `deploy` | 用绑定的 `DeployGoJob3` 部署新合约，打印合约地址 | 是 |
| `getCount` | 只读调用 `getCount()` 查询当前计数 | 否 |
| `increment` / `decrement` / `reset` | 发交易修改计数器状态 | 是 |
| `CountChanged` | 按交易哈希解析 `CountChanged` 事件 | 否 |

### 参数与环境变量

| 参数 / 变量 | 说明 |
| --- | --- |
| `-mode` | 操作模式，默认 `getCount` |
| `-contract` / `CONTRACT_ADDRESS` | Job3 合约地址 |
| `-tx` / `TX_HASH` | 交易哈希（`CountChanged` 模式） |
| `-private-key` / `PRIVATE_KEY` | 签名私钥（deploy 和写交易必填；注意本目录没有默认测试私钥） |
| `ETH_RPC_URL` | 节点 RPC 地址，默认 `http://127.0.0.1:8545` |

命令行参数优先于环境变量。

## 运行步骤

1. 启动本地链：`anvil` 或 `npx hardhat node`；
2. 进入 `job3` 目录，部署合约：

   ```bash
   go run main.go -mode deploy -private-key <PRIVATE_KEY>
   # 输出中的"合约地址"即后续所有操作的 -contract 参数
   ```

3. 常用命令：

```bash
# 查询计数（不需要私钥）
go run main.go -mode getCount -contract <JOB3_CONTRACT_ADDRESS>

# 写交易（需要私钥；decrement 在 count=0 时会失败）
go run main.go -mode increment -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
go run main.go -mode decrement -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
go run main.go -mode reset     -contract <JOB3_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>

# 解析某笔交易的 CountChanged 事件
go run main.go -mode CountChanged -contract <JOB3_CONTRACT_ADDRESS> -tx <TX_HASH>
```

用环境变量简化重复参数：

```bash
export ETH_RPC_URL=http://127.0.0.1:8545
export PRIVATE_KEY=<PRIVATE_KEY>
export CONTRACT_ADDRESS=<JOB3_CONTRACT_ADDRESS>

go run main.go -mode deploy
go run main.go -mode increment
```

## 实现要点

- **绑定即对象**：`NewGoJob3(addr, client)` 把已部署合约包装成 Go 对象；只读调用走 `GetCount(&bind.CallOpts{})`，写交易走 `Increment/Decrement/Reset(auth)`；
- **签名配置**：`bind.NewKeyedTransactorWithChainID(privateKey, chainID)` 一次构造，部署与写交易共用；
- **部署返回地址**：`DeployGoJob3` 返回的地址要等交易上链后才有效，代码里用 `bind.WaitMined` 等待后从回执的 `ContractAddress` 取最终地址；
- **事件解析**：`contract.ParseCountChanged(log)` 直接把日志解码为结构体字段（`NewCount`、`Caller`），非本事件日志解析失败会被跳过，不影响同一交易内的其它日志。

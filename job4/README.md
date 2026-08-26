# job4: 同一合约、两套工具链产物生成的 abigen 绑定

使用 `abigen` 为同一个 `Job4` 计数器合约生成**两套** Go 绑定——分别来自 Foundry 和 Hardhat 两条工具链的编译产物，并用一个 CLI 演示：两套 ABI 完全一致时，业务代码可以面向统一接口编程，通过参数切换实际使用的绑定。

## 项目结构

```text
job4/
├── main.go                            CLI：部署、调用合约并解析事件
├── go.mod / go.sum                    独立 Go module（module job4）
├── foundry-contract/                  Foundry 工具链合约项目（详见其 README.md）
│   ├── src/Job4.sol                   合约源码
│   ├── Job4.abi / Job4.bin            从 forge 编译产物导出的 ABI 与字节码
│   └── ...
├── hardhat3-contract/                 Hardhat 3 合约项目（详见其 README.md）
│   ├── contracts/Job4.sol             合约源码（与 Foundry 版完全相同）
│   ├── contracts_Job4_sol_Job4.abi    从 Hardhat 编译产物导出的 ABI 与字节码
│   └── ...
├── foundry_bind_go/foundry-job4.go    abigen 生成的绑定（pkg=foundry_bind_go, type=FoundryBindGo）
└── hardhat_bind_go/hardhat-job4.go    abigen 生成的绑定（pkg=hardhat_bind_go, type=HardhatBindGo）
```

## 两套绑定的来历

两个合约目录中的 `Job4.sol` 内容一致（`getCount`/`increment`/`decrement`/`reset` + `CountChanged` 事件），因此两套 ABI 完全相同，只是生成绑定时使用了不同工具链的产物：

| | foundry_bind_go | hardhat_bind_go |
| --- | --- | --- |
| ABI/BIN 来源 | `foundry-contract/Job4.abi`、`Job4.bin`（forge 导出） | `hardhat3-contract/contracts_Job4_sol_Job4.abi`、`.bin`（Hardhat 编译产物） |
| 包名 / 类型 | `foundry_bind_go` / `FoundryBindGo` | `hardhat_bind_go` / `HardhatBindGo` |

在 `job4/` 目录下用 `abigen` 重新生成（合约或 ABI 变更后需重新执行）：

```bash
abigen --abi foundry-contract/Job4.abi \
       --bin foundry-contract/Job4.bin \
       --pkg  foundry_bind_go \
       --type FoundryBindGo \
       --out  foundry_bind_go/foundry-job4.go

abigen --abi hardhat3-contract/contracts_Job4_sol_Job4.abi \
       --bin hardhat3-contract/contracts_Job4_sol_Job4.bin \
       --pkg  hardhat_bind_go \
       --type HardhatBindGo \
       --out  hardhat_bind_go/hardhat-job4.go
```

## 功能与模式

`main.go` 定义了统一的 `counter` 接口，把两套绑定各自的 `GetCount`/`Increment`/`Decrement`/`Reset`/`ParseCountChanged` 收敛到一起；`-bind` 参数决定底层用的是哪一套。

| `-mode` | 说明 | 需要私钥 |
| --- | --- | --- |
| `deploy` | 调用所选绑定的 `DeployXxx` 部署新合约，打印合约地址 | 是 |
| `getCount` | 只读调用 `getCount()` 查询当前计数 | 否 |
| `increment` / `decrement` / `reset` | 发交易修改计数器状态 | 是 |
| `CountChanged` | 按交易哈希解析回执中属于该合约的 `CountChanged` 事件 | 否 |

### 参数与环境变量

| 参数 / 变量 | 说明 |
| --- | --- |
| `-mode` | 操作模式，默认 `getCount` |
| `-bind` | 使用哪套生成绑定：`foundry`（默认）或 `hardhat` |
| `-contract` / `CONTRACT_ADDRESS` | Job4 合约地址 |
| `-tx` / `TX_HASH` | 交易哈希（`CountChanged` 模式） |
| `-private-key` / `PRIVATE_KEY` | 签名私钥（deploy 和写交易必填） |
| `ETH_RPC_URL` | 节点 RPC 地址，默认 `http://127.0.0.1:8545` |

命令行参数优先于环境变量。

## 运行步骤

1. 启动本地链：`anvil` 或 `npx hardhat node`；
2. 进入 `job4/` 目录（本目录是独立 module），部署合约：

   ```bash
   go run main.go -bind foundry -mode deploy -private-key <PRIVATE_KEY>
   # 输出中的"合约地址"即后续所有操作的 -contract 参数
   ```

3. 常用命令：

```bash
# 查询计数（不需要私钥）
go run main.go -bind foundry -mode getCount -contract <JOB4_CONTRACT_ADDRESS>

# 写交易（需要私钥；decrement 在 count=0 时会失败）
go run main.go -bind hardhat -mode increment -contract <JOB4_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
go run main.go -bind hardhat -mode decrement -contract <JOB4_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>
go run main.go -bind hardhat -mode reset     -contract <JOB4_CONTRACT_ADDRESS> -private-key <PRIVATE_KEY>

# 解析某笔交易的 CountChanged 事件
go run main.go -bind foundry -mode CountChanged -contract <JOB4_CONTRACT_ADDRESS> -tx <TX_HASH>
```

用环境变量简化重复参数：

```bash
export ETH_RPC_URL=http://127.0.0.1:8545
export PRIVATE_KEY=<PRIVATE_KEY>
export CONTRACT_ADDRESS=<JOB4_CONTRACT_ADDRESS>

go run main.go -mode deploy
go run main.go -mode increment
```

> 注意：两套绑定只是"生成来源"不同，部署出来的合约字节码一致。用 `-bind foundry` 部署的合约，同样可以用 `-bind hardhat` 来查询和调用。

## 实现要点

- **接口收敛差异**：两套绑定各自生成的类型（`FoundryBindGo` / `HardhatBindGo`）互不相同，main.go 用 `counter` 接口 + 两个薄包装（`foundryCounter`、`hardhatCounter`）适配 `ParseCountChanged` 的返回类型，业务逻辑完全不感知绑定差异；
- **binding 注册表**：`bindings` map 把「如何部署」和「如何按地址包装成 counter」登记在一起，`-bind` 取值即 map 的 key，新增绑定只需加一条记录；
- **签名配置**：`bind.NewKeyedTransactorWithChainID(privateKey, chainID)` 一次构造，部署与写交易共用；只读的 `getCount` 不需要私钥；
- **部署返回地址**：`DeployXxx` 返回的地址要等交易上链后才有效，代码用 `bind.WaitMined` 等待后从回执的 `ContractAddress` 取最终地址；
- **事件解析**：`ParseCountChanged(log)` 直接把日志解码为结构体字段（`NewCount`、`Caller`），非本事件的日志解析失败会被跳过，同一交易里有其它合约日志也不影响结果。

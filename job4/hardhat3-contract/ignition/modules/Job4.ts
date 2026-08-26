import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("Job4Module", (m) => {
  const job4 = m.contract("Job4");

  return { job4 };
});
// 用常驻节点 部署命令:
// npx hardhat ignition deploy ./ignition/modules/Job4.ts --network localhost

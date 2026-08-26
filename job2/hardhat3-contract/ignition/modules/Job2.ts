import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("Job2Module", (m) => {
  const job2 = m.contract("Job2");

  return { job2 };
});
// 用常驻节点 部署命令:
// npx hardhat ignition deploy ./ignition/modules/Job2.ts --network localhost

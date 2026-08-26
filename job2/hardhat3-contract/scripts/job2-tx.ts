import { network } from "hardhat";
// npx hardhat run scripts/job2-tx.ts --network hardhatOp
const { ethers } = await network.create();

console.log("Interacting with the Job2 contract");

const [sender] = await ethers.getSigners();

const job2 = await ethers.deployContract("Job2");
await job2.waitForDeployment();
console.log("Job2 deployed at", await job2.getAddress());

console.log("Calling increment()");
let tx = await job2.connect(sender).increment();
await tx.wait();
console.log("count =", (await job2.getCount()).toString());

console.log("Calling decrement()");
tx = await job2.decrement();
await tx.wait();
console.log("count =", (await job2.getCount()).toString());

console.log("Calling reset()");
tx = await job2.reset();
await tx.wait();
console.log("count =", (await job2.getCount()).toString());

console.log("Done");

import { network } from "hardhat";
// npx hardhat run scripts/job4-tx.ts --network hardhatOp
const { ethers } = await network.create();

console.log("Interacting with the Job4 contract");

const [sender] = await ethers.getSigners();

const job4 = await ethers.deployContract("Job4");
await job4.waitForDeployment();
console.log("Job4 deployed at", await job4.getAddress());

console.log("Calling increment()");
let tx = await job4.connect(sender).increment();
await tx.wait();
console.log("count =", (await job4.getCount()).toString());

console.log("Calling decrement()");
tx = await job4.decrement();
await tx.wait();
console.log("count =", (await job4.getCount()).toString());

console.log("Calling reset()");
tx = await job4.reset();
await tx.wait();
console.log("count =", (await job4.getCount()).toString());

console.log("Done");

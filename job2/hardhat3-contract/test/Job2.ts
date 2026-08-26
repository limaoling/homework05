import { expect } from "chai";
import { network } from "hardhat";

const { ethers } = await network.create();

describe("Job2", function () {
  it("Should start with count = 0", async function () {
    const job2 = await ethers.deployContract("Job2");

    expect(await job2.getCount()).to.equal(0n);
  });

  it("Should emit the CountChanged event when calling the increment() function", async function () {
    const [owner] = await ethers.getSigners();
    const job2 = await ethers.deployContract("Job2");

    await expect(job2.increment())
      .to.emit(job2, "CountChanged")
      .withArgs(1n, owner.address);
  });

  it("increment() and decrement() should update getCount()", async function () {
    const job2 = await ethers.deployContract("Job2");

    await job2.increment();
    expect(await job2.getCount()).to.equal(1n);

    await job2.decrement();
    expect(await job2.getCount()).to.equal(0n);
  });

  it("decrement() should revert when count is zero", async function () {
    const job2 = await ethers.deployContract("Job2");

    await expect(job2.decrement()).to.be.revertedWith("count is zero");
  });

  it("reset() should set count back to 0 and emit CountChanged", async function () {
    const [owner] = await ethers.getSigners();
    const job2 = await ethers.deployContract("Job2");

    await job2.increment();
    await expect(job2.reset())
      .to.emit(job2, "CountChanged")
      .withArgs(0n, owner.address);

    expect(await job2.getCount()).to.equal(0n);
  });
});

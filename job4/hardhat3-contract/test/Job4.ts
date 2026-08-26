import { expect } from "chai";
import { network } from "hardhat";

const { ethers } = await network.create();

describe("Job4", function () {
  it("Should start with count = 0", async function () {
    const job4 = await ethers.deployContract("Job4");

    expect(await job4.getCount()).to.equal(0n);
  });

  it("Should emit the CountChanged event when calling the increment() function", async function () {
    const [owner] = await ethers.getSigners();
    const job4 = await ethers.deployContract("Job4");

    await expect(job4.increment())
      .to.emit(job4, "CountChanged")
      .withArgs(1n, owner.address);
  });

  it("increment() and decrement() should update getCount()", async function () {
    const job4 = await ethers.deployContract("Job4");

    await job4.increment();
    expect(await job4.getCount()).to.equal(1n);

    await job4.decrement();
    expect(await job4.getCount()).to.equal(0n);
  });

  it("decrement() should revert when count is zero", async function () {
    const job4 = await ethers.deployContract("Job4");

    await expect(job4.decrement()).to.be.revertedWith("count is zero");
  });

  it("reset() should set count back to 0 and emit CountChanged", async function () {
    const [owner] = await ethers.getSigners();
    const job4 = await ethers.deployContract("Job4");

    await job4.increment();
    await expect(job4.reset())
      .to.emit(job4, "CountChanged")
      .withArgs(0n, owner.address);

    expect(await job4.getCount()).to.equal(0n);
  });
});

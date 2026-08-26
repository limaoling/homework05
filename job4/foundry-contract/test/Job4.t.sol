// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Job4} from "../src/Job4.sol";
import {Test} from "forge-std/Test.sol";

contract Job4Test is Test {
    Job4 job4;

    function setUp() public {
        job4 = new Job4();
    }

    function test_InitialValue() public view {
        require(job4.getCount() == 0, "Initial value should be 0");
    }

    function test_Increment() public {
        job4.increment();
        require(job4.getCount() == 1, "Value after calling increment once should be 1");
    }

    function test_Decrement() public {
        job4.increment();
        job4.decrement();
        require(job4.getCount() == 0, "Value after increment then decrement should be 0");
    }

    function test_Reset() public {
        job4.increment();
        job4.increment();
        job4.reset();
        require(job4.getCount() == 0, "Value after reset should be 0");
    }

    function testFuzz_Increment(uint8 x) public {
        for (uint8 i = 0; i < x; i++) {
            job4.increment();
        }
        require(job4.getCount() == x, "Value after calling increment x times should be x");
    }

    function test_DecrementByZero() public {
        vm.expectRevert(bytes("count is zero"));
        job4.decrement();
    }

    function test_CountChangedEvent() public {
        vm.expectEmit(true, true, true, true);
        emit Job4.CountChanged(1, address(this));
        job4.increment();
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Job2} from "./Job2.sol";
import {Test} from "forge-std/Test.sol";

// Solidity tests are compatible with foundry, so they
// use the same syntax and offer the same functionality.

contract Job2Test is Test {
    Job2 job2;

    function setUp() public {
        job2 = new Job2();
    }

    function test_InitialValue() public view {
        require(job2.getCount() == 0, "Initial value should be 0");
    }

    function test_Increment() public {
        job2.increment();
        require(job2.getCount() == 1, "Value after calling increment once should be 1");
    }

    function test_Decrement() public {
        job2.increment();
        job2.decrement();
        require(job2.getCount() == 0, "Value after increment then decrement should be 0");
    }

    function test_Reset() public {
        job2.increment();
        job2.increment();
        job2.reset();
        require(job2.getCount() == 0, "Value after reset should be 0");
    }

    function testFuzz_Increment(uint8 x) public {
        for (uint8 i = 0; i < x; i++) {
            job2.increment();
        }
        require(job2.getCount() == x, "Value after calling increment x times should be x");
    }

    function test_DecrementByZero() public {
        vm.expectRevert();
        job2.decrement();
    }

    function test_CountChangedEvent() public {
        vm.expectEmit(true, true, true, true);
        emit Job2.CountChanged(1, address(this));
        job2.increment();
    }
}

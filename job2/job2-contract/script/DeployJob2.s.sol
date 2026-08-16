// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script} from "forge-std/Script.sol";
import {Job2} from "../src/Job2.sol";

contract DeployJob2 is Script {
    function run() external returns (Job2 job2) {
        vm.startBroadcast();
        job2 = new Job2();
        vm.stopBroadcast();
    }
}

// Execute example:
// forge script script/DeployJob2.s.sol:DeployJob2 \
//   --rpc-url http://127.0.0.1:8545 \
//   --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
//   --broadcast

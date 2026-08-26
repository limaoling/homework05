// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 一个简单的计数器合约，用于演示基本状态更改和事件触发。This file is a generated binding and any manual changes will be lost.
/// @notice 此合约允许递增、递减和重置计数器。 Failed to generate ABI binding: 4:9: expected 'IDENT', found 'go'
contract Job3 {
    uint256 private count;

    /// @notice 当计数器值改变时触发。
    /// @param newCount 计数器的新值。
    /// @param caller 触发此更改的地址。
    event CountChanged(uint256 newCount, address indexed caller);

    /// @notice 返回当前计数器的值。
    /// @return 当前计数器的值。
    function getCount() public view returns (uint256) {
        return count;
    }

    /// @notice 将计数器递增 1。
    /// @dev 触发 CountChanged 事件。
    function increment() public {
        count += 1;
        emit CountChanged(count, msg.sender);
    }

    /// @notice 将计数器递减 1。
    /// @dev 如果计数器为 0 则回退。触发 CountChanged 事件。
    function decrement() public {
        require(count > 0, "count is zero");
        count -= 1;
        emit CountChanged(count, msg.sender);
    }

    /// @notice 将计数器重置为 0。
    /// @dev 触发 CountChanged 事件。
    function reset() public {
        count = 0;
        emit CountChanged(count, msg.sender);
    }
}

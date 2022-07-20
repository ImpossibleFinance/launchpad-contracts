// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract MockMessageBusSender {

    event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee);

    constructor() {}

    function sendMessage(
        address _receiver,
        uint256 _dstChainId,
        bytes memory _message
    ) external payable {
        emit Message(msg.sender, _receiver, _dstChainId, _message, msg.value);
    }
}

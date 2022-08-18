// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract MockMessageBusSender {
    uint256 public feeBase = 1;
    uint256 public feePerByte = 1;

    event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee);

    constructor() {}

    function sendMessage(
        address _receiver,
        uint256 _dstChainId,
        bytes memory _message
    ) external payable {
        emit Message(msg.sender, _receiver, _dstChainId, _message, msg.value);
    }

    function calcFee(bytes calldata _message) public view returns (uint256) {
        uint256 fee = feeBase + _message.length * feePerByte;
        return fee;
    }
}

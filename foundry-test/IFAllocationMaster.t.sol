// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../contracts/IFAllocationMaster.sol";
import '../contracts/interfaces/IIFBridgableStakeWeight.sol';
import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import 'sgn-v2-contracts/contracts/message/libraries/MessageSenderLib.sol';
import 'sgn-v2-contracts/contracts/message/interfaces/IMessageBus.sol';
import './helper-contracts/MockMessageBusSender.sol';

contract ContractTest is Test {
    IFAllocationMaster internal ifAllocationMaster;
    address internal idiaAddr = 0x0b15Ddf19D47E6a86A56148fb4aFFFc6929BcB89;
    MockMessageBusSender internal mockMessageBusSender;
    IMessageBus internal messageBus;
    ExpectEmit emitter;

    constructor() {
        mockMessageBusSender = new MockMessageBusSender();
        ifAllocationMaster = new IFAllocationMaster(address(mockMessageBusSender));
        ifAllocationMaster.addTrack(
            "Test Track",ERC20(idiaAddr),1, 1, 1, 1000
        );
        messageBus = IMessageBus(address(mockMessageBusSender));
        emitter = new ExpectEmit();
    }
    function setUp() public {
    }

    receive() external payable {
    }

    function syncUser(address[] memory users) internal {
        uint80 timestamp = uint80(block.timestamp);
        uint192[] memory userStakeWeights = new uint192[](users.length);

        bytes memory message = abi.encode(
            IIFBridgableStakeWeight.MessageRequest({
                bridgeType: IIFBridgableStakeWeight.BridgeType.UserWeight,
                users: users,
                timestamp: timestamp,
                weights: userStakeWeights,
                trackId: 0
            })
        );

        uint256 fee = 100000;
        vm.expectEmit(true, false, false, false);
        ifAllocationMaster.syncUserWeight{value: fee}(
            address(0),
            users,
            0,
            timestamp,
            1
        );
        emitter.syncUserWeight(
            address(ifAllocationMaster),
            address(0),
            1,
            message,
            0
        );
    }

    function testSyncOneUser() public {
        address[] memory users = new address[](1);
        users[0] = 0x1f1BDFE288a8C9ac31F1f7C70dfEE6c82EDF77f6;
        syncUser(users);
    }

    function testSyncHundredUsers() public {
        address[] memory users = new address[](100);
        for (uint i = 0; i < 100; i += 1) {
            users[i] = 0x1f1BDFE288a8C9ac31F1f7C70dfEE6c82EDF77f6;
        }
        syncUser(users);
    }

    function testSyncThousandUsers() public {
        address[] memory users = new address[](1000);
        for (uint i = 0; i < 1000; i += 1) {
            users[i] = 0x1f1BDFE288a8C9ac31F1f7C70dfEE6c82EDF77f6;
        }
        syncUser(users);
    }

    function testSyncTotalWeight() public {
        address[] memory users = new address[](1);
        users[0] = address(this);
        uint80 timestamp = uint80(block.timestamp);
        uint192[] memory userStakeWeights = new uint192[](1);
        userStakeWeights[0] = ifAllocationMaster.getTotalStakeWeight(0, timestamp);

        bytes memory message = abi.encode(
            IIFBridgableStakeWeight.MessageRequest({
                bridgeType: IIFBridgableStakeWeight.BridgeType.TotalWeight,
                users: users,
                timestamp: timestamp,
                weights: userStakeWeights,
                trackId: 0
            })
        );

        vm.expectEmit(true, false, false, false);
        ifAllocationMaster.syncTotalWeight{value: 100000}(
            address(0),
            0,
            timestamp,
            1
        );
        emitter.syncUserWeight(
            address(ifAllocationMaster),
            address(0),
            1,
            message,
            0
        );
    }
}

contract ExpectEmit {
    event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee);

    function syncUserWeight(
        address sender,
        address receiver,
        uint256 dstChainId,
        bytes memory message,
        uint256 fee
    ) public {
        emit Message(sender, receiver, dstChainId, message, fee);
    }
}
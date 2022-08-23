// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.9;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../contracts/LoyaltyCardMaster.sol";
import "../contracts/LoyaltyCardRewarder.sol";
import "../contracts/LoyaltyRewardsLookup.sol";
import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import "@openzeppelin/contracts/utils/Strings.sol";
import 'sgn-v2-contracts/contracts/message/libraries/MessageSenderLib.sol';
import 'sgn-v2-contracts/contracts/message/interfaces/IMessageBus.sol';

contract ContractTest is Test {
    LoyaltyCardMaster internal loyaltyCardMaster;
    LoyaltyCardRewarder internal loyaltyCardRewarder;
    LoyaltyRewardsLookup internal loyaltyRewardsLookup;
    address internal user = 0x1f1BDFE288a8C9ac31F1f7C70dfEE6c82EDF77f6;
    uint256 internal numUsers = 100;
    address[] internal users = new address[](numUsers);
    mapping(uint256 => string) internal testCodes;
    uint256[][] internal credCodes = new uint256[][](numUsers);
    string[][] internal credNames = new string[][](numUsers);


    constructor() {
        loyaltyCardMaster = new LoyaltyCardMaster("IFCard", "IFCard");
        loyaltyRewardsLookup = new LoyaltyRewardsLookup();
        loyaltyCardRewarder = new LoyaltyCardRewarder(address(loyaltyCardMaster), address(loyaltyRewardsLookup));
        loyaltyCardMaster.addOperator(address(loyaltyCardRewarder));
        loyaltyCardMaster.setMinter(address(this));
        for (uint i = 1; i < 11; ++i) {
            testCodes[i] = Strings.toString(i);
            loyaltyRewardsLookup.setCredential(i, i, testCodes[i]);
        }
        loyaltyCardMaster.mint(user);
        for (uint i = 0; i < numUsers; i += 1) {
            users[i] = user;
        }
        for (uint i = 0; i < users.length; ++i) {
            credCodes[i] = [1,2,3,4,5,6,7,8,9, 10];
            credNames[i] = ["1","2","3","4","5","6","7","8","9", "10"];
        }
    }

    function testSingleUser() public {
        loyaltyCardRewarder.rewardAccount(user, 1, "1");
    }

    function testBatch() public {
        loyaltyCardRewarder.rewardBatchSingleCredential(users, 1, "1");
    }


    function testSingleUserMultiCreds() public {
        for (uint code = 1; code < 11; ++code) {
            loyaltyCardRewarder.rewardAccount(user, code, testCodes[code]);
        }
    }

    function testBundled() public {
        for (uint i; i < users.length; ++i) {
            loyaltyCardRewarder.rewardAccount(user, 1, "1");
        }
    }

    function testBundledMultiCreds() public {
        for (uint i; i < users.length; ++i) {
            for (uint code = 1; code < 11; ++code) {
                loyaltyCardRewarder.rewardAccount(user, code, testCodes[code]);
            }
        }
    }

    function testBatchMultiCreds() public {
        loyaltyCardRewarder.rewardBatchMultiCredentials(users, credCodes, credNames);
    }
}
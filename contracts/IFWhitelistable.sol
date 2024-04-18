// SPDX-License-Identifier: MIT
pragma solidity 0.8.9;

import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/security/ReentrancyGuard.sol';

abstract contract IFWhitelistable is Ownable, ReentrancyGuard {
    // optional whitelist setter (settable by owner)
    address public whitelistSetter;

    // whitelist merkle root; if not set, then sale is open to everyone that has allocation
    bytes32 public whitelistRootHash;

    event SetWhitelistSetter(address indexed whitelistSetter);
    event SetWhitelist(bytes32 indexed whitelistRootHash);

    constructor() {
        whitelistSetter = _msgSender();
    }


    // Throws if called by any account other than the whitelist setter.
    modifier onlyWhitelistSetterOrOwner() {
        require(
            _msgSender() == whitelistSetter || _msgSender() == owner(),
            'caller not whitelist setter or owner'
        );
        _;
    }

    // Function for owner to set an optional, separate whitelist setter
    function setWhitelistSetter(address _whitelistSetter) public onlyOwner {
        whitelistSetter = _whitelistSetter;

        emit SetWhitelistSetter(_whitelistSetter);
    }

    // Function for owner or whitelist setter to set a whitelist; if not set, then everyone allowed
    function setWhitelist(bytes32 _whitelistRootHash)
        public
        onlyWhitelistSetterOrOwner
    {
        whitelistRootHash = _whitelistRootHash;

        emit SetWhitelist(_whitelistRootHash);
    }

    // purchase function when there is a whitelist
    function whitelistedPurchase(uint256 paymentAmount, bytes32[] calldata merkleProof) virtual public {}

    function withdrawGiveaway(bytes32[] calldata merkleProof) virtual public nonReentrant {}
}
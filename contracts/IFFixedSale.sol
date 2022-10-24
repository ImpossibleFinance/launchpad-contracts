// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/utils/cryptography/MerkleProof.sol';
import './IFSale.sol';

contract IFFixedSale is IFSale {
    // CONSTRUCTOR

    constructor(
        uint256 _salePrice,
        address _funder,
        ERC20 _paymentToken,
        ERC20 _saleToken,
        uint24 _trackId,
        uint256 _startTime,
        uint256 _endTime,
        uint256 _maxTotalPayment
    ) 
        IFSale(
            _funder,
            _salePrice,
            _paymentToken,
            _saleToken,
            _trackId,
            _startTime,
            _endTime,
            _maxTotalPayment
        )
    {}

    // FUNCTIONS

    function purchase(uint256) virtual override public {
        revert("Use purchase(uint256 paymentAmount, bytes32[] calldata merkleProof, uint256 allocation)");
    }

    function whitelistedPurchase(uint256, bytes32[] calldata) override public {
        revert("Use purchase(uint256 paymentAmount, bytes32[] calldata merkleProof, uint256 allocation)");
    }

    function withdrawGiveaway(bytes32[] calldata) override public {
        revert("Use withdrawGiveaway(bytes32[] calldata merkleProof, uint256 allocation)");
    }

    // purchase function when there is a whitelist
    function whitelistedPurchase(
        uint256 paymentAmount,
        bytes32[] calldata merkleProof,
        uint256 allocation
    ) public {
        // require that user is whitelisted by checking proof
        require(checkWhitelist(_msgSender(), merkleProof, allocation), 'proof invalid');

        uint256 remaining = getMaxPayment(_msgSender(), allocation);
        _purchase(paymentAmount, remaining);
    }

    // Function for withdrawing purchased sale token after sale end
    function withdraw() override public nonReentrant {
        address user = _msgSender()
        // must not be a zero price sale
        require(salePrice != 0, 'use withdrawGiveaway');

        // send token and update states
        uint256 tokenOwed = getCurrentClaimableToken(claimable[user], totalPurchased[user], user);
        _withdraw(tokenOwed);
        // sale token owed must be greater than 0
        require(tokenOwed != 0, 'no token to be withdrawn');
    }   


    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveaway(bytes32[] calldata merkleProof, uint256 allocation)
        external
        nonReentrant
    {
        address user = _msgSender();
        // must be a zero price sale
        require(salePrice == 0, 'not a giveaway');
        // if there is whitelist, require that user is whitelisted by checking proof
        require(
            whitelistRootHash == 0 || checkWhitelist(user, merkleProof, allocation),
            'proof invalid'
        );

        uint256 saleTokenOwed = 0;
        // initialize claimable before the first time of withdrawal
        if (!hasWithdrawn[user]) {
            // each participant in the zero cost "giveaway" gets a flat amount of sale token
            // claimable[_msgSender()] = getUserStakeValue(_msgSender());
            saleTokenOwed = getCurrentClaimableToken(claimable[user], totalPurchased[user], user);
            claimable[user] = saleTokenOwed;
            totalPurchased[user] = saleTokenOwed;
        }

        // send token and update states
        _withdraw(saleTokenOwed);
        // sale token owed must be greater than 0
        require(saleTokenOwed != 0, 'withdraw giveaway amount 0');
    }

    // Returns true if user's allocation matches the one in merkle root, otherwise false
    function checkWhitelist(address user, bytes32[] calldata merkleProof, uint256 allocation)
        public
        view
        returns (bool)
    {
        // compute merkle leaf from input
        bytes32 leaf = keccak256(abi.encodePacked(user, allocation));

        // verify merkle proof
        return MerkleProof.verify(merkleProof, whitelistRootHash, leaf);
    }


    function getMaxPayment(address user, uint256 allocation) public view returns (uint256) {
        // get the maximum total payment for a user
        uint256 max = (salePrice * allocation) / SALE_PRICE_DECIMALS;
        if (maxTotalPayment < max) {
            max = maxTotalPayment;
        }

        // calculate and return remaining
        return max - paymentReceived[user];
    }
}

// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.9;

import '@openzeppelin/contracts/utils/cryptography/MerkleProof.sol';
import './IFSale.sol';


/**
  @title Sale contract with user's allocation stored in merkle root
  @notice Regular purchase from IFSale is disabled here
  @notice Apart from merkle proof, users or the frontend has to supply allocation amount for verification
  @notice That means functions having `merkleProof` as param will also need `allocation`
 */
contract IFFixedSale is IFSale {
    // --- CONSTRUCTOR

    constructor(
        uint256 _salePrice,
        address _funder,
        ERC20 _paymentToken,
        ERC20 _saleToken,
        uint256 _startTime,
        uint256 _endTime,
        uint256 _maxTotalPayment
    ) 
        IFSale(
            _funder,
            _salePrice,
            _paymentToken,
            _saleToken,
            _startTime,
            _endTime,
            _maxTotalPayment
        )
    {}

    bool public isVestedGiveaway = false;

    // allocation when the user is not whitelisted
    // note that any user can get publicAllocation regardless of their whitelisted allocation
    uint256 public publicAllocation = 0;

    // --- SETTER FUNCTIONS
    function setVestedGiveaway(bool _isVestedGiveaway) public onlyOwner onlyBeforeSale {
        isVestedGiveaway = _isVestedGiveaway;
    }

    function setPublicAllocation(uint256 _publicAllocation) public onlyWhitelistSetterOrOwner onlyBeforeSale {
        publicAllocation = _publicAllocation;
    }

    function setMaxTotalPurchasable(uint256 _maxTotalPurchasable) override public onlyWhitelistSetterOrOwner {
        maxTotalPurchasable = _maxTotalPurchasable * salePrice;

        require(maxTotalPurchasable >= saleTokenPurchased, 'Max purchasable should not be lower than the amount of token purchased');

        emit SetMaxTotalPurchasable(_maxTotalPurchasable);
    }

    // --- DISABLED FUNCTIONS

    function purchase(uint256) virtual override public {
        revert("Use whitelistedPurchase(uint256 paymentAmount, bytes32[] calldata merkleProof, uint256 allocation)");
    }

    function whitelistedPurchase(uint256, bytes32[] calldata) override public pure {
        revert("Use whitelistedPurchase(uint256 paymentAmount, bytes32[] calldata merkleProof, uint256 allocation)");
    }

    function withdrawGiveaway(bytes32[] calldata) override public pure {
        revert("Use withdrawGiveaway(bytes32[] calldata merkleProof, uint256 allocation)");
    }

    // --- WHITELISTED ACTIONS

    // purchase with code function when there is a whitelist
    function whitelistedPurchaseWithCode(
        uint256 paymentAmount,
        bytes32[] calldata merkleProof,
        uint256 _allocation,
        string calldata code
    ) public onlyDuringSale {
        uint256 allocation = publicAllocation;
        if (merkleProof.length > 0) {
            // require that user is whitelisted by checking proof
            require(checkWhitelist(_msgSender(), merkleProof, _allocation), 'proof invalid');
            if (_allocation > publicAllocation) {
                allocation = _allocation;
            }
        }

        uint256 remaining = getMaxPayment(_msgSender(), allocation);
        _purchaseWithCode(paymentAmount, remaining, code);
    }

    // purchase function when there is a whitelist
    function whitelistedPurchase(
        uint256 paymentAmount,
        bytes32[] calldata merkleProof,
        uint256 _allocation
    ) public onlyDuringSale {
        uint256 allocation = publicAllocation;
        if (merkleProof.length > 0) {
            // require that user is whitelisted by checking proof
            require(checkWhitelist(_msgSender(), merkleProof, _allocation), 'proof invalid');
            if (_allocation > publicAllocation) {
                allocation = _allocation;
            }
        }

        uint256 remaining = getMaxPayment(_msgSender(), allocation);
        _purchase(paymentAmount, remaining);
    }

    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveaway(bytes32[] calldata merkleProof, uint256 allocation)
        external
        nonReentrant
        onlyAfterSale
    {
        address user = _msgSender();
        // not vested giveaway
        require(isVestedGiveaway == false, 'use withdrawGiveawayVested');
        // must be a zero price sale
        require(salePrice == 0, 'not a giveaway');
        // can withdraw only once
        require(hasWithdrawn[user] == false, 'already withdrawn');
        // require that user is whitelisted by checking proof
        require(checkWhitelist(user, merkleProof, allocation), 'proof invalid');

        uint256 saleTokenOwed = 0;
        // initialize claimable before the first time of withdrawal
        if (!hasWithdrawn[user]) {
            // each participant in the zero cost "giveaway" gets a flat amount of sale token
            saleTokenOwed = allocation;
            claimable[user] = allocation;
            totalPurchased[user] = allocation;
        }

        // sale token owed must be greater than 0
        require(saleTokenOwed != 0, 'withdraw giveaway amount 0');

        // send token and update states
        _withdraw(saleTokenOwed);
    }

    // Function to withdraw (redeem) tokens from a zero cost "giveaway" sale
    function withdrawGiveawayVested(bytes32[] calldata merkleProof, uint256 allocation)
        external
        nonReentrant
        onlyAfterSale
    {
        address user = _msgSender();

        // not vested giveaway
        require(isVestedGiveaway == true, 'use withdrawGiveaway');
        // must be a zero price sale
        require(salePrice == 0, 'not a giveaway');
        // require that user is whitelisted by checking proof
        require(checkWhitelist(user, merkleProof, allocation), 'proof invalid');

        // initialize claimable before the first time of withdrawal
        if (!hasWithdrawn[user]) {
            claimable[user] = allocation;
            totalPurchased[user] = allocation;
        }

        uint256 tokenOwed = getCurrentClaimableToken(user);

        // sale token owed must be greater than 0
        require(tokenOwed != 0, 'withdraw giveaway amount 0');

        // send token and update states
        _withdraw(tokenOwed);
    }

    // --- HELPER FUNCTIONS

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

    // @dev get max payment from allocation
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

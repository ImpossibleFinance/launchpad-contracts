// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import 'hardhat/console.sol';
import '../../contracts/IFMerkleAllocationSale.sol';

// Contract to set max allocation on all buyers
contract MockIFMerkleAllocationSale is IFMerkleAllocationSale {
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
        IFMerkleAllocationSale(
            _salePrice,
            _funder,
            _paymentToken,
            _saleToken,
            _trackId,
            _startTime,
            _endTime,
            _maxTotalPayment
        )
    {}

    function purchase(uint256 paymentAmount) external {
        // Skip merkle check and set max allocation
        _purchase(paymentAmount, type(uint256).max);
    }
}

// SPDX-License-Identifier: MIT
pragma solidity 0.8.9;

import '../../contracts/IFFixedSale.sol';
import '../interfaces/IIFRetrievableStakeWeight.sol';

// Contract to set max allocation on all buyers
contract MockIFFixedSale is IFFixedSale {
    constructor(
        uint256 _salePrice,
        address _funder,
        ERC20 _paymentToken,
        ERC20 _saleToken,
        IIFRetrievableStakeWeight _allocationMaster,
        uint24 _trackId,
        uint80 _allocSnapshotTimestamp,
        uint256 _startTime,
        uint256 _endTime,
        uint256 _maxTotalPayment
    )
        IFFixedSale(
            _salePrice,
            _funder,
            _paymentToken,
            _saleToken,
            _startTime,
            _endTime,
            _maxTotalPayment
        )
    {}

    function purchase(uint256 paymentAmount) public override {
        // Skip merkle check and set max allocation
        _purchase(paymentAmount, type(uint256).max);
    }

    function purchaseWithCode(uint256 paymentAmount, string memory code)
        public
    {
        // Skip merkle check and set max allocation
        _purchaseWithCode(paymentAmount, type(uint256).max, code);
    }
}

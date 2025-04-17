// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import '@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol';
import 'hardhat/console.sol';

contract GenericTokenWithDecimals is ERC20PresetMinterPauser {
    uint8 private _decimals;

    constructor(
        string memory _name,
        string memory _symbol,
        uint256 startingSupply,
        uint8 decimalsValue
    ) ERC20PresetMinterPauser(_name, _symbol) {
        _decimals = decimalsValue;
        _mint(msg.sender, startingSupply);
    }
    
    /**
     * @dev Returns the number of decimals used to get its user representation.
     * Override the default value of 18 from ERC20.
     */
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }
}
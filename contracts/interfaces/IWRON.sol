// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IWRON {
  function deposit() external payable;

  function transfer(address to, uint256 value) external returns (bool);

  function transferFrom(address src, address dst, uint256 wad) external returns (bool);

  function withdraw(uint256) external;

  function balanceOf(address) external view returns (uint256);

  function approve(address guy, uint256 wad) external returns (bool);

  function allowance(address owner, address spender) external view returns (uint256);
}
